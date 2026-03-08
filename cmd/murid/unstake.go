package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/MuriData/muri-node/chain"
	"github.com/MuriData/muri-node/config"
	"github.com/MuriData/muri-node/storage"
)

func runUnstake(args []string) {
	fs := flag.NewFlagSet("unstake", flag.ExitOnError)
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

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	client, err := chain.NewClient(ctx, cfg.Chain, privKey)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: connect to chain: %v\n", err)
		os.Exit(1)
	}
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

func runIncreaseCapacity(args []string) {
	fs := flag.NewFlagSet("increase-capacity", flag.ExitOnError)
	configPath := fs.String("config", "murid.toml", "path to config file")
	chunks := fs.Uint64("chunks", 0, "additional chunks to add")
	gb := fs.Float64("gb", 0, "additional capacity in GB (overrides -chunks)")
	fs.Parse(args)

	additional := *chunks
	if *gb > 0 {
		additional = uint64(*gb * 1024 * 1024 * 1024 / 16384)
	}
	if additional == 0 {
		fmt.Fprintf(os.Stderr, "error: specify -chunks or -gb\n")
		os.Exit(1)
	}

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

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	client, err := chain.NewClient(ctx, cfg.Chain, privKey)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: connect to chain: %v\n", err)
		os.Exit(1)
	}
	defer client.Close()

	fmt.Printf("Increasing capacity by %d chunks (%s)...\n", additional, formatChunkSize(additional))
	receipt, err := client.IncreaseCapacity(ctx, additional)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: increase capacity failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Capacity increased! tx: %s\n", receipt.TxHash.Hex())
}

func runDecreaseCapacity(args []string) {
	fs := flag.NewFlagSet("decrease-capacity", flag.ExitOnError)
	configPath := fs.String("config", "murid.toml", "path to config file")
	chunks := fs.Uint64("chunks", 0, "chunks to remove")
	gb := fs.Float64("gb", 0, "capacity to remove in GB (overrides -chunks)")
	fs.Parse(args)

	reduce := *chunks
	if *gb > 0 {
		reduce = uint64(*gb * 1024 * 1024 * 1024 / 16384)
	}
	if reduce == 0 {
		fmt.Fprintf(os.Stderr, "error: specify -chunks or -gb\n")
		os.Exit(1)
	}

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

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	client, err := chain.NewClient(ctx, cfg.Chain, privKey)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: connect to chain: %v\n", err)
		os.Exit(1)
	}
	defer client.Close()

	fmt.Printf("Decreasing capacity by %d chunks (%s)...\n", reduce, formatChunkSize(reduce))
	receipt, err := client.DecreaseCapacity(ctx, reduce)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: decrease capacity failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Capacity decreased! tx: %s\n", receipt.TxHash.Hex())
}
