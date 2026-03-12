package chain

import (
	"context"
	"fmt"
	"math/big"
	"sync"

	"github.com/MuriData/muri-node/chain/bindings"
	"github.com/MuriData/muri-node/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

// GetSlotInfo returns a single challenge slot's full state (including randomness).
func (c *Client) GetSlotInfo(ctx context.Context, slotIndex int) (types.ChallengeSlotInfo, error) {
	result, err := c.Market.GetSlotInfo(c.callOpts(ctx), big.NewInt(int64(slotIndex)))
	if err != nil {
		return types.ChallengeSlotInfo{}, fmt.Errorf("getSlotInfo(%d): %w", slotIndex, err)
	}
	return types.ChallengeSlotInfo{
		Index:          slotIndex,
		OrderID:        result.OrderId,
		ChallengedNode: result.ChallengedNode,
		Randomness:     result.Randomness,
		DeadlineBlock:  result.DeadlineBlock,
		IsExpired:      result.IsExpired,
	}, nil
}

// GetSlotInfoFresh queries slot info via the WebSocket client when available,
// bypassing HTTP reverse-proxy caching. Falls back to the HTTP client if WS
// is not configured. Used for the pre-submit freshness check where stale data
// would cause the node to submit to the wrong slot.
func (c *Client) GetSlotInfoFresh(ctx context.Context, slotIndex int) (types.ChallengeSlotInfo, error) {
	caller := c.wsMarketCaller
	if caller == nil {
		// No WS client — fall back to HTTP
		return c.GetSlotInfo(ctx, slotIndex)
	}
	result, err := caller.GetSlotInfo(c.callOpts(ctx), big.NewInt(int64(slotIndex)))
	if err != nil {
		// WS read failed — fall back to HTTP rather than aborting
		return c.GetSlotInfo(ctx, slotIndex)
	}
	return types.ChallengeSlotInfo{
		Index:          slotIndex,
		OrderID:        result.OrderId,
		ChallengedNode: result.ChallengedNode,
		Randomness:     result.Randomness,
		DeadlineBlock:  result.DeadlineBlock,
		IsExpired:      result.IsExpired,
	}, nil
}

// GetAllSlotInfo returns all challenge slot states (dynamic count via sqrt scaling).
func (c *Client) GetAllSlotInfo(ctx context.Context) ([]types.ChallengeSlotInfo, error) {
	return c.getAllSlotInfoFrom(ctx, c.Market)
}

// GetAllSlotInfoFresh queries all slot info via the WebSocket client when available,
// bypassing HTTP caching. Falls back to HTTP if WS is not configured.
func (c *Client) GetAllSlotInfoFresh(ctx context.Context) ([]types.ChallengeSlotInfo, error) {
	if c.wsMarketCaller != nil {
		slots, err := c.getAllSlotInfoFromCaller(ctx, c.wsMarketCaller)
		if err == nil {
			return slots, nil
		}
		// WS read failed — fall back to HTTP
	}
	return c.GetAllSlotInfo(ctx)
}

func (c *Client) getAllSlotInfoFrom(ctx context.Context, market *bindings.FileMarket) ([]types.ChallengeSlotInfo, error) {
	result, err := market.GetAllSlotInfo(c.callOpts(ctx))
	if err != nil {
		return nil, fmt.Errorf("getAllSlotInfo: %w", err)
	}
	return parseAllSlotInfo(result), nil
}

func (c *Client) getAllSlotInfoFromCaller(ctx context.Context, caller *bindings.FileMarketCaller) ([]types.ChallengeSlotInfo, error) {
	result, err := caller.GetAllSlotInfo(c.callOpts(ctx))
	if err != nil {
		return nil, fmt.Errorf("getAllSlotInfo(ws): %w", err)
	}
	return parseAllSlotInfo(result), nil
}

func parseAllSlotInfo(result struct {
	OrderIds        []*big.Int
	ChallengedNodes []common.Address
	Randomnesses    []*big.Int
	DeadlineBlocks  []*big.Int
	IsExpired       []bool
}) []types.ChallengeSlotInfo {
	n := len(result.OrderIds)
	slots := make([]types.ChallengeSlotInfo, n)
	for i := 0; i < n; i++ {
		slots[i] = types.ChallengeSlotInfo{
			Index:          i,
			OrderID:        result.OrderIds[i],
			ChallengedNode: result.ChallengedNodes[i],
			Randomness:     result.Randomnesses[i],
			DeadlineBlock:  result.DeadlineBlocks[i],
			IsExpired:      result.IsExpired[i],
		}
	}
	return slots
}

// GetNodeOrders returns all order IDs assigned to a node.
func (c *Client) GetNodeOrders(ctx context.Context) ([]*big.Int, error) {
	return c.Market.GetNodeOrders(c.callOpts(ctx), c.addr)
}

// GetActiveOrdersPage returns a page of active order IDs starting at offset.
func (c *Client) GetActiveOrdersPage(ctx context.Context, offset, limit uint64) ([]*big.Int, uint64, error) {
	result, err := c.Market.GetActiveOrdersPage(c.callOpts(ctx), new(big.Int).SetUint64(offset), new(big.Int).SetUint64(limit))
	if err != nil {
		return nil, 0, fmt.Errorf("getActiveOrdersPage(%d,%d): %w", offset, limit, err)
	}
	return result.OrderIds, result.Total.Uint64(), nil
}

// GetOrderDetails returns order details for the given order ID.
// The three independent RPC calls run in parallel for lower latency.
func (c *Client) GetOrderDetails(ctx context.Context, orderID *big.Int) (*types.OrderInfo, error) {
	var (
		wg       sync.WaitGroup
		dErr     error
		fErr     error
		priceErr error

		info  types.OrderInfo
		price *big.Int
	)

	wg.Add(3)
	go func() {
		defer wg.Done()
		d, err := c.Market.GetOrderDetails(c.callOpts(ctx), orderID)
		if err != nil {
			dErr = fmt.Errorf("getOrderDetails(%s): %w", orderID, err)
			return
		}
		info.Owner = d.Owner
		info.URI = d.Uri
		info.RootHash = d.Root
		info.NumChunks = d.NumChunks
		info.Periods = d.Periods
		info.Replicas = d.Replicas
		info.Filled = d.Filled
	}()
	go func() {
		defer wg.Done()
		f, err := c.Market.GetOrderFinancials(c.callOpts(ctx), orderID)
		if err != nil {
			fErr = fmt.Errorf("getOrderFinancials(%s): %w", orderID, err)
			return
		}
		info.Escrow = f.Escrow
		info.StartPeriod = f.StartPeriod
	}()
	go func() {
		defer wg.Done()
		var err error
		price, err = c.Market.GetOrderPrice(c.callOpts(ctx), orderID)
		if err != nil {
			priceErr = fmt.Errorf("getOrderPrice(%s): %w", orderID, err)
		}
	}()
	wg.Wait()

	if dErr != nil {
		return nil, dErr
	}
	if fErr != nil {
		return nil, fErr
	}
	if priceErr != nil {
		return nil, priceErr
	}

	info.ID = orderID
	info.Price = price
	return &info, nil
}

// GetClaimableRewards returns the claimable reward amount for this node.
func (c *Client) GetClaimableRewards(ctx context.Context) (*big.Int, error) {
	return c.Market.GetClaimableRewards(c.callOpts(ctx), c.addr)
}

// GetGlobalStats returns marketplace-wide stats.
func (c *Client) GetGlobalStats(ctx context.Context) (struct {
	TotalOrders              *big.Int
	ActiveOrdersCount        *big.Int
	TotalEscrowLocked        *big.Int
	TotalNodes               *big.Int
	TotalCapacityStaked      *big.Int
	TotalCapacityUsed        *big.Int
	ActiveChallengeSlots     *big.Int
	CurrentPeriod            *big.Int
	CurrentBlock             *big.Int
	ChallengeableOrdersCount *big.Int
}, error) {
	return c.Market.GetGlobalStats(c.callOpts(ctx))
}

// ExecuteOrder claims a replica slot on the given order with a compressed PoI proof.
func (c *Client) ExecuteOrder(ctx context.Context, orderID *big.Int, proof [4]*big.Int, commitment [32]byte) (*ethtypes.Receipt, error) {
	return c.SendTx(ctx, func(opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
		return c.Market.ExecuteOrder(opts, orderID, proof, commitment)
	})
}

// SubmitProof submits a compressed ZK proof for a challenge slot.
func (c *Client) SubmitProof(ctx context.Context, slotIndex int, proof [4]*big.Int, commitment [32]byte) (*ethtypes.Receipt, error) {
	return c.SendTx(ctx, func(opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
		return c.Market.SubmitProof(opts, big.NewInt(int64(slotIndex)), proof, commitment)
	})
}

// ClaimRewards withdraws accumulated node rewards.
func (c *Client) ClaimRewards(ctx context.Context) (*ethtypes.Receipt, error) {
	return c.SendTx(ctx, func(opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
		return c.Market.ClaimRewards(opts)
	})
}

// ProcessExpiredSlots triggers slash processing on expired challenge slots.
func (c *Client) ProcessExpiredSlots(ctx context.Context) (*ethtypes.Receipt, error) {
	return c.SendTx(ctx, func(opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
		return c.Market.ProcessExpiredSlots(opts)
	})
}

// ActivateSlots bootstraps or refills idle challenge slots.
func (c *Client) ActivateSlots(ctx context.Context) (*ethtypes.Receipt, error) {
	return c.SendTx(ctx, func(opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
		return c.Market.ActivateSlots(opts)
	})
}

// QuitOrder exits from an order (with slashing penalty).
func (c *Client) QuitOrder(ctx context.Context, orderID *big.Int) (*ethtypes.Receipt, error) {
	return c.SendTx(ctx, func(opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
		return c.Market.QuitOrder(opts, orderID)
	})
}

// PlaceOrder creates a new file storage order on-chain. The caller must provide
// the FSP proof (compressed [4]*big.Int) and send sufficient value to cover
// numChunks * periods * replicas * pricePerChunkPerPeriod.
func (c *Client) PlaceOrder(ctx context.Context, fileRoot *big.Int, fileUri string, numChunks uint32, periods uint16, replicas uint8, pricePerChunkPerPeriod *big.Int, fspProof [4]*big.Int) (*big.Int, *ethtypes.Receipt, error) {
	totalCost := new(big.Int).SetUint64(uint64(numChunks))
	totalCost.Mul(totalCost, new(big.Int).SetUint64(uint64(periods)))
	totalCost.Mul(totalCost, pricePerChunkPerPeriod)
	totalCost.Mul(totalCost, new(big.Int).SetUint64(uint64(replicas)))

	receipt, err := c.SendTx(ctx, func(opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
		opts.Value = totalCost
		return c.Market.PlaceOrder(opts, fileRoot, fileUri, numChunks, periods, replicas, pricePerChunkPerPeriod, fspProof)
	})
	if err != nil {
		return nil, nil, err
	}

	// Parse OrderPlaced event to extract the order ID
	for _, vLog := range receipt.Logs {
		parsed, err := c.Market.ParseOrderPlaced(*vLog)
		if err == nil {
			return parsed.OrderId, receipt, nil
		}
	}

	return nil, receipt, fmt.Errorf("OrderPlaced event not found in receipt")
}

// HasUnresolvedProofObligation checks if this node has a pending proof.
func (c *Client) HasUnresolvedProofObligation(ctx context.Context) (bool, error) {
	return c.Market.HasUnresolvedProofObligation(c.callOpts(ctx), c.addr)
}
