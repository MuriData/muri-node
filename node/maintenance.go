package node

import (
	"context"
	"math/big"
	"sync/atomic"
	"time"

	"github.com/rs/zerolog/log"
)

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
	cachedIDs, err := n.store.ListCachedOrderIDs()
	if err != nil {
		log.Warn().Err(err).Msg("orphan pin check: failed to list cached order IDs")
		return
	}

	activeSet := make(map[string]bool, len(n.prevOrders))
	for oid := range n.prevOrders {
		activeSet[oid] = true
	}

	// Check if any pinned CIDs are NOT in our active set.
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
	// they belonged to this node.
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
