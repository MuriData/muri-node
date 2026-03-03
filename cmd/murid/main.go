package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/MuriData/muri-node/config"
	"github.com/MuriData/muri-node/node"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	configPath := flag.String("config", "murid.toml", "path to config file")
	flag.Parse()

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
