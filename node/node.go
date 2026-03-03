package node

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/MuriData/muri-node/chain"
	"github.com/MuriData/muri-node/config"
	"github.com/MuriData/muri-node/ipfs"
	"github.com/MuriData/muri-node/prover"
	"github.com/MuriData/muri-node/storage"
	"github.com/MuriData/muri-zkproof/circuits/poi"
	"github.com/MuriData/muri-zkproof/pkg/merkle"
	"github.com/rs/zerolog/log"
)

// Node is the MuriData storage provider daemon.
type Node struct {
	cfg       *config.Config
	chain     *chain.Client
	ipfs      *ipfs.Client
	prover    *prover.Prover
	store     *storage.Store
	secretKey *big.Int
}

// New initializes a node: loads keys, connects to chain/IPFS, initializes prover.
func New(ctx context.Context, cfg *config.Config) (*Node, error) {
	// Load EVM private key
	privKey, err := storage.LoadPrivateKey(cfg.Node.PrivateKeyPath)
	if err != nil {
		return nil, fmt.Errorf("load private key: %w", err)
	}

	// Load ZK secret key
	sk, err := storage.LoadSecretKey(cfg.Node.SecretKeyPath)
	if err != nil {
		return nil, fmt.Errorf("load secret key: %w", err)
	}

	// Connect to chain
	chainClient, err := chain.NewClient(ctx, cfg.Chain, privKey)
	if err != nil {
		return nil, fmt.Errorf("chain client: %w", err)
	}
	log.Info().Str("addr", chainClient.Address().Hex()).Msg("connected to chain")

	// Connect to IPFS
	ipfsClient := ipfs.NewClient(cfg.IPFS)
	if err := ipfsClient.Ping(ctx); err != nil {
		return nil, fmt.Errorf("ipfs ping: %w", err)
	}
	log.Info().Str("url", cfg.IPFS.APIURL).Msg("connected to IPFS")

	// Initialize prover
	p, err := prover.NewProver(cfg.Node.KeysDir)
	if err != nil {
		return nil, fmt.Errorf("init prover: %w", err)
	}

	// Initialize store
	store, err := storage.NewStore(cfg.Node.DataDir)
	if err != nil {
		return nil, fmt.Errorf("init store: %w", err)
	}

	// Verify on-chain registration
	derivedPK := prover.PublicKeyFromSecret(sk)
	info, err := chainClient.GetNodeInfo(ctx)
	if err != nil {
		log.Warn().Err(err).Msg("could not fetch node info (node may not be registered)")
	} else if info.PublicKey != nil && info.PublicKey.Cmp(derivedPK) != 0 {
		return nil, fmt.Errorf("on-chain publicKey mismatch: expected %s, got %s",
			fmt.Sprintf("0x%x", derivedPK), fmt.Sprintf("0x%x", info.PublicKey))
	} else if info.Stake != nil && info.Stake.Sign() > 0 {
		log.Info().
			Uint64("capacity", info.Capacity).
			Uint64("used", info.Used).
			Str("publicKey", fmt.Sprintf("0x%x", info.PublicKey)).
			Msg("node registered on-chain")
	}

	return &Node{
		cfg:       cfg,
		chain:     chainClient,
		ipfs:      ipfsClient,
		prover:    p,
		store:     store,
		secretKey: sk,
	}, nil
}

// Run starts the daemon loops. Blocks until context is cancelled.
func (n *Node) Run(ctx context.Context) error {
	log.Info().Msg("starting daemon loops")

	errCh := make(chan error, 3)

	// Challenge response loop (highest priority)
	go func() {
		errCh <- n.challengeLoop(ctx)
	}()

	// Order execution loop (if enabled)
	if n.cfg.AutoExecute.Enabled {
		go func() {
			errCh <- n.orderLoop(ctx)
		}()
	}

	// Maintenance loop
	go func() {
		errCh <- n.maintenanceLoop(ctx)
	}()

	// Wait for context cancellation or fatal error
	select {
	case <-ctx.Done():
		log.Info().Msg("shutting down daemon loops")
		return nil
	case err := <-errCh:
		return fmt.Errorf("daemon loop failed: %w", err)
	}
}

// challengeLoop polls challenge slots and responds to challenges targeting this node.
func (n *Node) challengeLoop(ctx context.Context) error {
	interval := n.cfg.Challenge.PollInterval.Duration
	if interval == 0 {
		interval = 4 * time.Second
	}

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	log.Info().Dur("interval", interval).Msg("challenge loop started")

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-ticker.C:
			if err := n.checkChallenges(ctx); err != nil {
				log.Error().Err(err).Msg("challenge check failed")
			}
		}
	}
}

// checkChallenges inspects all 5 slots and responds to any targeting this node.
func (n *Node) checkChallenges(ctx context.Context) error {
	slots, err := n.chain.GetAllSlotInfo(ctx)
	if err != nil {
		return fmt.Errorf("get slot info: %w", err)
	}

	myAddr := n.chain.Address()
	blockNum, err := n.chain.BlockNumber(ctx)
	if err != nil {
		return fmt.Errorf("get block number: %w", err)
	}

	for _, slot := range slots {
		if slot.ChallengedNode != myAddr {
			continue
		}
		if slot.OrderID == nil || slot.OrderID.Sign() == 0 {
			continue
		}

		// Even if the view reports expired, attempt submission anyway — the on-chain
		// submitProof sweeps expired slots first, and there may be a block boundary
		// race where submission still succeeds. Let the contract be the arbiter.
		if slot.IsExpired {
			log.Warn().Int("slot", slot.Index).Msg("challenge appears expired, attempting proof anyway")
		}

		deadline := slot.DeadlineBlock.Uint64()
		remaining := int64(deadline) - int64(blockNum)
		if remaining <= int64(n.cfg.Challenge.SafetyMargin) && !slot.IsExpired {
			log.Warn().
				Int("slot", slot.Index).
				Int64("blocksRemaining", remaining).
				Msg("close to deadline, attempting anyway")
		}

		log.Info().
			Int("slot", slot.Index).
			Str("orderID", slot.OrderID.String()).
			Int64("blocksRemaining", remaining).
			Msg("responding to challenge")

		if err := n.respondToChallenge(ctx, slot.Index, slot.OrderID, slot.Randomness); err != nil {
			log.Error().Err(err).Int("slot", slot.Index).Msg("challenge response failed")
		}
	}

	return nil
}

// respondToChallenge fetches the file, generates a proof, and submits it on-chain.
// Uses cached SMT when available to avoid rebuilding the Merkle tree.
func (n *Node) respondToChallenge(ctx context.Context, slotIndex int, orderID, randomness *big.Int) error {
	start := time.Now()

	// 1. Get order details to find the file CID
	order, err := n.chain.GetOrderDetails(ctx, orderID)
	if err != nil {
		return fmt.Errorf("get order details: %w", err)
	}

	cid := extractCID(order.URI)
	if cid == "" {
		return fmt.Errorf("no CID found in URI: %s", order.URI)
	}

	// 2. Fetch file from IPFS
	fetchStart := time.Now()
	fileData, err := n.ipfs.Cat(ctx, cid)
	if err != nil {
		return fmt.Errorf("ipfs cat %s: %w", cid, err)
	}
	fetchDur := time.Since(fetchStart)

	// 3. Build or load cached SMT, then generate proof
	proveStart := time.Now()
	smt, chunks, err := n.loadOrBuildSMT(orderID, fileData)
	if err != nil {
		return fmt.Errorf("build smt: %w", err)
	}

	result, err := n.prover.GenerateProofFromSMT(n.secretKey, randomness, chunks, smt)
	if err != nil {
		return fmt.Errorf("generate proof: %w", err)
	}
	proveDur := time.Since(proveStart)

	// 4. Submit proof on-chain
	receipt, err := n.chain.SubmitProof(ctx, slotIndex, result.SolidityProof, result.Commitment)
	if err != nil {
		return fmt.Errorf("submit proof: %w", err)
	}

	totalDur := time.Since(start)
	log.Info().
		Int("slot", slotIndex).
		Str("orderID", orderID.String()).
		Dur("fetch", fetchDur).
		Dur("prove", proveDur).
		Dur("total", totalDur).
		Str("tx", receipt.TxHash.Hex()).
		Msg("proof submitted successfully")

	return nil
}

// orderLoop polls for new orders and auto-executes matching ones.
func (n *Node) orderLoop(ctx context.Context) error {
	interval := n.cfg.AutoExecute.PollInterval.Duration
	if interval == 0 {
		interval = 30 * time.Second
	}

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	log.Info().Dur("interval", interval).Msg("order loop started")

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-ticker.C:
			if err := n.checkOrders(ctx); err != nil {
				log.Error().Err(err).Msg("order check failed")
			}
		}
	}
}

// checkOrders finds and claims eligible orders.
func (n *Node) checkOrders(ctx context.Context) error {
	// Check node capacity
	info, err := n.chain.GetNodeInfo(ctx)
	if err != nil {
		return fmt.Errorf("get node info: %w", err)
	}

	// Guard against uint64 underflow after slashing reduces capacity below used
	if info.Used > info.Capacity {
		log.Warn().
			Uint64("capacity", info.Capacity).
			Uint64("used", info.Used).
			Msg("used exceeds capacity (slashed?), skipping order check")
		return nil
	}
	availableCapacity := info.Capacity - info.Used

	if availableCapacity == 0 {
		log.Debug().Msg("no available capacity, skipping order check")
		return nil
	}

	// Get existing node orders to avoid duplicates
	existingOrders, err := n.chain.GetNodeOrders(ctx)
	if err != nil {
		return fmt.Errorf("get node orders: %w", err)
	}
	existing := make(map[string]bool, len(existingOrders))
	for _, oid := range existingOrders {
		existing[oid.String()] = true
	}

	// Get active orders
	activeOrders, err := n.chain.GetActiveOrders(ctx)
	if err != nil {
		return fmt.Errorf("get active orders: %w", err)
	}

	filled := 0
	maxToFill := n.cfg.AutoExecute.MaxOrdersToFill
	if maxToFill == 0 {
		maxToFill = 5
	}

	for _, orderID := range activeOrders {
		if filled >= maxToFill {
			break
		}
		if existing[orderID.String()] {
			continue
		}

		order, err := n.chain.GetOrderDetails(ctx, orderID)
		if err != nil {
			log.Warn().Err(err).Str("orderID", orderID.String()).Msg("skip: can't get order details")
			continue
		}

		// Check if order is full
		if order.Filled >= order.Replicas {
			continue
		}

		// Check capacity
		if uint64(order.NumChunks) > availableCapacity {
			continue
		}

		// Check price threshold (Price is populated by GetOrderDetails)
		if n.cfg.Storage.MinPrice > 0 && order.Price != nil && order.Price.Uint64() < n.cfg.Storage.MinPrice {
			continue
		}

		// Verify file integrity before committing
		cid := extractCID(order.URI)
		if cid == "" {
			log.Warn().Str("orderID", orderID.String()).Str("uri", order.URI).Msg("skip: no CID in URI")
			continue
		}

		fileData, err := n.ipfs.Cat(ctx, cid)
		if err != nil {
			log.Warn().Err(err).Str("orderID", orderID.String()).Str("cid", cid).Msg("skip: can't fetch file")
			continue
		}

		// Build SMT and verify root matches on-chain
		smt, _ := n.prover.BuildSMT(fileData)
		if smt.Root.Cmp(order.RootHash) != 0 {
			log.Warn().
				Str("orderID", orderID.String()).
				Str("expected", fmt.Sprintf("0x%x", order.RootHash)).
				Str("computed", fmt.Sprintf("0x%x", smt.Root)).
				Msg("skip: root hash mismatch")
			continue
		}

		// Execute order on-chain
		receipt, err := n.chain.ExecuteOrder(ctx, orderID)
		if err != nil {
			log.Warn().Err(err).Str("orderID", orderID.String()).Msg("execute order failed")
			continue
		}

		// Pin file in IPFS
		if n.cfg.IPFS.PinFiles {
			if err := n.ipfs.Pin(ctx, cid); err != nil {
				log.Warn().Err(err).Str("cid", cid).Msg("pin failed (non-fatal)")
			}
		}

		// Cache Merkle tree for future challenge responses
		if err := n.store.SaveTree(orderID, smt); err != nil {
			log.Warn().Err(err).Str("orderID", orderID.String()).Msg("cache tree failed (non-fatal)")
		}

		availableCapacity -= uint64(order.NumChunks)
		filled++

		log.Info().
			Str("orderID", orderID.String()).
			Str("tx", receipt.TxHash.Hex()).
			Uint32("chunks", order.NumChunks).
			Msg("order claimed")
	}

	return nil
}

// maintenanceLoop performs periodic housekeeping.
func (n *Node) maintenanceLoop(ctx context.Context) error {
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()

	log.Info().Msg("maintenance loop started")

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-ticker.C:
			n.runMaintenance(ctx)
		}
	}
}

// runMaintenance claims rewards, processes expired slots, activates idle slots.
// Checks on-chain state before submitting transactions to avoid wasting gas.
func (n *Node) runMaintenance(ctx context.Context) {
	// Claim rewards if any are available
	claimable, err := n.chain.GetClaimableRewards(ctx)
	if err != nil {
		log.Error().Err(err).Msg("get claimable rewards failed")
	} else if claimable != nil && claimable.Sign() > 0 {
		receipt, err := n.chain.ClaimRewards(ctx)
		if err != nil {
			log.Error().Err(err).Msg("claim rewards failed")
		} else {
			log.Info().
				Str("amount", claimable.String()).
				Str("tx", receipt.TxHash.Hex()).
				Msg("rewards claimed")
		}
	}

	// Check slot state before submitting maintenance transactions
	slots, err := n.chain.GetAllSlotInfo(ctx)
	if err != nil {
		log.Error().Err(err).Msg("get slot info for maintenance failed")
		return
	}

	hasExpired := false
	hasIdle := false
	for _, slot := range slots {
		if slot.IsExpired {
			hasExpired = true
		}
		if slot.OrderID == nil || slot.OrderID.Sign() == 0 {
			hasIdle = true
		}
	}

	// Process expired slots for reporter rewards (only if there are expired ones)
	if hasExpired {
		if receipt, err := n.chain.ProcessExpiredSlots(ctx); err != nil {
			log.Warn().Err(err).Msg("process expired slots failed")
		} else {
			log.Info().Str("tx", receipt.TxHash.Hex()).Msg("processed expired slots")
		}
	}

	// Activate idle slots (only if there are idle ones)
	if hasIdle {
		if receipt, err := n.chain.ActivateSlots(ctx); err != nil {
			log.Debug().Err(err).Msg("activate slots failed (may have no challengeable orders)")
		} else {
			log.Info().Str("tx", receipt.TxHash.Hex()).Msg("activated idle slots")
		}
	}
}

// extractCID extracts an IPFS CID from a URI.
// Supports formats: "ipfs://CID", "ipfs://CID/path", or raw CID.
func extractCID(uri string) string {
	uri = strings.TrimSpace(uri)
	if strings.HasPrefix(uri, "ipfs://") {
		uri = strings.TrimPrefix(uri, "ipfs://")
	}
	// Take first path segment as CID
	if idx := strings.Index(uri, "/"); idx > 0 {
		uri = uri[:idx]
	}
	if uri == "" {
		return ""
	}
	return uri
}

// loadOrBuildSMT attempts to load a cached SMT from disk, falling back to
// building one from the file data. Always returns chunks consistent with the SMT.
func (n *Node) loadOrBuildSMT(orderID *big.Int, fileData []byte) (*merkle.SparseMerkleTree, [][]byte, error) {
	// Always split chunks from the current file data — these chunks must be
	// consistent with whichever SMT we return (cached or freshly built).
	chunks := merkle.SplitIntoChunks(fileData, poi.FileSize)

	// Try loading from cache
	smt, err := n.store.LoadTree(orderID, n.prover.ZeroLeafHash())
	if err != nil {
		log.Warn().Err(err).Str("orderID", orderID.String()).Msg("cache load failed, rebuilding")
		smt = nil
	}

	if smt != nil {
		// Validate cached tree is consistent with current file data
		freshSMT := merkle.GenerateSparseMerkleTree(chunks, poi.MaxTreeDepth, poi.HashChunk, n.prover.ZeroLeafHash())
		if smt.Root.Cmp(freshSMT.Root) == 0 {
			log.Debug().Str("orderID", orderID.String()).Msg("using cached SMT")
			return smt, chunks, nil
		}
		log.Warn().Str("orderID", orderID.String()).Msg("cached SMT root mismatch, rebuilding")
		smt = freshSMT
	} else {
		// Build from scratch
		smt = merkle.GenerateSparseMerkleTree(chunks, poi.MaxTreeDepth, poi.HashChunk, n.prover.ZeroLeafHash())
	}

	// Cache for next time
	if err := n.store.SaveTree(orderID, smt); err != nil {
		log.Warn().Err(err).Msg("cache save failed")
	}

	return smt, chunks, nil
}
