package config

import (
	"os"
	"path/filepath"
	"testing"
)

// baseConfig returns a minimal valid config with optional extra chain fields appended
// inside the [chain] section.
func baseConfig(extraChainFields string) string {
	return `
[chain]
rpc_url = "http://127.0.0.1:9650/ext/bc/C/rpc"
chain_id = 43113
market_address = "0x1234567890abcdef1234567890abcdef12345678"
gas_priority = 2
` + extraChainFields + `

[node]
private_key_path = "./keys/node.key"
secret_key_path = "./keys/secret.key"
keys_dir = "./keys"

[challenge]
safety_margin = 10
`
}

func loadFromString(t *testing.T, tomlStr string) (*Config, error) {
	t.Helper()
	dir := t.TempDir()
	p := filepath.Join(dir, "test.toml")
	if err := os.WriteFile(p, []byte(tomlStr), 0o644); err != nil {
		t.Fatalf("write temp config: %v", err)
	}
	return Load(p)
}

func TestDefaultListenMode(t *testing.T) {
	cfg, err := loadFromString(t, baseConfig(""))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.Chain.ListenMode != "poll" {
		t.Errorf("expected default listen_mode=poll, got %q", cfg.Chain.ListenMode)
	}
}

func TestListenModePoll(t *testing.T) {
	cfg, err := loadFromString(t, baseConfig(`listen_mode = "poll"`))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.Chain.ListenMode != "poll" {
		t.Errorf("expected listen_mode=poll, got %q", cfg.Chain.ListenMode)
	}
}

func TestListenModeEvents(t *testing.T) {
	cfg, err := loadFromString(t, baseConfig(`
listen_mode = "events"
ws_url = "ws://127.0.0.1:9650/ext/bc/C/ws"
`))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.Chain.ListenMode != "events" {
		t.Errorf("expected listen_mode=events, got %q", cfg.Chain.ListenMode)
	}
	if cfg.Chain.WSURL != "ws://127.0.0.1:9650/ext/bc/C/ws" {
		t.Errorf("unexpected ws_url: %q", cfg.Chain.WSURL)
	}
}

func TestListenModeEventsWSSURL(t *testing.T) {
	_, err := loadFromString(t, baseConfig(`
listen_mode = "events"
ws_url = "wss://avax-node.example.com/ws"
`))
	if err != nil {
		t.Fatalf("unexpected error for wss URL: %v", err)
	}
}

func TestListenModeInvalid(t *testing.T) {
	_, err := loadFromString(t, baseConfig(`listen_mode = "websocket"`))
	if err == nil {
		t.Fatal("expected error for invalid listen_mode")
	}
}

func TestListenModeEventsMissingWSURL(t *testing.T) {
	_, err := loadFromString(t, baseConfig(`listen_mode = "events"`))
	if err == nil {
		t.Fatal("expected error when listen_mode=events without ws_url")
	}
}

func TestListenModeEventsInvalidWSURL(t *testing.T) {
	_, err := loadFromString(t, baseConfig(`
listen_mode = "events"
ws_url = "http://127.0.0.1:9650/ext/bc/C/ws"
`))
	if err == nil {
		t.Fatal("expected error for http:// ws_url")
	}
}

func TestListenModePollWithWSURL(t *testing.T) {
	// ws_url is allowed in poll mode (ignored, no error)
	_, err := loadFromString(t, baseConfig(`
listen_mode = "poll"
ws_url = "ws://127.0.0.1:9650/ext/bc/C/ws"
`))
	if err != nil {
		t.Fatalf("ws_url in poll mode should be allowed: %v", err)
	}
}
