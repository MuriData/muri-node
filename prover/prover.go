package prover

import (
	"fmt"
	"math/big"
	"sync"

	"github.com/MuriData/muri-zkproof/circuits/poi"
	muricrypto "github.com/MuriData/muri-zkproof/pkg/crypto"
	"github.com/MuriData/muri-zkproof/pkg/merkle"
	"github.com/MuriData/muri-zkproof/pkg/setup"
	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	groth16bn254 "github.com/consensys/gnark/backend/groth16/bn254"
	"github.com/consensys/gnark/constraint"
	"github.com/consensys/gnark/frontend"
	"github.com/rs/zerolog/log"
)

// Prover generates PoI ZK proofs for challenge responses.
type Prover struct {
	mu  sync.Mutex
	ccs constraint.ConstraintSystem
	pk  groth16.ProvingKey
	vk  groth16.VerifyingKey

	zeroLeafHash *big.Int
}

// ProofResult holds the compressed proof output for on-chain submission.
type ProofResult struct {
	SolidityProof [4]*big.Int
	Commitment    [32]byte
}

// NewProver compiles the PoI circuit and loads prover/verifier keys.
func NewProver(keysDir string) (*Prover, error) {
	log.Info().Msg("compiling PoI circuit...")
	ccs, err := setup.CompileCircuit(&poi.PoICircuit{})
	if err != nil {
		return nil, fmt.Errorf("compile circuit: %w", err)
	}

	log.Info().Str("dir", keysDir).Msg("loading prover keys...")
	pk, vk, err := setup.LoadKeys(keysDir, "poi")
	if err != nil {
		return nil, fmt.Errorf("load keys: %w", err)
	}

	zeroLeaf := muricrypto.ComputeZeroLeafHash(poi.ElementSize, poi.NumChunks)

	log.Info().Int("constraints", ccs.GetNbConstraints()).Msg("prover initialized")
	return &Prover{
		ccs:          ccs,
		pk:           pk,
		vk:           vk,
		zeroLeafHash: zeroLeaf,
	}, nil
}

// DeriveLeafIndices extracts the 8 leaf indices that the PoI circuit will
// open for a given randomness and tree size. This allows pre-fetching only
// the needed chunks instead of downloading the entire file.
func DeriveLeafIndices(randomness *big.Int, numLeaves int) [poi.OpeningsCount]int {
	var indices [poi.OpeningsCount]int
	numLeavesBig := big.NewInt(int64(numLeaves))
	for k := 0; k < poi.OpeningsCount; k++ {
		bitOffset := k * poi.MaxTreeDepth
		var rawIndex int64
		for i := 0; i < poi.MaxTreeDepth; i++ {
			bit := randomness.Bit(bitOffset + i)
			rawIndex |= int64(bit) << i
		}
		leafIndexBig := new(big.Int).Mod(big.NewInt(rawIndex), numLeavesBig)
		indices[k] = int(leafIndexBig.Int64())
	}
	return indices
}

// GenerateProof builds an SMT from file data and generates a Groth16 proof.
// Returns the serialized proof points and commitment for on-chain submission.
func (p *Prover) GenerateProof(secretKey, randomness *big.Int, fileData []byte) (*ProofResult, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	chunks := merkle.SplitIntoChunks(fileData, poi.FileSize)
	smt, err := merkle.GenerateSparseMerkleTree(chunks, poi.MaxTreeDepth, poi.HashChunk, p.zeroLeafHash)
	if err != nil {
		return nil, fmt.Errorf("build SMT: %w", err)
	}

	log.Debug().
		Int("chunks", len(chunks)).
		Str("root", fmt.Sprintf("0x%x", smt.Root)).
		Msg("built merkle tree")

	result, err := poi.PrepareWitness(secretKey, randomness, chunks, smt)
	if err != nil {
		return nil, fmt.Errorf("prepare witness: %w", err)
	}

	return p.proveFromWitness(result)
}

// GenerateProofFromSMT generates a proof using a pre-built SMT and chunk data.
func (p *Prover) GenerateProofFromSMT(secretKey, randomness *big.Int, chunks [][]byte, smt *merkle.SparseMerkleTree) (*ProofResult, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	result, err := poi.PrepareWitness(secretKey, randomness, chunks, smt)
	if err != nil {
		return nil, fmt.Errorf("prepare witness: %w", err)
	}

	return p.proveFromWitness(result)
}

// proveFromWitness generates and verifies a Groth16 proof from a prepared witness.
// Caller must hold p.mu.
func (p *Prover) proveFromWitness(result *poi.WitnessResult) (*ProofResult, error) {
	witness, err := frontend.NewWitness(&result.Assignment, ecc.BN254.ScalarField())
	if err != nil {
		return nil, fmt.Errorf("create witness: %w", err)
	}

	proof, err := groth16.Prove(p.ccs, p.pk, witness)
	if err != nil {
		return nil, fmt.Errorf("prove: %w", err)
	}

	publicWitness, err := witness.Public()
	if err != nil {
		return nil, fmt.Errorf("extract public witness: %w", err)
	}
	if err := groth16.Verify(proof, p.vk, publicWitness); err != nil {
		return nil, fmt.Errorf("local verify failed: %w", err)
	}

	compressedProof, err := extractAndCompressProof(proof)
	if err != nil {
		return nil, fmt.Errorf("extract/compress proof: %w", err)
	}

	var commitment [32]byte
	commitBytes := result.Commitment.Bytes()
	copy(commitment[32-len(commitBytes):], commitBytes)

	return &ProofResult{
		SolidityProof: compressedProof,
		Commitment:    commitment,
	}, nil
}

// BuildSMT builds a sparse Merkle tree from file data.
func (p *Prover) BuildSMT(fileData []byte) (*merkle.SparseMerkleTree, [][]byte, error) {
	chunks := merkle.SplitIntoChunks(fileData, poi.FileSize)
	smt, err := merkle.GenerateSparseMerkleTree(chunks, poi.MaxTreeDepth, poi.HashChunk, p.zeroLeafHash)
	if err != nil {
		return nil, chunks, fmt.Errorf("build SMT: %w", err)
	}
	return smt, chunks, nil
}

// ZeroLeafHash returns the precomputed zero leaf hash for tree operations.
func (p *Prover) ZeroLeafHash() *big.Int {
	return p.zeroLeafHash
}

// PublicKeyFromSecret derives the public key from a ZK secret key.
func PublicKeyFromSecret(sk *big.Int) *big.Int {
	return muricrypto.DerivePublicKey(sk)
}

// extractAndCompressProof extracts a Groth16 BN254 proof, then compresses it
// from 8 uint256 (uncompressed) to 4 uint256 (compressed) matching the on-chain
// verifier's compressProof format.
func extractAndCompressProof(proof groth16.Proof) ([4]*big.Int, error) {
	bn254Proof, ok := proof.(*groth16bn254.Proof)
	if !ok {
		return [4]*big.Int{}, fmt.Errorf("expected bn254 proof, got %T", proof)
	}

	aX, aY := new(big.Int), new(big.Int)
	bn254Proof.Ar.X.BigInt(aX)
	bn254Proof.Ar.Y.BigInt(aY)

	bX0, bX1 := new(big.Int), new(big.Int)
	bY0, bY1 := new(big.Int), new(big.Int)
	bn254Proof.Bs.X.A0.BigInt(bX0)
	bn254Proof.Bs.X.A1.BigInt(bX1)
	bn254Proof.Bs.Y.A0.BigInt(bY0)
	bn254Proof.Bs.Y.A1.BigInt(bY1)

	cX, cY := new(big.Int), new(big.Int)
	bn254Proof.Krs.X.BigInt(cX)
	bn254Proof.Krs.Y.BigInt(cY)

	// Uncompressed format: [A.x, A.y, B.x1, B.x0, B.y1, B.y0, C.x, C.y]
	uncompressed := [8]*big.Int{aX, aY, bX1, bX0, bY1, bY0, cX, cY}

	return muricrypto.CompressProof(uncompressed)
}
