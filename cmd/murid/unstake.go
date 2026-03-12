package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func runUnstake(args []string) {
	fs := flag.NewFlagSet("unstake", flag.ExitOnError)
	configPath := fs.String("config", "murid.toml", "path to config file")
	fs.Parse(args)

	_, client, ctx, cancel, err := loadClient(*configPath, 2*time.Minute)
	if err != nil {
		fatal("%v", err)
	}
	defer cancel()
	defer client.Close()

	info, err := client.GetNodeInfo(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: get node info: %v\n", err)
		os.Exit(1)
	}

	if info.Stake.Sign() == 0 {
		fmt.Println("Node is not registered.")
		return
	}

	if info.Used > 0 {
		fmt.Fprintf(os.Stderr, "error: node has %d active order chunks — cannot unstake while serving orders\n", info.Used)
		os.Exit(1)
	}

	fmt.Printf("Node address: %s\n", client.Address())
	fmt.Printf("Current stake: %s MURI\n", formatWei(info.Stake))
	fmt.Printf("Capacity: %d chunks (%s)\n", info.Capacity, formatChunkSize(info.Capacity))
	fmt.Println()
	fmt.Println("Submitting unstakeNode transaction...")

	receipt, err := client.UnstakeNode(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: unstake failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Unstaked successfully! tx: %s\n", receipt.TxHash.Hex())
	fmt.Printf("Stake of %s MURI has been returned.\n", formatWei(info.Stake))
}
