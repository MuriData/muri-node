package main

import (
	"context"
	"flag"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/MuriData/muri-node/chain"
	"github.com/MuriData/muri-node/config"
	"github.com/MuriData/muri-node/storage"
)

func runStatus(args []string) {
	fs := flag.NewFlagSet("status", flag.ExitOnError)
	configPath := fs.String("config", "murid.toml", "path to config file")
	fs.Parse(args)

	cfg, err := config.Load(*configPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: load config: %v\n", err)
		os.Exit(1)
	}

	privKey, err := storage.LoadPrivateKey(cfg.Node.PrivateKeyPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: load private key: %v\n", err)
		os.Exit(1)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	client, err := chain.NewClient(ctx, cfg.Chain, privKey)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: connect to chain: %v\n", err)
		os.Exit(1)
	}
	defer client.Close()

	fmt.Printf("Node address: %s\n", client.Address())
	fmt.Printf("Chain: %d | RPC: %s\n", cfg.Chain.ChainID, cfg.Chain.RPCURL)

	balance, err := client.GetBalance(ctx)
	if err == nil {
		fmt.Printf("Balance: %s MURI\n", formatWei(balance))
	}
	fmt.Println()

	// Check registration
	isValid, err := client.IsValidNode(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: check registration: %v\n", err)
		os.Exit(1)
	}

	if !isValid {
		fmt.Println("Status: NOT REGISTERED")
		fmt.Println("  Run \"murid stake -capacity-gb 10\" to register and stake.")
		return
	}

	// Fetch node info
	info, err := client.GetNodeInfo(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: get node info: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Status: REGISTERED")
	fmt.Printf("  Stake:      %s MURI\n", formatWei(info.Stake))
	fmt.Printf("  Public Key: 0x%s\n", info.PublicKey.Text(16))
	fmt.Printf("  Capacity:   %d chunks (%s)\n", info.Capacity, formatChunkSize(info.Capacity))
	fmt.Printf("  Used:       %d chunks (%s)\n", info.Used, formatChunkSize(info.Used))
	if info.Capacity > 0 {
		pct := float64(info.Used) / float64(info.Capacity) * 100
		fmt.Printf("  Utilization: %.1f%%\n", pct)
	}

	// Orders
	orders, err := client.GetNodeOrders(ctx)
	if err == nil {
		fmt.Printf("  Orders:     %d active\n", len(orders))
	}

	// Claimable rewards
	rewards, err := client.GetClaimableRewards(ctx)
	if err == nil {
		fmt.Printf("  Claimable:  %s MURI\n", formatWei(rewards))
	}

	// Challenge status
	hasObligation, err := client.HasUnresolvedProofObligation(ctx)
	if err == nil && hasObligation {
		fmt.Println("  ⚠ Unresolved proof obligation — respond ASAP!")
	}

	fmt.Println()
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
