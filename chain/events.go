package chain

import (
	"context"
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

// SubscribeChallenges returns a channel that emits fully-populated ChallengeSlotInfo
// for challenges targeting this node. Events are enriched via GetSlotInfo to include
// randomness (which the event itself does not contain).
func (el *EventListener) SubscribeChallenges(ctx context.Context) (<-chan types.ChallengeSlotInfo, error) {
	ch := make(chan types.ChallengeSlotInfo, 8)

	go el.challengeSubscriptionLoop(ctx, ch)

	return ch, nil
}

func (el *EventListener) challengeSubscriptionLoop(ctx context.Context, out chan<- types.ChallengeSlotInfo) {
	defer close(out)

	for attempt := 0; ; attempt++ {
		if ctx.Err() != nil {
			return
		}

		if attempt > 0 {
			if attempt > maxReconnectAttempts {
				log.Error().Int("attempts", attempt).Msg("challenge subscription: max reconnect attempts reached, falling back to polling")
				return
			}
			log.Warn().Int("attempt", attempt).Msg("challenge subscription: reconnecting")
			select {
			case <-ctx.Done():
				return
			case <-time.After(reconnectBackoff):
			}
		}

		sink := make(chan *bindings.FileMarketSlotChallengeIssued, 8)
		sub, err := el.client.Filterer.WatchSlotChallengeIssued(&bind.WatchOpts{Context: ctx}, sink, nil)
		if err != nil {
			log.Warn().Err(err).Msg("challenge subscription: watch failed")
			continue
		}

		// Reset attempt counter on successful subscribe
		attempt = 0

		if err := el.processChallengeEvents(ctx, sink, sub, out); err != nil {
			log.Warn().Err(err).Msg("challenge subscription: error, will reconnect")
			sub.Unsubscribe()
			continue
		}

		// Context cancelled — clean exit
		sub.Unsubscribe()
		return
	}
}

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

			// Enrich with full slot data (includes randomness)
			slot, err := el.client.GetSlotInfo(ctx, int(ev.SlotIndex.Int64()))
			if err != nil {
				log.Error().Err(err).Int64("slot", ev.SlotIndex.Int64()).Msg("failed to fetch slot info for event, skipping (will be caught by fallback poll)")
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

// SubscribeNewOrders returns a channel that emits order IDs when new orders are placed.
func (el *EventListener) SubscribeNewOrders(ctx context.Context) (<-chan *big.Int, error) {
	ch := make(chan *big.Int, 16)

	go el.orderSubscriptionLoop(ctx, ch)

	return ch, nil
}

func (el *EventListener) orderSubscriptionLoop(ctx context.Context, out chan<- *big.Int) {
	defer close(out)

	for attempt := 0; ; attempt++ {
		if ctx.Err() != nil {
			return
		}

		if attempt > 0 {
			if attempt > maxReconnectAttempts {
				log.Error().Int("attempts", attempt).Msg("order subscription: max reconnect attempts reached, falling back to polling")
				return
			}
			log.Warn().Int("attempt", attempt).Msg("order subscription: reconnecting")
			select {
			case <-ctx.Done():
				return
			case <-time.After(reconnectBackoff):
			}
		}

		sink := make(chan *bindings.FileMarketOrderPlaced, 16)
		sub, err := el.client.Filterer.WatchOrderPlaced(
			&bind.WatchOpts{Context: ctx},
			sink,
			nil, // no orderId filter
			nil, // no owner filter
		)
		if err != nil {
			log.Warn().Err(err).Msg("order subscription: watch failed")
			continue
		}

		// Reset attempt counter on successful subscribe
		attempt = 0

		if err := el.processOrderEvents(ctx, sink, sub, out); err != nil {
			log.Warn().Err(err).Msg("order subscription: error, will reconnect")
			sub.Unsubscribe()
			continue
		}

		sub.Unsubscribe()
		return
	}
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
