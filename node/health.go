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
// snapshot to detect cancellations and removals. Tree cache cleanup is
// handled by the existing pruneStaleCache in the same maintenance cycle.
func (n *Node) detectOrderChanges(ctx context.Context) {
	orders, err := n.chain.GetNodeOrders(ctx)
	if err != nil {
		log.Warn().Err(err).Msg("health: failed to get node orders")
		return
	}

	currentSet := make(map[string]bool, len(orders))
	for _, oid := range orders {
		currentSet[oid.String()] = true
	}

	// First run: initialize snapshot without alerting
	if len(n.prevOrders) == 0 && len(orders) > 0 {
		n.prevOrders = currentSet
		return
	}

	// Detect removed orders (were in prev, not in current)
	for oid := range n.prevOrders {
		if currentSet[oid] {
			continue
		}

		log.Warn().Str("orderID", oid).Msg("health: order removed (cancelled, completed, or force-exited)")

		// Try to unpin from IPFS if we can still read the order details
		orderID, ok := new(big.Int).SetString(oid, 10)
		if !ok {
			continue
		}
		order, err := n.chain.GetOrderDetails(ctx, orderID)
		if err != nil {
			log.Debug().Err(err).Str("orderID", oid).Msg("health: could not fetch removed order details (already purged)")
			continue
		}
		cid := extractCID(order.URI)
		if cid != "" {
			if err := n.ipfs.Unpin(ctx, cid); err != nil {
				log.Debug().Err(err).Str("cid", cid).Msg("health: unpin failed (may already be unpinned)")
			} else {
				log.Info().Str("cid", cid).Str("orderID", oid).Msg("health: unpinned file for removed order")
			}
		}
	}

	// Detect new orders we didn't execute ourselves
	for oid := range currentSet {
		if n.prevOrders[oid] {
			continue
		}
		log.Info().Str("orderID", oid).Msg("health: new order detected in active set")
	}

	n.prevOrders = currentSet
}
