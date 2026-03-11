package node

import (
	"context"
	"sync"
	"time"

	"github.com/rs/zerolog/log"
)

// repinFailureEscalationAge is how long a re-pin failure must persist
// before it is escalated from WARN to ERROR level.
const repinFailureEscalationAge = 1 * time.Hour

// verifyPins checks that all active order CIDs are still pinned locally.
// For any missing pin, it attempts to re-pin (Kubo re-fetches from the network).
// Runs synchronously — IsPinned is fast (~10ms per CID).
func (n *Node) verifyPins(ctx context.Context, activeCIDs map[string]string) {
	checked, missing, repinned, failed := 0, 0, 0, 0

	for orderID, cid := range activeCIDs {
		if cid == "" {
			continue
		}
		if ctx.Err() != nil {
			break
		}
		// Don't touch CIDs with in-flight challenge responses.
		if n.isOrderInFlightChallenge(orderID) {
			continue
		}

		pinned, err := n.ipfs.IsPinned(ctx, cid)
		if err != nil {
			log.Warn().Err(err).Str("cid", cid).Str("orderID", orderID).
				Msg("pin verify: IsPinned check failed (skipping)")
			continue
		}
		checked++

		if pinned {
			n.repinFailures.Delete(cid)
			continue
		}

		missing++
		log.Warn().Str("cid", cid).Str("orderID", orderID).
			Msg("pin verify: CID is not pinned, attempting re-pin")

		pinCtx, cancel := context.WithTimeout(ctx, 5*time.Minute)
		err = n.ipfs.Pin(pinCtx, cid)
		cancel()

		if err != nil {
			failed++
			firstSeen := time.Now()
			if v, loaded := n.repinFailures.LoadOrStore(cid, firstSeen); loaded {
				firstSeen = v.(time.Time)
			}

			age := time.Since(firstSeen)
			evt := log.Warn()
			if age >= repinFailureEscalationAge {
				evt = log.Error()
			}
			evt.Err(err).
				Str("cid", cid).
				Str("orderID", orderID).
				Dur("failureDuration", age).
				Msg("pin verify: re-pin failed — file may be unavailable on network")
		} else {
			repinned++
			n.repinFailures.Delete(cid)
			log.Info().Str("cid", cid).Str("orderID", orderID).
				Msg("pin verify: successfully re-pinned missing CID")
		}
	}

	if missing > 0 || failed > 0 {
		log.Warn().
			Int("checked", checked).
			Int("missing", missing).
			Int("repinned", repinned).
			Int("failed", failed).
			Msg("pin verify complete")
	} else if checked > 0 {
		log.Debug().Int("checked", checked).Msg("pin verify: all CIDs healthy")
	}
}

// provideWorkers is the concurrency limit for parallel DHT provide calls.
const provideWorkers = 3

// provideAll re-announces all active order CIDs to the IPFS DHT.
// Supplements Kubo's built-in reprovider to keep DHT records fresh for data
// the node is contractually obligated to serve.
//
// Runs in a background goroutine — DHT provide is slow (multi-second
// network round-trip per CID). activeCIDs must be a snapshot copy.
func (n *Node) provideAll(ctx context.Context, activeCIDs map[string]string) {
	// Deduplicate: multiple orders may reference the same root CID.
	unique := make(map[string]struct{}, len(activeCIDs))
	for _, cid := range activeCIDs {
		if cid != "" {
			unique[cid] = struct{}{}
		}
	}
	if len(unique) == 0 {
		return
	}

	log.Info().Int("cids", len(unique)).Msg("dht provide: starting re-announce")

	cidCh := make(chan string, len(unique))
	for cid := range unique {
		cidCh <- cid
	}
	close(cidCh)

	var (
		wg        sync.WaitGroup
		mu        sync.Mutex
		succeeded int
		failed    int
	)

	workers := provideWorkers
	if len(unique) < workers {
		workers = len(unique)
	}
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			for cid := range cidCh {
				if ctx.Err() != nil {
					return
				}
				if err := n.ipfs.Provide(ctx, cid); err != nil {
					log.Warn().Err(err).Str("cid", cid).Msg("dht provide: re-announce failed")
					mu.Lock()
					failed++
					mu.Unlock()
				} else {
					mu.Lock()
					succeeded++
					mu.Unlock()
				}
			}
		}()
	}
	wg.Wait()

	log.Info().Int("succeeded", succeeded).Int("failed", failed).Msg("dht provide: re-announce complete")
}
