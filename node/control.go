package node

import (
	"bufio"
	"context"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strings"

	"github.com/rs/zerolog/log"
)

// ControlSocketPath returns the default control socket path based on data dir.
func ControlSocketPath(dataDir string) string {
	return filepath.Join(dataDir, "murid.sock")
}

// startControlSocket launches a Unix domain socket listener that accepts
// pause/resume/status commands. Returns a cleanup function.
func (n *Node) startControlSocket(ctx context.Context) func() {
	sockPath := ControlSocketPath(n.cfg.Node.DataDir)

	// Clean up stale socket
	os.Remove(sockPath)

	ln, err := net.Listen("unix", sockPath)
	if err != nil {
		log.Warn().Err(err).Str("path", sockPath).Msg("control socket: failed to listen (CLI commands unavailable)")
		return func() {}
	}

	log.Info().Str("path", sockPath).Msg("control socket listening")

	// Accept connections in background
	go func() {
		for {
			conn, err := ln.Accept()
			if err != nil {
				if ctx.Err() != nil {
					return
				}
				log.Debug().Err(err).Msg("control socket: accept error")
				return
			}
			go n.handleControlConn(conn)
		}
	}()

	// Close listener on context cancel
	go func() {
		<-ctx.Done()
		ln.Close()
		os.Remove(sockPath)
	}()

	return func() {
		ln.Close()
		os.Remove(sockPath)
	}
}

func (n *Node) handleControlConn(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	if !scanner.Scan() {
		return
	}
	cmd := strings.TrimSpace(scanner.Text())

	var resp string
	switch cmd {
	case "pause":
		n.Pause()
		resp = "OK: paused\n"
	case "resume":
		n.Resume()
		resp = "OK: resumed\n"
	case "status":
		paused := "running"
		if n.IsPaused() {
			paused = "paused"
		}
		dereg := ""
		if n.deregistered.Load() {
			dereg = " [DEREGISTERED]"
		}
		resp = fmt.Sprintf("OK: %s%s\n", paused, dereg)
	default:
		resp = fmt.Sprintf("ERR: unknown command: %s\n", cmd)
	}
	conn.Write([]byte(resp))
}

// SendControlCommand connects to a running daemon's control socket and sends a command.
// Returns the daemon's response.
func SendControlCommand(dataDir, command string) (string, error) {
	sockPath := ControlSocketPath(dataDir)
	conn, err := net.Dial("unix", sockPath)
	if err != nil {
		return "", fmt.Errorf("connect to daemon (is it running?): %w", err)
	}
	defer conn.Close()

	fmt.Fprintf(conn, "%s\n", command)

	scanner := bufio.NewScanner(conn)
	if scanner.Scan() {
		return scanner.Text(), nil
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return "", fmt.Errorf("no response from daemon")
}
