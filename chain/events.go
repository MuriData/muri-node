package chain

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/MuriData/muri-node/chain/bindings"
	"github.com/MuriData/muri-node/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog/log"
)

const (
	maxReconnectAttempts = 5
	reconnectBackoff     = 5 * time.Second
)

// EventListener manages WebSocket event subscriptions with automatic reconnection.
type EventListener struct {
	client *Client
	myAddr common.Address
}

// NewEventListener creates a new EventListener.
func NewEventListener(client *Client) *EventListener {
	return &EventListener{
		client: client,
		myAddr: client.Address(),
	}
}

// retrySubscription runs connectAndProcess in a retry-with-backoff loop.
// connectAndProcess should establish a subscription, process events, and return
// nil on clean exit (context cancelled) or an error to trigger reconnection.
// The loop gives up after maxReconnectAttempts consecutive quick failures.
// If a connection was alive longer than reconnectBackoff before failing,
// the attempt counter resets (transient disconnection, not a persistent issue).
func retrySubscription(ctx context.Context, name string, connectAndProcess func(ctx context.Context) error) {
	for attempt := 0; ; attempt++ {
		if ctx.Err() != nil {
			return
		}

		if attempt > 0 {
			if attempt > maxReconnectAttempts {
				log.Error().Int("attempts", attempt).Msgf("%s: max reconnect attempts reached, falling back to polling", name)
				return
			}
			log.Warn().Int("attempt", attempt).Msgf("%s: reconnecting", name)
			select {
			case <-ctx.Done():
				return
			case <-time.After(reconnectBackoff):
			}
		}

		start := time.Now()
		err := connectAndProcess(ctx)
		if err == nil {
			return // clean exit (context cancelled)
		}
		log.Warn().Err(err).Msgf("%s: error, will reconnect", name)

		// Reset attempts if the connection was alive for a while — the
		// subscribe succeeded and then failed later (transient disconnect).
		if time.Since(start) > reconnectBackoff {
			attempt = 0
		}
	}
}

// SubscribeChallenges returns a channel that emits fully-populated ChallengeSlotInfo
// for challenges targeting this node. Events are enriched via GetSlotInfo to include
// randomness (which the event itself does not contain).
func (el *EventListener) SubscribeChallenges(ctx context.Context) (<-chan types.ChallengeSlotInfo, error) {
	ch := make(chan types.ChallengeSlotInfo, 8)

	go func() {
		defer close(ch)
		retrySubscription(ctx, "challenge subscription", func(ctx context.Context) error {
			sink := make(chan *bindings.FileMarketSlotChallengeIssued, 8)
			sub, err := el.client.Filterer.WatchSlotChallengeIssued(&bind.WatchOpts{Context: ctx}, sink, nil)
			if err != nil {
				return fmt.Errorf("watch: %w", err)
			}
			defer sub.Unsubscribe()
			return el.processChallengeEvents(ctx, sink, sub, ch)
		})
	}()

	return ch, nil
}

// SubscribeNewOrders returns a channel that emits order IDs when new orders are placed.
func (el *EventListener) SubscribeNewOrders(ctx context.Context) (<-chan *big.Int, error) {
	ch := make(chan *big.Int, 16)

	go func() {
		defer close(ch)
		retrySubscription(ctx, "order subscription", func(ctx context.Context) error {
			sink := make(chan *bindings.FileMarketOrderPlaced, 16)
			sub, err := el.client.Filterer.WatchOrderPlaced(
				&bind.WatchOpts{Context: ctx},
				sink,
				nil, // no orderId filter
				nil, // no owner filter
			)
			if err != nil {
				return fmt.Errorf("watch: %w", err)
			}
			defer sub.Unsubscribe()
			return el.processOrderEvents(ctx, sink, sub, ch)
		})
	}()

	return ch, nil
}

// enrichRetries is how many times we retry GetSlotInfo when the RPC returns
// stale data that doesn't match the event (HTTP endpoint lagging behind WS).
const enrichRetries = 3

func (el *EventListener) processChallengeEvents(
	ctx context.Context,
	sink <-chan *bindings.FileMarketSlotChallengeIssued,
	sub interface{ Err() <-chan error },
	out chan<- types.ChallengeSlotInfo,
) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		case err := <-sub.Err():
			return err
		case ev := <-sink:
			if ev == nil {
				continue
			}
			// Filter: only process challenges targeting this node
			if ev.ChallengedNode != el.myAddr {
				continue
			}

			log.Info().
				Int64("slot", ev.SlotIndex.Int64()).
				Str("orderID", ev.OrderId.String()).
				Uint64("deadline", ev.DeadlineBlock.Uint64()).
				Msg("challenge event received")

			// Enrich with full slot data (includes randomness).
			// The HTTP RPC may lag behind the WS event, so retry briefly
			// if the enriched data doesn't match the event.
			slot, err := el.enrichSlotFromEvent(ctx, ev)
			if err != nil {
				log.Error().Err(err).Int64("slot", ev.SlotIndex.Int64()).
					Msg("failed to enrich slot from event, skipping (will be caught by fallback poll)")
				continue
			}

			select {
			case out <- slot:
			case <-ctx.Done():
				return nil
			}
		}
	}
}

// enrichSlotFromEvent fetches full slot data via GetSlotInfo and validates
// it matches the event fields. Retries with a short delay if the HTTP RPC
// returns stale data (e.g., hasn't indexed the block that emitted the event yet).
func (el *EventListener) enrichSlotFromEvent(ctx context.Context, ev *bindings.FileMarketSlotChallengeIssued) (types.ChallengeSlotInfo, error) {
	slotIndex := int(ev.SlotIndex.Int64())

	for attempt := 0; attempt < enrichRetries; attempt++ {
		if attempt > 0 {
			select {
			case <-ctx.Done():
				return types.ChallengeSlotInfo{}, ctx.Err()
			case <-time.After(2 * time.Second):
			}
		}

		slot, err := el.client.GetSlotInfo(ctx, slotIndex)
		if err != nil {
			log.Warn().Err(err).Int("slot", slotIndex).Int("attempt", attempt+1).
				Msg("GetSlotInfo failed during enrichment")
			continue
		}

		// Validate enriched data matches the event — detects stale RPC responses.
		if slot.ChallengedNode != ev.ChallengedNode {
			log.Warn().Int("slot", slotIndex).Int("attempt", attempt+1).
				Str("event_node", ev.ChallengedNode.Hex()).
				Str("rpc_node", slot.ChallengedNode.Hex()).
				Msg("enrichment mismatch: RPC returned stale slot data, retrying")
			continue
		}
		if slot.OrderID == nil || slot.OrderID.Cmp(ev.OrderId) != 0 {
			log.Warn().Int("slot", slotIndex).Int("attempt", attempt+1).
				Msg("enrichment mismatch: order ID differs, retrying")
			continue
		}

		return slot, nil
	}

	return types.ChallengeSlotInfo{}, fmt.Errorf("slot %d enrichment failed after %d attempts (stale RPC)", slotIndex, enrichRetries)
}

func (el *EventListener) processOrderEvents(
	ctx context.Context,
	sink <-chan *bindings.FileMarketOrderPlaced,
	sub interface{ Err() <-chan error },
	out chan<- *big.Int,
) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		case err := <-sub.Err():
			return err
		case ev := <-sink:
			if ev == nil {
				continue
			}

			log.Info().
				Str("orderID", ev.OrderId.String()).
				Uint32("chunks", ev.NumChunks).
				Uint8("replicas", ev.Replicas).
				Msg("new order event received")

			select {
			case out <- ev.OrderId:
			case <-ctx.Done():
				return nil
			}
		}
	}
}
