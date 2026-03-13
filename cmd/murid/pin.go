package main

import (
	"context"
	"flag"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/MuriData/muri-node/ipfs"
	"github.com/MuriData/muri-node/prover"
	"github.com/rs/zerolog/log"
)

const pinUsage = `Usage: murid pin <subcommand> [flags]

Subcommands:
  file    Pin an existing IPFS file to MuriData
  dir     Pin an existing IPFS directory (and all contents) to MuriData

Run "murid pin <subcommand> -h" for subcommand-specific help.
`

const pinFileUsage = `Usage: murid pin file [flags]

Pin an existing IPFS file to MuriData. Downloads the file from IPFS,
generates an FSP proof, and places an on-chain storage order.

Flags:
  -config string    Path to config file (default "murid.toml")
  -cid string       IPFS CID of the file to pin (required)
  -periods uint     Number of storage periods (default 1)
  -replicas uint    Number of replicas (default 1)
  -price string     Price per chunk per period in wei (required)

Example:
  murid pin file -cid bafybeigdyrzt5sfp7udm7hu76uh7y26nf3efuylqabf3oepc -periods 4 -replicas 3 -price 1000000000000
`

const pinDirUsage = `Usage: murid pin dir [flags]

Pin an existing IPFS directory to MuriData. Recursively enumerates all
files and directory DAG nodes, downloads each, generates FSP proofs,
and places on-chain storage orders for every entry.

Directory DAG nodes (the raw blocks that define the directory structure)
are also pinned with ?type=raw URIs so the full directory is recoverable.

Flags:
  -config string    Path to config file (default "murid.toml")
  -cid string       IPFS CID of the directory to pin (required)
  -periods uint     Number of storage periods (default 1)
  -replicas uint    Number of replicas (default 1)
  -price string     Price per chunk per period in wei (required)

Example:
  murid pin dir -cid bafybeigdyrzt5sfp7udm7hu76uh7y26nf3efuylqabf3oepc -periods 4 -replicas 3 -price 1000000000000
`

func runPin(args []string) {
	if len(args) == 0 {
		fmt.Print(pinUsage)
		os.Exit(1)
	}

	switch args[0] {
	case "file":
		runPinFile(args[1:])
	case "dir":
		runPinDir(args[1:])
	case "-h", "--help", "help":
		fmt.Print(pinUsage)
		os.Exit(0)
	default:
		fmt.Fprintf(os.Stderr, "unknown pin subcommand: %s\n\n%s", args[0], pinUsage)
		os.Exit(1)
	}
}

func parsePinFlags(name, usageText string, args []string) (configPath string, cid string, periods uint, replicas uint, price *big.Int) {
	fs := flag.NewFlagSet("pin "+name, flag.ExitOnError)
	cfgPtr := fs.String("config", "murid.toml", "path to config file")
	cidPtr := fs.String("cid", "", "IPFS CID to pin")
	periodsPtr := fs.Uint("periods", 1, "number of storage periods")
	replicasPtr := fs.Uint("replicas", 1, "number of replicas")
	priceStr := fs.String("price", "", "price per chunk per period in wei")
	fs.Parse(args)

	if *cidPtr == "" || *priceStr == "" {
		fmt.Print(usageText)
		os.Exit(1)
	}

	p, ok := new(big.Int).SetString(*priceStr, 10)
	if !ok || p.Sign() <= 0 {
		fatal("invalid price: %s", *priceStr)
	}
	if *periodsPtr == 0 || *periodsPtr > 65535 {
		fatal("periods must be 1-65535")
	}
	if *replicasPtr == 0 || *replicasPtr > 10 {
		fatal("replicas must be 1-10")
	}

	return *cfgPtr, *cidPtr, *periodsPtr, *replicasPtr, p
}

// pinEntry represents a single IPFS entry to be ordered on MuriData.
type pinEntry struct {
	cid   string
	name  string // display path (relative)
	isDir bool
}

// walkIPFSDir recursively enumerates all entries under an IPFS directory CID.
// Returns all files and directory DAG nodes (including the root itself).
func walkIPFSDir(ctx context.Context, ipfsClient *ipfs.Client, rootCID string) ([]pinEntry, error) {
	type walkItem struct {
		cid  string
		path string // relative display path
	}

	var entries []pinEntry
	stack := []walkItem{{cid: rootCID, path: ""}}

	for len(stack) > 0 {
		if ctx.Err() != nil {
			return nil, ctx.Err()
		}

		// Pop from stack
		item := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// This directory node itself needs an order
		entries = append(entries, pinEntry{cid: item.cid, name: item.path, isDir: true})

		// List children
		links, err := ipfsClient.LsWithRetry(ctx, item.cid)
		if err != nil {
			return nil, fmt.Errorf("ls %s: %w", item.cid, err)
		}

		for _, link := range links {
			childPath := link.Name
			if item.path != "" {
				childPath = item.path + "/" + link.Name
			}

			if link.Type == 2 {
				// Subdirectory — push onto stack for recursive enumeration
				stack = append(stack, walkItem{cid: link.Hash, path: childPath})
			} else {
				// File
				entries = append(entries, pinEntry{cid: link.Hash, name: childPath, isDir: false})
			}
		}
	}

	return entries, nil
}

func runPinFile(args []string) {
	configPath, cid, periods, replicas, price := parsePinFlags("file", pinFileUsage, args)

	cfg, client, ctx, cancel, err := loadClient(configPath, 30*time.Minute)
	if err != nil {
		fatal("%v", err)
	}
	defer cancel()
	defer client.Close()

	ipfsClient := ipfs.NewClient(cfg.IPFS)
	if err := ipfsClient.Ping(ctx); err != nil {
		fatal("IPFS not reachable: %v", err)
	}

	fmt.Println("Initializing FSP prover...")
	fspProver, err := prover.NewFSPProver(cfg.Node.KeysDir)
	if err != nil {
		fatal("init FSP prover: %v", err)
	}

	fmt.Printf("Downloading %s from IPFS...\n", cid)
	data, err := ipfsClient.CatChunked(ctx, cid)
	if err != nil {
		fatal("IPFS download: %v", err)
	}
	fmt.Printf("Downloaded %d bytes, generating FSP proof...\n", len(data))

	fspResult, err := fspProver.GenerateProof(data)
	if err != nil {
		fatal("FSP proof: %v", err)
	}

	uri := "ipfs://" + cid

	cost := new(big.Int).SetUint64(uint64(fspResult.NumChunks))
	cost.Mul(cost, new(big.Int).SetUint64(uint64(periods)))
	cost.Mul(cost, price)
	cost.Mul(cost, new(big.Int).SetUint64(uint64(replicas)))

	fmt.Printf("Chunks: %d, Root: 0x%x\n", fspResult.NumChunks, fspResult.RootHash)
	fmt.Printf("Placing order (cost: %s wei)...\n", cost.String())

	orderID, receipt, err := client.PlaceOrder(
		ctx,
		fspResult.RootHash,
		uri,
		fspResult.NumChunks,
		uint16(periods),
		uint8(replicas),
		price,
		fspResult.CompressedProof,
	)
	if err != nil {
		fatal("placeOrder: %v", err)
	}

	fmt.Printf("\nOrder #%s placed successfully\n", orderID.String())
	fmt.Printf("TX: %s\n", receipt.TxHash.Hex())
	fmt.Printf("URI: %s\n", uri)
	fmt.Printf("Cost: %s wei (%s MURI)\n", cost.String(), formatWei(cost))
}

func runPinDir(args []string) {
	configPath, cid, periods, replicas, price := parsePinFlags("dir", pinDirUsage, args)

	cfg, client, ctx, cancel, err := loadClient(configPath, 60*time.Minute)
	if err != nil {
		fatal("%v", err)
	}
	defer cancel()
	defer client.Close()

	ipfsClient := ipfs.NewClient(cfg.IPFS)
	if err := ipfsClient.Ping(ctx); err != nil {
		fatal("IPFS not reachable: %v", err)
	}

	fmt.Println("Initializing FSP prover...")
	fspProver, err := prover.NewFSPProver(cfg.Node.KeysDir)
	if err != nil {
		fatal("init FSP prover: %v", err)
	}

	// Recursively enumerate the directory
	fmt.Printf("Enumerating directory %s...\n", cid)
	toOrder, err := walkIPFSDir(ctx, ipfsClient, cid)
	if err != nil {
		fatal("enumerate directory: %v", err)
	}

	var fileCount, dirCount int
	for _, o := range toOrder {
		if o.isDir {
			dirCount++
		} else {
			fileCount++
		}
	}
	fmt.Printf("\nFiles: %d, Directory nodes: %d\n", fileCount, dirCount)
	fmt.Printf("Periods: %d, Replicas: %d, Price: %s wei/chunk/period\n", periods, replicas, price.String())
	fmt.Println()

	type result struct {
		name    string
		cid     string
		orderID *big.Int
		uri     string
		cost    *big.Int
		err     error
	}
	var results []result
	var totalCost big.Int

	for i, o := range toOrder {
		displayName := o.name
		if displayName == "" {
			displayName = "(root directory)"
		}
		fmt.Printf("[%d/%d] %s (CID: %s)\n", i+1, len(toOrder), displayName, o.cid)

		// Fetch data from IPFS
		var data []byte
		if o.isDir {
			data, err = ipfsClient.BlockGetWithRetry(ctx, o.cid)
		} else {
			data, err = ipfsClient.CatChunked(ctx, o.cid)
		}
		if err != nil {
			log.Error().Err(err).Str("cid", o.cid).Msg("failed to fetch from IPFS")
			results = append(results, result{name: o.name, cid: o.cid, err: err})
			continue
		}

		fmt.Printf("  %d bytes, generating FSP proof...\n", len(data))

		fspResult, err := fspProver.GenerateProof(data)
		if err != nil {
			log.Error().Err(err).Str("cid", o.cid).Msg("FSP proof failed")
			results = append(results, result{name: o.name, cid: o.cid, err: err})
			continue
		}

		// Build URI: directory DAGs get ?type=raw
		uri := "ipfs://" + o.cid
		if o.isDir {
			uri += "?type=raw"
		}

		// Calculate cost
		cost := new(big.Int).SetUint64(uint64(fspResult.NumChunks))
		cost.Mul(cost, new(big.Int).SetUint64(uint64(periods)))
		cost.Mul(cost, price)
		cost.Mul(cost, new(big.Int).SetUint64(uint64(replicas)))

		fmt.Printf("  Chunks: %d, Root: 0x%x\n", fspResult.NumChunks, fspResult.RootHash)
		fmt.Printf("  Placing order (cost: %s wei)...\n", cost.String())

		orderID, receipt, err := client.PlaceOrder(
			ctx,
			fspResult.RootHash,
			uri,
			fspResult.NumChunks,
			uint16(periods),
			uint8(replicas),
			price,
			fspResult.CompressedProof,
		)
		if err != nil {
			log.Error().Err(err).Str("cid", o.cid).Msg("placeOrder failed")
			results = append(results, result{name: o.name, cid: o.cid, err: err})
			continue
		}

		totalCost.Add(&totalCost, cost)
		fmt.Printf("  Order #%s placed (tx: %s)\n", orderID.String(), receipt.TxHash.Hex())
		results = append(results, result{
			name:    o.name,
			cid:     o.cid,
			orderID: orderID,
			uri:     uri,
			cost:    cost,
		})
	}

	// Print summary
	fmt.Println("\n========== PIN SUMMARY ==========")
	successCount := 0
	failCount := 0
	for _, r := range results {
		displayName := r.name
		if displayName == "" {
			displayName = "(root)"
		}
		if r.err != nil {
			fmt.Printf("  FAIL  %s: %v\n", displayName, r.err)
			failCount++
		} else {
			fmt.Printf("  OK    %s -> order #%s\n", displayName, r.orderID.String())
			successCount++
		}
	}

	fmt.Printf("\nOrders: %d succeeded, %d failed\n", successCount, failCount)
	fmt.Printf("Total cost: %s wei (%s MURI)\n", totalCost.String(), formatWei(&totalCost))
	fmt.Printf("\nRoot CID: %s\n", cid)
	fmt.Printf("Access via IPFS gateway: https://ipfs.io/ipfs/%s\n", cid)

	if failCount > 0 {
		os.Exit(1)
	}
}
