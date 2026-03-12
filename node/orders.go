package node

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/MuriData/muri-node/prover"
	"github.com/MuriData/muri-node/types"
	"github.com/MuriData/muri-zkproof/pkg/merkle"
	"github.com/rs/zerolog/log"
)

// fetchCooldown is how long a CID is skipped after a fetch failure in the order loop.
// Prevents repeatedly blocking the loop on the same unreachable file.
const fetchCooldown = 10 * time.Minute

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

// orderLoop dispatches to event-based or poll-based order listening.
func (n *Node) orderLoop(ctx context.Context) error {
	if n.cfg.Chain.ListenMode == "events" && n.chain.HasEventSupport() {
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
	orderCh, err := n.chain.SubscribeNewOrders(ctx)
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
	// and will be picked up on the next poll tick.
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
		case n.orderSem <- struct{}{}:
			dispatched++
			go func(c orderCandidate) {
				defer func() { <-n.orderSem }()
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
	// Fetch file and build SMT.
	fetchCtx, fetchCancel := context.WithTimeout(ctx, fetchTimeout(cand.order.NumChunks))
	var smt *merkle.SparseMerkleTree
	var err error
	if isRawBlockURI(cand.order.URI) {
		smt, err = n.buildSMTFromRawBlock(fetchCtx, cand.ref)
	} else {
		smt, err = n.buildSMTStreaming(fetchCtx, cand.ref, cand.order.NumChunks)
	}
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
	proofResult, err := n.proveSelective(ctx, randomness, smt, cand.ref, isRawBlockURI(cand.order.URI))
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
