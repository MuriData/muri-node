package node

import (
	"context"
	"math/big"

	"github.com/MuriData/muri-node/prover"
	"github.com/MuriData/muri-node/types"
	"github.com/MuriData/muri-zkproof/pkg/merkle"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

// ChainService abstracts the on-chain client used by the node daemon.
// *chain.Client satisfies this interface.
type ChainService interface {
	// Identity
	Address() common.Address
	BlockNumber(ctx context.Context) (uint64, error)

	// Event support
	HasEventSupport() bool
	SubscribeChallenges(ctx context.Context) (<-chan types.ChallengeSlotInfo, error)
	SubscribeNewOrders(ctx context.Context) (<-chan *big.Int, error)

	// Slot queries
	GetAllSlotInfo(ctx context.Context) ([]types.ChallengeSlotInfo, error)
	GetAllSlotInfoFresh(ctx context.Context) ([]types.ChallengeSlotInfo, error)
	GetSlotInfoFresh(ctx context.Context, slotIndex int) (types.ChallengeSlotInfo, error)

	// Node queries
	GetNodeInfo(ctx context.Context) (*types.NodeInfo, error)
	GetNodeOrders(ctx context.Context) ([]*big.Int, error)
	IsValidNode(ctx context.Context) (bool, error)

	// Order queries
	GetOrderDetails(ctx context.Context, orderID *big.Int) (*types.OrderInfo, error)
	GetActiveOrdersPage(ctx context.Context, offset, limit uint64) ([]*big.Int, uint64, error)
	GetClaimableRewards(ctx context.Context) (*big.Int, error)

	// Transactions
	SubmitProof(ctx context.Context, slotIndex int, proof [4]*big.Int, commitment [32]byte) (*ethtypes.Receipt, error)
	ExecuteOrder(ctx context.Context, orderID *big.Int, proof [4]*big.Int, commitment [32]byte) (*ethtypes.Receipt, error)
	ClaimRewards(ctx context.Context) (*ethtypes.Receipt, error)
	ProcessExpiredSlots(ctx context.Context) (*ethtypes.Receipt, error)
	ActivateSlots(ctx context.Context) (*ethtypes.Receipt, error)
}

// IPFSService abstracts the IPFS client used by the node daemon.
// *ipfs.Client satisfies this interface.
type IPFSService interface {
	Ping(ctx context.Context) error
	CatChunkedTo(ctx context.Context, cid string, fileChunkSize int, fn func(index int, data []byte)) (int, error)
	CatRangeWithRetry(ctx context.Context, cid string, offset, length int64) ([]byte, error)
	BlockGetWithRetry(ctx context.Context, cid string) ([]byte, error)
	Pin(ctx context.Context, cid string) error
	Unpin(ctx context.Context, cid string) error
	IsPinned(ctx context.Context, cid string) (bool, error)
	ListPins(ctx context.Context) ([]string, error)
	Provide(ctx context.Context, cid string) error
}

// ProverService abstracts the ZK proof generator used by the node daemon.
// *prover.Prover satisfies this interface.
type ProverService interface {
	GenerateProofFromSMT(secretKey, randomness *big.Int, chunks [][]byte, smt *merkle.SparseMerkleTree) (*prover.ProofResult, error)
	ZeroLeafHash() fr.Element
}

// StorageService abstracts the local persistence layer used by the node daemon.
// *storage.Store satisfies this interface.
type StorageService interface {
	SaveTree(orderID *big.Int, smt *merkle.SparseMerkleTree) error
	LoadTree(orderID *big.Int, zeroLeafHash fr.Element) (*merkle.SparseMerkleTree, error)
	DeleteTree(orderID *big.Int) error
	ListCachedOrderIDs() ([]*big.Int, error)
	SaveOrderMapAtomic(orders map[string]string) error
	LoadOrderMap() (map[string]string, error)
}
