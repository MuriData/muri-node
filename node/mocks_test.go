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

// mockChain is a test double for ChainService.
type mockChain struct {
	addr common.Address

	eventSupport bool

	// Configurable return values
	blockNum       uint64
	blockNumErr    error
	slotInfo       []types.ChallengeSlotInfo
	slotInfoErr    error
	nodeInfo       *types.NodeInfo
	nodeInfoErr    error
	nodeOrders     []*big.Int
	nodeOrdersErr  error
	orderDetails   map[string]*types.OrderInfo
	orderDetailErr error
	isValid        bool
	isValidErr     error
	claimable      *big.Int
	claimableErr   error

	// Transaction receipts
	submitProofReceipt *ethtypes.Receipt
	submitProofErr     error
	executeReceipt     *ethtypes.Receipt
	executeErr         error
	claimReceipt       *ethtypes.Receipt
	claimErr           error
	expiredReceipt     *ethtypes.Receipt
	expiredErr         error
	activateReceipt    *ethtypes.Receipt
	activateErr        error

	// Active orders page
	activeOrders    []*big.Int
	activeTotal     uint64
	activeOrdersErr error

	// Subscription channels (set by test)
	challengeCh <-chan types.ChallengeSlotInfo
	orderCh     <-chan *big.Int
}

func (m *mockChain) Address() common.Address                    { return m.addr }
func (m *mockChain) BlockNumber(_ context.Context) (uint64, error) { return m.blockNum, m.blockNumErr }
func (m *mockChain) HasEventSupport() bool                      { return m.eventSupport }

func (m *mockChain) SubscribeChallenges(_ context.Context) (<-chan types.ChallengeSlotInfo, error) {
	return m.challengeCh, nil
}
func (m *mockChain) SubscribeNewOrders(_ context.Context) (<-chan *big.Int, error) {
	return m.orderCh, nil
}

func (m *mockChain) GetAllSlotInfo(_ context.Context) ([]types.ChallengeSlotInfo, error) {
	return m.slotInfo, m.slotInfoErr
}
func (m *mockChain) GetAllSlotInfoFresh(_ context.Context) ([]types.ChallengeSlotInfo, error) {
	return m.slotInfo, m.slotInfoErr
}
func (m *mockChain) GetSlotInfoFresh(_ context.Context, idx int) (types.ChallengeSlotInfo, error) {
	for _, s := range m.slotInfo {
		if s.Index == idx {
			return s, nil
		}
	}
	return types.ChallengeSlotInfo{}, m.slotInfoErr
}
func (m *mockChain) GetNodeInfo(_ context.Context) (*types.NodeInfo, error) {
	return m.nodeInfo, m.nodeInfoErr
}
func (m *mockChain) GetNodeOrders(_ context.Context) ([]*big.Int, error) {
	return m.nodeOrders, m.nodeOrdersErr
}
func (m *mockChain) IsValidNode(_ context.Context) (bool, error) {
	return m.isValid, m.isValidErr
}
func (m *mockChain) GetOrderDetails(_ context.Context, orderID *big.Int) (*types.OrderInfo, error) {
	if m.orderDetails != nil {
		if d, ok := m.orderDetails[orderID.String()]; ok {
			return d, nil
		}
	}
	return nil, m.orderDetailErr
}
func (m *mockChain) GetActiveOrdersPage(_ context.Context, _, _ uint64) ([]*big.Int, uint64, error) {
	return m.activeOrders, m.activeTotal, m.activeOrdersErr
}
func (m *mockChain) GetClaimableRewards(_ context.Context) (*big.Int, error) {
	return m.claimable, m.claimableErr
}

func (m *mockChain) SubmitProof(_ context.Context, _ int, _ [4]*big.Int, _ [32]byte) (*ethtypes.Receipt, error) {
	return m.submitProofReceipt, m.submitProofErr
}
func (m *mockChain) ExecuteOrder(_ context.Context, _ *big.Int, _ [4]*big.Int, _ [32]byte) (*ethtypes.Receipt, error) {
	return m.executeReceipt, m.executeErr
}
func (m *mockChain) ClaimRewards(_ context.Context) (*ethtypes.Receipt, error) {
	return m.claimReceipt, m.claimErr
}
func (m *mockChain) ProcessExpiredSlots(_ context.Context) (*ethtypes.Receipt, error) {
	return m.expiredReceipt, m.expiredErr
}
func (m *mockChain) ActivateSlots(_ context.Context) (*ethtypes.Receipt, error) {
	return m.activateReceipt, m.activateErr
}

// mockIPFS is a test double for IPFSService.
type mockIPFS struct {
	pinErr     error
	unpinErr   error
	isPinned   bool
	isPinnedErr error
	pins       []string
	pinsErr    error
	provideErr error
	pingErr    error

	catChunkedData  []byte
	catChunkedErr   error
	catRangeData    []byte
	catRangeErr     error
	blockGetData    []byte
	blockGetErr     error
}

func (m *mockIPFS) Ping(_ context.Context) error { return m.pingErr }
func (m *mockIPFS) CatChunkedTo(_ context.Context, _ string, chunkSize int, fn func(int, []byte)) (int, error) {
	if m.catChunkedErr != nil {
		return 0, m.catChunkedErr
	}
	idx := 0
	for off := 0; off < len(m.catChunkedData); off += chunkSize {
		end := off + chunkSize
		if end > len(m.catChunkedData) {
			end = len(m.catChunkedData)
		}
		fn(idx, m.catChunkedData[off:end])
		idx++
	}
	return idx, nil
}
func (m *mockIPFS) CatRangeWithRetry(_ context.Context, _ string, _, _ int64) ([]byte, error) {
	return m.catRangeData, m.catRangeErr
}
func (m *mockIPFS) BlockGetWithRetry(_ context.Context, _ string) ([]byte, error) {
	return m.blockGetData, m.blockGetErr
}
func (m *mockIPFS) Pin(_ context.Context, _ string) error          { return m.pinErr }
func (m *mockIPFS) Unpin(_ context.Context, _ string) error        { return m.unpinErr }
func (m *mockIPFS) IsPinned(_ context.Context, _ string) (bool, error) { return m.isPinned, m.isPinnedErr }
func (m *mockIPFS) ListPins(_ context.Context) ([]string, error)   { return m.pins, m.pinsErr }
func (m *mockIPFS) Provide(_ context.Context, _ string) error      { return m.provideErr }

// mockProver is a test double for ProverService.
type mockProver struct {
	proofResult *prover.ProofResult
	proofErr    error
	zeroLeaf    fr.Element
}

func (m *mockProver) GenerateProofFromSMT(_, _ *big.Int, _ [][]byte, _ *merkle.SparseMerkleTree) (*prover.ProofResult, error) {
	return m.proofResult, m.proofErr
}
func (m *mockProver) ZeroLeafHash() fr.Element { return m.zeroLeaf }

// mockStore is a test double for StorageService.
type mockStore struct {
	trees    map[string]*merkle.SparseMerkleTree
	orderMap map[string]string

	saveTreeErr     error
	loadTreeErr     error
	deleteTreeErr   error
	listIDsErr      error
	saveMapErr      error
	loadMapErr      error
}

func newMockStore() *mockStore {
	return &mockStore{
		trees:    make(map[string]*merkle.SparseMerkleTree),
		orderMap: make(map[string]string),
	}
}

func (m *mockStore) SaveTree(orderID *big.Int, smt *merkle.SparseMerkleTree) error {
	if m.saveTreeErr != nil {
		return m.saveTreeErr
	}
	m.trees[orderID.String()] = smt
	return nil
}

func (m *mockStore) LoadTree(orderID *big.Int, _ fr.Element) (*merkle.SparseMerkleTree, error) {
	if m.loadTreeErr != nil {
		return nil, m.loadTreeErr
	}
	return m.trees[orderID.String()], nil
}

func (m *mockStore) DeleteTree(orderID *big.Int) error {
	if m.deleteTreeErr != nil {
		return m.deleteTreeErr
	}
	delete(m.trees, orderID.String())
	return nil
}

func (m *mockStore) ListCachedOrderIDs() ([]*big.Int, error) {
	if m.listIDsErr != nil {
		return nil, m.listIDsErr
	}
	var ids []*big.Int
	for k := range m.trees {
		id, _ := new(big.Int).SetString(k, 10)
		if id != nil {
			ids = append(ids, id)
		}
	}
	return ids, nil
}

func (m *mockStore) SaveOrderMapAtomic(orders map[string]string) error {
	if m.saveMapErr != nil {
		return m.saveMapErr
	}
	m.orderMap = make(map[string]string, len(orders))
	for k, v := range orders {
		m.orderMap[k] = v
	}
	return nil
}

func (m *mockStore) LoadOrderMap() (map[string]string, error) {
	if m.loadMapErr != nil {
		return nil, m.loadMapErr
	}
	cp := make(map[string]string, len(m.orderMap))
	for k, v := range m.orderMap {
		cp[k] = v
	}
	return cp, nil
}
