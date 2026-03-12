package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

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

	_, client, ctx, cancel, err := loadClient(*configPath, 2*time.Minute)
	if err != nil {
		fatal("%v", err)
	}
	defer cancel()
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

	_, client, ctx, cancel, err := loadClient(*configPath, 2*time.Minute)
	if err != nil {
		fatal("%v", err)
	}
	defer cancel()
	defer client.Close()

	fmt.Printf("Decreasing capacity by %d chunks (%s)...\n", reduce, formatChunkSize(reduce))
	receipt, err := client.DecreaseCapacity(ctx, reduce)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: decrease capacity failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Capacity decreased! tx: %s\n", receipt.TxHash.Hex())
}
