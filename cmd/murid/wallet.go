package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/MuriData/muri-node/storage"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
)

const walletUsage = `Usage: murid wallet <subcommand> [flags]

Subcommands:
  new       Generate a new EVM wallet and save the key
  import    Import an existing private key from hex
  balance   Show wallet address and MURI balance

Examples:
  murid wallet new -out ./keys/node.key
  murid wallet import -key 0xabc123... -out ./keys/node.key
  murid wallet balance -config murid.toml
`

func runWallet(args []string) {
	if len(args) == 0 {
		fmt.Print(walletUsage)
		os.Exit(1)
	}

	switch args[0] {
	case "new":
		runWalletNew(args[1:])
	case "import":
		runWalletImport(args[1:])
	case "balance":
		runWalletBalance(args[1:])
	case "-h", "--help", "help":
		fmt.Print(walletUsage)
	default:
		fmt.Fprintf(os.Stderr, "unknown wallet subcommand: %s\n\n%s", args[0], walletUsage)
		os.Exit(1)
	}
}

func runWalletNew(args []string) {
	fs := flag.NewFlagSet("wallet new", flag.ExitOnError)
	outPath := fs.String("out", "./keys/node.key", "path to save the private key")
	force := fs.Bool("force", false, "overwrite existing key file")
	fs.Parse(args)

	if !*force {
		if _, err := os.Stat(*outPath); err == nil {
			fmt.Fprintf(os.Stderr, "error: %s already exists (use -force to overwrite)\n", *outPath)
			os.Exit(1)
		}
	}

	key, err := ethcrypto.GenerateKey()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: generate key: %v\n", err)
		os.Exit(1)
	}

	keyHex := hex.EncodeToString(ethcrypto.FromECDSA(key))
	if err := storage.SavePrivateKey(*outPath, keyHex); err != nil {
		fmt.Fprintf(os.Stderr, "error: save key: %v\n", err)
		os.Exit(1)
	}

	addr := ethcrypto.PubkeyToAddress(key.PublicKey)
	fmt.Printf("Private key saved to %s\n", *outPath)
	fmt.Printf("Address: %s\n", addr.Hex())
	fmt.Printf("\nFund this address with MURI to stake.\n")
}

func runWalletImport(args []string) {
	fs := flag.NewFlagSet("wallet import", flag.ExitOnError)
	keyHex := fs.String("key", "", "private key in hex (with or without 0x prefix)")
	outPath := fs.String("out", "./keys/node.key", "path to save the private key")
	force := fs.Bool("force", false, "overwrite existing key file")
	fs.Parse(args)

	if *keyHex == "" {
		fmt.Fprintf(os.Stderr, "error: -key is required\n")
		os.Exit(1)
	}

	clean := strings.TrimPrefix(strings.TrimSpace(*keyHex), "0x")

	// Validate the key parses correctly
	key, err := ethcrypto.HexToECDSA(clean)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: invalid private key: %v\n", err)
		os.Exit(1)
	}

	if !*force {
		if _, err := os.Stat(*outPath); err == nil {
			fmt.Fprintf(os.Stderr, "error: %s already exists (use -force to overwrite)\n", *outPath)
			os.Exit(1)
		}
	}

	if err := storage.SavePrivateKey(*outPath, clean); err != nil {
		fmt.Fprintf(os.Stderr, "error: save key: %v\n", err)
		os.Exit(1)
	}

	addr := ethcrypto.PubkeyToAddress(key.PublicKey)
	fmt.Printf("Private key saved to %s\n", *outPath)
	fmt.Printf("Address: %s\n", addr.Hex())
}

func runWalletBalance(args []string) {
	fs := flag.NewFlagSet("wallet balance", flag.ExitOnError)
	configPath := fs.String("config", "murid.toml", "path to config file")
	fs.Parse(args)

	_, client, ctx, cancel, err := loadClient(*configPath, 30*time.Second)
	if err != nil {
		fatal("%v", err)
	}
	defer cancel()
	defer client.Close()

	balance, err := client.GetBalance(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: get balance: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Address: %s\n", client.Address())
	fmt.Printf("Balance: %s MURI\n", formatWei(balance))
}
