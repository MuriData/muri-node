package storage

import (
	"math/big"
	"os"
	"path/filepath"
	"testing"

	muricrypto "github.com/MuriData/muri-zkproof/pkg/crypto"
	"github.com/MuriData/muri-zkproof/pkg/merkle"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr"
)

func testStore(t *testing.T) *Store {
	t.Helper()
	dir := t.TempDir()
	s, err := NewStore(dir)
	if err != nil {
		t.Fatalf("NewStore: %v", err)
	}
	return s
}

func testSMT(t *testing.T, numLeaves int) (*merkle.SparseMerkleTree, fr.Element) {
	t.Helper()
	zeroLeaf := muricrypto.ComputeZeroLeafHashFr(31, 528)
	hashes := make([]fr.Element, numLeaves)
	for i := range hashes {
		hashes[i].SetUint64(uint64(i + 1))
	}
	smt, err := merkle.BuildSMTFromLeafHashes(hashes, 20, zeroLeaf)
	if err != nil {
		t.Fatalf("BuildSMTFromLeafHashes: %v", err)
	}
	return smt, zeroLeaf
}

func TestSaveLoadTreeRoundTrip(t *testing.T) {
	s := testStore(t)
	smt, zeroLeaf := testSMT(t, 8)

	orderID := big.NewInt(42)
	if err := s.SaveTree(orderID, smt); err != nil {
		t.Fatalf("SaveTree: %v", err)
	}

	loaded, err := s.LoadTree(orderID, zeroLeaf)
	if err != nil {
		t.Fatalf("LoadTree: %v", err)
	}
	if loaded == nil {
		t.Fatal("LoadTree returned nil")
	}
	if loaded.RootBigInt().Cmp(smt.RootBigInt()) != 0 {
		t.Errorf("root mismatch: got %s, want %s", loaded.RootBigInt(), smt.RootBigInt())
	}
	if loaded.NumLeaves != smt.NumLeaves {
		t.Errorf("numLeaves mismatch: got %d, want %d", loaded.NumLeaves, smt.NumLeaves)
	}
}

func TestLoadTree_NotFound(t *testing.T) {
	s := testStore(t)
	var zeroLeaf fr.Element

	smt, err := s.LoadTree(big.NewInt(999), zeroLeaf)
	if err != nil {
		t.Fatalf("LoadTree should not error for missing file: %v", err)
	}
	if smt != nil {
		t.Fatal("LoadTree should return nil for missing file")
	}
}

func TestDeleteTree(t *testing.T) {
	s := testStore(t)
	smt, zeroLeaf := testSMT(t, 4)

	orderID := big.NewInt(10)
	if err := s.SaveTree(orderID, smt); err != nil {
		t.Fatalf("SaveTree: %v", err)
	}

	if err := s.DeleteTree(orderID); err != nil {
		t.Fatalf("DeleteTree: %v", err)
	}

	loaded, err := s.LoadTree(orderID, zeroLeaf)
	if err != nil {
		t.Fatalf("LoadTree after delete: %v", err)
	}
	if loaded != nil {
		t.Fatal("LoadTree should return nil after delete")
	}
}

func TestDeleteTree_NotFound(t *testing.T) {
	s := testStore(t)
	if err := s.DeleteTree(big.NewInt(999)); err != nil {
		t.Fatalf("DeleteTree should not error for missing file: %v", err)
	}
}

func TestListCachedOrderIDs(t *testing.T) {
	s := testStore(t)
	smt, _ := testSMT(t, 4)

	ids := []*big.Int{big.NewInt(1), big.NewInt(42), big.NewInt(100)}
	for _, id := range ids {
		if err := s.SaveTree(id, smt); err != nil {
			t.Fatalf("SaveTree(%s): %v", id, err)
		}
	}

	listed, err := s.ListCachedOrderIDs()
	if err != nil {
		t.Fatalf("ListCachedOrderIDs: %v", err)
	}
	if len(listed) != len(ids) {
		t.Fatalf("expected %d IDs, got %d", len(ids), len(listed))
	}

	listedSet := make(map[string]bool)
	for _, id := range listed {
		listedSet[id.String()] = true
	}
	for _, id := range ids {
		if !listedSet[id.String()] {
			t.Errorf("missing order ID %s", id)
		}
	}
}

func TestListCachedOrderIDs_Empty(t *testing.T) {
	s := testStore(t)
	listed, err := s.ListCachedOrderIDs()
	if err != nil {
		t.Fatalf("ListCachedOrderIDs: %v", err)
	}
	if len(listed) != 0 {
		t.Errorf("expected 0 IDs, got %d", len(listed))
	}
}

func TestListCachedOrderIDs_IgnoresNonSMTFiles(t *testing.T) {
	s := testStore(t)
	// Create a non-SMT file in the data directory
	f, err := os.Create(filepath.Join(s.dataDir, "some_other_file.txt"))
	if err != nil {
		t.Fatalf("create file: %v", err)
	}
	f.Close()

	listed, err := s.ListCachedOrderIDs()
	if err != nil {
		t.Fatalf("ListCachedOrderIDs: %v", err)
	}
	if len(listed) != 0 {
		t.Errorf("expected 0 IDs, got %d", len(listed))
	}
}

func TestSaveLoadOrderMapAtomic(t *testing.T) {
	s := testStore(t)

	orders := map[string]string{
		"100": "QmABC123",
		"200": "QmDEF456",
	}

	if err := s.SaveOrderMapAtomic(orders); err != nil {
		t.Fatalf("SaveOrderMapAtomic: %v", err)
	}

	loaded, err := s.LoadOrderMap()
	if err != nil {
		t.Fatalf("LoadOrderMap: %v", err)
	}
	if len(loaded) != len(orders) {
		t.Fatalf("expected %d entries, got %d", len(orders), len(loaded))
	}
	for k, v := range orders {
		if loaded[k] != v {
			t.Errorf("key %s: got %q, want %q", k, loaded[k], v)
		}
	}
}

func TestLoadOrderMap_NotFound(t *testing.T) {
	s := testStore(t)
	m, err := s.LoadOrderMap()
	if err != nil {
		t.Fatalf("LoadOrderMap: %v", err)
	}
	if len(m) != 0 {
		t.Errorf("expected empty map, got %d entries", len(m))
	}
}

func TestSaveOrderMapAtomic_Overwrite(t *testing.T) {
	s := testStore(t)

	v1 := map[string]string{"1": "cidA"}
	if err := s.SaveOrderMapAtomic(v1); err != nil {
		t.Fatalf("SaveOrderMapAtomic v1: %v", err)
	}

	v2 := map[string]string{"2": "cidB", "3": "cidC"}
	if err := s.SaveOrderMapAtomic(v2); err != nil {
		t.Fatalf("SaveOrderMapAtomic v2: %v", err)
	}

	loaded, err := s.LoadOrderMap()
	if err != nil {
		t.Fatalf("LoadOrderMap: %v", err)
	}
	if len(loaded) != 2 {
		t.Fatalf("expected 2 entries, got %d", len(loaded))
	}
	if _, ok := loaded["1"]; ok {
		t.Error("old key '1' should not be present after overwrite")
	}
}

func TestHasTree(t *testing.T) {
	s := testStore(t)
	smt, _ := testSMT(t, 4)

	orderID := big.NewInt(77)
	if s.HasTree(orderID) {
		t.Error("HasTree should return false before save")
	}

	if err := s.SaveTree(orderID, smt); err != nil {
		t.Fatalf("SaveTree: %v", err)
	}

	if !s.HasTree(orderID) {
		t.Error("HasTree should return true after save")
	}
}
