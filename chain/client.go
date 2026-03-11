package chain

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"sync"
	"time"

	"encoding/hex"
	"strings"

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

	// WebSocket client and filterer for event subscriptions (nil if listen_mode == "poll").
	wsEth          *ethclient.Client
	Filterer       *bindings.FileMarketFilterer
	wsMarketCaller *bindings.FileMarketCaller // read-only caller via WS (bypasses HTTP caching)

	// txMu serializes SendTx calls so concurrent callers (challenge loop,
	// maintenance loop) don't race on nonce management.
	txMu sync.Mutex
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

	// Set up WebSocket client for event subscriptions if configured
	if cfg.WSURL != "" {
		wsEth, wsErr := ethclient.DialContext(ctx, cfg.WSURL)
		if wsErr != nil {
			log.Warn().Err(wsErr).Str("ws_url", cfg.WSURL).Msg("failed to connect WebSocket client, events will be unavailable")
		} else {
			c.wsEth = wsEth
			filterer, fErr := bindings.NewFileMarketFilterer(marketAddr, wsEth)
			if fErr != nil {
				log.Warn().Err(fErr).Msg("failed to create event filterer, events will be unavailable")
				wsEth.Close()
				c.wsEth = nil
			} else {
				c.Filterer = filterer
				// Create a read-only market caller via WS for cache-bypassing reads
				wsCaller, wcErr := bindings.NewFileMarketCaller(marketAddr, wsEth)
				if wcErr != nil {
					log.Warn().Err(wcErr).Msg("failed to create WS market caller (will use HTTP for reads)")
				} else {
					c.wsMarketCaller = wsCaller
				}
				log.Info().Str("ws_url", cfg.WSURL).Msg("WebSocket event subscriptions enabled")
			}
		}
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

// Close shuts down the client connections.
func (c *Client) Close() {
	if c.wsEth != nil {
		c.wsEth.Close()
	}
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
// Retries reuse the same nonce so the replacement tx overwrites the stuck
// original in the mempool instead of queuing behind it.
// Serialized via txMu to prevent concurrent callers from racing on nonces.
//
// When a retry gets "nonce too low", the original tx was mined during the
// receipt timeout. SendTx recovers by fetching the original tx's receipt
// instead of reporting failure.
func (c *Client) SendTx(ctx context.Context, fn func(*bind.TransactOpts) (*types.Transaction, error)) (*types.Receipt, error) {
	c.txMu.Lock()
	defer c.txMu.Unlock()

	var lastErr error
	var firstTxHash common.Hash // track original tx for nonce-too-low recovery
	maxRetries := c.cfg.MaxRetries
	if maxRetries == 0 {
		maxRetries = 1
	}

	// Pin the nonce on the first attempt so retries replace rather than enqueue.
	nonce, err := c.eth.PendingNonceAt(ctx, c.addr)
	if err != nil {
		return nil, fmt.Errorf("get pending nonce: %w", err)
	}

	for attempt := 0; attempt < maxRetries; attempt++ {
		opts, err := c.transactOpts(ctx)
		if err != nil {
			return nil, err
		}
		opts.Nonce = new(big.Int).SetUint64(nonce)

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
			// "nonce too low" during retry means the original tx was mined
			// while we were waiting. Recover its receipt instead of failing.
			if attempt > 0 && isNonceTooLow(err) && firstTxHash != (common.Hash{}) {
				log.Info().Str("tx", firstTxHash.Hex()).Msg("nonce consumed — recovering receipt for original tx")
				if receipt, recErr := c.recoverReceipt(ctx, firstTxHash); recErr == nil {
					return receipt, nil
				} else {
					log.Warn().Err(recErr).Str("tx", firstTxHash.Hex()).Msg("receipt recovery failed")
				}
			}
			lastErr = err
			log.Warn().Err(err).Int("attempt", attempt+1).Msg("tx submission failed")
			continue
		}

		// Track the first successfully submitted tx hash
		if firstTxHash == (common.Hash{}) {
			firstTxHash = tx.Hash()
		}

		receipt, err := c.waitForReceipt(ctx, tx.Hash())
		if err != nil {
			lastErr = err
			log.Warn().Err(err).Str("tx", tx.Hash().Hex()).Msg("waiting for receipt failed")
			continue
		}

		if receipt.Status == 0 {
			reason := c.getRevertReason(ctx, tx)
			return nil, fmt.Errorf("tx reverted: %s (reason: %s)", tx.Hash().Hex(), reason)
		}

		return receipt, nil
	}

	// Final recovery: if we exhausted retries but have an original tx hash,
	// try one more time to get its receipt (it may have been mined during
	// the last retry attempt).
	if firstTxHash != (common.Hash{}) {
		if receipt, err := c.recoverReceipt(ctx, firstTxHash); err == nil {
			return receipt, nil
		}
	}

	return nil, fmt.Errorf("all %d attempts failed: %w", maxRetries, lastErr)
}

// isNonceTooLow checks if an error indicates the nonce has been consumed.
func isNonceTooLow(err error) bool {
	return strings.Contains(strings.ToLower(err.Error()), "nonce too low")
}

// recoverReceipt polls briefly for a transaction receipt that was mined
// while the retry loop was running. Returns the receipt or an error if
// not found after a short wait.
func (c *Client) recoverReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	for i := 0; i < 5; i++ {
		receipt, err := c.eth.TransactionReceipt(ctx, txHash)
		if err == nil {
			if receipt.Status == 0 {
				return nil, fmt.Errorf("original tx reverted: %s", txHash.Hex())
			}
			log.Info().Str("tx", txHash.Hex()).Uint64("block", receipt.BlockNumber.Uint64()).Msg("recovered receipt for original tx (mined during retry)")
			return receipt, nil
		}
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(2 * time.Second):
		}
	}
	return nil, fmt.Errorf("could not recover receipt for %s", txHash.Hex())
}

// receiptTimeout bounds how long waitForReceipt will poll for a pending tx
// before giving up. Set to 4 minutes to accommodate Avalanche L1's on-demand
// block production (~150s between blocks in low-traffic periods).
const receiptTimeout = 4 * time.Minute

// waitForReceipt polls for a transaction receipt, distinguishing
// "not yet mined" (NotFound) from real RPC errors.
func (c *Client) waitForReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	const maxErrors = 10
	consecutiveErrors := 0

	deadline := time.After(receiptTimeout)

	for {
		receipt, err := c.eth.TransactionReceipt(ctx, txHash)
		if err == nil {
			return receipt, nil
		}

		// NotFound means tx is pending — keep waiting (up to timeout)
		if errors.Is(err, ethereum.NotFound) {
			consecutiveErrors = 0
			select {
			case <-ctx.Done():
				return nil, ctx.Err()
			case <-deadline:
				return nil, fmt.Errorf("receipt timeout after %s: tx %s still pending", receiptTimeout, txHash.Hex())
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
		case <-deadline:
			return nil, fmt.Errorf("receipt timeout after %s: tx %s still pending", receiptTimeout, txHash.Hex())
		case <-time.After(2 * time.Second):
		}
	}
}

// getRevertReason replays a failed transaction as a call to extract the revert data,
// then decodes it into a human-readable string.
func (c *Client) getRevertReason(ctx context.Context, tx *types.Transaction) string {
	// Replay the tx as an eth_call at the block it was mined
	msg := ethereum.CallMsg{
		From:      c.addr,
		To:        tx.To(),
		Gas:       tx.Gas(),
		GasFeeCap: tx.GasFeeCap(),
		GasTipCap: tx.GasTipCap(),
		Value:     tx.Value(),
		Data:      tx.Data(),
	}
	result, err := c.eth.CallContract(ctx, msg, nil)
	if err != nil {
		// Try multiple extraction methods
		if data := extractRevertData(err); len(data) > 0 {
			return decodeRevertData(data)
		}
		// Try parsing hex from the error string itself (some RPCs embed it)
		if data := extractHexFromError(err.Error()); len(data) > 0 {
			return decodeRevertData(data)
		}
		// Return the raw error as the reason
		return err.Error()
	}
	// If call succeeded (shouldn't happen for a reverted tx), return raw output
	if len(result) > 0 {
		return "0x" + hex.EncodeToString(result)
	}
	return "unknown (replay succeeded — state may have changed)"
}

// extractRevertData tries to pull hex revert data from an error via ErrorData() interface.
func extractRevertData(err error) []byte {
	// go-ethereum wraps revert data in various error types
	var rpcErr interface{ ErrorData() interface{} }
	if errors.As(err, &rpcErr) {
		switch data := rpcErr.ErrorData().(type) {
		case string:
			if len(data) > 2 && data[:2] == "0x" {
				b, _ := hex.DecodeString(data[2:])
				return b
			}
		case []byte:
			return data
		}
	}
	return nil
}

// extractHexFromError scans an error string for a 0x-prefixed hex revert payload.
func extractHexFromError(s string) []byte {
	// Look for patterns like "0x08c379a0..." or "revert: 0x..." in the error string
	for i := 0; i < len(s)-3; i++ {
		if s[i] == '0' && s[i+1] == 'x' {
			end := i + 2
			for end < len(s) && isHexChar(s[end]) {
				end++
			}
			hexStr := s[i+2 : end]
			if len(hexStr) >= 8 { // at least a 4-byte selector
				b, err := hex.DecodeString(hexStr)
				if err == nil {
					return b
				}
			}
		}
	}
	return nil
}

func isHexChar(c byte) bool {
	return (c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F')
}

// decodeRevertData decodes ABI-encoded revert data into a human-readable reason.
// Handles: Error(string) from require(), and known custom errors.
func decodeRevertData(data []byte) string {
	if len(data) == 0 {
		return "unknown"
	}

	// Error(string) — selector 0x08c379a0
	if len(data) >= 68 && data[0] == 0x08 && data[1] == 0xc3 && data[2] == 0x79 && data[3] == 0xa0 {
		// ABI decode: offset (32 bytes) + length (32 bytes) + string data
		strLen := new(big.Int).SetBytes(data[36:68]).Uint64()
		if 68+strLen <= uint64(len(data)) {
			return string(data[68 : 68+strLen])
		}
	}

	// Known custom errors (from verifier contract)
	selector := hex.EncodeToString(data[:4])
	switch selector {
	case "7fcdd1f4":
		return "ProofInvalid()"
	case "a54f8e27":
		return "PublicInputNotInField()"
	}

	return "0x" + hex.EncodeToString(data)
}

// GetBalance returns the native token balance in wei.
func (c *Client) GetBalance(ctx context.Context) (*big.Int, error) {
	return c.eth.BalanceAt(ctx, c.addr, nil)
}

// callOpts returns CallOpts with the given context.
func (c *Client) callOpts(ctx context.Context) *bind.CallOpts {
	return &bind.CallOpts{Context: ctx, From: c.addr}
}
