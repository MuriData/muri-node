package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/MuriData/muri-node/config"
	"github.com/MuriData/muri-node/node"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const usage = `Usage: murid <command> [flags]

Commands:
  init                Interactive setup — generate keys, write config
  keygen              Generate a ZK secret key and derive its public key
  wallet              Wallet management (new, import, balance)
  stake               Register node and stake MURI on-chain
  unstake             Unregister node and withdraw stake
  increase-capacity   Add storage capacity (requires additional stake)
  decrease-capacity   Remove storage capacity (unlocks stake)
  status              Show node registration and on-chain status
  run                 Start the storage provider daemon (default)
  pause               Stop accepting new orders (daemon keeps responding to challenges)
  resume              Resume accepting new orders
  download-keys       Download PoI prover/verifier keys from GitHub
  web publish         Upload a directory to IPFS and place orders for all files
  pin file            Pin an existing IPFS file to MuriData
  pin dir             Pin an existing IPFS directory to MuriData

Flags:
  -config string   Path to config file (default "murid.toml")

Run "murid <command> -h" for command-specific help.
`

func main() {
	if len(os.Args) < 2 {
		runDaemon(os.Args[1:])
		return
	}

	switch os.Args[1] {
	case "init":
		runInit(os.Args[2:])
	case "keygen":
		runKeygen(os.Args[2:])
	case "wallet":
		runWallet(os.Args[2:])
	case "stake":
		runStake(os.Args[2:])
	case "unstake":
		runUnstake(os.Args[2:])
	case "increase-capacity":
		runIncreaseCapacity(os.Args[2:])
	case "decrease-capacity":
		runDecreaseCapacity(os.Args[2:])
	case "status":
		runStatus(os.Args[2:])
	case "run":
		runDaemon(os.Args[2:])
	case "pause":
		runPause(os.Args[2:])
	case "resume":
		runResume(os.Args[2:])
	case "download-keys":
		runDownloadKeys(os.Args[2:])
	case "pin":
		runPin(os.Args[2:])
	case "web":
		if len(os.Args) < 3 {
			fmt.Fprintf(os.Stderr, "Usage: murid web <subcommand>\n\nSubcommands:\n  publish   Upload a directory to IPFS and place orders\n")
			os.Exit(1)
		}
		switch os.Args[2] {
		case "publish":
			runWebPublish(os.Args[3:])
		default:
			fmt.Fprintf(os.Stderr, "unknown web subcommand: %s\n", os.Args[2])
			os.Exit(1)
		}
	case "-h", "--help", "help":
		fmt.Print(usage)
		os.Exit(0)
	default:
		// If first arg looks like a flag, treat as "run" with all args
		if len(os.Args[1]) > 0 && os.Args[1][0] == '-' {
			runDaemon(os.Args[1:])
			return
		}
		fmt.Fprintf(os.Stderr, "unknown command: %s\n\n%s", os.Args[1], usage)
		os.Exit(1)
	}
}

func runDaemon(args []string) {
	fs := flag.NewFlagSet("run", flag.ExitOnError)
	configPath := fs.String("config", "murid.toml", "path to config file")
	fs.Parse(args)

	// Load config
	cfg, err := config.Load(*configPath)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to load config")
	}

	// Setup logging
	level, err := zerolog.ParseLevel(cfg.Log.Level)
	if err != nil {
		level = zerolog.InfoLevel
	}
	zerolog.SetGlobalLevel(level)
	if cfg.Log.Pretty {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	log.Info().Str("config", *configPath).Msg("starting murid")

	// Context with signal-driven cancellation
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigCh
		log.Info().Str("signal", sig.String()).Msg("shutting down")
		cancel()
	}()

	// Create and run node
	n, err := node.New(ctx, cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to initialize node")
	}

	if err := n.Run(ctx); err != nil {
		log.Fatal().Err(err).Msg("node exited with error")
	}

	log.Info().Msg("murid stopped")
}
