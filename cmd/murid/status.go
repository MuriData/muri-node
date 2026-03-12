package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func runStatus(args []string) {
	fs := flag.NewFlagSet("status", flag.ExitOnError)
	configPath := fs.String("config", "murid.toml", "path to config file")
	fs.Parse(args)

	cfg, client, ctx, cancel, err := loadClient(*configPath, 30*time.Second)
	if err != nil {
		fatal("%v", err)
	}
	defer cancel()
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
