package node

import (
	"context"
	"fmt"
	"math/big"
	"sort"
	"sync"
	"time"

	"github.com/MuriData/muri-node/prover"
	"github.com/MuriData/muri-node/types"
	"github.com/rs/zerolog/log"
)

// challengeWorkers controls how many challenges are handled concurrently.
// 2 provides a pipeline benefit (IPFS fetch for slot B overlaps with proof
// generation for slot A) without overwhelming CPU, since the prover itself
// is mutex-serialized.
const challengeWorkers = 2

// challengeLoop dispatches to event-based or poll-based challenge listening.
func (n *Node) challengeLoop(ctx context.Context) error {
	if n.cfg.Chain.ListenMode == "events" && n.chain.HasEventSupport() {
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
	challengeCh, err := n.chain.SubscribeChallenges(ctx)
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
			if n.isRecentlyProved(slot.Index, slot.Randomness) {
				log.Debug().Int("slot", slot.Index).Msg("event for already-proved challenge, skipping")
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

// isRecentlyProved checks if the given (slot, randomness) pair was recently
// proved, which means re-dispatching it would be redundant. Expired entries
// are cleaned up on access.
func (n *Node) isRecentlyProved(slotIndex int, randomness *big.Int) bool {
	proofKey := fmt.Sprintf("%d:%s", slotIndex, randomness.Text(16))
	ts, ok := n.recentProofs.Load(proofKey)
	if !ok {
		return false
	}
	if time.Since(ts.(time.Time)) < 10*time.Minute {
		return true
	}
	n.recentProofs.Delete(proofKey)
	return false
}

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
		if n.isRecentlyProved(slot.Index, slot.Randomness) {
			log.Debug().Int("slot", slot.Index).Msg("skipping already-proved challenge")
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
	rawBlock := isRawBlockURI(order.URI)

	// 2. Try fast path: cached SMT + selective chunk fetch
	var result *prover.ProofResult
	var path string
	smt, err := n.store.LoadTree(orderID, n.prover.ZeroLeafHash())
	if err == nil && smt != nil && order.RootHash != nil && smt.RootBigInt().Cmp(order.RootHash) == 0 {
		result, err = n.proveSelective(ctx, randomness, smt, ref, rawBlock)
		if err != nil {
			log.Warn().Err(err).Str("orderID", orderID.String()).Msg("selective proof failed, falling back to full download")
			result = nil
		} else {
			path = "selective"
		}
	}

	// 3. Slow path: stream download → hash → build SMT → prove selectively.
	if result == nil {
		fetchCtx, fetchCancel := context.WithTimeout(ctx, fetchTimeout(order.NumChunks))
		if isRawBlockURI(order.URI) {
			smt, err = n.buildSMTFromRawBlock(fetchCtx, ref)
		} else {
			smt, err = n.buildSMTStreaming(fetchCtx, ref, order.NumChunks)
		}
		fetchCancel()
		if err != nil {
			return fmt.Errorf("build smt %s: %w", ref, err)
		}

		// Cache for future challenge responses
		if saveErr := n.store.SaveTree(orderID, smt); saveErr != nil {
			log.Warn().Err(saveErr).Str("orderID", orderID.String()).Msg("cache tree failed (non-fatal)")
		}

		result, err = n.proveSelective(ctx, randomness, smt, ref, rawBlock)
		if err != nil {
			return fmt.Errorf("generate proof (streaming+selective): %w", err)
		}
		path = "streaming+selective"
	}

	// 4. Verify slot hasn't been re-advanced while we were proving (~20-40s).
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
