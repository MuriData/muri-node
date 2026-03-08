package node

import (
	"context"

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
// snapshot to detect cancellations and removals. Tree cache cleanup is
// handled by the existing pruneStaleCache in the same maintenance cycle.
//
// prevOrders maps orderID → rootCID so we can unpin even after the order
// has been purged from chain state.
func (n *Node) detectOrderChanges(ctx context.Context) {
	orders, err := n.chain.GetNodeOrders(ctx)
	if err != nil {
		log.Warn().Err(err).Msg("health: failed to get node orders")
		return
	}

	// Build current snapshot: orderID → rootCID.
	// Fetch details for any order we haven't seen before so we capture its CID
	// while the chain still has the data.
	currentSet := make(map[string]string, len(orders))
	for _, oid := range orders {
		key := oid.String()
		if cid, ok := n.prevOrders[key]; ok {
			// Carry forward the CID we already know.
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
		if err := n.store.SaveOrderMap(currentSet); err != nil {
			log.Warn().Err(err).Msg("health: failed to persist initial order map")
		}
		return
	}

	// Detect removed orders (were in prev, not in current) and unpin using
	// the CID we stored when the order was first seen.
	for oid, rootCID := range n.prevOrders {
		if _, ok := currentSet[oid]; ok {
			continue
		}

		log.Warn().Str("orderID", oid).Msg("health: order removed (cancelled, completed, or force-exited)")

		if rootCID == "" {
			continue
		}
		if err := n.ipfs.Unpin(ctx, rootCID); err != nil {
			log.Debug().Err(err).Str("cid", rootCID).Msg("health: unpin failed (may already be unpinned)")
		} else {
			log.Info().Str("cid", rootCID).Str("orderID", oid).Msg("health: unpinned file for removed order")
		}
	}

	n.prevOrders = currentSet

	// Persist to disk so restarts don't lose CID mappings.
	if err := n.store.SaveOrderMap(currentSet); err != nil {
		log.Warn().Err(err).Msg("health: failed to persist order map")
	}
}
