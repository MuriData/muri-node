package storage

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"strings"

	"github.com/MuriData/muri-zkproof/pkg/merkle"
)

// Store manages local persistence of Merkle trees per order.
type Store struct {
	dataDir string
}

// NewStore creates a store backed by the given directory.
func NewStore(dataDir string) (*Store, error) {
	if err := os.MkdirAll(dataDir, 0o755); err != nil {
		return nil, fmt.Errorf("create data dir: %w", err)
	}
	return &Store{dataDir: dataDir}, nil
}

// treePath returns the file path for an order's SMT cache.
func (s *Store) treePath(orderID *big.Int) string {
	return filepath.Join(s.dataDir, fmt.Sprintf("order_%s.smt", orderID.Text(10)))
}

// SaveTree persists an SMT to disk for the given order.
func (s *Store) SaveTree(orderID *big.Int, smt *merkle.SparseMerkleTree) error {
	f, err := os.Create(s.treePath(orderID))
	if err != nil {
		return fmt.Errorf("create tree file: %w", err)
	}
	defer f.Close()

	if err := smt.Save(f); err != nil {
		return fmt.Errorf("save tree: %w", err)
	}
	return nil
}

// LoadTree loads a cached SMT from disk. Returns nil if not found.
func (s *Store) LoadTree(orderID *big.Int, zeroLeafHash *big.Int) (*merkle.SparseMerkleTree, error) {
	f, err := os.Open(s.treePath(orderID))
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, fmt.Errorf("open tree file: %w", err)
	}
	defer f.Close()

	smt, err := merkle.LoadSparseMerkleTree(f, zeroLeafHash)
	if err != nil {
		return nil, fmt.Errorf("load tree: %w", err)
	}
	return smt, nil
}

// HasTree checks if a cached SMT exists for the order.
func (s *Store) HasTree(orderID *big.Int) bool {
	_, err := os.Stat(s.treePath(orderID))
	return err == nil
}

// DeleteTree removes the cached SMT for an order.
func (s *Store) DeleteTree(orderID *big.Int) error {
	err := os.Remove(s.treePath(orderID))
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("delete tree: %w", err)
	}
	return nil
}

// ListCachedOrderIDs scans the data directory for cached SMT files and returns their order IDs.
func (s *Store) ListCachedOrderIDs() ([]*big.Int, error) {
	entries, err := os.ReadDir(s.dataDir)
	if err != nil {
		return nil, fmt.Errorf("read data dir: %w", err)
	}

	var ids []*big.Int
	for _, e := range entries {
		name := e.Name()
		if !strings.HasPrefix(name, "order_") || !strings.HasSuffix(name, ".smt") {
			continue
		}
		idStr := strings.TrimPrefix(name, "order_")
		idStr = strings.TrimSuffix(idStr, ".smt")
		id, ok := new(big.Int).SetString(idStr, 10)
		if !ok {
			continue
		}
		ids = append(ids, id)
	}
	return ids, nil
}

// SaveSecretKey writes a secret key to a hex file.
func SaveSecretKey(path string, sk *big.Int) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		return fmt.Errorf("create key dir: %w", err)
	}
	hexStr := hex.EncodeToString(sk.Bytes())
	return os.WriteFile(path, []byte(hexStr), 0o600)
}

// LoadSecretKey reads a secret key from a hex file.
func LoadSecretKey(path string) (*big.Int, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read secret key: %w", err)
	}
	hexStr := strings.TrimSpace(string(data))
	b, err := hex.DecodeString(hexStr)
	if err != nil {
		return nil, fmt.Errorf("decode hex: %w", err)
	}
	return new(big.Int).SetBytes(b), nil
}

// SavePrivateKey writes a hex-encoded EVM private key to a file.
func SavePrivateKey(path string, keyHex string) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		return fmt.Errorf("create key dir: %w", err)
	}
	return os.WriteFile(path, []byte(keyHex), 0o600)
}

// LoadPrivateKey reads a hex-encoded EVM private key from a file.
func LoadPrivateKey(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("read private key: %w", err)
	}
	key := strings.TrimSpace(string(data))
	// Strip 0x prefix if present
	key = strings.TrimPrefix(key, "0x")
	return key, nil
}
