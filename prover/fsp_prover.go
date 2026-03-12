package prover

import (
	"fmt"
	"math/big"
	"sync"

	"github.com/MuriData/muri-zkproof/circuits/fsp"
	muricrypto "github.com/MuriData/muri-zkproof/pkg/crypto"
	"github.com/MuriData/muri-zkproof/pkg/merkle"
	"github.com/MuriData/muri-zkproof/pkg/setup"
	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr"
	"github.com/consensys/gnark/backend/groth16"
	groth16bn254 "github.com/consensys/gnark/backend/groth16/bn254"
	"github.com/consensys/gnark/constraint"
	"github.com/consensys/gnark/frontend"
	"github.com/rs/zerolog/log"
)

// FSPProver generates File-Size Proof (FSP) ZK proofs for placeOrder.
type FSPProver struct {
	mu  sync.Mutex
	ccs constraint.ConstraintSystem
	pk  groth16.ProvingKey
	vk  groth16.VerifyingKey

	zeroLeafHash fr.Element
}

// FSPProofResult holds the compressed FSP proof and public inputs for placeOrder.
type FSPProofResult struct {
	CompressedProof [4]*big.Int
	RootHash        *big.Int
	NumChunks       uint32
}

// NewFSPProver compiles the FSP circuit and loads prover/verifier keys.
func NewFSPProver(keysDir string) (*FSPProver, error) {
	log.Info().Msg("compiling FSP circuit...")
	ccs, err := setup.CompileCircuit(&fsp.FSPCircuit{})
	if err != nil {
		return nil, fmt.Errorf("compile FSP circuit: %w", err)
	}

	log.Info().Str("dir", keysDir).Msg("loading FSP prover keys...")
	pk, vk, err := setup.LoadKeys(keysDir, "fsp")
	if err != nil {
		return nil, fmt.Errorf("load FSP keys: %w", err)
	}

	zeroLeaf := muricrypto.ComputeZeroLeafHashFr(fsp.ElementSize, fsp.NumChunks)

	log.Info().Int("constraints", ccs.GetNbConstraints()).Msg("FSP prover initialized")
	return &FSPProver{
		ccs:          ccs,
		pk:           pk,
		vk:           vk,
		zeroLeafHash: zeroLeaf,
	}, nil
}

// GenerateProof builds an SMT from file data and generates an FSP Groth16 proof.
func (p *FSPProver) GenerateProof(fileData []byte) (*FSPProofResult, error) {
	chunks := merkle.SplitIntoChunks(fileData, fsp.FileSize)
	smt, err := merkle.GenerateSparseMerkleTree(chunks, fsp.MaxTreeDepth, fsp.HashChunk, p.zeroLeafHash)
	if err != nil {
		return nil, fmt.Errorf("build SMT: %w", err)
	}

	return p.GenerateProofFromSMT(smt)
}

// GenerateProofFromSMT generates an FSP proof from a pre-built SMT.
func (p *FSPProver) GenerateProofFromSMT(smt *merkle.SparseMerkleTree) (*FSPProofResult, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	result, err := fsp.PrepareWitness(smt)
	if err != nil {
		return nil, fmt.Errorf("prepare FSP witness: %w", err)
	}

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

	compressed, err := extractAndCompressFSPProof(proof)
	if err != nil {
		return nil, fmt.Errorf("extract/compress proof: %w", err)
	}

	rootBig := smt.RootBigInt()

	return &FSPProofResult{
		CompressedProof: compressed,
		RootHash:        rootBig,
		NumChunks:       uint32(result.NumLeaves),
	}, nil
}

// extractAndCompressFSPProof extracts proof points and compresses to [4]*big.Int.
func extractAndCompressFSPProof(proof groth16.Proof) ([4]*big.Int, error) {
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

	uncompressed := [8]*big.Int{aX, aY, bX1, bX0, bY1, bY0, cX, cY}
	return muricrypto.CompressProof(uncompressed)
}
