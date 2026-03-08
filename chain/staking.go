package chain

import (
	"context"
	"fmt"
	"math/big"

	"github.com/MuriData/muri-node/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

// StakePerChunk is the collateral per chunk in wei (4 * 10^14 = 0.0004 MURI).
const StakePerChunk = 400_000_000_000_000

// GetNodeInfo returns the on-chain node registration info.
func (c *Client) GetNodeInfo(ctx context.Context) (*types.NodeInfo, error) {
	info, err := c.Staking.GetNodeInfo(c.callOpts(ctx), c.addr)
	if err != nil {
		return nil, fmt.Errorf("getNodeInfo: %w", err)
	}
	return &types.NodeInfo{
		Stake:     info.Stake,
		Capacity:  info.Capacity,
		Used:      info.Used,
		PublicKey: info.PublicKey,
	}, nil
}

// IsValidNode checks if this node is registered and has available capacity.
func (c *Client) IsValidNode(ctx context.Context) (bool, error) {
	return c.Staking.IsValidNode(c.callOpts(ctx), c.addr)
}

// StakeNode registers the node with the given capacity and public key.
// Value must be capacity * StakePerChunk.
func (c *Client) StakeNode(ctx context.Context, capacity uint64, publicKey *big.Int) (*ethtypes.Receipt, error) {
	value := new(big.Int).Mul(
		new(big.Int).SetUint64(capacity),
		big.NewInt(StakePerChunk),
	)
	return c.SendTx(ctx, func(opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
		opts.Value = value
		return c.Staking.StakeNode(opts, capacity, publicKey)
	})
}

// IncreaseCapacity adds capacity by staking additional tokens.
func (c *Client) IncreaseCapacity(ctx context.Context, additionalCapacity uint64) (*ethtypes.Receipt, error) {
	value := new(big.Int).Mul(
		new(big.Int).SetUint64(additionalCapacity),
		big.NewInt(StakePerChunk),
	)
	return c.SendTx(ctx, func(opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
		opts.Value = value
		return c.Staking.IncreaseCapacity(opts, additionalCapacity)
	})
}

// DecreaseCapacity removes capacity and unlocks proportional stake.
func (c *Client) DecreaseCapacity(ctx context.Context, reduceCapacity uint64) (*ethtypes.Receipt, error) {
	return c.SendTx(ctx, func(opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
		return c.Staking.DecreaseCapacity(opts, reduceCapacity)
	})
}

// UnstakeNode fully exits the node and withdraws all stake.
func (c *Client) UnstakeNode(ctx context.Context) (*ethtypes.Receipt, error) {
	return c.SendTx(ctx, func(opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
		return c.Staking.UnstakeNode(opts)
	})
}
