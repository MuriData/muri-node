package node

import (
	"math/big"

	"github.com/ethereum/go-ethereum/crypto"
)

// snarkScalarField is the BN254 scalar field modulus.
var snarkScalarField *big.Int

func init() {
	var ok bool
	snarkScalarField, ok = new(big.Int).SetString("21888242871839275222246405745257275088548364400416034343698204186575808495617", 10)
	if !ok {
		panic("failed to parse BN254 scalar field modulus")
	}
}

// deriveExecutionRandomness computes deterministic randomness for executeOrder
// PoI proof as keccak256(fileRoot, publicKey) % SNARK_SCALAR_FIELD.
func deriveExecutionRandomness(fileRoot, publicKey *big.Int) *big.Int {
	rootBytes := make([]byte, 32)
	pkBytes := make([]byte, 32)
	fileRoot.FillBytes(rootBytes)
	publicKey.FillBytes(pkBytes)
	hash := crypto.Keccak256(append(rootBytes, pkBytes...))
	r := new(big.Int).SetBytes(hash)
	r.Mod(r, snarkScalarField)
	return r
}
