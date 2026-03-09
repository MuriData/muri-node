package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// OrderInfo mirrors the on-chain FileOrder struct.
// Price is fetched via the getOrderPrice view (pricePerChunkPerPeriod set at order placement).
type OrderInfo struct {
	ID          *big.Int
	Owner       common.Address
	URI         string
	RootHash    *big.Int
	NumChunks   uint32
	Periods     uint16
	Replicas    uint8
	Filled      uint8
	Price       *big.Int // from getOrderPrice: pricePerChunkPerPeriod
	Escrow      *big.Int
	StartPeriod uint32
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
