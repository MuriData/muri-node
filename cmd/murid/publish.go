package main

import (
	"flag"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"time"

	"github.com/MuriData/muri-node/ipfs"
	"github.com/MuriData/muri-node/prover"
	"github.com/rs/zerolog/log"
)

const publishUsage = `Usage: murid web publish [flags]

Upload a local directory to IPFS and place MuriData orders for every file
and directory DAG node. This is the easiest way to host a static website
on IPFS with MuriData storage guarantees.

Flags:
  -config string    Path to config file (default "murid.toml")
  -dir string       Local directory to publish (required)
  -periods uint     Number of storage periods (default 1)
  -replicas uint    Number of replicas per file (default 1)
  -price string     Price per chunk per period in wei (required)

Example:
  murid web publish -dir ./dist -periods 4 -replicas 3 -price 1000000000000
`

func runWebPublish(args []string) {
	fs := flag.NewFlagSet("web publish", flag.ExitOnError)
	configPath := fs.String("config", "murid.toml", "path to config file")
	dir := fs.String("dir", "", "local directory to publish")
	periods := fs.Uint("periods", 1, "number of storage periods")
	replicas := fs.Uint("replicas", 1, "number of replicas per file")
	priceStr := fs.String("price", "", "price per chunk per period in wei")
	fs.Parse(args)

	if *dir == "" || *priceStr == "" {
		fmt.Print(publishUsage)
		os.Exit(1)
	}

	price, ok := new(big.Int).SetString(*priceStr, 10)
	if !ok || price.Sign() <= 0 {
		fatal("invalid price: %s", *priceStr)
	}
	if *periods == 0 || *periods > 65535 {
		fatal("periods must be 1-65535")
	}
	if *replicas == 0 || *replicas > 10 {
		fatal("replicas must be 1-10")
	}

	absDir, err := filepath.Abs(*dir)
	if err != nil {
		fatal("resolve directory: %v", err)
	}
	info, err := os.Stat(absDir)
	if err != nil || !info.IsDir() {
		fatal("%s is not a directory", absDir)
	}

	// Load config and connect to chain (30 min timeout for large sites)
	cfg, client, ctx, cancel, err := loadClient(*configPath, 30*time.Minute)
	if err != nil {
		fatal("%v", err)
	}
	defer cancel()
	defer client.Close()

	ipfsClient := ipfs.NewClient(cfg.IPFS)

	// Check IPFS connectivity
	if err := ipfsClient.Ping(ctx); err != nil {
		fatal("IPFS not reachable: %v", err)
	}

	// Initialize FSP prover
	fmt.Println("Initializing FSP prover...")
	fspProver, err := prover.NewFSPProver(cfg.Node.KeysDir)
	if err != nil {
		fatal("init FSP prover: %v", err)
	}

	// Upload directory to IPFS
	fmt.Printf("Uploading %s to IPFS...\n", absDir)
	entries, err := ipfsClient.AddDirectory(ctx, absDir)
	if err != nil {
		fatal("IPFS upload: %v", err)
	}
	fmt.Printf("Uploaded %d entries to IPFS\n", len(entries))

	// Build set of local directory paths to classify IPFS entries.
	// Kubo returns files with relative paths and directories with their path
	// (empty string for the wrap root).
	dirSet := map[string]bool{"": true}
	filepath.Walk(absDir, func(path string, fi os.FileInfo, walkErr error) error {
		if walkErr != nil || !fi.IsDir() {
			return nil
		}
		rel, _ := filepath.Rel(absDir, path)
		rel = filepath.ToSlash(rel)
		if rel == "." {
			rel = ""
		}
		dirSet[rel] = true
		return nil
	})

	type entry struct {
		cid   string
		name  string
		isDir bool
	}
	var toOrder []entry
	for _, e := range entries {
		toOrder = append(toOrder, entry{cid: e.Hash, name: e.Name, isDir: dirSet[e.Name]})
	}

	// Show summary
	var fileCount, dirCount int
	for _, o := range toOrder {
		if o.isDir {
			dirCount++
		} else {
			fileCount++
		}
	}
	fmt.Printf("\nFiles: %d, Directory nodes: %d\n", fileCount, dirCount)
	fmt.Printf("Periods: %d, Replicas: %d, Price: %s wei/chunk/period\n", *periods, *replicas, price.String())
	fmt.Println()

	// Process each entry: fetch data -> FSP proof -> place order
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

		// Generate FSP proof
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
		cost.Mul(cost, new(big.Int).SetUint64(uint64(*periods)))
		cost.Mul(cost, price)
		cost.Mul(cost, new(big.Int).SetUint64(uint64(*replicas)))

		fmt.Printf("  Chunks: %d, Root: 0x%x\n", fspResult.NumChunks, fspResult.RootHash)
		fmt.Printf("  Placing order (cost: %s wei)...\n", cost.String())

		// Place order on-chain
		orderID, receipt, err := client.PlaceOrder(
			ctx,
			fspResult.RootHash,
			uri,
			fspResult.NumChunks,
			uint16(*periods),
			uint8(*replicas),
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
	fmt.Println("\n========== PUBLISH SUMMARY ==========")
	rootCID := ""
	successCount := 0
	failCount := 0
	for _, r := range results {
		displayName := r.name
		if displayName == "" {
			displayName = "(root)"
			if r.err == nil {
				rootCID = r.cid
			}
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

	if rootCID != "" {
		fmt.Printf("\nRoot CID: %s\n", rootCID)
		fmt.Printf("Access via IPFS gateway: https://ipfs.io/ipfs/%s\n", rootCID)
	}

	if failCount > 0 {
		os.Exit(1)
	}
}
