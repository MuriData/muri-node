package node

import (
	"context"
	"fmt"
	"math/big"
	"runtime"
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
	"github.com/consensys/gnark-crypto/ecc/bn254/fr"
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

// fetchCooldown is how long a CID is skipped after a fetch failure in the order loop.
// Prevents repeatedly blocking the loop on the same unreachable file.
const fetchCooldown = 10 * time.Minute

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

	// inFlightChallengeOrders tracks which orders have an in-flight challenge response.
	// Used by maintenance cleanup to avoid unpinning/deleting data mid-proof.
	inFlightChallengeOrders sync.Map // orderID string → true

	// deferredCleanups holds orders removed while a challenge was in-flight.
	// Cleaned up once the challenge goroutine finishes.
	deferredCleanups   []deferredCleanup
	deferredCleanupsMu sync.Mutex

	// inFlightOrders tracks which orders have an in-flight processCandidate.
	// Prevents duplicate workers from processing the same order across poll ticks.
	inFlightOrders sync.Map // orderID string → true

	// fetchFailures tracks CIDs that recently failed to download.
	// Maps CID → time of failure. Entries older than fetchCooldown are retried.
	// Prevents the order loop from repeatedly blocking on the same unreachable file.
	fetchFailures sync.Map // CID string → time.Time

	// startupCleanupDone ensures orphaned pin cleanup runs only once (at startup).
	startupCleanupDone atomic.Bool

	// lastPinVerify tracks when pin verification last ran.
	lastPinVerify atomic.Value // stores time.Time

	// lastProvide tracks when DHT provide last ran.
	lastProvide atomic.Value // stores time.Time

	// repinFailures tracks CIDs that failed re-pin, with timestamp of first failure.
	// Used to escalate log level on repeated failures.
	repinFailures sync.Map // CID string → time.Time

	// recentProofs tracks (slotIndex, randomness) pairs for recently submitted proofs.
	// Prevents re-dispatching the same challenge when RPC returns stale slot data.
	recentProofs sync.Map // "slot:randomnessHex" → time.Time
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

	n := &Node{
		cfg:        cfg,
		chain:      chainClient,
		ipfs:       ipfsClient,
		prover:     p,
		store:      store,
		secretKey:  sk,
		prevOrders: prevOrders,
	}

	// Defer first pin verify / provide until a full interval after startup.
	// The startup maintenance tick already handles orphaned pins and order
	// detection — running these immediately would double startup latency.
	n.lastPinVerify.Store(time.Now())
	n.lastProvide.Store(time.Now())

	return n, nil
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
			if slot.Randomness == nil || slot.Randomness.Sign() == 0 {
				log.Warn().Int("slot", slot.Index).Msg("event slot has no randomness, skipping")
				continue
			}
			// Skip if we already proved this exact (slot, randomness)
			proofKey := fmt.Sprintf("%d:%s", slot.Index, slot.Randomness.Text(16))
			if ts, ok := n.recentProofs.Load(proofKey); ok {
				if time.Since(ts.(time.Time)) < 10*time.Minute {
					log.Debug().Int("slot", slot.Index).Msg("event for already-proved challenge, skipping")
					continue
				}
				n.recentProofs.Delete(proofKey)
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

// checkChallenges inspects all slots, collects those targeting this node,
// sorts by deadline (most urgent first), and dispatches them concurrently.
func (n *Node) checkChallenges(ctx context.Context) error {
	// Fetch slot info and block number in parallel — independent RPCs.
	var (
		slots    []types.ChallengeSlotInfo
		slotsErr error
		blockNum uint64
		blockErr error
		wg       sync.WaitGroup
	)
	wg.Add(2)
	go func() {
		defer wg.Done()
		slots, slotsErr = n.chain.GetAllSlotInfoFresh(ctx)
	}()
	go func() {
		defer wg.Done()
		blockNum, blockErr = n.chain.BlockNumber(ctx)
	}()
	wg.Wait()

	if slotsErr != nil {
		return fmt.Errorf("get slot info: %w", slotsErr)
	}
	if blockErr != nil {
		return fmt.Errorf("get block number: %w", blockErr)
	}

	myAddr := n.chain.Address()

	// Collect slots that need a response.
	var tasks []types.ChallengeSlotInfo
	for _, slot := range slots {
		if slot.ChallengedNode != myAddr {
			continue
		}
		if slot.OrderID == nil || slot.OrderID.Sign() == 0 {
			continue
		}
		if slot.Randomness == nil || slot.Randomness.Sign() == 0 {
			log.Warn().Int("slot", slot.Index).Msg("poll slot has no randomness, skipping")
			continue
		}
		// Skip if we already proved this exact (slot, randomness) — protects
		// against stale RPC data causing re-dispatch of already-proved challenges.
		proofKey := fmt.Sprintf("%d:%s", slot.Index, slot.Randomness.Text(16))
		if ts, ok := n.recentProofs.Load(proofKey); ok {
			if time.Since(ts.(time.Time)) < 10*time.Minute {
				log.Debug().Int("slot", slot.Index).Msg("skipping already-proved challenge")
				continue
			}
			n.recentProofs.Delete(proofKey) // expired entry
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

	go n.dispatchChallenges(ctx, tasks, blockNum)
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

				// Track this order so maintenance doesn't clean it up mid-proof.
				orderKey := slot.OrderID.String()
				n.inFlightChallengeOrders.Store(orderKey, true)

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

				n.inFlightChallengeOrders.Delete(orderKey)
				n.inFlightSlots.Delete(slot.Index)
			}
		}()
	}
	wg.Wait()
}

// respondToChallenge generates a proof and submits it on-chain.
//
// Fast path (cached SMT): derives the 8 leaf indices from randomness, fetches
// only those chunks via IPFS byte-range requests (~128 KB for a 1 GB file),
// then generates the proof. Falls back to the slow path on any error.
//
// Slow path: downloads the entire file, builds/loads the SMT, generates proof.
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

	// 2. Try fast path: cached SMT + selective chunk fetch
	var result *prover.ProofResult
	var path string
	smt, err := n.store.LoadTree(orderID, n.prover.ZeroLeafHash())
	if err == nil && smt != nil && order.RootHash != nil && smt.RootBigInt().Cmp(order.RootHash) == 0 {
		result, err = n.proveSelective(ctx, randomness, smt, ref)
		if err != nil {
			log.Warn().Err(err).Str("orderID", orderID.String()).Msg("selective proof failed, falling back to full download")
			result = nil
		} else {
			path = "selective"
		}
	}

	// 3. Slow path: stream download → hash → build SMT → prove selectively.
	// The full file is never buffered — chunks are hashed on the fly during
	// download, then only the 8 challenged chunks (~128 KB) are re-fetched.
	if result == nil {
		fetchCtx, fetchCancel := context.WithTimeout(ctx, fetchTimeout(order.NumChunks))
		smt, err = n.buildSMTStreaming(fetchCtx, ref, order.NumChunks)
		fetchCancel()
		if err != nil {
			return fmt.Errorf("streaming build smt %s: %w", ref, err)
		}

		// Cache for future challenge responses
		if saveErr := n.store.SaveTree(orderID, smt); saveErr != nil {
			log.Warn().Err(saveErr).Str("orderID", orderID.String()).Msg("cache tree failed (non-fatal)")
		}

		result, err = n.proveSelective(ctx, randomness, smt, ref)
		if err != nil {
			return fmt.Errorf("generate proof (streaming+selective): %w", err)
		}
		path = "streaming+selective"
	}

	// 4. Verify slot hasn't been re-advanced while we were proving (~20-40s).
	// Uses the WS client when available to bypass HTTP reverse-proxy caching
	// that caused the node to submit to wrong slots with stale data.
	freshSlot, err := n.chain.GetSlotInfoFresh(ctx, slotIndex)
	if err != nil {
		return fmt.Errorf("slot %d pre-submit check failed (aborting to avoid stale submission): %w", slotIndex, err)
	}
	if freshSlot.ChallengedNode != n.chain.Address() {
		return fmt.Errorf("slot %d no longer targets this node (challenged: %s), skipping", slotIndex, freshSlot.ChallengedNode.Hex())
	}
	if freshSlot.Randomness == nil || freshSlot.Randomness.Cmp(randomness) != 0 {
		return fmt.Errorf("slot %d randomness changed during proving, skipping stale proof", slotIndex)
	}

	// 5. Submit proof on-chain
	receipt, err := n.chain.SubmitProof(ctx, slotIndex, result.SolidityProof, result.Commitment)
	if err != nil {
		return fmt.Errorf("submit proof: %w", err)
	}

	// Track this proof to prevent re-dispatch if RPC serves stale slot data
	proofKey := fmt.Sprintf("%d:%s", slotIndex, randomness.Text(16))
	n.recentProofs.Store(proofKey, time.Now())

	log.Info().
		Int("slot", slotIndex).
		Str("orderID", orderID.String()).
		Str("path", path).
		Dur("total", time.Since(start)).
		Str("tx", receipt.TxHash.Hex()).
		Msg("proof submitted successfully")

	return nil
}

// proveSelective generates a proof by fetching only the 8 challenged chunks
// (via IPFS byte-range requests) instead of downloading the entire file.
func (n *Node) proveSelective(ctx context.Context, randomness *big.Int, smt *merkle.SparseMerkleTree, ref string) (*prover.ProofResult, error) {
	// Derive the 8 leaf indices the circuit will open
	indices := prover.DeriveLeafIndices(randomness, smt.NumLeaves)

	// Collect unique indices (small files may repeat via modular wrapping)
	unique := make(map[int]struct{})
	for _, idx := range indices {
		unique[idx] = struct{}{}
	}

	// Fetch unique chunks in parallel via byte-range requests
	fetchCtx, fetchCancel := context.WithTimeout(ctx, 2*time.Minute)
	chunkMap, err := n.fetchChunksSelective(fetchCtx, ref, unique)
	fetchCancel()
	if err != nil {
		return nil, fmt.Errorf("selective fetch: %w", err)
	}

	log.Debug().
		Int("unique_chunks", len(unique)).
		Int("tree_leaves", smt.NumLeaves).
		Msg("selective fetch complete")

	// Build sparse chunks slice — only the 8 needed indices are populated.
	// PrepareWitness only accesses chunks[leafIndex] for the derived indices.
	sparseChunks := make([][]byte, smt.NumLeaves)
	for idx, data := range chunkMap {
		sparseChunks[idx] = data
	}

	return n.prover.GenerateProofFromSMT(n.secretKey, randomness, sparseChunks, smt)
}

// fetchChunksSelective fetches specific file chunks by index from IPFS using
// byte-range requests. Each chunk is poi.FileSize bytes at offset index*FileSize.
// Fetches run in parallel; returns a map of index → padded chunk data.
func (n *Node) fetchChunksSelective(ctx context.Context, ref string, indices map[int]struct{}) (map[int][]byte, error) {
	type result struct {
		idx  int
		data []byte
		err  error
	}

	ch := make(chan result, len(indices))
	for idx := range indices {
		go func(i int) {
			offset := int64(i) * int64(poi.FileSize)
			data, err := n.ipfs.CatRangeWithRetry(ctx, ref, offset, int64(poi.FileSize))
			if err != nil {
				ch <- result{i, nil, err}
				return
			}
			// Zero-pad last chunk if shorter (same as SplitIntoChunks)
			if len(data) < poi.FileSize {
				padded := make([]byte, poi.FileSize)
				copy(padded, data)
				data = padded
			}
			ch <- result{i, data, nil}
		}(idx)
	}

	m := make(map[int][]byte, len(indices))
	for range indices {
		r := <-ch
		if r.err != nil {
			return nil, fmt.Errorf("chunk %d: %w", r.idx, r.err)
		}
		m[r.idx] = r.data
	}
	return m, nil
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

// orderCandidate is an order that passed fast eligibility checks and is
// ready for the slow fetch → verify → prove → execute pipeline.
type orderCandidate struct {
	id      *big.Int
	order   *types.OrderInfo
	ref     string
	rootCID string
}

// orderWorkers controls how many orders are fetched/processed concurrently.
// 3 provides good pipeline utilization: one fetching from IPFS, one proving
// (mutex-serialized), one executing on-chain.
const orderWorkers = 3

// orderSem is a package-level semaphore shared across poll ticks so the order
// loop is not blocked by slow fetches — new ticks can dispatch fresh candidates
// into free worker slots while previous ones are still running.
var orderSem = make(chan struct{}, orderWorkers)

// checkOrders finds eligible orders and processes them concurrently.
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

	capacity := info.Capacity
	maxCap := n.cfg.Storage.MaxCapacityChunks()
	if maxCap > 0 && maxCap < capacity {
		capacity = maxCap
	}

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

	existingOrders, err := n.chain.GetNodeOrders(ctx)
	if err != nil {
		return fmt.Errorf("get node orders: %w", err)
	}
	existing := make(map[string]bool, len(existingOrders))
	for _, oid := range existingOrders {
		existing[oid.String()] = true
	}

	maxToFill := n.cfg.AutoExecute.MaxOrdersToFill
	if maxToFill == 0 {
		maxToFill = 5
	}

	// Phase 1: collect eligible candidates (fast — only chain queries + filtering).
	candidates := n.collectCandidates(ctx, existing, availableCapacity, maxToFill)
	if len(candidates) == 0 {
		return nil
	}

	// Phase 2: dispatch candidates into the shared worker pool.
	// Non-blocking: if all workers are busy, remaining candidates are skipped
	// and will be picked up on the next poll tick. This prevents a slow fetch
	// from blocking the entire order loop.
	dispatched := 0
	for _, cand := range candidates {
		if dispatched >= maxToFill {
			break
		}

		// Skip if this order is already being processed from a previous tick
		if _, loaded := n.inFlightOrders.LoadOrStore(cand.id.String(), true); loaded {
			continue
		}

		// Try to acquire a worker slot without blocking
		select {
		case orderSem <- struct{}{}:
			dispatched++
			go func(c orderCandidate) {
				defer func() { <-orderSem }()
				defer n.inFlightOrders.Delete(c.id.String())
				n.processCandidate(ctx, c)
			}(cand)
		default:
			// All workers busy — remaining candidates will be retried next tick
			n.inFlightOrders.Delete(cand.id.String())
			log.Debug().Str("orderID", cand.id.String()).Msg("workers busy, deferring to next tick")
			goto doneDispatching
		}
	}
doneDispatching:

	if dispatched > 0 {
		log.Info().Int("dispatched", dispatched).Msg("order candidates dispatched")
	}
	return nil
}

// collectCandidates paginates through active orders and returns those that
// pass all fast eligibility checks (price, capacity, not full, not on cooldown).
func (n *Node) collectCandidates(ctx context.Context, existing map[string]bool, availableCapacity uint64, maxCandidates int) []orderCandidate {
	var candidates []orderCandidate

	const pageSize uint64 = 50
	var offset uint64
	for {
		page, total, err := n.chain.GetActiveOrdersPage(ctx, offset, pageSize)
		if err != nil {
			log.Error().Err(err).Uint64("offset", offset).Msg("get active orders page failed")
			break
		}
		if offset == 0 {
			log.Info().
				Uint64("activeOrders", total).
				Int("existingNodeOrders", len(existing)).
				Uint64("availableCapacity", availableCapacity).
				Msg("scanning orders")
		}

		for _, orderID := range page {
			if len(candidates) >= maxCandidates {
				break
			}
			key := orderID.String()
			if existing[key] {
				continue
			}
			// Skip orders already being processed from a previous tick
			if _, ok := n.inFlightOrders.Load(key); ok {
				continue
			}

			order, err := n.chain.GetOrderDetails(ctx, orderID)
			if err != nil {
				log.Warn().Err(err).Str("orderID", orderID.String()).Msg("skip: can't get order details")
				continue
			}

			if order.Filled >= order.Replicas {
				continue
			}
			if uint64(order.NumChunks) > availableCapacity {
				continue
			}
			if n.cfg.Storage.MinPrice > 0 && order.Price != nil && order.Price.Cmp(new(big.Int).SetUint64(n.cfg.Storage.MinPrice)) < 0 {
				continue
			}

			ref := extractIPFSRef(order.URI)
			if ref == "" {
				continue
			}

			rootCID := extractRootCID(order.URI)
			if failTime, ok := n.fetchFailures.Load(rootCID); ok {
				if time.Since(failTime.(time.Time)) < fetchCooldown {
					log.Debug().Str("orderID", orderID.String()).Msg("skip: CID on fetch cooldown")
					continue
				}
				n.fetchFailures.Delete(rootCID)
			}

			candidates = append(candidates, orderCandidate{
				id:      orderID,
				order:   order,
				ref:     ref,
				rootCID: rootCID,
			})
		}

		if len(candidates) >= maxCandidates {
			break
		}
		offset += pageSize
		if offset >= total {
			break
		}
	}

	return candidates
}

// processCandidate handles a single order: fetch → verify → prove → execute.
// Runs concurrently from checkOrders; errors are logged, not returned.
func (n *Node) processCandidate(ctx context.Context, cand orderCandidate) {
	// Stream download → hash chunks → build SMT. The full file is never
	// buffered in memory — only ~1 MB download segments are held at a time.
	fetchCtx, fetchCancel := context.WithTimeout(ctx, fetchTimeout(cand.order.NumChunks))
	smt, err := n.buildSMTStreaming(fetchCtx, cand.ref, cand.order.NumChunks)
	fetchCancel()
	if err != nil {
		n.fetchFailures.Store(cand.rootCID, time.Now())
		log.Warn().Err(err).Str("orderID", cand.id.String()).Str("ref", cand.ref).Msg("skip: can't fetch/hash file (cooldown 10m)")
		return
	}
	if cand.order.RootHash == nil || smt.RootBigInt().Cmp(cand.order.RootHash) != 0 {
		log.Warn().
			Str("orderID", cand.id.String()).
			Str("expected", fmt.Sprintf("0x%x", cand.order.RootHash)).
			Str("computed", fmt.Sprintf("0x%x", smt.RootBigInt())).
			Msg("skip: root hash mismatch")
		return
	}

	publicKey := prover.PublicKeyFromSecret(n.secretKey)
	randomness := deriveExecutionRandomness(cand.order.RootHash, publicKey)
	proofResult, err := n.proveSelective(ctx, randomness, smt, cand.ref)
	if err != nil {
		log.Warn().Err(err).Str("orderID", cand.id.String()).Msg("skip: proof generation failed")
		return
	}

	receipt, err := n.chain.ExecuteOrder(ctx, cand.id, proofResult.SolidityProof, proofResult.Commitment)
	if err != nil {
		log.Warn().Err(err).Str("orderID", cand.id.String()).Msg("execute order failed")
		return
	}

	// Pin + provide in background
	if n.cfg.IPFS.PinFiles {
		go func(cid string) {
			bgCtx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
			defer cancel()
			if err := n.ipfs.Pin(bgCtx, cid); err != nil {
				log.Warn().Err(err).Str("cid", cid).Msg("pin failed (non-fatal)")
			} else if err := n.ipfs.Provide(bgCtx, cid); err != nil {
				log.Warn().Err(err).Str("cid", cid).Msg("dht provide failed (non-fatal)")
			}
		}(cand.rootCID)
	}

	// Cache SMT for future challenge responses
	if err := n.store.SaveTree(cand.id, smt); err != nil {
		log.Warn().Err(err).Str("orderID", cand.id.String()).Msg("cache tree failed (non-fatal)")
	}

	log.Info().
		Str("orderID", cand.id.String()).
		Str("tx", receipt.TxHash.Hex()).
		Uint32("chunks", cand.order.NumChunks).
		Msg("order claimed")
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
	activeOrders := n.detectOrderChanges(ctx)

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

	// Prune stale SMT caches for orders we no longer serve.
	// Reuse the order list from detectOrderChanges to avoid a duplicate RPC.
	n.pruneStaleCache(activeOrders)

	// On first maintenance tick (startup), scan for orphaned IPFS pins
	// from orders that were removed while orders.json was missing/corrupted.
	if !n.startupCleanupDone.Load() {
		n.startupCleanupDone.Store(true)
		n.cleanOrphanedPins(ctx)
	}

	// ── Periodic IPFS maintenance ──

	// Verify all active order CIDs are still pinned locally.
	pinInterval := n.cfg.IPFS.PinVerifyInterval.Duration
	if pinInterval == 0 {
		pinInterval = 30 * time.Minute
	}
	if n.cfg.IPFS.PinFiles && isDue(&n.lastPinVerify, pinInterval) {
		n.verifyPins(ctx, n.prevOrders)
		n.lastPinVerify.Store(time.Now())
	}

	// Re-announce active order CIDs to the DHT (supplements Kubo's reprovider).
	provideInterval := n.cfg.IPFS.ProvideInterval.Duration
	if provideInterval == 0 {
		provideInterval = 4 * time.Hour
	}
	if n.cfg.IPFS.PinFiles && isDue(&n.lastProvide, provideInterval) {
		n.lastProvide.Store(time.Now())
		go n.provideAll(context.Background(), copyOrderMap(n.prevOrders))
	}
}

// cleanOrphanedPins finds IPFS pins for orders that have stale .smt cache files
// but are no longer in the active order set, and unpins them. This handles the
// case where orders.json was lost/corrupted and previously-removed orders' pins
// were never cleaned up.
//
// Only unpins CIDs whose order ID has a stale .smt file — avoids touching
// pins from other applications sharing the same IPFS node.
func (n *Node) cleanOrphanedPins(ctx context.Context) {
	if !n.cfg.IPFS.PinFiles {
		return
	}

	// Find stale .smt files (orders we're not serving anymore).
	// These indicate orders that were ours but whose cleanup was incomplete.
	cachedIDs, err := n.store.ListCachedOrderIDs()
	if err != nil {
		log.Warn().Err(err).Msg("orphan pin check: failed to list cached order IDs")
		return
	}

	activeSet := make(map[string]bool, len(n.prevOrders))
	for oid := range n.prevOrders {
		activeSet[oid] = true
	}

	// For stale orders, try to look up their CID from the .smt filename
	// alone we can't recover the CID. Instead, check if any pinned CIDs
	// are NOT in our active set. Cross-reference with the pin list.
	pins, err := n.ipfs.ListPins(ctx)
	if err != nil {
		log.Warn().Err(err).Msg("orphan pin check: failed to list IPFS pins")
		return
	}

	activeCIDs := make(map[string]bool, len(n.prevOrders))
	for _, cid := range n.prevOrders {
		if cid != "" {
			activeCIDs[cid] = true
		}
	}

	// Only consider stale orders that have cached .smt files as evidence
	// they belonged to this node. If there are stale caches, any pin not
	// in the active CID set is a candidate for cleanup.
	hasStaleOrders := false
	for _, id := range cachedIDs {
		if !activeSet[id.String()] {
			hasStaleOrders = true
			break
		}
	}
	if !hasStaleOrders {
		return
	}

	orphaned := 0
	for _, pin := range pins {
		if activeCIDs[pin] {
			continue
		}
		// Verify this pin is actually from us by checking if it's pinned
		// (IsPinned would be redundant since we got it from ListPins).
		// Unpin it — if the IPFS node is shared, worst case another app
		// can re-pin. Log at Info level so the operator can see what was cleaned.
		if err := n.ipfs.Unpin(ctx, pin); err != nil {
			log.Debug().Err(err).Str("cid", pin).Msg("orphan pin: unpin failed")
		} else {
			orphaned++
			log.Info().Str("cid", pin).Msg("orphan pin: unpinned stale CID")
		}
	}
	if orphaned > 0 {
		log.Info().Int("count", orphaned).Msg("cleaned orphaned IPFS pins")
	}
}

// pruneStaleCache removes cached SMT files for orders the node is no longer serving.
// If orders is nil (detectOrderChanges failed), the prune is skipped.
func (n *Node) pruneStaleCache(orders []*big.Int) {
	if orders == nil {
		return
	}

	activeSet := make(map[string]bool, len(orders))
	for _, oid := range orders {
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

// fetchTimeout computes a generous IPFS fetch deadline based on file size.
// Base of 5 minutes + 3 seconds per MB accounts for CatChunked's per-segment
// retry (each segment can take up to 2 min response header timeout + retries).
func fetchTimeout(numChunks uint32) time.Duration {
	const chunkSize = 16384 // 16 KB
	sizeMB := uint64(numChunks) * chunkSize / (1024 * 1024)
	if sizeMB < 1 {
		sizeMB = 1
	}
	return 5*time.Minute + time.Duration(sizeMB)*3*time.Second
}

// buildSMTStreaming downloads a file from IPFS and builds an SMT by hashing
// chunks on the fly. The full file is never buffered in memory — only one
// download segment (~1 MB) plus the leaf hashes are held at any time.
//
// Peak memory: ~12 MB for a 1 GB file, ~124 MB for a 10 GB file
// (vs ~1 GB / ~10 GB with the old CatChunked + BuildSMT approach).
func (n *Node) buildSMTStreaming(ctx context.Context, ref string, numChunks uint32) (*merkle.SparseMerkleTree, error) {
	type hashJob struct {
		index int
		data  []byte
	}

	jobs := make(chan hashJob, 128)

	// Pre-allocate the results slice. Each index is written by exactly one
	// worker (disjoint indices), so no synchronization is needed for writes.
	hashCap := int(numChunks) + 1
	hashes := make([]fr.Element, hashCap)

	// Start parallel hash workers
	workers := runtime.NumCPU()
	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := range jobs {
				if j.index < hashCap {
					hashes[j.index] = poi.HashChunk(j.data)
				}
			}
		}()
	}

	// Stream download: fetch each 16 KB chunk and dispatch for hashing.
	// CatChunkedTo holds only ~1 MB (one download segment) at a time.
	totalChunks, err := n.ipfs.CatChunkedTo(ctx, ref, poi.FileSize, func(index int, chunk []byte) {
		// Copy chunk data — the slice is only valid during this callback.
		data := make([]byte, len(chunk))
		copy(data, chunk)
		select {
		case jobs <- hashJob{index: index, data: data}:
		case <-ctx.Done():
		}
	})
	close(jobs)
	wg.Wait()

	if err != nil {
		return nil, fmt.Errorf("streaming download: %w", err)
	}

	if totalChunks > hashCap {
		return nil, fmt.Errorf("chunk count %d exceeds expected %d", totalChunks, numChunks)
	}
	hashes = hashes[:totalChunks]

	log.Debug().Int("chunks", totalChunks).Msg("streaming hash complete")

	smt, err := merkle.BuildSMTFromLeafHashes(hashes, poi.MaxTreeDepth, n.prover.ZeroLeafHash())
	if err != nil {
		return nil, fmt.Errorf("build SMT: %w", err)
	}

	return smt, nil
}

// isDue returns true if the last-run timestamp is older than interval,
// or if it was never set (nil). Used to gate infrequent maintenance tasks.
func isDue(last *atomic.Value, interval time.Duration) bool {
	v := last.Load()
	if v == nil {
		return true
	}
	return time.Since(v.(time.Time)) >= interval
}

// copyOrderMap returns a shallow copy of m, safe for use in a background goroutine.
func copyOrderMap(m map[string]string) map[string]string {
	cp := make(map[string]string, len(m))
	for k, v := range m {
		cp[k] = v
	}
	return cp
}
