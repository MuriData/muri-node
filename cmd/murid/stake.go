package main

import (
	"flag"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/MuriData/muri-node/chain"
	"github.com/MuriData/muri-node/storage"
	muricrypto "github.com/MuriData/muri-zkproof/pkg/crypto"
)

const stakeUsage = `Usage: murid stake [flags]

Register and stake the node on-chain. Reads the ZK secret key from config
to derive the public key automatically.

Flags:
  -config string    Path to config file (default "murid.toml")
  -capacity uint    Storage capacity in chunks (1 chunk = 16 KB)
  -capacity-gb float  Storage capacity in GB (overrides -capacity)

Examples:
  murid stake -capacity 65536              # stake 65536 chunks (~1 GB)
  murid stake -capacity-gb 10              # stake ~10 GB
  murid stake -capacity-gb 0.5             # stake ~500 MB
`

func runStake(args []string) {
	fs := flag.NewFlagSet("stake", flag.ExitOnError)
	configPath := fs.String("config", "murid.toml", "path to config file")
	capacityChunks := fs.Uint64("capacity", 0, "capacity in chunks (1 chunk = 16 KB)")
	capacityGB := fs.Float64("capacity-gb", 0, "capacity in GB (overrides -capacity)")
	fs.Parse(args)

	if *capacityChunks == 0 && *capacityGB == 0 {
		fmt.Print(stakeUsage)
		os.Exit(1)
	}

	// Convert GB to chunks if specified
	chunks := *capacityChunks
	if *capacityGB > 0 {
		chunks = uint64(*capacityGB * 1024 * 1024 * 1024 / 16384)
		if chunks == 0 {
			fmt.Fprintf(os.Stderr, "error: capacity too small (minimum 1 chunk = 16 KB)\n")
			os.Exit(1)
		}
	}

	cfg, client, ctx, cancel, err := loadClient(*configPath, 2*time.Minute)
	if err != nil {
		fatal("%v", err)
	}
	defer cancel()
	defer client.Close()

	sk, err := storage.LoadSecretKey(cfg.Node.SecretKeyPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: load ZK secret key: %v\n", err)
		os.Exit(1)
	}
	publicKey := muricrypto.DerivePublicKey(sk)

	// Check if already registered
	isValid, err := client.IsValidNode(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: check registration: %v\n", err)
		os.Exit(1)
	}
	if isValid {
		fmt.Fprintf(os.Stderr, "error: node is already registered — use 'murid increase-capacity' to add more\n")
		os.Exit(1)
	}

	// Show what we're about to do
	stakeAmount := new(big.Int).Mul(
		new(big.Int).SetUint64(chunks),
		big.NewInt(chain.StakePerChunk),
	)

	balance, err := client.GetBalance(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: get balance: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Node address:  %s\n", client.Address())
	fmt.Printf("Balance:       %s MURI\n", formatWei(balance))
	fmt.Printf("Capacity:      %d chunks (%s)\n", chunks, formatChunkSize(chunks))
	fmt.Printf("Stake required: %s MURI\n", formatWei(stakeAmount))
	fmt.Printf("Public key:    0x%s\n", publicKey.Text(16))
	fmt.Println()

	if balance.Cmp(stakeAmount) < 0 {
		deficit := new(big.Int).Sub(stakeAmount, balance)
		fmt.Fprintf(os.Stderr, "error: insufficient balance — need %s more MURI\n", formatWei(deficit))
		fmt.Fprintf(os.Stderr, "Send MURI to %s and try again.\n", client.Address())
		os.Exit(1)
	}

	fmt.Println("Submitting stakeNode transaction...")
	receipt, err := client.StakeNode(ctx, chunks, publicKey)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: stake transaction failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Staked successfully! tx: %s\n", receipt.TxHash.Hex())
	fmt.Printf("Node is now registered with %d chunks (%s).\n", chunks, formatChunkSize(chunks))
	fmt.Printf("\nRun \"murid run\" to start serving storage.\n")
}
