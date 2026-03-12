package main

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/MuriData/muri-node/chain"
	"github.com/MuriData/muri-node/config"
	"github.com/MuriData/muri-node/storage"
)

// loadClient loads config, reads the EVM private key, and connects to the chain.
// Returns the config, client, context, cancel function, and any error.
// The caller must defer cancel() and client.Close().
func loadClient(configPath string, timeout time.Duration) (*config.Config, *chain.Client, context.Context, context.CancelFunc, error) {
	cfg, err := config.Load(configPath)
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("load config: %w", err)
	}

	privKey, err := storage.LoadPrivateKey(cfg.Node.PrivateKeyPath)
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("load private key: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)

	client, err := chain.NewClient(ctx, cfg.Chain, privKey)
	if err != nil {
		cancel()
		return nil, nil, nil, nil, fmt.Errorf("connect to chain: %w", err)
	}

	return cfg, client, ctx, cancel, nil
}

// formatWei formats a wei amount as a human-readable MURI string.
func formatWei(wei *big.Int) string {
	if wei == nil || wei.Sign() == 0 {
		return "0"
	}
	ether := new(big.Float).SetInt(wei)
	ether.Quo(ether, new(big.Float).SetFloat64(1e18))
	return ether.Text('f', 6)
}

// formatChunkSize formats a chunk count as a human-readable size.
func formatChunkSize(chunks uint64) string {
	bytes := chunks * 16384
	switch {
	case bytes >= 1<<30:
		return fmt.Sprintf("%.2f GB", float64(bytes)/float64(1<<30))
	case bytes >= 1<<20:
		return fmt.Sprintf("%.1f MB", float64(bytes)/float64(1<<20))
	default:
		return fmt.Sprintf("%d KB", bytes/1024)
	}
}

// fatal prints an error message to stderr and exits.
func fatal(format string, args ...any) {
	fmt.Fprintf(os.Stderr, "error: "+format+"\n", args...)
	os.Exit(1)
}
