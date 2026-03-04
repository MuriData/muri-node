package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

const (
	defaultRepo   = "MuriData/muri-artifacts"
	defaultBranch = "main"
	defaultCircuit = "poi"
)

// artifact describes a single key file to download.
type artifact struct {
	name string // e.g. "poi_prover.key"
	desc string // human-readable description
}

func runDownloadKeys(args []string) {
	fs := flag.NewFlagSet("download-keys", flag.ExitOnError)
	outDir := fs.String("out", "./keys", "directory to save keys to")
	repo := fs.String("repo", defaultRepo, "GitHub org/repo for artifacts")
	branch := fs.String("branch", defaultBranch, "git branch or tag")
	circuit := fs.String("circuit", defaultCircuit, "circuit name (e.g. poi, fsp)")
	fs.Parse(args)

	artifacts := []artifact{
		{name: fmt.Sprintf("%s_prover.key", *circuit), desc: "proving key"},
		{name: fmt.Sprintf("%s_verifier.key", *circuit), desc: "verifying key"},
	}

	if err := os.MkdirAll(*outDir, 0o755); err != nil {
		fmt.Fprintf(os.Stderr, "error: create output dir: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Downloading %s keys from github.com/%s (branch: %s)\n", *circuit, *repo, *branch)
	fmt.Printf("Output directory: %s\n\n", *outDir)

	for _, a := range artifacts {
		dest := filepath.Join(*outDir, a.name)
		url := fmt.Sprintf("https://github.com/%s/raw/%s/%s/%s", *repo, *branch, *circuit, a.name)

		fmt.Printf("  %s (%s)\n", a.name, a.desc)
		fmt.Printf("    URL: %s\n", url)

		if err := downloadFile(url, dest); err != nil {
			fmt.Fprintf(os.Stderr, "    error: %v\n", err)
			os.Exit(1)
		}
	}

	fmt.Printf("\nDone. Set keys_dir = %q in your config.\n", *outDir)
}

// downloadFile fetches a URL to a local file with progress reporting.
func downloadFile(url, dest string) error {
	client := &http.Client{
		Timeout: 10 * time.Minute,
	}

	resp, err := client.Get(url)
	if err != nil {
		return fmt.Errorf("GET %s: %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("GET %s: status %d", url, resp.StatusCode)
	}

	// Write to temp file, rename on success for atomicity
	tmp := dest + ".tmp"
	f, err := os.Create(tmp)
	if err != nil {
		return fmt.Errorf("create %s: %w", tmp, err)
	}

	totalBytes := resp.ContentLength
	var written int64
	buf := make([]byte, 256*1024) // 256 KB buffer

	lastReport := time.Now()
	var writeErr error
	for {
		n, readErr := resp.Body.Read(buf)
		if n > 0 {
			nw, wErr := f.Write(buf[:n])
			if wErr != nil {
				writeErr = fmt.Errorf("write: %w", wErr)
				break
			}
			written += int64(nw)

			// Report progress at most once per second
			if time.Since(lastReport) > time.Second {
				printProgress(written, totalBytes)
				lastReport = time.Now()
			}
		}
		if readErr == io.EOF {
			break
		}
		if readErr != nil {
			writeErr = fmt.Errorf("read: %w", readErr)
			break
		}
	}

	f.Close()
	if writeErr != nil {
		os.Remove(tmp)
		return writeErr
	}

	if err := os.Rename(tmp, dest); err != nil {
		os.Remove(tmp)
		return fmt.Errorf("rename %s → %s: %w", tmp, dest, err)
	}

	printProgress(written, totalBytes)
	fmt.Println()

	return nil
}

func printProgress(written, total int64) {
	if total > 0 {
		pct := float64(written) / float64(total) * 100
		fmt.Printf("\r    %s / %s (%.0f%%)", formatBytes(written), formatBytes(total), pct)
	} else {
		fmt.Printf("\r    %s downloaded", formatBytes(written))
	}
}

func formatBytes(b int64) string {
	switch {
	case b >= 1<<30:
		return fmt.Sprintf("%.1f GB", float64(b)/(1<<30))
	case b >= 1<<20:
		return fmt.Sprintf("%.1f MB", float64(b)/(1<<20))
	case b >= 1<<10:
		return fmt.Sprintf("%.1f KB", float64(b)/(1<<10))
	default:
		return fmt.Sprintf("%d B", b)
	}
}
