package node

import (
	"math/big"
	"testing"
)

func TestSnarkScalarFieldInit(t *testing.T) {
	if snarkScalarField == nil {
		t.Fatal("snarkScalarField not initialized")
	}
	if snarkScalarField.Sign() <= 0 {
		t.Fatal("snarkScalarField must be positive")
	}
}

func TestDeriveExecutionRandomness_Deterministic(t *testing.T) {
	root := big.NewInt(42)
	pk := big.NewInt(12345)

	r1 := deriveExecutionRandomness(root, pk)
	r2 := deriveExecutionRandomness(root, pk)

	if r1.Cmp(r2) != 0 {
		t.Errorf("same inputs produced different randomness: %s vs %s", r1, r2)
	}
}

func TestDeriveExecutionRandomness_DifferentInputs(t *testing.T) {
	root := big.NewInt(42)
	pk1 := big.NewInt(100)
	pk2 := big.NewInt(200)

	r1 := deriveExecutionRandomness(root, pk1)
	r2 := deriveExecutionRandomness(root, pk2)

	if r1.Cmp(r2) == 0 {
		t.Error("different public keys should produce different randomness")
	}
}

func TestDeriveExecutionRandomness_BoundedByScalarField(t *testing.T) {
	root := new(big.Int).SetBytes(make([]byte, 32)) // zero
	root.SetBit(root, 255, 1)                        // large value

	pk := new(big.Int).SetBytes(make([]byte, 32))
	pk.SetBit(pk, 254, 1) // large value

	r := deriveExecutionRandomness(root, pk)

	if r.Cmp(snarkScalarField) >= 0 {
		t.Errorf("randomness %s >= scalar field %s", r, snarkScalarField)
	}
	if r.Sign() < 0 {
		t.Error("randomness must be non-negative")
	}
}
