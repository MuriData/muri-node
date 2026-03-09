package main

import (
	"bufio"
	"crypto/ecdsa"
	"encoding/hex"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	muricrypto "github.com/MuriData/muri-zkproof/pkg/crypto"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"

	"github.com/MuriData/muri-node/storage"
)

var scanner *bufio.Scanner

// prompt prints a label with a default value and reads user input.
// Returns the default if the user presses Enter without typing.
func prompt(label, defaultVal string) string {
	if defaultVal != "" {
		fmt.Printf("  %s [%s]: ", label, defaultVal)
	} else {
		fmt.Printf("  %s: ", label)
	}
	scanner.Scan()
	val := strings.TrimSpace(scanner.Text())
	if val == "" {
		return defaultVal
	}
	return val
}

// promptYN asks a yes/no question. Default is indicated by capitalization.
func promptYN(label string, defaultYes bool) bool {
	suffix := "[Y/n]"
	if !defaultYes {
		suffix = "[y/N]"
	}
	fmt.Printf("  %s %s: ", label, suffix)
	scanner.Scan()
	val := strings.ToLower(strings.TrimSpace(scanner.Text()))
	if val == "" {
		return defaultYes
	}
	return val == "y" || val == "yes"
}

// envDefault returns the environment variable value if set, otherwise the fallback.
func envDefault(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func runInit(args []string) {
	fs := flag.NewFlagSet("init", flag.ExitOnError)
	outPath := fs.String("config", "murid.toml", "config file to write")
	fs.Parse(args)

	scanner = bufio.NewScanner(os.Stdin)

	// Check if config already exists
	if _, err := os.Stat(*outPath); err == nil {
		if !promptYN(fmt.Sprintf("%s already exists. Overwrite?", *outPath), false) {
			fmt.Println("Aborted.")
			return
		}
	}

	fmt.Println()
	fmt.Println("Welcome to MuriData Node Setup!")
	fmt.Println("Press Enter to accept defaults shown in [brackets].")
	fmt.Println()

	// ── Chain ──
	fmt.Println("── Chain Configuration ──")
	rpcURL := prompt("RPC URL", envDefault("MURID_RPC_URL", "https://testnet-rpc.muri.moe/ext/bc/inP2vNhcVSABGmq39UHwuB9tDxUUWp3g6gpRwdE6TqtAtAWmu/rpc"))
	chainIDStr := prompt("Chain ID", envDefault("MURID_CHAIN_ID", "97981"))
	chainID, _ := strconv.ParseInt(chainIDStr, 10, 64)
	if chainID == 0 {
		chainID = 97981
	}
	marketAddr := prompt("FileMarket contract address", envDefault("MURID_MARKET_ADDRESS", "0xaab9f94671d6b22eee60509b5c3149e90a78fb54"))
	listenMode := prompt("Listen mode (poll/events)", "poll")
	wsURL := ""
	if listenMode == "events" {
		wsURL = prompt("WebSocket URL", "ws://127.0.0.1:9650/ext/bc/inP2vNhcVSABGmq39UHwuB9tDxUUWp3g6gpRwdE6TqtAtAWmu/ws")
	}
	fmt.Println()

	// ── Node Identity ──
	fmt.Println("── Node Identity ──")
	keysDir := prompt("Keys directory", envDefault("MURID_KEYS_DIR", "./keys"))
	dataDir := prompt("Data directory", envDefault("MURID_DATA_DIR", "./data"))

	// EVM private key
	privKeyPath := filepath.Join(keysDir, "node.key")
	var evmAddr string
	if promptYN("Generate new EVM private key?", true) {
		key, err := ethcrypto.GenerateKey()
		if err != nil {
			fmt.Fprintf(os.Stderr, "  error: %v\n", err)
			os.Exit(1)
		}
		if err := saveEVMKey(privKeyPath, key); err != nil {
			fmt.Fprintf(os.Stderr, "  error: %v\n", err)
			os.Exit(1)
		}
		evmAddr = ethcrypto.PubkeyToAddress(key.PublicKey).Hex()
		fmt.Printf("  ✓ EVM private key saved to %s\n", privKeyPath)
		fmt.Printf("  ✓ Node address: %s\n", evmAddr)
	} else {
		privKeyPath = prompt("Path to existing EVM private key", privKeyPath)
	}

	// ZK secret key
	secretKeyPath := filepath.Join(keysDir, "secret.key")
	if promptYN("Generate new ZK secret key?", true) {
		sk, err := muricrypto.GenerateSecretKey()
		if err != nil {
			fmt.Fprintf(os.Stderr, "  error: %v\n", err)
			os.Exit(1)
		}
		pk := muricrypto.DerivePublicKey(sk)
		if err := storage.SaveSecretKey(secretKeyPath, sk); err != nil {
			fmt.Fprintf(os.Stderr, "  error: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("  ✓ ZK secret key saved to %s\n", secretKeyPath)
		fmt.Printf("  ✓ Public key (for dashboard): 0x%s\n", pk.Text(16))
	} else {
		secretKeyPath = prompt("Path to existing ZK secret key", secretKeyPath)
	}
	fmt.Println()

	// ── IPFS ──
	fmt.Println("── IPFS ──")
	ipfsURL := prompt("Kubo API URL", envDefault("MURID_IPFS_URL", "http://127.0.0.1:5001"))
	pinFiles := promptYN("Pin files after download?", true)
	fmt.Println()

	// ── Storage ──
	fmt.Println("── Storage ──")
	maxCapGBStr := prompt("Max capacity (GB, 0=unlimited)", "10")
	maxCapGB, _ := strconv.ParseFloat(maxCapGBStr, 64)
	if maxCapGB < 0 {
		maxCapGB = 0
	}
	minPriceStr := prompt("Min price (wei/chunk/period)", "1000")
	minPrice, _ := strconv.ParseUint(minPriceStr, 10, 64)
	fmt.Println()

	// ── Auto Execute ──
	fmt.Println("── Order Execution ──")
	autoExec := promptYN("Auto-execute orders?", true)
	fmt.Println()

	// ── Download Keys ──
	downloadKeys := promptYN("Download PoI prover keys now? (~200 MB)", true)
	fmt.Println()

	// ── Write config ──
	cfg := buildTOML(tomlParams{
		rpcURL:        rpcURL,
		chainID:       chainID,
		marketAddr:    marketAddr,
		listenMode:    listenMode,
		wsURL:         wsURL,
		ipfsURL:       ipfsURL,
		pinFiles:      pinFiles,
		privKeyPath:   privKeyPath,
		dataDir:       dataDir,
		keysDir:       keysDir,
		secretKeyPath: secretKeyPath,
		maxCapGB:      maxCapGB,
		minPrice:      minPrice,
		autoExec:      autoExec,
	})

	if err := os.WriteFile(*outPath, []byte(cfg), 0o644); err != nil {
		fmt.Fprintf(os.Stderr, "error: write config: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("✓ Config written to %s\n", *outPath)

	if evmAddr != "" {
		fmt.Printf("✓ Fund %s with MURI before staking\n", evmAddr)
	}

	// Download keys if requested
	if downloadKeys {
		fmt.Println()
		runDownloadKeys([]string{"-out", keysDir})
	}

	fmt.Println()
	fmt.Println("Next steps:")
	fmt.Printf("  1. Send MURI to %s\n", evmAddr)
	fmt.Printf("  2. murid stake -capacity-gb 10      # register and stake\n")
	fmt.Printf("  3. murid run -config %s             # start the daemon\n", *outPath)
}

// saveEVMKey writes an ECDSA private key to a hex file.
func saveEVMKey(path string, key *ecdsa.PrivateKey) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		return fmt.Errorf("create key dir: %w", err)
	}
	hexStr := hex.EncodeToString(ethcrypto.FromECDSA(key))
	return os.WriteFile(path, []byte(hexStr), 0o600)
}

type tomlParams struct {
	rpcURL, marketAddr, listenMode, wsURL    string
	ipfsURL, privKeyPath, dataDir, keysDir   string
	secretKeyPath                            string
	chainID                                  int64
	maxCapGB                                 float64
	minPrice                                 uint64
	pinFiles, autoExec                       bool
}

func buildTOML(p tomlParams) string {
	var b strings.Builder
	b.WriteString("# murid — generated by murid init\n\n")

	b.WriteString("[chain]\n")
	fmt.Fprintf(&b, "rpc_url = %q\n", p.rpcURL)
	fmt.Fprintf(&b, "chain_id = %d\n", p.chainID)
	fmt.Fprintf(&b, "market_address = %q\n", p.marketAddr)
	b.WriteString("staking_address = \"\"   # auto-resolved from market contract\n")
	b.WriteString("gas_limit = 2000000\n")
	b.WriteString("max_gas_price = 100\n")
	b.WriteString("gas_priority = 2\n")
	b.WriteString("gas_escalation = 1.25\n")
	b.WriteString("max_retries = 3\n")
	b.WriteString("confirmation_blocks = 1\n")
	fmt.Fprintf(&b, "listen_mode = %q\n", p.listenMode)
	if p.wsURL != "" {
		fmt.Fprintf(&b, "ws_url = %q\n", p.wsURL)
	} else {
		b.WriteString("ws_url = \"\"\n")
	}

	b.WriteString("\n[ipfs]\n")
	fmt.Fprintf(&b, "api_url = %q\n", p.ipfsURL)
	b.WriteString("timeout = \"30s\"\n")
	fmt.Fprintf(&b, "pin_files = %t\n", p.pinFiles)

	b.WriteString("\n[node]\n")
	fmt.Fprintf(&b, "private_key_path = %q\n", p.privKeyPath)
	fmt.Fprintf(&b, "data_dir = %q\n", p.dataDir)
	fmt.Fprintf(&b, "keys_dir = %q\n", p.keysDir)
	fmt.Fprintf(&b, "secret_key_path = %q\n", p.secretKeyPath)

	b.WriteString("\n[storage]\n")
	fmt.Fprintf(&b, "max_capacity_gb = %g\n", p.maxCapGB)
	fmt.Fprintf(&b, "min_price = %d\n", p.minPrice)

	b.WriteString("\n[auto_execute]\n")
	fmt.Fprintf(&b, "enabled = %t\n", p.autoExec)
	b.WriteString("poll_interval = \"30s\"\n")
	b.WriteString("max_orders_to_fill = 5\n")

	b.WriteString("\n[challenge]\n")
	b.WriteString("poll_interval = \"4s\"\n")
	b.WriteString("safety_margin = 10\n")

	b.WriteString("\n[log]\n")
	b.WriteString("level = \"info\"\n")
	b.WriteString("pretty = true\n")

	return b.String()
}
