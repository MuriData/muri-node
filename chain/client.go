package chain

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/MuriData/muri-node/chain/bindings"
	"github.com/MuriData/muri-node/config"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog/log"
)

// Client wraps an EVM client with contract instances and transaction helpers.
type Client struct {
	cfg     config.ChainConfig
	eth     *ethclient.Client
	privKey *ecdsa.PrivateKey
	chainID *big.Int
	addr    common.Address

	Market  *bindings.FileMarket
	Staking *bindings.NodeStaking
}

// NewClient connects to the chain and instantiates contract bindings.
func NewClient(ctx context.Context, cfg config.ChainConfig, privKeyHex string) (*Client, error) {
	eth, err := ethclient.DialContext(ctx, cfg.RPCURL)
	if err != nil {
		return nil, fmt.Errorf("dial rpc: %w", err)
	}

	privKey, err := crypto.HexToECDSA(privKeyHex)
	if err != nil {
		return nil, fmt.Errorf("parse private key: %w", err)
	}

	addr := crypto.PubkeyToAddress(privKey.PublicKey)
	chainID := big.NewInt(cfg.ChainID)

	marketAddr := common.HexToAddress(cfg.MarketAddress)
	market, err := bindings.NewFileMarket(marketAddr, eth)
	if err != nil {
		return nil, fmt.Errorf("bind FileMarket: %w", err)
	}

	c := &Client{
		cfg:     cfg,
		eth:     eth,
		privKey: privKey,
		chainID: chainID,
		addr:    addr,
		Market:  market,
	}

	// Resolve staking address from contract if not configured
	stakingAddr := common.HexToAddress(cfg.StakingAddress)
	if stakingAddr == (common.Address{}) {
		stakingAddr, err = market.NodeStaking(&bind.CallOpts{Context: ctx})
		if err != nil {
			return nil, fmt.Errorf("resolve staking address: %w", err)
		}
		log.Info().Str("staking", stakingAddr.Hex()).Msg("resolved staking address from market")
	}

	staking, err := bindings.NewNodeStaking(stakingAddr, eth)
	if err != nil {
		return nil, fmt.Errorf("bind NodeStaking: %w", err)
	}
	c.Staking = staking

	return c, nil
}

// Address returns the node's EVM address.
func (c *Client) Address() common.Address {
	return c.addr
}

// BlockNumber returns the current block number.
func (c *Client) BlockNumber(ctx context.Context) (uint64, error) {
	return c.eth.BlockNumber(ctx)
}

// Close shuts down the client connection.
func (c *Client) Close() {
	c.eth.Close()
}

// transactOpts creates signed TransactOpts for sending transactions.
func (c *Client) transactOpts(ctx context.Context) (*bind.TransactOpts, error) {
	opts, err := bind.NewKeyedTransactorWithChainID(c.privKey, c.chainID)
	if err != nil {
		return nil, fmt.Errorf("create transactor: %w", err)
	}
	opts.Context = ctx

	if c.cfg.GasLimit > 0 {
		opts.GasLimit = c.cfg.GasLimit
	}

	// EIP-1559 gas pricing
	tip := new(big.Int).Mul(
		new(big.Int).SetUint64(c.cfg.GasPriority),
		big.NewInt(1_000_000_000), // gwei → wei
	)
	opts.GasTipCap = tip

	if c.cfg.MaxGasPrice > 0 {
		maxFee := new(big.Int).Mul(
			new(big.Int).SetUint64(c.cfg.MaxGasPrice),
			big.NewInt(1_000_000_000), // gwei → wei
		)
		opts.GasFeeCap = maxFee
	}

	return opts, nil
}

// SendTx sends a transaction with retry and gas escalation.
// fn is called with TransactOpts and should return the transaction.
func (c *Client) SendTx(ctx context.Context, fn func(*bind.TransactOpts) (*types.Transaction, error)) (*types.Receipt, error) {
	var lastErr error
	maxRetries := c.cfg.MaxRetries
	if maxRetries == 0 {
		maxRetries = 1
	}

	for attempt := 0; attempt < maxRetries; attempt++ {
		opts, err := c.transactOpts(ctx)
		if err != nil {
			return nil, err
		}

		// Escalate gas on retries
		if attempt > 0 {
			escalation := c.cfg.GasEscalation
			if escalation < 1.0 {
				escalation = 1.25
			}
			multiplier := 1.0
			for i := 0; i < attempt; i++ {
				multiplier *= escalation
			}
			if opts.GasFeeCap != nil {
				fee := new(big.Float).SetInt(opts.GasFeeCap)
				fee.Mul(fee, big.NewFloat(multiplier))
				opts.GasFeeCap, _ = fee.Int(nil)
			}
			if opts.GasTipCap != nil {
				tip := new(big.Float).SetInt(opts.GasTipCap)
				tip.Mul(tip, big.NewFloat(multiplier))
				opts.GasTipCap, _ = tip.Int(nil)
			}
			// Clamp tip to not exceed fee cap
			if opts.GasTipCap != nil && opts.GasFeeCap != nil && opts.GasTipCap.Cmp(opts.GasFeeCap) > 0 {
				opts.GasTipCap.Set(opts.GasFeeCap)
			}
			log.Warn().Int("attempt", attempt+1).Float64("multiplier", multiplier).Msg("retrying with escalated gas")
		}

		tx, err := fn(opts)
		if err != nil {
			lastErr = err
			log.Warn().Err(err).Int("attempt", attempt+1).Msg("tx submission failed")
			continue
		}

		receipt, err := c.waitForReceipt(ctx, tx.Hash())
		if err != nil {
			lastErr = err
			log.Warn().Err(err).Str("tx", tx.Hash().Hex()).Msg("waiting for receipt failed")
			continue
		}

		if receipt.Status == 0 {
			return nil, fmt.Errorf("tx reverted: %s", tx.Hash().Hex())
		}

		return receipt, nil
	}

	return nil, fmt.Errorf("all %d attempts failed: %w", maxRetries, lastErr)
}

// waitForReceipt polls for a transaction receipt, distinguishing
// "not yet mined" (NotFound) from real RPC errors.
func (c *Client) waitForReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	const maxErrors = 10
	consecutiveErrors := 0

	for {
		receipt, err := c.eth.TransactionReceipt(ctx, txHash)
		if err == nil {
			return receipt, nil
		}

		// NotFound means tx is pending — keep waiting
		if errors.Is(err, ethereum.NotFound) {
			consecutiveErrors = 0
			select {
			case <-ctx.Done():
				return nil, ctx.Err()
			case <-time.After(2 * time.Second):
			}
			continue
		}

		// Real RPC error — retry a few times then give up
		consecutiveErrors++
		if consecutiveErrors >= maxErrors {
			return nil, fmt.Errorf("receipt poll failed after %d errors: %w", maxErrors, err)
		}
		log.Warn().Err(err).Str("tx", txHash.Hex()).Int("errors", consecutiveErrors).Msg("receipt poll error, retrying")

		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(2 * time.Second):
		}
	}
}

// callOpts returns CallOpts with the given context.
func (c *Client) callOpts(ctx context.Context) *bind.CallOpts {
	return &bind.CallOpts{Context: ctx, From: c.addr}
}
