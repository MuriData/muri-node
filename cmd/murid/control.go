package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/MuriData/muri-node/node"
)

func runPause(args []string) {
	fs := flag.NewFlagSet("pause", flag.ExitOnError)
	dataDir := fs.String("data-dir", "./data", "daemon data directory (for control socket)")
	fs.Parse(args)

	resp, err := node.SendControlCommand(*dataDir, "pause")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(resp)
}

func runResume(args []string) {
	fs := flag.NewFlagSet("resume", flag.ExitOnError)
	dataDir := fs.String("data-dir", "./data", "daemon data directory (for control socket)")
	fs.Parse(args)

	resp, err := node.SendControlCommand(*dataDir, "resume")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(resp)
}
