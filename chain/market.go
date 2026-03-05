package chain

import (
	"context"
	"fmt"
	"math/big"

	"github.com/MuriData/muri-node/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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

// GetAllSlotInfo returns all challenge slot states (dynamic count via sqrt scaling).
func (c *Client) GetAllSlotInfo(ctx context.Context) ([]types.ChallengeSlotInfo, error) {
	result, err := c.Market.GetAllSlotInfo(c.callOpts(ctx))
	if err != nil {
		return nil, fmt.Errorf("getAllSlotInfo: %w", err)
	}

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
	return slots, nil
}

// GetNodeOrders returns all order IDs assigned to a node.
func (c *Client) GetNodeOrders(ctx context.Context) ([]*big.Int, error) {
	return c.Market.GetNodeOrders(c.callOpts(ctx), c.addr)
}

// GetActiveOrders returns all active order IDs.
func (c *Client) GetActiveOrders(ctx context.Context) ([]*big.Int, error) {
	return c.Market.GetActiveOrders(c.callOpts(ctx))
}

// GetOrderDetails returns order details for the given order ID.
func (c *Client) GetOrderDetails(ctx context.Context, orderID *big.Int) (*types.OrderInfo, error) {
	d, err := c.Market.GetOrderDetails(c.callOpts(ctx), orderID)
	if err != nil {
		return nil, fmt.Errorf("getOrderDetails(%s): %w", orderID, err)
	}

	f, err := c.Market.GetOrderFinancials(c.callOpts(ctx), orderID)
	if err != nil {
		return nil, fmt.Errorf("getOrderFinancials(%s): %w", orderID, err)
	}

	// Price is not returned by getOrderDetails/getOrderFinancials views;
	// fetch from the orders mapping directly.
	raw, err := c.Market.Orders(c.callOpts(ctx), orderID)
	if err != nil {
		return nil, fmt.Errorf("orders(%s): %w", orderID, err)
	}

	return &types.OrderInfo{
		ID:          orderID,
		Owner:       d.Owner,
		URI:         d.Uri,
		RootHash:    d.Root,
		NumChunks:   d.NumChunks,
		Periods:     d.Periods,
		Replicas:    d.Replicas,
		Filled:      d.Filled,
		Price:       raw.Price,
		Escrow:      f.Escrow,
		StartPeriod: f.StartPeriod,
	}, nil
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
	CurrentRandomnessValue   *big.Int
	ActiveChallengeSlots     *big.Int
	CurrentPeriod            *big.Int
	CurrentBlock             *big.Int
	ChallengeableOrdersCount *big.Int
}, error) {
	return c.Market.GetGlobalStats(c.callOpts(ctx))
}

// ExecuteOrder claims a replica slot on the given order with a PoI proof of data possession.
func (c *Client) ExecuteOrder(ctx context.Context, orderID *big.Int, proof [8]*big.Int, commitment [32]byte) (*ethtypes.Receipt, error) {
	return c.SendTx(ctx, func(opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
		return c.Market.ExecuteOrder(opts, orderID, proof, commitment)
	})
}

// SubmitProof submits a ZK proof for a challenge slot.
func (c *Client) SubmitProof(ctx context.Context, slotIndex int, proof [8]*big.Int, commitment [32]byte) (*ethtypes.Receipt, error) {
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

// HasUnresolvedProofObligation checks if this node has a pending proof.
func (c *Client) HasUnresolvedProofObligation(ctx context.Context) (bool, error) {
	return c.Market.HasUnresolvedProofObligation(c.callOpts(ctx), c.addr)
}
