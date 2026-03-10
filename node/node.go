package node

import (
	"context"
	"fmt"
	"math/big"
	"sort"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/MuriData/muri-node/chain"
	"github.com/MuriData/muri-node/config"
	"github.com/MuriData/muri-node/ipfs"
	"github.com/MuriData/muri-node/prover"
	"github.com/MuriData/muri-node/storage"
	"github.com/MuriData/muri-node/types"
	"github.com/MuriData/muri-zkproof/circuits/poi"
	"github.com/MuriData/muri-zkproof/pkg/merkle"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/rs/zerolog/log"
)

// snarkScalarField is the BN254 scalar field modulus.
var snarkScalarField, _ = new(big.Int).SetString("21888242871839275222246405745257275088548364400416034343698204186575808495617", 10)

// deriveExecutionRandomness computes deterministic randomness for executeOrder
// PoI proof as keccak256(fileRoot, publicKey) % SNARK_SCALAR_FIELD.
func deriveExecutionRandomness(fileRoot, publicKey *big.Int) *big.Int {
	rootBytes := make([]byte, 32)
	pkBytes := make([]byte, 32)
	fileRoot.FillBytes(rootBytes)
	publicKey.FillBytes(pkBytes)
	hash := crypto.Keccak256(append(rootBytes, pkBytes...))
	r := new(big.Int).SetBytes(hash)
	r.Mod(r, snarkScalarField)
	return r
}

// Node is the MuriData storage provider daemon.
type Node struct {
	cfg       *config.Config
	chain     *chain.Client
	ipfs      *ipfs.Client
	prover    *prover.Prover
	store     *storage.Store
	secretKey *big.Int

	// paused controls whether the order loop accepts new orders.
	// Challenge responses continue regardless — pausing only stops new order intake.
	paused atomic.Bool

	// deregistered is set when on-chain checks detect the node is no longer valid.
	// Stops both order and maintenance loops; challenge loop continues until context cancelled.
	deregistered atomic.Bool

	// prevOrders tracks known orders for cancellation/removal detection.
	// Maps order ID string → root CID (for unpinning when the order disappears).
	prevOrders map[string]string

	// inFlightSlots tracks which challenge slots have an in-flight response.
	// Prevents duplicate goroutines from responding to the same slot concurrently.
	inFlightSlots sync.Map // slot index (int) → true
}

// Pause stops accepting new orders. Challenge responses continue.
func (n *Node) Pause() {
	n.paused.Store(true)
	log.Warn().Msg("node PAUSED — no longer accepting new orders")
}

// Resume re-enables order acceptance.
func (n *Node) Resume() {
	n.paused.Store(false)
	log.Info().Msg("node RESUMED — accepting orders again")
}

// IsPaused returns whether the node is paused.
func (n *Node) IsPaused() bool {
	return n.paused.Load()
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

	// Restore persisted order map (orderID → rootCID) for unpin tracking.
	// If no persisted state, rebuild from on-chain data.
	prevOrders, err := store.LoadOrderMap()
	if err != nil {
		log.Warn().Err(err).Msg("failed to load persisted order map, rebuilding from chain")
		prevOrders = make(map[string]string)
	}
	if len(prevOrders) == 0 {
		prevOrders = rebuildOrderMap(ctx, chainClient)
	}

	return &Node{
		cfg:        cfg,
		chain:      chainClient,
		ipfs:       ipfsClient,
		prover:     p,
		store:      store,
		secretKey:  sk,
		prevOrders: prevOrders,
	}, nil
}

// rebuildOrderMap fetches current orders from chain and builds the
// orderID → rootCID mapping. Used on first run or when persisted state is lost.
func rebuildOrderMap(ctx context.Context, c *chain.Client) map[string]string {
	orders, err := c.GetNodeOrders(ctx)
	if err != nil {
		log.Warn().Err(err).Msg("could not fetch node orders for order map rebuild")
		return make(map[string]string)
	}

	m := make(map[string]string, len(orders))
	for _, oid := range orders {
		detail, err := c.GetOrderDetails(ctx, oid)
		if err != nil {
			log.Warn().Err(err).Str("orderID", oid.String()).Msg("could not fetch order details during rebuild")
			m[oid.String()] = ""
			continue
		}
		m[oid.String()] = extractRootCID(detail.URI)
	}

	if len(m) > 0 {
		log.Info().Int("orders", len(m)).Msg("rebuilt order map from chain")
	}
	return m
}

// Run starts the daemon loops. Blocks until context is cancelled.
func (n *Node) Run(ctx context.Context) error {
	log.Info().Msg("starting daemon loops")

	// Start control socket for pause/resume commands
	ctrlCleanup := n.startControlSocket(ctx)
	defer ctrlCleanup()

	errCh := make(chan error, 4)

	// Challenge response loop (highest priority — never paused)
	go func() {
		errCh <- n.challengeLoop(ctx)
	}()

	// Order execution loop (if enabled)
	if n.cfg.AutoExecute.Enabled {
		go func() {
			errCh <- n.orderLoop(ctx)
		}()
	}

	// Maintenance + health monitoring loop
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

// challengeLoop dispatches to event-based or poll-based challenge listening.
func (n *Node) challengeLoop(ctx context.Context) error {
	if n.cfg.Chain.ListenMode == "events" && n.chain.Filterer != nil {
		return n.challengeLoopEvents(ctx)
	}
	return n.challengeLoopPoll(ctx)
}

// challengeLoopPoll polls challenge slots at a fixed interval.
func (n *Node) challengeLoopPoll(ctx context.Context) error {
	interval := n.cfg.Challenge.PollInterval.Duration
	if interval == 0 {
		interval = 4 * time.Second
	}

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	log.Info().Dur("interval", interval).Msg("challenge loop started (poll mode)")

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

// challengeLoopEvents listens for SlotChallengeIssued events via WebSocket,
// with a fallback poll ticker to catch anything events may miss.
func (n *Node) challengeLoopEvents(ctx context.Context) error {
	listener := chain.NewEventListener(n.chain)
	challengeCh, err := listener.SubscribeChallenges(ctx)
	if err != nil {
		log.Warn().Err(err).Msg("failed to subscribe challenges, falling back to poll mode")
		return n.challengeLoopPoll(ctx)
	}

	// Fallback poll at a much slower rate — events handle the fast path
	fallbackInterval := 30 * time.Second
	fallbackTicker := time.NewTicker(fallbackInterval)
	defer fallbackTicker.Stop()

	log.Info().Dur("fallback_interval", fallbackInterval).Msg("challenge loop started (event mode)")

	for {
		select {
		case <-ctx.Done():
			return nil

		case slot, ok := <-challengeCh:
			if !ok {
				// Subscription closed after max retries — fall through to poll mode
				log.Warn().Msg("challenge event channel closed, switching to poll mode")
				return n.challengeLoopPoll(ctx)
			}
			if slot.OrderID == nil || slot.OrderID.Sign() == 0 {
				continue
			}
			// Dispatch through concurrent handler so events don't block each other.
			blockNum, _ := n.chain.BlockNumber(ctx)
			go func(s types.ChallengeSlotInfo) {
				n.dispatchChallenges(ctx, []types.ChallengeSlotInfo{s}, blockNum)
			}(slot)

		case <-fallbackTicker.C:
			if err := n.checkChallenges(ctx); err != nil {
				log.Error().Err(err).Msg("challenge fallback check failed")
			}
		}
	}
}

// challengeWorkers controls how many challenges are handled concurrently.
// 2 provides a pipeline benefit (IPFS fetch for slot B overlaps with proof
// generation for slot A) without overwhelming CPU, since the prover itself
// is mutex-serialized.
const challengeWorkers = 2

// checkChallenges inspects all 5 slots, collects those targeting this node,
// sorts by deadline (most urgent first), and dispatches them concurrently.
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

	// Collect slots that need a response.
	var tasks []types.ChallengeSlotInfo
	for _, slot := range slots {
		if slot.ChallengedNode != myAddr {
			continue
		}
		if slot.OrderID == nil || slot.OrderID.Sign() == 0 {
			continue
		}
		tasks = append(tasks, slot)
	}

	if len(tasks) == 0 {
		return nil
	}

	// Sort by deadline — most urgent (lowest block) first.
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].DeadlineBlock.Cmp(tasks[j].DeadlineBlock) < 0
	})

	n.dispatchChallenges(ctx, tasks, blockNum)
	return nil
}

// dispatchChallenges runs challenge responses concurrently through a bounded
// worker pool, ordered by deadline urgency.
func (n *Node) dispatchChallenges(ctx context.Context, tasks []types.ChallengeSlotInfo, blockNum uint64) {
	taskCh := make(chan types.ChallengeSlotInfo, len(tasks))
	for _, t := range tasks {
		taskCh <- t
	}
	close(taskCh)

	workers := challengeWorkers
	if len(tasks) < workers {
		workers = len(tasks)
	}

	var wg sync.WaitGroup
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			for slot := range taskCh {
				if ctx.Err() != nil {
					return
				}

				// Skip if another goroutine is already handling this slot
				if _, loaded := n.inFlightSlots.LoadOrStore(slot.Index, true); loaded {
					log.Debug().Int("slot", slot.Index).Msg("slot already in-flight, skipping")
					continue
				}

				deadline := slot.DeadlineBlock.Uint64()
				remaining := int64(deadline) - int64(blockNum)

				if slot.IsExpired {
					log.Warn().Int("slot", slot.Index).Msg("challenge appears expired, attempting proof anyway")
				} else if remaining <= int64(n.cfg.Challenge.SafetyMargin) {
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

				n.inFlightSlots.Delete(slot.Index)
			}
		}()
	}
	wg.Wait()
}

// respondToChallenge fetches the file, generates a proof, and submits it on-chain.
// Uses cached SMT when available to avoid rebuilding the Merkle tree.
func (n *Node) respondToChallenge(ctx context.Context, slotIndex int, orderID, randomness *big.Int) error {
	if randomness == nil || randomness.Sign() == 0 {
		return fmt.Errorf("slot %d: nil or zero randomness, cannot generate proof", slotIndex)
	}

	start := time.Now()

	// 1. Get order details to find the file CID
	order, err := n.chain.GetOrderDetails(ctx, orderID)
	if err != nil {
		return fmt.Errorf("get order details: %w", err)
	}

	ref := extractIPFSRef(order.URI)
	if ref == "" {
		return fmt.Errorf("no CID found in URI: %s", order.URI)
	}

	// 2. Fetch file from IPFS (with exponential backoff retry)
	fetchStart := time.Now()
	fileData, err := n.ipfs.CatWithRetry(ctx, ref)
	if err != nil {
		return fmt.Errorf("ipfs cat %s: %w", ref, err)
	}
	fetchDur := time.Since(fetchStart)

	// 3. Build or load cached SMT, then generate proof
	proveStart := time.Now()
	smt, chunks, err := n.loadOrBuildSMT(orderID, fileData, order.RootHash)
	if err != nil {
		return fmt.Errorf("build smt: %w", err)
	}

	result, err := n.prover.GenerateProofFromSMT(n.secretKey, randomness, chunks, smt)
	if err != nil {
		return fmt.Errorf("generate proof: %w", err)
	}
	proveDur := time.Since(proveStart)

	// 4. Verify slot hasn't been re-advanced while we were proving (~20-40s)
	freshSlot, err := n.chain.GetSlotInfo(ctx, slotIndex)
	if err != nil {
		log.Warn().Err(err).Int("slot", slotIndex).Msg("pre-submit slot check failed, submitting anyway")
	} else if freshSlot.Randomness == nil || freshSlot.Randomness.Cmp(randomness) != 0 {
		return fmt.Errorf("slot %d randomness changed during proving, skipping stale proof", slotIndex)
	}

	// 5. Submit proof on-chain
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

// orderLoop dispatches to event-based or poll-based order listening.
func (n *Node) orderLoop(ctx context.Context) error {
	if n.cfg.Chain.ListenMode == "events" && n.chain.Filterer != nil {
		return n.orderLoopEvents(ctx)
	}
	return n.orderLoopPoll(ctx)
}

// orderLoopPoll polls for new orders at a fixed interval.
func (n *Node) orderLoopPoll(ctx context.Context) error {
	interval := n.cfg.AutoExecute.PollInterval.Duration
	if interval == 0 {
		interval = 30 * time.Second
	}

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	log.Info().Dur("interval", interval).Msg("order loop started (poll mode)")

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

// orderLoopEvents listens for OrderPlaced events via WebSocket,
// with a fallback poll ticker.
func (n *Node) orderLoopEvents(ctx context.Context) error {
	listener := chain.NewEventListener(n.chain)
	orderCh, err := listener.SubscribeNewOrders(ctx)
	if err != nil {
		log.Warn().Err(err).Msg("failed to subscribe orders, falling back to poll mode")
		return n.orderLoopPoll(ctx)
	}

	// Fallback poll at a slower rate than pure poll mode
	fallbackInterval := 60 * time.Second
	fallbackTicker := time.NewTicker(fallbackInterval)
	defer fallbackTicker.Stop()

	log.Info().Dur("fallback_interval", fallbackInterval).Msg("order loop started (event mode)")

	for {
		select {
		case <-ctx.Done():
			return nil

		case _, ok := <-orderCh:
			if !ok {
				log.Warn().Msg("order event channel closed, switching to poll mode")
				return n.orderLoopPoll(ctx)
			}
			// New order placed — re-evaluate all candidates
			if err := n.checkOrders(ctx); err != nil {
				log.Error().Err(err).Msg("order check failed (event-triggered)")
			}

		case <-fallbackTicker.C:
			if err := n.checkOrders(ctx); err != nil {
				log.Error().Err(err).Msg("order fallback check failed")
			}
		}
	}
}

// checkOrders finds and claims eligible orders.
func (n *Node) checkOrders(ctx context.Context) error {
	if n.paused.Load() {
		log.Debug().Msg("order check skipped (paused)")
		return nil
	}
	if n.deregistered.Load() {
		log.Debug().Msg("order check skipped (deregistered)")
		return nil
	}

	// Check node capacity
	info, err := n.chain.GetNodeInfo(ctx)
	if err != nil {
		return fmt.Errorf("get node info: %w", err)
	}

	// Apply local capacity cap if configured
	capacity := info.Capacity
	maxCap := n.cfg.Storage.MaxCapacityChunks()
	if maxCap > 0 && maxCap < capacity {
		capacity = maxCap
	}

	// Guard against uint64 underflow after slashing reduces capacity below used
	if info.Used > capacity {
		log.Warn().
			Uint64("capacity", capacity).
			Uint64("used", info.Used).
			Msg("used exceeds capacity (slashed?), skipping order check")
		return nil
	}
	availableCapacity := capacity - info.Used

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

	filled := 0
	maxToFill := n.cfg.AutoExecute.MaxOrdersToFill
	if maxToFill == 0 {
		maxToFill = 5
	}

	// Paginate through active orders
	const pageSize uint64 = 50
	var offset uint64
	for {
		page, total, err := n.chain.GetActiveOrdersPage(ctx, offset, pageSize)
		if err != nil {
			return fmt.Errorf("get active orders page(%d): %w", offset, err)
		}
		if offset == 0 {
			log.Info().
				Uint64("activeOrders", total).
				Int("existingNodeOrders", len(existing)).
				Uint64("availableCapacity", availableCapacity).
				Msg("scanning orders")
		}

		for _, orderID := range page {
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
				log.Debug().Str("orderID", orderID.String()).Uint8("filled", order.Filled).Uint8("replicas", order.Replicas).Msg("skip: order full")
				continue
			}

			// Check capacity
			if uint64(order.NumChunks) > availableCapacity {
				log.Debug().Str("orderID", orderID.String()).Uint64("chunks", uint64(order.NumChunks)).Uint64("available", availableCapacity).Msg("skip: insufficient capacity")
				continue
			}

			// Check price threshold (Price is populated by GetOrderDetails)
			if n.cfg.Storage.MinPrice > 0 && order.Price != nil && order.Price.Cmp(new(big.Int).SetUint64(n.cfg.Storage.MinPrice)) < 0 {
				log.Debug().Str("orderID", orderID.String()).Str("price", order.Price.String()).Uint64("minPrice", n.cfg.Storage.MinPrice).Msg("skip: price below threshold")
				continue
			}

			// Verify file integrity before committing
			ref := extractIPFSRef(order.URI)
			if ref == "" {
				log.Warn().Str("orderID", orderID.String()).Str("uri", order.URI).Msg("skip: no CID in URI")
				continue
			}

			// Use a bounded timeout so a slow/unreachable file doesn't stall the loop
			fetchCtx, fetchCancel := context.WithTimeout(ctx, 60*time.Second)
			fileData, err := n.ipfs.CatWithRetry(fetchCtx, ref)
			fetchCancel()
			if err != nil {
				log.Warn().Err(err).Str("orderID", orderID.String()).Str("ref", ref).Msg("skip: can't fetch file after retries")
				continue
			}

			// Build SMT and verify root matches on-chain
			smt, chunks, err := n.prover.BuildSMT(fileData)
			if err != nil {
				log.Warn().Err(err).Str("orderID", orderID.String()).Msg("skip: can't build SMT")
				continue
			}
			if order.RootHash == nil || smt.Root.Cmp(order.RootHash) != 0 {
				log.Warn().
					Str("orderID", orderID.String()).
					Str("expected", fmt.Sprintf("0x%x", order.RootHash)).
					Str("computed", fmt.Sprintf("0x%x", smt.Root)).
					Msg("skip: root hash mismatch")
				continue
			}

			// Generate PoI proof for execution (proves data possession)
			publicKey := prover.PublicKeyFromSecret(n.secretKey)
			randomness := deriveExecutionRandomness(order.RootHash, publicKey)
			proofResult, err := n.prover.GenerateProofFromSMT(n.secretKey, randomness, chunks, smt)
			if err != nil {
				log.Warn().Err(err).Str("orderID", orderID.String()).Msg("skip: proof generation failed")
				continue
			}

			// Execute order on-chain with proof
			receipt, err := n.chain.ExecuteOrder(ctx, orderID, proofResult.SolidityProof, proofResult.Commitment)
			if err != nil {
				log.Warn().Err(err).Str("orderID", orderID.String()).Msg("execute order failed")
				continue
			}

			// Pin root CID in IPFS (pins entire DAG even if order uses a subpath)
			// then advertise as provider on the DHT so other peers can discover us.
			// Run async — Pin is fast but Provide streams NDJSON and can take 30s+.
			if n.cfg.IPFS.PinFiles {
				rootCID := extractRootCID(order.URI)
				go func(cid string) {
					bgCtx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
					defer cancel()
					if err := n.ipfs.Pin(bgCtx, cid); err != nil {
						log.Warn().Err(err).Str("cid", cid).Msg("pin failed (non-fatal)")
					} else if err := n.ipfs.Provide(bgCtx, cid); err != nil {
						log.Warn().Err(err).Str("cid", cid).Msg("dht provide failed (non-fatal)")
					}
				}(rootCID)
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

		// Stop paginating if we've filled enough or reached the end
		if filled >= maxToFill {
			break
		}
		offset += pageSize
		if offset >= total {
			break
		}
	}

	return nil
}

// maintenanceLoop performs periodic housekeeping.
func (n *Node) maintenanceLoop(ctx context.Context) error {
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()

	log.Info().Msg("maintenance loop started")

	// Run immediately on startup (activate slots, claim rewards, etc.)
	n.runMaintenance(ctx)

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-ticker.C:
			n.runMaintenance(ctx)
		}
	}
}

// runMaintenance claims rewards, processes expired slots, activates idle slots,
// and monitors on-chain health (deregistration, slashing, order changes).
func (n *Node) runMaintenance(ctx context.Context) {
	// ── Health check: detect deregistration / slashing ──
	n.checkNodeHealth(ctx)
	n.detectOrderChanges(ctx)

	if n.deregistered.Load() {
		return
	}

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

	// Prune stale SMT caches for orders we no longer serve
	n.pruneStaleCache(ctx)
}

// pruneStaleCache removes cached SMT files for orders the node is no longer serving.
func (n *Node) pruneStaleCache(ctx context.Context) {
	activeOrders, err := n.chain.GetNodeOrders(ctx)
	if err != nil {
		log.Warn().Err(err).Msg("prune: failed to get node orders")
		return
	}

	activeSet := make(map[string]bool, len(activeOrders))
	for _, oid := range activeOrders {
		activeSet[oid.String()] = true
	}

	cachedIDs, err := n.store.ListCachedOrderIDs()
	if err != nil {
		log.Warn().Err(err).Msg("prune: failed to list cached order IDs")
		return
	}

	pruned := 0
	for _, id := range cachedIDs {
		if activeSet[id.String()] {
			continue
		}
		if err := n.store.DeleteTree(id); err != nil {
			log.Warn().Err(err).Str("orderID", id.String()).Msg("prune: failed to delete cache")
		} else {
			pruned++
		}
	}
	if pruned > 0 {
		log.Info().Int("pruned", pruned).Msg("pruned stale SMT caches")
	}
}

// extractIPFSRef extracts the full IPFS reference (CID or CID/path) from a URI.
// Supports formats: "ipfs://CID", "ipfs://CID/path/to/file", or raw CID.
// The returned value can be passed directly to IPFS Cat.
func extractIPFSRef(uri string) string {
	uri = strings.TrimSpace(uri)
	if strings.HasPrefix(uri, "ipfs://") {
		uri = strings.TrimPrefix(uri, "ipfs://")
	}
	uri = strings.TrimRight(uri, "/")
	if uri == "" {
		return ""
	}
	return uri
}

// extractRootCID extracts just the root CID from a URI, stripping any subpath.
// Use this for Pin/Unpin operations that apply to the whole DAG.
func extractRootCID(uri string) string {
	ref := extractIPFSRef(uri)
	if idx := strings.Index(ref, "/"); idx > 0 {
		return ref[:idx]
	}
	return ref
}

// loadOrBuildSMT attempts to load a cached SMT from disk, falling back to
// building one from the file data. If expectedRoot is non-nil, the cached tree's
// root is validated against it instead of rebuilding the full tree.
func (n *Node) loadOrBuildSMT(orderID *big.Int, fileData []byte, expectedRoot *big.Int) (*merkle.SparseMerkleTree, [][]byte, error) {
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
		// Validate cached tree root against the expected on-chain root
		if expectedRoot != nil && smt.Root.Cmp(expectedRoot) == 0 {
			log.Debug().Str("orderID", orderID.String()).Msg("using cached SMT")
			return smt, chunks, nil
		}
		log.Warn().Str("orderID", orderID.String()).Msg("cached SMT root mismatch, rebuilding")
	}

	// Build from scratch
	smt, err = merkle.GenerateSparseMerkleTree(chunks, poi.MaxTreeDepth, poi.HashChunk, n.prover.ZeroLeafHash())
	if err != nil {
		return nil, nil, fmt.Errorf("build SMT: %w", err)
	}

	// Cache for next time
	if err := n.store.SaveTree(orderID, smt); err != nil {
		log.Warn().Err(err).Msg("cache save failed")
	}

	return smt, chunks, nil
}
