package node

import (
	"context"
	"fmt"
	"math/big"
	"sync"
	"sync/atomic"
	"time"

	"github.com/MuriData/muri-node/chain"
	"github.com/MuriData/muri-node/config"
	"github.com/MuriData/muri-node/ipfs"
	"github.com/MuriData/muri-node/prover"
	"github.com/MuriData/muri-node/storage"
	"github.com/rs/zerolog/log"
)

// Node is the MuriData storage provider daemon.
type Node struct {
	cfg       *config.Config
	chain     ChainService
	ipfs      IPFSService
	prover    ProverService
	store     StorageService
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

	// orderSem is a semaphore shared across poll ticks so the order loop is not
	// blocked by slow fetches — new ticks can dispatch fresh candidates into free
	// worker slots while previous ones are still running.
	orderSem chan struct{}
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
		orderSem:   make(chan struct{}, orderWorkers),
	}

	// Defer first pin verify / provide until a full interval after startup.
	n.lastPinVerify.Store(time.Now())
	n.lastProvide.Store(time.Now())

	return n, nil
}

// rebuildOrderMap fetches current orders from chain and builds the
// orderID → rootCID mapping. Used on first run or when persisted state is lost.
func rebuildOrderMap(ctx context.Context, c ChainService) map[string]string {
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
