// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// NodeStakingMetaData contains all meta data concerning the NodeStaking contract.
var NodeStakingMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"BURN_ADDRESS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"addresspayable\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MIN_CAPACITY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"STAKE_PER_CHUNK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decreaseCapacity\",\"inputs\":[{\"name\":\"_reduceCapacity\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forceReduceUsed\",\"inputs\":[{\"name\":\"node\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"newUsed\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getMaxSlashable\",\"inputs\":[{\"name\":\"node\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"maxSlashable\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNetworkStats\",\"inputs\":[],\"outputs\":[{\"name\":\"totalNodes\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalCapacityStaked\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalCapacityUsed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNodeInfo\",\"inputs\":[{\"name\":\"node\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"stake\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"capacity\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"used\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"publicKey\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNodeListPage\",\"inputs\":[{\"name\":\"offset\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"addresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"total\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"globalTotalCapacity\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"globalTotalUsed\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasCapacity\",\"inputs\":[{\"name\":\"node\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"requiredChunks\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"increaseCapacity\",\"inputs\":[{\"name\":\"_additionalCapacity\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_market\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isValidNode\",\"inputs\":[{\"name\":\"node\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"market\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nodeIndexInList\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nodeList\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nodes\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"stake\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"capacity\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"used\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"publicKey\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"publicKeyOwner\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"simulateSlash\",\"inputs\":[{\"name\":\"node\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"slashAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"newCapacity\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"willForceExit\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"slashNode\",\"inputs\":[{\"name\":\"node\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"slashAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"forcedOrderExit\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"totalSlashed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakeNode\",\"inputs\":[{\"name\":\"_capacity\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"_publicKey\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"unstakeNode\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateNodeUsed\",\"inputs\":[{\"name\":\"node\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"newUsed\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"ForcedOrderExit\",\"inputs\":[{\"name\":\"node\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"orderIds\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"},{\"name\":\"additionalSlash\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NodeCapacityDecreased\",\"inputs\":[{\"name\":\"node\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"releasedStake\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newCapacity\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NodeCapacityIncreased\",\"inputs\":[{\"name\":\"node\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"additionalStake\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newCapacity\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NodeSlashed\",\"inputs\":[{\"name\":\"node\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"slashAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newCapacity\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"forcedOrderExit\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NodeStaked\",\"inputs\":[{\"name\":\"node\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"stake\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"capacity\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NodeUnstaked\",\"inputs\":[{\"name\":\"node\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"stakeReturned\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
}

// NodeStakingABI is the input ABI used to generate the binding from.
// Deprecated: Use NodeStakingMetaData.ABI instead.
var NodeStakingABI = NodeStakingMetaData.ABI

// NodeStaking is an auto generated Go binding around an Ethereum contract.
type NodeStaking struct {
	NodeStakingCaller     // Read-only binding to the contract
	NodeStakingTransactor // Write-only binding to the contract
	NodeStakingFilterer   // Log filterer for contract events
}

// NodeStakingCaller is an auto generated read-only Go binding around an Ethereum contract.
type NodeStakingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeStakingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NodeStakingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeStakingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NodeStakingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeStakingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NodeStakingSession struct {
	Contract     *NodeStaking      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NodeStakingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NodeStakingCallerSession struct {
	Contract *NodeStakingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// NodeStakingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NodeStakingTransactorSession struct {
	Contract     *NodeStakingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// NodeStakingRaw is an auto generated low-level Go binding around an Ethereum contract.
type NodeStakingRaw struct {
	Contract *NodeStaking // Generic contract binding to access the raw methods on
}

// NodeStakingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NodeStakingCallerRaw struct {
	Contract *NodeStakingCaller // Generic read-only contract binding to access the raw methods on
}

// NodeStakingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NodeStakingTransactorRaw struct {
	Contract *NodeStakingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNodeStaking creates a new instance of NodeStaking, bound to a specific deployed contract.
func NewNodeStaking(address common.Address, backend bind.ContractBackend) (*NodeStaking, error) {
	contract, err := bindNodeStaking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NodeStaking{NodeStakingCaller: NodeStakingCaller{contract: contract}, NodeStakingTransactor: NodeStakingTransactor{contract: contract}, NodeStakingFilterer: NodeStakingFilterer{contract: contract}}, nil
}

// NewNodeStakingCaller creates a new read-only instance of NodeStaking, bound to a specific deployed contract.
func NewNodeStakingCaller(address common.Address, caller bind.ContractCaller) (*NodeStakingCaller, error) {
	contract, err := bindNodeStaking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NodeStakingCaller{contract: contract}, nil
}

// NewNodeStakingTransactor creates a new write-only instance of NodeStaking, bound to a specific deployed contract.
func NewNodeStakingTransactor(address common.Address, transactor bind.ContractTransactor) (*NodeStakingTransactor, error) {
	contract, err := bindNodeStaking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NodeStakingTransactor{contract: contract}, nil
}

// NewNodeStakingFilterer creates a new log filterer instance of NodeStaking, bound to a specific deployed contract.
func NewNodeStakingFilterer(address common.Address, filterer bind.ContractFilterer) (*NodeStakingFilterer, error) {
	contract, err := bindNodeStaking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NodeStakingFilterer{contract: contract}, nil
}

// bindNodeStaking binds a generic wrapper to an already deployed contract.
func bindNodeStaking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NodeStakingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodeStaking *NodeStakingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NodeStaking.Contract.NodeStakingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodeStaking *NodeStakingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeStaking.Contract.NodeStakingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodeStaking *NodeStakingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodeStaking.Contract.NodeStakingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodeStaking *NodeStakingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NodeStaking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodeStaking *NodeStakingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeStaking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodeStaking *NodeStakingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodeStaking.Contract.contract.Transact(opts, method, params...)
}

// BURNADDRESS is a free data retrieval call binding the contract method 0xfccc2813.
//
// Solidity: function BURN_ADDRESS() view returns(address)
func (_NodeStaking *NodeStakingCaller) BURNADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NodeStaking.contract.Call(opts, &out, "BURN_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BURNADDRESS is a free data retrieval call binding the contract method 0xfccc2813.
//
// Solidity: function BURN_ADDRESS() view returns(address)
func (_NodeStaking *NodeStakingSession) BURNADDRESS() (common.Address, error) {
	return _NodeStaking.Contract.BURNADDRESS(&_NodeStaking.CallOpts)
}

// BURNADDRESS is a free data retrieval call binding the contract method 0xfccc2813.
//
// Solidity: function BURN_ADDRESS() view returns(address)
func (_NodeStaking *NodeStakingCallerSession) BURNADDRESS() (common.Address, error) {
	return _NodeStaking.Contract.BURNADDRESS(&_NodeStaking.CallOpts)
}

// MINCAPACITY is a free data retrieval call binding the contract method 0x4256cae6.
//
// Solidity: function MIN_CAPACITY() view returns(uint64)
func (_NodeStaking *NodeStakingCaller) MINCAPACITY(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _NodeStaking.contract.Call(opts, &out, "MIN_CAPACITY")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MINCAPACITY is a free data retrieval call binding the contract method 0x4256cae6.
//
// Solidity: function MIN_CAPACITY() view returns(uint64)
func (_NodeStaking *NodeStakingSession) MINCAPACITY() (uint64, error) {
	return _NodeStaking.Contract.MINCAPACITY(&_NodeStaking.CallOpts)
}

// MINCAPACITY is a free data retrieval call binding the contract method 0x4256cae6.
//
// Solidity: function MIN_CAPACITY() view returns(uint64)
func (_NodeStaking *NodeStakingCallerSession) MINCAPACITY() (uint64, error) {
	return _NodeStaking.Contract.MINCAPACITY(&_NodeStaking.CallOpts)
}

// STAKEPERCHUNK is a free data retrieval call binding the contract method 0x889ed6cd.
//
// Solidity: function STAKE_PER_CHUNK() view returns(uint256)
func (_NodeStaking *NodeStakingCaller) STAKEPERCHUNK(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NodeStaking.contract.Call(opts, &out, "STAKE_PER_CHUNK")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// STAKEPERCHUNK is a free data retrieval call binding the contract method 0x889ed6cd.
//
// Solidity: function STAKE_PER_CHUNK() view returns(uint256)
func (_NodeStaking *NodeStakingSession) STAKEPERCHUNK() (*big.Int, error) {
	return _NodeStaking.Contract.STAKEPERCHUNK(&_NodeStaking.CallOpts)
}

// STAKEPERCHUNK is a free data retrieval call binding the contract method 0x889ed6cd.
//
// Solidity: function STAKE_PER_CHUNK() view returns(uint256)
func (_NodeStaking *NodeStakingCallerSession) STAKEPERCHUNK() (*big.Int, error) {
	return _NodeStaking.Contract.STAKEPERCHUNK(&_NodeStaking.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_NodeStaking *NodeStakingCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _NodeStaking.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_NodeStaking *NodeStakingSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _NodeStaking.Contract.UPGRADEINTERFACEVERSION(&_NodeStaking.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_NodeStaking *NodeStakingCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _NodeStaking.Contract.UPGRADEINTERFACEVERSION(&_NodeStaking.CallOpts)
}

// GetMaxSlashable is a free data retrieval call binding the contract method 0xdb052446.
//
// Solidity: function getMaxSlashable(address node) view returns(uint256 maxSlashable)
func (_NodeStaking *NodeStakingCaller) GetMaxSlashable(opts *bind.CallOpts, node common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NodeStaking.contract.Call(opts, &out, "getMaxSlashable", node)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxSlashable is a free data retrieval call binding the contract method 0xdb052446.
//
// Solidity: function getMaxSlashable(address node) view returns(uint256 maxSlashable)
func (_NodeStaking *NodeStakingSession) GetMaxSlashable(node common.Address) (*big.Int, error) {
	return _NodeStaking.Contract.GetMaxSlashable(&_NodeStaking.CallOpts, node)
}

// GetMaxSlashable is a free data retrieval call binding the contract method 0xdb052446.
//
// Solidity: function getMaxSlashable(address node) view returns(uint256 maxSlashable)
func (_NodeStaking *NodeStakingCallerSession) GetMaxSlashable(node common.Address) (*big.Int, error) {
	return _NodeStaking.Contract.GetMaxSlashable(&_NodeStaking.CallOpts, node)
}

// GetNetworkStats is a free data retrieval call binding the contract method 0x4a662f48.
//
// Solidity: function getNetworkStats() view returns(uint256 totalNodes, uint256 totalCapacityStaked, uint256 totalCapacityUsed)
func (_NodeStaking *NodeStakingCaller) GetNetworkStats(opts *bind.CallOpts) (struct {
	TotalNodes          *big.Int
	TotalCapacityStaked *big.Int
	TotalCapacityUsed   *big.Int
}, error) {
	var out []interface{}
	err := _NodeStaking.contract.Call(opts, &out, "getNetworkStats")

	outstruct := new(struct {
		TotalNodes          *big.Int
		TotalCapacityStaked *big.Int
		TotalCapacityUsed   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalNodes = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TotalCapacityStaked = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.TotalCapacityUsed = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetNetworkStats is a free data retrieval call binding the contract method 0x4a662f48.
//
// Solidity: function getNetworkStats() view returns(uint256 totalNodes, uint256 totalCapacityStaked, uint256 totalCapacityUsed)
func (_NodeStaking *NodeStakingSession) GetNetworkStats() (struct {
	TotalNodes          *big.Int
	TotalCapacityStaked *big.Int
	TotalCapacityUsed   *big.Int
}, error) {
	return _NodeStaking.Contract.GetNetworkStats(&_NodeStaking.CallOpts)
}

// GetNetworkStats is a free data retrieval call binding the contract method 0x4a662f48.
//
// Solidity: function getNetworkStats() view returns(uint256 totalNodes, uint256 totalCapacityStaked, uint256 totalCapacityUsed)
func (_NodeStaking *NodeStakingCallerSession) GetNetworkStats() (struct {
	TotalNodes          *big.Int
	TotalCapacityStaked *big.Int
	TotalCapacityUsed   *big.Int
}, error) {
	return _NodeStaking.Contract.GetNetworkStats(&_NodeStaking.CallOpts)
}

// GetNodeInfo is a free data retrieval call binding the contract method 0x582115fb.
//
// Solidity: function getNodeInfo(address node) view returns(uint256 stake, uint64 capacity, uint64 used, uint256 publicKey)
func (_NodeStaking *NodeStakingCaller) GetNodeInfo(opts *bind.CallOpts, node common.Address) (struct {
	Stake     *big.Int
	Capacity  uint64
	Used      uint64
	PublicKey *big.Int
}, error) {
	var out []interface{}
	err := _NodeStaking.contract.Call(opts, &out, "getNodeInfo", node)

	outstruct := new(struct {
		Stake     *big.Int
		Capacity  uint64
		Used      uint64
		PublicKey *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Stake = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Capacity = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.Used = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.PublicKey = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetNodeInfo is a free data retrieval call binding the contract method 0x582115fb.
//
// Solidity: function getNodeInfo(address node) view returns(uint256 stake, uint64 capacity, uint64 used, uint256 publicKey)
func (_NodeStaking *NodeStakingSession) GetNodeInfo(node common.Address) (struct {
	Stake     *big.Int
	Capacity  uint64
	Used      uint64
	PublicKey *big.Int
}, error) {
	return _NodeStaking.Contract.GetNodeInfo(&_NodeStaking.CallOpts, node)
}

// GetNodeInfo is a free data retrieval call binding the contract method 0x582115fb.
//
// Solidity: function getNodeInfo(address node) view returns(uint256 stake, uint64 capacity, uint64 used, uint256 publicKey)
func (_NodeStaking *NodeStakingCallerSession) GetNodeInfo(node common.Address) (struct {
	Stake     *big.Int
	Capacity  uint64
	Used      uint64
	PublicKey *big.Int
}, error) {
	return _NodeStaking.Contract.GetNodeInfo(&_NodeStaking.CallOpts, node)
}

// GetNodeListPage is a free data retrieval call binding the contract method 0xb101939d.
//
// Solidity: function getNodeListPage(uint256 offset, uint256 limit) view returns(address[] addresses, uint256 total)
func (_NodeStaking *NodeStakingCaller) GetNodeListPage(opts *bind.CallOpts, offset *big.Int, limit *big.Int) (struct {
	Addresses []common.Address
	Total     *big.Int
}, error) {
	var out []interface{}
	err := _NodeStaking.contract.Call(opts, &out, "getNodeListPage", offset, limit)

	outstruct := new(struct {
		Addresses []common.Address
		Total     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Addresses = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Total = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetNodeListPage is a free data retrieval call binding the contract method 0xb101939d.
//
// Solidity: function getNodeListPage(uint256 offset, uint256 limit) view returns(address[] addresses, uint256 total)
func (_NodeStaking *NodeStakingSession) GetNodeListPage(offset *big.Int, limit *big.Int) (struct {
	Addresses []common.Address
	Total     *big.Int
}, error) {
	return _NodeStaking.Contract.GetNodeListPage(&_NodeStaking.CallOpts, offset, limit)
}

// GetNodeListPage is a free data retrieval call binding the contract method 0xb101939d.
//
// Solidity: function getNodeListPage(uint256 offset, uint256 limit) view returns(address[] addresses, uint256 total)
func (_NodeStaking *NodeStakingCallerSession) GetNodeListPage(offset *big.Int, limit *big.Int) (struct {
	Addresses []common.Address
	Total     *big.Int
}, error) {
	return _NodeStaking.Contract.GetNodeListPage(&_NodeStaking.CallOpts, offset, limit)
}

// GlobalTotalCapacity is a free data retrieval call binding the contract method 0x9b7553fa.
//
// Solidity: function globalTotalCapacity() view returns(uint256)
func (_NodeStaking *NodeStakingCaller) GlobalTotalCapacity(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NodeStaking.contract.Call(opts, &out, "globalTotalCapacity")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GlobalTotalCapacity is a free data retrieval call binding the contract method 0x9b7553fa.
//
// Solidity: function globalTotalCapacity() view returns(uint256)
func (_NodeStaking *NodeStakingSession) GlobalTotalCapacity() (*big.Int, error) {
	return _NodeStaking.Contract.GlobalTotalCapacity(&_NodeStaking.CallOpts)
}

// GlobalTotalCapacity is a free data retrieval call binding the contract method 0x9b7553fa.
//
// Solidity: function globalTotalCapacity() view returns(uint256)
func (_NodeStaking *NodeStakingCallerSession) GlobalTotalCapacity() (*big.Int, error) {
	return _NodeStaking.Contract.GlobalTotalCapacity(&_NodeStaking.CallOpts)
}

// GlobalTotalUsed is a free data retrieval call binding the contract method 0x59469156.
//
// Solidity: function globalTotalUsed() view returns(uint256)
func (_NodeStaking *NodeStakingCaller) GlobalTotalUsed(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NodeStaking.contract.Call(opts, &out, "globalTotalUsed")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GlobalTotalUsed is a free data retrieval call binding the contract method 0x59469156.
//
// Solidity: function globalTotalUsed() view returns(uint256)
func (_NodeStaking *NodeStakingSession) GlobalTotalUsed() (*big.Int, error) {
	return _NodeStaking.Contract.GlobalTotalUsed(&_NodeStaking.CallOpts)
}

// GlobalTotalUsed is a free data retrieval call binding the contract method 0x59469156.
//
// Solidity: function globalTotalUsed() view returns(uint256)
func (_NodeStaking *NodeStakingCallerSession) GlobalTotalUsed() (*big.Int, error) {
	return _NodeStaking.Contract.GlobalTotalUsed(&_NodeStaking.CallOpts)
}

// HasCapacity is a free data retrieval call binding the contract method 0xba48c453.
//
// Solidity: function hasCapacity(address node, uint32 requiredChunks) view returns(bool)
func (_NodeStaking *NodeStakingCaller) HasCapacity(opts *bind.CallOpts, node common.Address, requiredChunks uint32) (bool, error) {
	var out []interface{}
	err := _NodeStaking.contract.Call(opts, &out, "hasCapacity", node, requiredChunks)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasCapacity is a free data retrieval call binding the contract method 0xba48c453.
//
// Solidity: function hasCapacity(address node, uint32 requiredChunks) view returns(bool)
func (_NodeStaking *NodeStakingSession) HasCapacity(node common.Address, requiredChunks uint32) (bool, error) {
	return _NodeStaking.Contract.HasCapacity(&_NodeStaking.CallOpts, node, requiredChunks)
}

// HasCapacity is a free data retrieval call binding the contract method 0xba48c453.
//
// Solidity: function hasCapacity(address node, uint32 requiredChunks) view returns(bool)
func (_NodeStaking *NodeStakingCallerSession) HasCapacity(node common.Address, requiredChunks uint32) (bool, error) {
	return _NodeStaking.Contract.HasCapacity(&_NodeStaking.CallOpts, node, requiredChunks)
}

// IsValidNode is a free data retrieval call binding the contract method 0x9ebd11ef.
//
// Solidity: function isValidNode(address node) view returns(bool)
func (_NodeStaking *NodeStakingCaller) IsValidNode(opts *bind.CallOpts, node common.Address) (bool, error) {
	var out []interface{}
	err := _NodeStaking.contract.Call(opts, &out, "isValidNode", node)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidNode is a free data retrieval call binding the contract method 0x9ebd11ef.
//
// Solidity: function isValidNode(address node) view returns(bool)
func (_NodeStaking *NodeStakingSession) IsValidNode(node common.Address) (bool, error) {
	return _NodeStaking.Contract.IsValidNode(&_NodeStaking.CallOpts, node)
}

// IsValidNode is a free data retrieval call binding the contract method 0x9ebd11ef.
//
// Solidity: function isValidNode(address node) view returns(bool)
func (_NodeStaking *NodeStakingCallerSession) IsValidNode(node common.Address) (bool, error) {
	return _NodeStaking.Contract.IsValidNode(&_NodeStaking.CallOpts, node)
}

// Market is a free data retrieval call binding the contract method 0x80f55605.
//
// Solidity: function market() view returns(address)
func (_NodeStaking *NodeStakingCaller) Market(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NodeStaking.contract.Call(opts, &out, "market")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Market is a free data retrieval call binding the contract method 0x80f55605.
//
// Solidity: function market() view returns(address)
func (_NodeStaking *NodeStakingSession) Market() (common.Address, error) {
	return _NodeStaking.Contract.Market(&_NodeStaking.CallOpts)
}

// Market is a free data retrieval call binding the contract method 0x80f55605.
//
// Solidity: function market() view returns(address)
func (_NodeStaking *NodeStakingCallerSession) Market() (common.Address, error) {
	return _NodeStaking.Contract.Market(&_NodeStaking.CallOpts)
}

// NodeIndexInList is a free data retrieval call binding the contract method 0x0af21e60.
//
// Solidity: function nodeIndexInList(address ) view returns(uint256)
func (_NodeStaking *NodeStakingCaller) NodeIndexInList(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NodeStaking.contract.Call(opts, &out, "nodeIndexInList", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NodeIndexInList is a free data retrieval call binding the contract method 0x0af21e60.
//
// Solidity: function nodeIndexInList(address ) view returns(uint256)
func (_NodeStaking *NodeStakingSession) NodeIndexInList(arg0 common.Address) (*big.Int, error) {
	return _NodeStaking.Contract.NodeIndexInList(&_NodeStaking.CallOpts, arg0)
}

// NodeIndexInList is a free data retrieval call binding the contract method 0x0af21e60.
//
// Solidity: function nodeIndexInList(address ) view returns(uint256)
func (_NodeStaking *NodeStakingCallerSession) NodeIndexInList(arg0 common.Address) (*big.Int, error) {
	return _NodeStaking.Contract.NodeIndexInList(&_NodeStaking.CallOpts, arg0)
}

// NodeList is a free data retrieval call binding the contract method 0x208f2a31.
//
// Solidity: function nodeList(uint256 ) view returns(address)
func (_NodeStaking *NodeStakingCaller) NodeList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _NodeStaking.contract.Call(opts, &out, "nodeList", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NodeList is a free data retrieval call binding the contract method 0x208f2a31.
//
// Solidity: function nodeList(uint256 ) view returns(address)
func (_NodeStaking *NodeStakingSession) NodeList(arg0 *big.Int) (common.Address, error) {
	return _NodeStaking.Contract.NodeList(&_NodeStaking.CallOpts, arg0)
}

// NodeList is a free data retrieval call binding the contract method 0x208f2a31.
//
// Solidity: function nodeList(uint256 ) view returns(address)
func (_NodeStaking *NodeStakingCallerSession) NodeList(arg0 *big.Int) (common.Address, error) {
	return _NodeStaking.Contract.NodeList(&_NodeStaking.CallOpts, arg0)
}

// Nodes is a free data retrieval call binding the contract method 0x189a5a17.
//
// Solidity: function nodes(address ) view returns(uint256 stake, uint64 capacity, uint64 used, uint256 publicKey)
func (_NodeStaking *NodeStakingCaller) Nodes(opts *bind.CallOpts, arg0 common.Address) (struct {
	Stake     *big.Int
	Capacity  uint64
	Used      uint64
	PublicKey *big.Int
}, error) {
	var out []interface{}
	err := _NodeStaking.contract.Call(opts, &out, "nodes", arg0)

	outstruct := new(struct {
		Stake     *big.Int
		Capacity  uint64
		Used      uint64
		PublicKey *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Stake = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Capacity = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.Used = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.PublicKey = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Nodes is a free data retrieval call binding the contract method 0x189a5a17.
//
// Solidity: function nodes(address ) view returns(uint256 stake, uint64 capacity, uint64 used, uint256 publicKey)
func (_NodeStaking *NodeStakingSession) Nodes(arg0 common.Address) (struct {
	Stake     *big.Int
	Capacity  uint64
	Used      uint64
	PublicKey *big.Int
}, error) {
	return _NodeStaking.Contract.Nodes(&_NodeStaking.CallOpts, arg0)
}

// Nodes is a free data retrieval call binding the contract method 0x189a5a17.
//
// Solidity: function nodes(address ) view returns(uint256 stake, uint64 capacity, uint64 used, uint256 publicKey)
func (_NodeStaking *NodeStakingCallerSession) Nodes(arg0 common.Address) (struct {
	Stake     *big.Int
	Capacity  uint64
	Used      uint64
	PublicKey *big.Int
}, error) {
	return _NodeStaking.Contract.Nodes(&_NodeStaking.CallOpts, arg0)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_NodeStaking *NodeStakingCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NodeStaking.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_NodeStaking *NodeStakingSession) ProxiableUUID() ([32]byte, error) {
	return _NodeStaking.Contract.ProxiableUUID(&_NodeStaking.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_NodeStaking *NodeStakingCallerSession) ProxiableUUID() ([32]byte, error) {
	return _NodeStaking.Contract.ProxiableUUID(&_NodeStaking.CallOpts)
}

// PublicKeyOwner is a free data retrieval call binding the contract method 0x6c175d10.
//
// Solidity: function publicKeyOwner(uint256 ) view returns(address)
func (_NodeStaking *NodeStakingCaller) PublicKeyOwner(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _NodeStaking.contract.Call(opts, &out, "publicKeyOwner", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PublicKeyOwner is a free data retrieval call binding the contract method 0x6c175d10.
//
// Solidity: function publicKeyOwner(uint256 ) view returns(address)
func (_NodeStaking *NodeStakingSession) PublicKeyOwner(arg0 *big.Int) (common.Address, error) {
	return _NodeStaking.Contract.PublicKeyOwner(&_NodeStaking.CallOpts, arg0)
}

// PublicKeyOwner is a free data retrieval call binding the contract method 0x6c175d10.
//
// Solidity: function publicKeyOwner(uint256 ) view returns(address)
func (_NodeStaking *NodeStakingCallerSession) PublicKeyOwner(arg0 *big.Int) (common.Address, error) {
	return _NodeStaking.Contract.PublicKeyOwner(&_NodeStaking.CallOpts, arg0)
}

// SimulateSlash is a free data retrieval call binding the contract method 0x80937ee8.
//
// Solidity: function simulateSlash(address node, uint256 slashAmount) view returns(uint64 newCapacity, bool willForceExit)
func (_NodeStaking *NodeStakingCaller) SimulateSlash(opts *bind.CallOpts, node common.Address, slashAmount *big.Int) (struct {
	NewCapacity   uint64
	WillForceExit bool
}, error) {
	var out []interface{}
	err := _NodeStaking.contract.Call(opts, &out, "simulateSlash", node, slashAmount)

	outstruct := new(struct {
		NewCapacity   uint64
		WillForceExit bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NewCapacity = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.WillForceExit = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// SimulateSlash is a free data retrieval call binding the contract method 0x80937ee8.
//
// Solidity: function simulateSlash(address node, uint256 slashAmount) view returns(uint64 newCapacity, bool willForceExit)
func (_NodeStaking *NodeStakingSession) SimulateSlash(node common.Address, slashAmount *big.Int) (struct {
	NewCapacity   uint64
	WillForceExit bool
}, error) {
	return _NodeStaking.Contract.SimulateSlash(&_NodeStaking.CallOpts, node, slashAmount)
}

// SimulateSlash is a free data retrieval call binding the contract method 0x80937ee8.
//
// Solidity: function simulateSlash(address node, uint256 slashAmount) view returns(uint64 newCapacity, bool willForceExit)
func (_NodeStaking *NodeStakingCallerSession) SimulateSlash(node common.Address, slashAmount *big.Int) (struct {
	NewCapacity   uint64
	WillForceExit bool
}, error) {
	return _NodeStaking.Contract.SimulateSlash(&_NodeStaking.CallOpts, node, slashAmount)
}

// DecreaseCapacity is a paid mutator transaction binding the contract method 0xc90cbfbe.
//
// Solidity: function decreaseCapacity(uint64 _reduceCapacity) returns()
func (_NodeStaking *NodeStakingTransactor) DecreaseCapacity(opts *bind.TransactOpts, _reduceCapacity uint64) (*types.Transaction, error) {
	return _NodeStaking.contract.Transact(opts, "decreaseCapacity", _reduceCapacity)
}

// DecreaseCapacity is a paid mutator transaction binding the contract method 0xc90cbfbe.
//
// Solidity: function decreaseCapacity(uint64 _reduceCapacity) returns()
func (_NodeStaking *NodeStakingSession) DecreaseCapacity(_reduceCapacity uint64) (*types.Transaction, error) {
	return _NodeStaking.Contract.DecreaseCapacity(&_NodeStaking.TransactOpts, _reduceCapacity)
}

// DecreaseCapacity is a paid mutator transaction binding the contract method 0xc90cbfbe.
//
// Solidity: function decreaseCapacity(uint64 _reduceCapacity) returns()
func (_NodeStaking *NodeStakingTransactorSession) DecreaseCapacity(_reduceCapacity uint64) (*types.Transaction, error) {
	return _NodeStaking.Contract.DecreaseCapacity(&_NodeStaking.TransactOpts, _reduceCapacity)
}

// ForceReduceUsed is a paid mutator transaction binding the contract method 0xd578bc14.
//
// Solidity: function forceReduceUsed(address node, uint64 newUsed) returns()
func (_NodeStaking *NodeStakingTransactor) ForceReduceUsed(opts *bind.TransactOpts, node common.Address, newUsed uint64) (*types.Transaction, error) {
	return _NodeStaking.contract.Transact(opts, "forceReduceUsed", node, newUsed)
}

// ForceReduceUsed is a paid mutator transaction binding the contract method 0xd578bc14.
//
// Solidity: function forceReduceUsed(address node, uint64 newUsed) returns()
func (_NodeStaking *NodeStakingSession) ForceReduceUsed(node common.Address, newUsed uint64) (*types.Transaction, error) {
	return _NodeStaking.Contract.ForceReduceUsed(&_NodeStaking.TransactOpts, node, newUsed)
}

// ForceReduceUsed is a paid mutator transaction binding the contract method 0xd578bc14.
//
// Solidity: function forceReduceUsed(address node, uint64 newUsed) returns()
func (_NodeStaking *NodeStakingTransactorSession) ForceReduceUsed(node common.Address, newUsed uint64) (*types.Transaction, error) {
	return _NodeStaking.Contract.ForceReduceUsed(&_NodeStaking.TransactOpts, node, newUsed)
}

// IncreaseCapacity is a paid mutator transaction binding the contract method 0xbb34a381.
//
// Solidity: function increaseCapacity(uint64 _additionalCapacity) payable returns()
func (_NodeStaking *NodeStakingTransactor) IncreaseCapacity(opts *bind.TransactOpts, _additionalCapacity uint64) (*types.Transaction, error) {
	return _NodeStaking.contract.Transact(opts, "increaseCapacity", _additionalCapacity)
}

// IncreaseCapacity is a paid mutator transaction binding the contract method 0xbb34a381.
//
// Solidity: function increaseCapacity(uint64 _additionalCapacity) payable returns()
func (_NodeStaking *NodeStakingSession) IncreaseCapacity(_additionalCapacity uint64) (*types.Transaction, error) {
	return _NodeStaking.Contract.IncreaseCapacity(&_NodeStaking.TransactOpts, _additionalCapacity)
}

// IncreaseCapacity is a paid mutator transaction binding the contract method 0xbb34a381.
//
// Solidity: function increaseCapacity(uint64 _additionalCapacity) payable returns()
func (_NodeStaking *NodeStakingTransactorSession) IncreaseCapacity(_additionalCapacity uint64) (*types.Transaction, error) {
	return _NodeStaking.Contract.IncreaseCapacity(&_NodeStaking.TransactOpts, _additionalCapacity)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _market) returns()
func (_NodeStaking *NodeStakingTransactor) Initialize(opts *bind.TransactOpts, _market common.Address) (*types.Transaction, error) {
	return _NodeStaking.contract.Transact(opts, "initialize", _market)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _market) returns()
func (_NodeStaking *NodeStakingSession) Initialize(_market common.Address) (*types.Transaction, error) {
	return _NodeStaking.Contract.Initialize(&_NodeStaking.TransactOpts, _market)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _market) returns()
func (_NodeStaking *NodeStakingTransactorSession) Initialize(_market common.Address) (*types.Transaction, error) {
	return _NodeStaking.Contract.Initialize(&_NodeStaking.TransactOpts, _market)
}

// SlashNode is a paid mutator transaction binding the contract method 0xd2f9f285.
//
// Solidity: function slashNode(address node, uint256 slashAmount) returns(bool forcedOrderExit, uint256 totalSlashed)
func (_NodeStaking *NodeStakingTransactor) SlashNode(opts *bind.TransactOpts, node common.Address, slashAmount *big.Int) (*types.Transaction, error) {
	return _NodeStaking.contract.Transact(opts, "slashNode", node, slashAmount)
}

// SlashNode is a paid mutator transaction binding the contract method 0xd2f9f285.
//
// Solidity: function slashNode(address node, uint256 slashAmount) returns(bool forcedOrderExit, uint256 totalSlashed)
func (_NodeStaking *NodeStakingSession) SlashNode(node common.Address, slashAmount *big.Int) (*types.Transaction, error) {
	return _NodeStaking.Contract.SlashNode(&_NodeStaking.TransactOpts, node, slashAmount)
}

// SlashNode is a paid mutator transaction binding the contract method 0xd2f9f285.
//
// Solidity: function slashNode(address node, uint256 slashAmount) returns(bool forcedOrderExit, uint256 totalSlashed)
func (_NodeStaking *NodeStakingTransactorSession) SlashNode(node common.Address, slashAmount *big.Int) (*types.Transaction, error) {
	return _NodeStaking.Contract.SlashNode(&_NodeStaking.TransactOpts, node, slashAmount)
}

// StakeNode is a paid mutator transaction binding the contract method 0x780a7750.
//
// Solidity: function stakeNode(uint64 _capacity, uint256 _publicKey) payable returns()
func (_NodeStaking *NodeStakingTransactor) StakeNode(opts *bind.TransactOpts, _capacity uint64, _publicKey *big.Int) (*types.Transaction, error) {
	return _NodeStaking.contract.Transact(opts, "stakeNode", _capacity, _publicKey)
}

// StakeNode is a paid mutator transaction binding the contract method 0x780a7750.
//
// Solidity: function stakeNode(uint64 _capacity, uint256 _publicKey) payable returns()
func (_NodeStaking *NodeStakingSession) StakeNode(_capacity uint64, _publicKey *big.Int) (*types.Transaction, error) {
	return _NodeStaking.Contract.StakeNode(&_NodeStaking.TransactOpts, _capacity, _publicKey)
}

// StakeNode is a paid mutator transaction binding the contract method 0x780a7750.
//
// Solidity: function stakeNode(uint64 _capacity, uint256 _publicKey) payable returns()
func (_NodeStaking *NodeStakingTransactorSession) StakeNode(_capacity uint64, _publicKey *big.Int) (*types.Transaction, error) {
	return _NodeStaking.Contract.StakeNode(&_NodeStaking.TransactOpts, _capacity, _publicKey)
}

// UnstakeNode is a paid mutator transaction binding the contract method 0x95d426ad.
//
// Solidity: function unstakeNode() returns()
func (_NodeStaking *NodeStakingTransactor) UnstakeNode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeStaking.contract.Transact(opts, "unstakeNode")
}

// UnstakeNode is a paid mutator transaction binding the contract method 0x95d426ad.
//
// Solidity: function unstakeNode() returns()
func (_NodeStaking *NodeStakingSession) UnstakeNode() (*types.Transaction, error) {
	return _NodeStaking.Contract.UnstakeNode(&_NodeStaking.TransactOpts)
}

// UnstakeNode is a paid mutator transaction binding the contract method 0x95d426ad.
//
// Solidity: function unstakeNode() returns()
func (_NodeStaking *NodeStakingTransactorSession) UnstakeNode() (*types.Transaction, error) {
	return _NodeStaking.Contract.UnstakeNode(&_NodeStaking.TransactOpts)
}

// UpdateNodeUsed is a paid mutator transaction binding the contract method 0xe5b8c891.
//
// Solidity: function updateNodeUsed(address node, uint64 newUsed) returns()
func (_NodeStaking *NodeStakingTransactor) UpdateNodeUsed(opts *bind.TransactOpts, node common.Address, newUsed uint64) (*types.Transaction, error) {
	return _NodeStaking.contract.Transact(opts, "updateNodeUsed", node, newUsed)
}

// UpdateNodeUsed is a paid mutator transaction binding the contract method 0xe5b8c891.
//
// Solidity: function updateNodeUsed(address node, uint64 newUsed) returns()
func (_NodeStaking *NodeStakingSession) UpdateNodeUsed(node common.Address, newUsed uint64) (*types.Transaction, error) {
	return _NodeStaking.Contract.UpdateNodeUsed(&_NodeStaking.TransactOpts, node, newUsed)
}

// UpdateNodeUsed is a paid mutator transaction binding the contract method 0xe5b8c891.
//
// Solidity: function updateNodeUsed(address node, uint64 newUsed) returns()
func (_NodeStaking *NodeStakingTransactorSession) UpdateNodeUsed(node common.Address, newUsed uint64) (*types.Transaction, error) {
	return _NodeStaking.Contract.UpdateNodeUsed(&_NodeStaking.TransactOpts, node, newUsed)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_NodeStaking *NodeStakingTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _NodeStaking.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_NodeStaking *NodeStakingSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _NodeStaking.Contract.UpgradeToAndCall(&_NodeStaking.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_NodeStaking *NodeStakingTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _NodeStaking.Contract.UpgradeToAndCall(&_NodeStaking.TransactOpts, newImplementation, data)
}

// NodeStakingForcedOrderExitIterator is returned from FilterForcedOrderExit and is used to iterate over the raw logs and unpacked data for ForcedOrderExit events raised by the NodeStaking contract.
type NodeStakingForcedOrderExitIterator struct {
	Event *NodeStakingForcedOrderExit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NodeStakingForcedOrderExitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeStakingForcedOrderExit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NodeStakingForcedOrderExit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NodeStakingForcedOrderExitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeStakingForcedOrderExitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeStakingForcedOrderExit represents a ForcedOrderExit event raised by the NodeStaking contract.
type NodeStakingForcedOrderExit struct {
	Node            common.Address
	OrderIds        []*big.Int
	AdditionalSlash *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterForcedOrderExit is a free log retrieval operation binding the contract event 0xd3dca24110f4fb1f953dd4d85238f61cef1f0e83bcf2b57bcd25e2ae5add98f8.
//
// Solidity: event ForcedOrderExit(address indexed node, uint256[] orderIds, uint256 additionalSlash)
func (_NodeStaking *NodeStakingFilterer) FilterForcedOrderExit(opts *bind.FilterOpts, node []common.Address) (*NodeStakingForcedOrderExitIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _NodeStaking.contract.FilterLogs(opts, "ForcedOrderExit", nodeRule)
	if err != nil {
		return nil, err
	}
	return &NodeStakingForcedOrderExitIterator{contract: _NodeStaking.contract, event: "ForcedOrderExit", logs: logs, sub: sub}, nil
}

// WatchForcedOrderExit is a free log subscription operation binding the contract event 0xd3dca24110f4fb1f953dd4d85238f61cef1f0e83bcf2b57bcd25e2ae5add98f8.
//
// Solidity: event ForcedOrderExit(address indexed node, uint256[] orderIds, uint256 additionalSlash)
func (_NodeStaking *NodeStakingFilterer) WatchForcedOrderExit(opts *bind.WatchOpts, sink chan<- *NodeStakingForcedOrderExit, node []common.Address) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _NodeStaking.contract.WatchLogs(opts, "ForcedOrderExit", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeStakingForcedOrderExit)
				if err := _NodeStaking.contract.UnpackLog(event, "ForcedOrderExit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseForcedOrderExit is a log parse operation binding the contract event 0xd3dca24110f4fb1f953dd4d85238f61cef1f0e83bcf2b57bcd25e2ae5add98f8.
//
// Solidity: event ForcedOrderExit(address indexed node, uint256[] orderIds, uint256 additionalSlash)
func (_NodeStaking *NodeStakingFilterer) ParseForcedOrderExit(log types.Log) (*NodeStakingForcedOrderExit, error) {
	event := new(NodeStakingForcedOrderExit)
	if err := _NodeStaking.contract.UnpackLog(event, "ForcedOrderExit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeStakingInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the NodeStaking contract.
type NodeStakingInitializedIterator struct {
	Event *NodeStakingInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NodeStakingInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeStakingInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NodeStakingInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NodeStakingInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeStakingInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeStakingInitialized represents a Initialized event raised by the NodeStaking contract.
type NodeStakingInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_NodeStaking *NodeStakingFilterer) FilterInitialized(opts *bind.FilterOpts) (*NodeStakingInitializedIterator, error) {

	logs, sub, err := _NodeStaking.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &NodeStakingInitializedIterator{contract: _NodeStaking.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_NodeStaking *NodeStakingFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *NodeStakingInitialized) (event.Subscription, error) {

	logs, sub, err := _NodeStaking.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeStakingInitialized)
				if err := _NodeStaking.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_NodeStaking *NodeStakingFilterer) ParseInitialized(log types.Log) (*NodeStakingInitialized, error) {
	event := new(NodeStakingInitialized)
	if err := _NodeStaking.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeStakingNodeCapacityDecreasedIterator is returned from FilterNodeCapacityDecreased and is used to iterate over the raw logs and unpacked data for NodeCapacityDecreased events raised by the NodeStaking contract.
type NodeStakingNodeCapacityDecreasedIterator struct {
	Event *NodeStakingNodeCapacityDecreased // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NodeStakingNodeCapacityDecreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeStakingNodeCapacityDecreased)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NodeStakingNodeCapacityDecreased)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NodeStakingNodeCapacityDecreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeStakingNodeCapacityDecreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeStakingNodeCapacityDecreased represents a NodeCapacityDecreased event raised by the NodeStaking contract.
type NodeStakingNodeCapacityDecreased struct {
	Node          common.Address
	ReleasedStake *big.Int
	NewCapacity   uint64
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNodeCapacityDecreased is a free log retrieval operation binding the contract event 0xf5e05c1f669287993dc39670285b277b2d24b60be6d7767c7db3c85172762765.
//
// Solidity: event NodeCapacityDecreased(address indexed node, uint256 releasedStake, uint64 newCapacity)
func (_NodeStaking *NodeStakingFilterer) FilterNodeCapacityDecreased(opts *bind.FilterOpts, node []common.Address) (*NodeStakingNodeCapacityDecreasedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _NodeStaking.contract.FilterLogs(opts, "NodeCapacityDecreased", nodeRule)
	if err != nil {
		return nil, err
	}
	return &NodeStakingNodeCapacityDecreasedIterator{contract: _NodeStaking.contract, event: "NodeCapacityDecreased", logs: logs, sub: sub}, nil
}

// WatchNodeCapacityDecreased is a free log subscription operation binding the contract event 0xf5e05c1f669287993dc39670285b277b2d24b60be6d7767c7db3c85172762765.
//
// Solidity: event NodeCapacityDecreased(address indexed node, uint256 releasedStake, uint64 newCapacity)
func (_NodeStaking *NodeStakingFilterer) WatchNodeCapacityDecreased(opts *bind.WatchOpts, sink chan<- *NodeStakingNodeCapacityDecreased, node []common.Address) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _NodeStaking.contract.WatchLogs(opts, "NodeCapacityDecreased", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeStakingNodeCapacityDecreased)
				if err := _NodeStaking.contract.UnpackLog(event, "NodeCapacityDecreased", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNodeCapacityDecreased is a log parse operation binding the contract event 0xf5e05c1f669287993dc39670285b277b2d24b60be6d7767c7db3c85172762765.
//
// Solidity: event NodeCapacityDecreased(address indexed node, uint256 releasedStake, uint64 newCapacity)
func (_NodeStaking *NodeStakingFilterer) ParseNodeCapacityDecreased(log types.Log) (*NodeStakingNodeCapacityDecreased, error) {
	event := new(NodeStakingNodeCapacityDecreased)
	if err := _NodeStaking.contract.UnpackLog(event, "NodeCapacityDecreased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeStakingNodeCapacityIncreasedIterator is returned from FilterNodeCapacityIncreased and is used to iterate over the raw logs and unpacked data for NodeCapacityIncreased events raised by the NodeStaking contract.
type NodeStakingNodeCapacityIncreasedIterator struct {
	Event *NodeStakingNodeCapacityIncreased // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NodeStakingNodeCapacityIncreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeStakingNodeCapacityIncreased)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NodeStakingNodeCapacityIncreased)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NodeStakingNodeCapacityIncreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeStakingNodeCapacityIncreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeStakingNodeCapacityIncreased represents a NodeCapacityIncreased event raised by the NodeStaking contract.
type NodeStakingNodeCapacityIncreased struct {
	Node            common.Address
	AdditionalStake *big.Int
	NewCapacity     uint64
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNodeCapacityIncreased is a free log retrieval operation binding the contract event 0x026b1e4d13deb2b6423daed34bf909bf7fdccc09f1ed246ba9e17aaf290df300.
//
// Solidity: event NodeCapacityIncreased(address indexed node, uint256 additionalStake, uint64 newCapacity)
func (_NodeStaking *NodeStakingFilterer) FilterNodeCapacityIncreased(opts *bind.FilterOpts, node []common.Address) (*NodeStakingNodeCapacityIncreasedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _NodeStaking.contract.FilterLogs(opts, "NodeCapacityIncreased", nodeRule)
	if err != nil {
		return nil, err
	}
	return &NodeStakingNodeCapacityIncreasedIterator{contract: _NodeStaking.contract, event: "NodeCapacityIncreased", logs: logs, sub: sub}, nil
}

// WatchNodeCapacityIncreased is a free log subscription operation binding the contract event 0x026b1e4d13deb2b6423daed34bf909bf7fdccc09f1ed246ba9e17aaf290df300.
//
// Solidity: event NodeCapacityIncreased(address indexed node, uint256 additionalStake, uint64 newCapacity)
func (_NodeStaking *NodeStakingFilterer) WatchNodeCapacityIncreased(opts *bind.WatchOpts, sink chan<- *NodeStakingNodeCapacityIncreased, node []common.Address) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _NodeStaking.contract.WatchLogs(opts, "NodeCapacityIncreased", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeStakingNodeCapacityIncreased)
				if err := _NodeStaking.contract.UnpackLog(event, "NodeCapacityIncreased", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNodeCapacityIncreased is a log parse operation binding the contract event 0x026b1e4d13deb2b6423daed34bf909bf7fdccc09f1ed246ba9e17aaf290df300.
//
// Solidity: event NodeCapacityIncreased(address indexed node, uint256 additionalStake, uint64 newCapacity)
func (_NodeStaking *NodeStakingFilterer) ParseNodeCapacityIncreased(log types.Log) (*NodeStakingNodeCapacityIncreased, error) {
	event := new(NodeStakingNodeCapacityIncreased)
	if err := _NodeStaking.contract.UnpackLog(event, "NodeCapacityIncreased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeStakingNodeSlashedIterator is returned from FilterNodeSlashed and is used to iterate over the raw logs and unpacked data for NodeSlashed events raised by the NodeStaking contract.
type NodeStakingNodeSlashedIterator struct {
	Event *NodeStakingNodeSlashed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NodeStakingNodeSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeStakingNodeSlashed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NodeStakingNodeSlashed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NodeStakingNodeSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeStakingNodeSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeStakingNodeSlashed represents a NodeSlashed event raised by the NodeStaking contract.
type NodeStakingNodeSlashed struct {
	Node            common.Address
	SlashAmount     *big.Int
	NewCapacity     uint64
	ForcedOrderExit bool
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNodeSlashed is a free log retrieval operation binding the contract event 0xd4b198736fc97da662d2e54a436006b5539a884842e0d264b12bbddcdb0566a8.
//
// Solidity: event NodeSlashed(address indexed node, uint256 slashAmount, uint64 newCapacity, bool forcedOrderExit)
func (_NodeStaking *NodeStakingFilterer) FilterNodeSlashed(opts *bind.FilterOpts, node []common.Address) (*NodeStakingNodeSlashedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _NodeStaking.contract.FilterLogs(opts, "NodeSlashed", nodeRule)
	if err != nil {
		return nil, err
	}
	return &NodeStakingNodeSlashedIterator{contract: _NodeStaking.contract, event: "NodeSlashed", logs: logs, sub: sub}, nil
}

// WatchNodeSlashed is a free log subscription operation binding the contract event 0xd4b198736fc97da662d2e54a436006b5539a884842e0d264b12bbddcdb0566a8.
//
// Solidity: event NodeSlashed(address indexed node, uint256 slashAmount, uint64 newCapacity, bool forcedOrderExit)
func (_NodeStaking *NodeStakingFilterer) WatchNodeSlashed(opts *bind.WatchOpts, sink chan<- *NodeStakingNodeSlashed, node []common.Address) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _NodeStaking.contract.WatchLogs(opts, "NodeSlashed", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeStakingNodeSlashed)
				if err := _NodeStaking.contract.UnpackLog(event, "NodeSlashed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNodeSlashed is a log parse operation binding the contract event 0xd4b198736fc97da662d2e54a436006b5539a884842e0d264b12bbddcdb0566a8.
//
// Solidity: event NodeSlashed(address indexed node, uint256 slashAmount, uint64 newCapacity, bool forcedOrderExit)
func (_NodeStaking *NodeStakingFilterer) ParseNodeSlashed(log types.Log) (*NodeStakingNodeSlashed, error) {
	event := new(NodeStakingNodeSlashed)
	if err := _NodeStaking.contract.UnpackLog(event, "NodeSlashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeStakingNodeStakedIterator is returned from FilterNodeStaked and is used to iterate over the raw logs and unpacked data for NodeStaked events raised by the NodeStaking contract.
type NodeStakingNodeStakedIterator struct {
	Event *NodeStakingNodeStaked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NodeStakingNodeStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeStakingNodeStaked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NodeStakingNodeStaked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NodeStakingNodeStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeStakingNodeStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeStakingNodeStaked represents a NodeStaked event raised by the NodeStaking contract.
type NodeStakingNodeStaked struct {
	Node     common.Address
	Stake    *big.Int
	Capacity uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNodeStaked is a free log retrieval operation binding the contract event 0x91f8b06ee53363aa910dd05b362fa962d55553b58626d79e136a24d2e26519aa.
//
// Solidity: event NodeStaked(address indexed node, uint256 stake, uint64 capacity)
func (_NodeStaking *NodeStakingFilterer) FilterNodeStaked(opts *bind.FilterOpts, node []common.Address) (*NodeStakingNodeStakedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _NodeStaking.contract.FilterLogs(opts, "NodeStaked", nodeRule)
	if err != nil {
		return nil, err
	}
	return &NodeStakingNodeStakedIterator{contract: _NodeStaking.contract, event: "NodeStaked", logs: logs, sub: sub}, nil
}

// WatchNodeStaked is a free log subscription operation binding the contract event 0x91f8b06ee53363aa910dd05b362fa962d55553b58626d79e136a24d2e26519aa.
//
// Solidity: event NodeStaked(address indexed node, uint256 stake, uint64 capacity)
func (_NodeStaking *NodeStakingFilterer) WatchNodeStaked(opts *bind.WatchOpts, sink chan<- *NodeStakingNodeStaked, node []common.Address) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _NodeStaking.contract.WatchLogs(opts, "NodeStaked", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeStakingNodeStaked)
				if err := _NodeStaking.contract.UnpackLog(event, "NodeStaked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNodeStaked is a log parse operation binding the contract event 0x91f8b06ee53363aa910dd05b362fa962d55553b58626d79e136a24d2e26519aa.
//
// Solidity: event NodeStaked(address indexed node, uint256 stake, uint64 capacity)
func (_NodeStaking *NodeStakingFilterer) ParseNodeStaked(log types.Log) (*NodeStakingNodeStaked, error) {
	event := new(NodeStakingNodeStaked)
	if err := _NodeStaking.contract.UnpackLog(event, "NodeStaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeStakingNodeUnstakedIterator is returned from FilterNodeUnstaked and is used to iterate over the raw logs and unpacked data for NodeUnstaked events raised by the NodeStaking contract.
type NodeStakingNodeUnstakedIterator struct {
	Event *NodeStakingNodeUnstaked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NodeStakingNodeUnstakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeStakingNodeUnstaked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NodeStakingNodeUnstaked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NodeStakingNodeUnstakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeStakingNodeUnstakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeStakingNodeUnstaked represents a NodeUnstaked event raised by the NodeStaking contract.
type NodeStakingNodeUnstaked struct {
	Node          common.Address
	StakeReturned *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNodeUnstaked is a free log retrieval operation binding the contract event 0x67722cd735ac17faf7e2ff2d57ac864a5a1f41405ee9113c16a413344ce31fbf.
//
// Solidity: event NodeUnstaked(address indexed node, uint256 stakeReturned)
func (_NodeStaking *NodeStakingFilterer) FilterNodeUnstaked(opts *bind.FilterOpts, node []common.Address) (*NodeStakingNodeUnstakedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _NodeStaking.contract.FilterLogs(opts, "NodeUnstaked", nodeRule)
	if err != nil {
		return nil, err
	}
	return &NodeStakingNodeUnstakedIterator{contract: _NodeStaking.contract, event: "NodeUnstaked", logs: logs, sub: sub}, nil
}

// WatchNodeUnstaked is a free log subscription operation binding the contract event 0x67722cd735ac17faf7e2ff2d57ac864a5a1f41405ee9113c16a413344ce31fbf.
//
// Solidity: event NodeUnstaked(address indexed node, uint256 stakeReturned)
func (_NodeStaking *NodeStakingFilterer) WatchNodeUnstaked(opts *bind.WatchOpts, sink chan<- *NodeStakingNodeUnstaked, node []common.Address) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _NodeStaking.contract.WatchLogs(opts, "NodeUnstaked", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeStakingNodeUnstaked)
				if err := _NodeStaking.contract.UnpackLog(event, "NodeUnstaked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNodeUnstaked is a log parse operation binding the contract event 0x67722cd735ac17faf7e2ff2d57ac864a5a1f41405ee9113c16a413344ce31fbf.
//
// Solidity: event NodeUnstaked(address indexed node, uint256 stakeReturned)
func (_NodeStaking *NodeStakingFilterer) ParseNodeUnstaked(log types.Log) (*NodeStakingNodeUnstaked, error) {
	event := new(NodeStakingNodeUnstaked)
	if err := _NodeStaking.contract.UnpackLog(event, "NodeUnstaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeStakingUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the NodeStaking contract.
type NodeStakingUpgradedIterator struct {
	Event *NodeStakingUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NodeStakingUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeStakingUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NodeStakingUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NodeStakingUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeStakingUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeStakingUpgraded represents a Upgraded event raised by the NodeStaking contract.
type NodeStakingUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_NodeStaking *NodeStakingFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*NodeStakingUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _NodeStaking.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &NodeStakingUpgradedIterator{contract: _NodeStaking.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_NodeStaking *NodeStakingFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *NodeStakingUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _NodeStaking.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeStakingUpgraded)
				if err := _NodeStaking.contract.UnpackLog(event, "Upgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_NodeStaking *NodeStakingFilterer) ParseUpgraded(log types.Log) (*NodeStakingUpgraded, error) {
	event := new(NodeStakingUpgraded)
	if err := _NodeStaking.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
