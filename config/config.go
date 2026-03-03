package config

import (
	"fmt"
	"os"
	"time"

	"github.com/BurntSushi/toml"
)

// Config is the top-level configuration for murid.
type Config struct {
	Chain       ChainConfig       `toml:"chain"`
	IPFS        IPFSConfig        `toml:"ipfs"`
	Node        NodeConfig        `toml:"node"`
	Storage     StorageConfig     `toml:"storage"`
	AutoExecute AutoExecuteConfig `toml:"auto_execute"`
	Challenge   ChallengeConfig   `toml:"challenge"`
	Log         LogConfig         `toml:"log"`
}

// ChainConfig holds EVM chain connection settings.
type ChainConfig struct {
	RPCURL             string  `toml:"rpc_url"`
	ChainID            int64   `toml:"chain_id"`
	MarketAddress      string  `toml:"market_address"`
	StakingAddress     string  `toml:"staking_address"`
	GasLimit           uint64  `toml:"gas_limit"`
	MaxGasPrice        uint64  `toml:"max_gas_price"`        // gwei
	GasEscalation      float64 `toml:"gas_escalation"`       // multiplier per retry
	MaxRetries         int     `toml:"max_retries"`
	ConfirmationBlocks uint64  `toml:"confirmation_blocks"`
}

// IPFSConfig holds IPFS Kubo API settings.
type IPFSConfig struct {
	APIURL   string        `toml:"api_url"`
	Timeout  tomlDuration  `toml:"timeout"`
	PinFiles bool          `toml:"pin_files"`
}

// NodeConfig holds node identity settings.
type NodeConfig struct {
	PrivateKeyPath string `toml:"private_key_path"`
	DataDir        string `toml:"data_dir"`
	KeysDir        string `toml:"keys_dir"`
	SecretKeyPath  string `toml:"secret_key_path"`
}

// StorageConfig holds capacity/price thresholds.
type StorageConfig struct {
	MaxCapacity uint64 `toml:"max_capacity"`
	MinPrice    uint64 `toml:"min_price"` // wei per chunk per period
}

// AutoExecuteConfig controls automatic order filling.
type AutoExecuteConfig struct {
	Enabled        bool         `toml:"enabled"`
	PollInterval   tomlDuration `toml:"poll_interval"`
	MaxOrdersToFill int         `toml:"max_orders_to_fill"`
}

// ChallengeConfig controls challenge response behavior.
type ChallengeConfig struct {
	PollInterval tomlDuration `toml:"poll_interval"`
	GasPriority  uint64       `toml:"gas_priority"` // gwei tip
	SafetyMargin uint64       `toml:"safety_margin"` // blocks before deadline to respond
}

// LogConfig holds logging settings.
type LogConfig struct {
	Level  string `toml:"level"`
	Pretty bool   `toml:"pretty"`
}

// tomlDuration wraps time.Duration for TOML parsing.
type tomlDuration struct {
	time.Duration
}

func (d *tomlDuration) UnmarshalText(text []byte) error {
	var err error
	d.Duration, err = time.ParseDuration(string(text))
	return err
}

// Load reads and parses a TOML config file.
func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read config: %w", err)
	}
	cfg := DefaultConfig()
	if err := toml.Unmarshal(data, cfg); err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}
	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("validate config: %w", err)
	}
	return cfg, nil
}

// DefaultConfig returns a config with sensible defaults.
func DefaultConfig() *Config {
	return &Config{
		Chain: ChainConfig{
			GasLimit:           500_000,
			MaxGasPrice:        100,
			GasEscalation:      1.25,
			MaxRetries:         3,
			ConfirmationBlocks: 1,
		},
		IPFS: IPFSConfig{
			APIURL:   "http://127.0.0.1:5001",
			Timeout:  tomlDuration{30 * time.Second},
			PinFiles: true,
		},
		Node: NodeConfig{
			DataDir: "./data",
			KeysDir: "./keys",
		},
		AutoExecute: AutoExecuteConfig{
			Enabled:         false,
			PollInterval:    tomlDuration{30 * time.Second},
			MaxOrdersToFill: 5,
		},
		Challenge: ChallengeConfig{
			PollInterval: tomlDuration{4 * time.Second},
			GasPriority:  2,
			SafetyMargin: 10,
		},
		Log: LogConfig{
			Level:  "info",
			Pretty: true,
		},
	}
}

// Validate checks required fields and constraints.
func (c *Config) Validate() error {
	if c.Chain.RPCURL == "" {
		return fmt.Errorf("chain.rpc_url is required")
	}
	if c.Chain.ChainID == 0 {
		return fmt.Errorf("chain.chain_id is required")
	}
	if c.Chain.MarketAddress == "" {
		return fmt.Errorf("chain.market_address is required")
	}
	if c.Node.PrivateKeyPath == "" {
		return fmt.Errorf("node.private_key_path is required")
	}
	if c.Node.SecretKeyPath == "" {
		return fmt.Errorf("node.secret_key_path is required")
	}
	if c.Node.KeysDir == "" {
		return fmt.Errorf("node.keys_dir is required")
	}
	if c.Chain.GasEscalation < 1.0 {
		return fmt.Errorf("chain.gas_escalation must be >= 1.0")
	}
	if c.Challenge.SafetyMargin == 0 {
		return fmt.Errorf("challenge.safety_margin must be > 0")
	}
	return nil
}
