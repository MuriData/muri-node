package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// OrderInfo mirrors the on-chain FileOrder struct.
type OrderInfo struct {
	ID        *big.Int
	Owner     common.Address
	URI       string
	RootHash  *big.Int
	NumChunks uint32
	Periods   uint16
	Replicas  uint8
	Filled    uint8
	Price     *big.Int
	Escrow    *big.Int
	StartPeriod uint64
}

// ChallengeSlotInfo mirrors the on-chain ChallengeSlot struct.
type ChallengeSlotInfo struct {
	Index          int
	OrderID        *big.Int
	ChallengedNode common.Address
	Randomness     *big.Int
	DeadlineBlock  *big.Int
	IsExpired      bool
}

// NodeInfo mirrors the on-chain NodeInfo struct from NodeStaking.
type NodeInfo struct {
	Stake     *big.Int
	Capacity  uint64
	Used      uint64
	PublicKey *big.Int
}

// ChallengeTask represents a pending challenge the node must respond to.
type ChallengeTask struct {
	SlotIndex  int
	OrderID    *big.Int
	Randomness *big.Int
	Deadline   *big.Int
}

// ProofResult holds the output from a ZK proof generation.
type ProofResult struct {
	SolidityProof [8]*big.Int
	Commitment    [32]byte
}
