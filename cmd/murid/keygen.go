package main

import (
	"flag"
	"fmt"
	"os"

	muricrypto "github.com/MuriData/muri-zkproof/pkg/crypto"
	"github.com/MuriData/muri-node/storage"
)

func runKeygen(args []string) {
	fs := flag.NewFlagSet("keygen", flag.ExitOnError)
	outPath := fs.String("out", "./keys/secret.key", "path to save the secret key")
	force := fs.Bool("force", false, "overwrite existing key file")
	fs.Parse(args)

	// Check if file already exists
	if !*force {
		if _, err := os.Stat(*outPath); err == nil {
			fmt.Fprintf(os.Stderr, "error: %s already exists (use -force to overwrite)\n", *outPath)
			os.Exit(1)
		}
	}

	// Generate secret key (random BN254 scalar)
	sk, err := muricrypto.GenerateSecretKey()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: generate secret key: %v\n", err)
		os.Exit(1)
	}

	// Derive public key
	pk := muricrypto.DerivePublicKey(sk)

	// Save to file
	if err := storage.SaveSecretKey(*outPath, sk); err != nil {
		fmt.Fprintf(os.Stderr, "error: save secret key: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("ZK secret key saved to %s\n", *outPath)
	fmt.Printf("Public key: 0x%s\n", pk.Text(16))
	fmt.Println("\nPaste the public key into the dashboard when registering your node.")
}
