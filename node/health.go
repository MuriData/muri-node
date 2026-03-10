package node

import (
	"context"
	"math/big"

	"github.com/rs/zerolog/log"
)

// checkNodeHealth verifies the node is still registered on-chain.
// If the node has been unstaked, force-exited, or fully slashed, it sets
// the deregistered flag and auto-pauses order intake.
func (n *Node) checkNodeHealth(ctx context.Context) {
	if n.deregistered.Load() {
		return
	}

	isValid, err := n.chain.IsValidNode(ctx)
	if err != nil {
		log.Warn().Err(err).Msg("health: failed to check node validity")
		return
	}

	if !isValid {
		log.Error().Msg("health: node is NO LONGER REGISTERED on-chain (unstaked, force-exited, or fully slashed)")
		n.deregistered.Store(true)
		n.paused.Store(true)
		return
	}

	// Check for slashing: compare current stake against expected
	info, err := n.chain.GetNodeInfo(ctx)
	if err != nil {
		log.Warn().Err(err).Msg("health: failed to get node info")
		return
	}

	if info.Stake != nil && info.Stake.Sign() == 0 {
		log.Error().Msg("health: node stake is zero — deregistering")
		n.deregistered.Store(true)
		n.paused.Store(true)
		return
	}

	// Warn about capacity issues after slashing
	if info.Used > info.Capacity {
		log.Warn().
			Uint64("capacity", info.Capacity).
			Uint64("used", info.Used).
			Msg("health: used exceeds capacity (likely slashed) — auto-pausing orders")
		n.paused.Store(true)
	}
}

// detectOrderChanges compares current on-chain orders against the previous
// snapshot to detect cancellations and removals. Returns the current order
// IDs so callers (e.g. pruneStaleCache) can reuse them without a duplicate RPC.
//
// prevOrders maps orderID → rootCID so we can unpin even after the order
// has been purged from chain state.
func (n *Node) detectOrderChanges(ctx context.Context) []*big.Int {
	orders, err := n.chain.GetNodeOrders(ctx)
	if err != nil {
		log.Warn().Err(err).Msg("health: failed to get node orders")
		return nil
	}

	// Build current snapshot: orderID → rootCID.
	// Fetch details for any order we haven't seen before so we capture its CID
	// while the chain still has the data.
	currentSet := make(map[string]string, len(orders))
	for _, oid := range orders {
		key := oid.String()
		if cid, ok := n.prevOrders[key]; ok {
			// Carry forward the CID we already know. If it was empty from
			// a previous failed fetch, retry now while the order is still active.
			if cid == "" {
				if order, err := n.chain.GetOrderDetails(ctx, oid); err == nil {
					cid = extractRootCID(order.URI)
				}
			}
			currentSet[key] = cid
		} else {
			// New order — fetch its CID now while it's still on-chain.
			order, err := n.chain.GetOrderDetails(ctx, oid)
			if err != nil {
				log.Warn().Err(err).Str("orderID", key).Msg("health: failed to fetch new order details")
				currentSet[key] = "" // track the order even without CID
			} else {
				currentSet[key] = extractRootCID(order.URI)
			}
			log.Info().Str("orderID", key).Msg("health: new order detected in active set")
		}
	}

	// First run: initialize snapshot without alerting about removals.
	if len(n.prevOrders) == 0 && len(orders) > 0 {
		n.prevOrders = currentSet
		if err := n.store.SaveOrderMapAtomic(currentSet); err != nil {
			log.Warn().Err(err).Msg("health: failed to persist initial order map")
		}
		return orders
	}

	// Detect removed orders (were in prev, not in current) and unpin using
	// the CID we stored when the order was first seen.
	for oid, rootCID := range n.prevOrders {
		if _, ok := currentSet[oid]; ok {
			continue
		}

		// Skip cleanup if a challenge goroutine is actively working on this order.
		// The deferred cleanup list will handle it after the challenge completes.
		if n.isOrderInFlightChallenge(oid) {
			log.Info().Str("orderID", oid).Msg("health: order removed but challenge in-flight, deferring cleanup")
			n.deferCleanup(oid, rootCID)
			continue
		}

		n.cleanupRemovedOrder(ctx, oid, rootCID)
	}

	n.prevOrders = currentSet

	// Persist to disk so restarts don't lose CID mappings.
	if err := n.store.SaveOrderMapAtomic(currentSet); err != nil {
		log.Warn().Err(err).Msg("health: failed to persist order map")
	}

	// Process any deferred cleanups whose challenges have completed.
	n.processDeferredCleanups(ctx)

	return orders
}

// cleanupRemovedOrder handles IPFS unpin and logging for a removed order.
func (n *Node) cleanupRemovedOrder(ctx context.Context, oid, rootCID string) {
	log.Warn().Str("orderID", oid).Msg("health: order removed (cancelled, completed, or force-exited)")

	if rootCID == "" || !n.cfg.IPFS.PinFiles {
		return
	}
	if err := n.ipfs.Unpin(ctx, rootCID); err != nil {
		log.Debug().Err(err).Str("cid", rootCID).Msg("health: unpin failed (may already be unpinned)")
	} else {
		log.Info().Str("cid", rootCID).Str("orderID", oid).Msg("health: unpinned file for removed order")
	}
}

// deferredCleanup holds info for orders removed while a challenge was in-flight.
type deferredCleanup struct {
	orderID string
	rootCID string
}

// deferCleanup adds an order to the deferred cleanup list.
func (n *Node) deferCleanup(orderID, rootCID string) {
	n.deferredCleanupsMu.Lock()
	defer n.deferredCleanupsMu.Unlock()
	// Avoid duplicates
	for _, dc := range n.deferredCleanups {
		if dc.orderID == orderID {
			return
		}
	}
	n.deferredCleanups = append(n.deferredCleanups, deferredCleanup{orderID, rootCID})
}

// processDeferredCleanups runs cleanup for orders whose challenges have finished.
func (n *Node) processDeferredCleanups(ctx context.Context) {
	n.deferredCleanupsMu.Lock()
	var remaining []deferredCleanup
	var ready []deferredCleanup
	for _, dc := range n.deferredCleanups {
		if n.isOrderInFlightChallenge(dc.orderID) {
			remaining = append(remaining, dc)
		} else {
			ready = append(ready, dc)
		}
	}
	n.deferredCleanups = remaining
	n.deferredCleanupsMu.Unlock()

	for _, dc := range ready {
		n.cleanupRemovedOrder(ctx, dc.orderID, dc.rootCID)
	}
}

// isOrderInFlightChallenge checks if any challenge goroutine is working on the given order.
func (n *Node) isOrderInFlightChallenge(orderID string) bool {
	found := false
	n.inFlightChallengeOrders.Range(func(key, _ any) bool {
		if key.(string) == orderID {
			found = true
			return false
		}
		return true
	})
	return found
}
