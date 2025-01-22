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

// NetworkStatsNodeInfo is an auto generated low-level Go binding around an user-defined struct.
type NetworkStatsNodeInfo struct {
	NodeAddress common.Address
	GPUModel    string
	VRAM        *big.Int
}

// NetworkStatsMetaData contains all meta data concerning the NetworkStats contract.
var NetworkStatsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"activeNodes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"availableNodes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"busyNodes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"getAllNodeInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"nodeAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"GPUModel\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"VRAM\",\"type\":\"uint256\"}],\"internalType\":\"structNetworkStats.NodeInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nodeAvailable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nodeAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"gpuModel\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"vRAM\",\"type\":\"uint256\"}],\"name\":\"nodeJoined\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nodeQuit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nodeTaskFinished\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nodeTaskStarted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nodeUnavailable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"queuedTasks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"runningTasks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"taskCreated\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"taskDequeue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"taskEnqueue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"taskFinished\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"taskStarted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalNodes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalTasks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nodeContract\",\"type\":\"address\"}],\"name\":\"updateNodeContractAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"taskContract\",\"type\":\"address\"}],\"name\":\"updateTaskContractAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// NetworkStatsABI is the input ABI used to generate the binding from.
// Deprecated: Use NetworkStatsMetaData.ABI instead.
var NetworkStatsABI = NetworkStatsMetaData.ABI

// NetworkStats is an auto generated Go binding around an Ethereum contract.
type NetworkStats struct {
	NetworkStatsCaller     // Read-only binding to the contract
	NetworkStatsTransactor // Write-only binding to the contract
	NetworkStatsFilterer   // Log filterer for contract events
}

// NetworkStatsCaller is an auto generated read-only Go binding around an Ethereum contract.
type NetworkStatsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NetworkStatsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NetworkStatsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NetworkStatsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NetworkStatsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NetworkStatsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NetworkStatsSession struct {
	Contract     *NetworkStats     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NetworkStatsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NetworkStatsCallerSession struct {
	Contract *NetworkStatsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// NetworkStatsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NetworkStatsTransactorSession struct {
	Contract     *NetworkStatsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// NetworkStatsRaw is an auto generated low-level Go binding around an Ethereum contract.
type NetworkStatsRaw struct {
	Contract *NetworkStats // Generic contract binding to access the raw methods on
}

// NetworkStatsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NetworkStatsCallerRaw struct {
	Contract *NetworkStatsCaller // Generic read-only contract binding to access the raw methods on
}

// NetworkStatsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NetworkStatsTransactorRaw struct {
	Contract *NetworkStatsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNetworkStats creates a new instance of NetworkStats, bound to a specific deployed contract.
func NewNetworkStats(address common.Address, backend bind.ContractBackend) (*NetworkStats, error) {
	contract, err := bindNetworkStats(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NetworkStats{NetworkStatsCaller: NetworkStatsCaller{contract: contract}, NetworkStatsTransactor: NetworkStatsTransactor{contract: contract}, NetworkStatsFilterer: NetworkStatsFilterer{contract: contract}}, nil
}

// NewNetworkStatsCaller creates a new read-only instance of NetworkStats, bound to a specific deployed contract.
func NewNetworkStatsCaller(address common.Address, caller bind.ContractCaller) (*NetworkStatsCaller, error) {
	contract, err := bindNetworkStats(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NetworkStatsCaller{contract: contract}, nil
}

// NewNetworkStatsTransactor creates a new write-only instance of NetworkStats, bound to a specific deployed contract.
func NewNetworkStatsTransactor(address common.Address, transactor bind.ContractTransactor) (*NetworkStatsTransactor, error) {
	contract, err := bindNetworkStats(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NetworkStatsTransactor{contract: contract}, nil
}

// NewNetworkStatsFilterer creates a new log filterer instance of NetworkStats, bound to a specific deployed contract.
func NewNetworkStatsFilterer(address common.Address, filterer bind.ContractFilterer) (*NetworkStatsFilterer, error) {
	contract, err := bindNetworkStats(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NetworkStatsFilterer{contract: contract}, nil
}

// bindNetworkStats binds a generic wrapper to an already deployed contract.
func bindNetworkStats(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NetworkStatsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NetworkStats *NetworkStatsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NetworkStats.Contract.NetworkStatsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NetworkStats *NetworkStatsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NetworkStats.Contract.NetworkStatsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NetworkStats *NetworkStatsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NetworkStats.Contract.NetworkStatsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NetworkStats *NetworkStatsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NetworkStats.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NetworkStats *NetworkStatsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NetworkStats.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NetworkStats *NetworkStatsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NetworkStats.Contract.contract.Transact(opts, method, params...)
}

// ActiveNodes is a free data retrieval call binding the contract method 0x07f19651.
//
// Solidity: function activeNodes() view returns(uint256)
func (_NetworkStats *NetworkStatsCaller) ActiveNodes(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NetworkStats.contract.Call(opts, &out, "activeNodes")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ActiveNodes is a free data retrieval call binding the contract method 0x07f19651.
//
// Solidity: function activeNodes() view returns(uint256)
func (_NetworkStats *NetworkStatsSession) ActiveNodes() (*big.Int, error) {
	return _NetworkStats.Contract.ActiveNodes(&_NetworkStats.CallOpts)
}

// ActiveNodes is a free data retrieval call binding the contract method 0x07f19651.
//
// Solidity: function activeNodes() view returns(uint256)
func (_NetworkStats *NetworkStatsCallerSession) ActiveNodes() (*big.Int, error) {
	return _NetworkStats.Contract.ActiveNodes(&_NetworkStats.CallOpts)
}

// AvailableNodes is a free data retrieval call binding the contract method 0x204401e7.
//
// Solidity: function availableNodes() view returns(uint256)
func (_NetworkStats *NetworkStatsCaller) AvailableNodes(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NetworkStats.contract.Call(opts, &out, "availableNodes")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AvailableNodes is a free data retrieval call binding the contract method 0x204401e7.
//
// Solidity: function availableNodes() view returns(uint256)
func (_NetworkStats *NetworkStatsSession) AvailableNodes() (*big.Int, error) {
	return _NetworkStats.Contract.AvailableNodes(&_NetworkStats.CallOpts)
}

// AvailableNodes is a free data retrieval call binding the contract method 0x204401e7.
//
// Solidity: function availableNodes() view returns(uint256)
func (_NetworkStats *NetworkStatsCallerSession) AvailableNodes() (*big.Int, error) {
	return _NetworkStats.Contract.AvailableNodes(&_NetworkStats.CallOpts)
}

// BusyNodes is a free data retrieval call binding the contract method 0x3d8c94ce.
//
// Solidity: function busyNodes() view returns(uint256)
func (_NetworkStats *NetworkStatsCaller) BusyNodes(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NetworkStats.contract.Call(opts, &out, "busyNodes")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BusyNodes is a free data retrieval call binding the contract method 0x3d8c94ce.
//
// Solidity: function busyNodes() view returns(uint256)
func (_NetworkStats *NetworkStatsSession) BusyNodes() (*big.Int, error) {
	return _NetworkStats.Contract.BusyNodes(&_NetworkStats.CallOpts)
}

// BusyNodes is a free data retrieval call binding the contract method 0x3d8c94ce.
//
// Solidity: function busyNodes() view returns(uint256)
func (_NetworkStats *NetworkStatsCallerSession) BusyNodes() (*big.Int, error) {
	return _NetworkStats.Contract.BusyNodes(&_NetworkStats.CallOpts)
}

// GetAllNodeInfo is a free data retrieval call binding the contract method 0xa2de2197.
//
// Solidity: function getAllNodeInfo(uint256 offset, uint256 length) view returns((address,string,uint256)[])
func (_NetworkStats *NetworkStatsCaller) GetAllNodeInfo(opts *bind.CallOpts, offset *big.Int, length *big.Int) ([]NetworkStatsNodeInfo, error) {
	var out []interface{}
	err := _NetworkStats.contract.Call(opts, &out, "getAllNodeInfo", offset, length)

	if err != nil {
		return *new([]NetworkStatsNodeInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]NetworkStatsNodeInfo)).(*[]NetworkStatsNodeInfo)

	return out0, err

}

// GetAllNodeInfo is a free data retrieval call binding the contract method 0xa2de2197.
//
// Solidity: function getAllNodeInfo(uint256 offset, uint256 length) view returns((address,string,uint256)[])
func (_NetworkStats *NetworkStatsSession) GetAllNodeInfo(offset *big.Int, length *big.Int) ([]NetworkStatsNodeInfo, error) {
	return _NetworkStats.Contract.GetAllNodeInfo(&_NetworkStats.CallOpts, offset, length)
}

// GetAllNodeInfo is a free data retrieval call binding the contract method 0xa2de2197.
//
// Solidity: function getAllNodeInfo(uint256 offset, uint256 length) view returns((address,string,uint256)[])
func (_NetworkStats *NetworkStatsCallerSession) GetAllNodeInfo(offset *big.Int, length *big.Int) ([]NetworkStatsNodeInfo, error) {
	return _NetworkStats.Contract.GetAllNodeInfo(&_NetworkStats.CallOpts, offset, length)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NetworkStats *NetworkStatsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NetworkStats.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NetworkStats *NetworkStatsSession) Owner() (common.Address, error) {
	return _NetworkStats.Contract.Owner(&_NetworkStats.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NetworkStats *NetworkStatsCallerSession) Owner() (common.Address, error) {
	return _NetworkStats.Contract.Owner(&_NetworkStats.CallOpts)
}

// QueuedTasks is a free data retrieval call binding the contract method 0xc7efb34f.
//
// Solidity: function queuedTasks() view returns(uint256)
func (_NetworkStats *NetworkStatsCaller) QueuedTasks(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NetworkStats.contract.Call(opts, &out, "queuedTasks")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QueuedTasks is a free data retrieval call binding the contract method 0xc7efb34f.
//
// Solidity: function queuedTasks() view returns(uint256)
func (_NetworkStats *NetworkStatsSession) QueuedTasks() (*big.Int, error) {
	return _NetworkStats.Contract.QueuedTasks(&_NetworkStats.CallOpts)
}

// QueuedTasks is a free data retrieval call binding the contract method 0xc7efb34f.
//
// Solidity: function queuedTasks() view returns(uint256)
func (_NetworkStats *NetworkStatsCallerSession) QueuedTasks() (*big.Int, error) {
	return _NetworkStats.Contract.QueuedTasks(&_NetworkStats.CallOpts)
}

// RunningTasks is a free data retrieval call binding the contract method 0x399fda5c.
//
// Solidity: function runningTasks() view returns(uint256)
func (_NetworkStats *NetworkStatsCaller) RunningTasks(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NetworkStats.contract.Call(opts, &out, "runningTasks")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RunningTasks is a free data retrieval call binding the contract method 0x399fda5c.
//
// Solidity: function runningTasks() view returns(uint256)
func (_NetworkStats *NetworkStatsSession) RunningTasks() (*big.Int, error) {
	return _NetworkStats.Contract.RunningTasks(&_NetworkStats.CallOpts)
}

// RunningTasks is a free data retrieval call binding the contract method 0x399fda5c.
//
// Solidity: function runningTasks() view returns(uint256)
func (_NetworkStats *NetworkStatsCallerSession) RunningTasks() (*big.Int, error) {
	return _NetworkStats.Contract.RunningTasks(&_NetworkStats.CallOpts)
}

// TotalNodes is a free data retrieval call binding the contract method 0x9592d424.
//
// Solidity: function totalNodes() view returns(uint256)
func (_NetworkStats *NetworkStatsCaller) TotalNodes(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NetworkStats.contract.Call(opts, &out, "totalNodes")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalNodes is a free data retrieval call binding the contract method 0x9592d424.
//
// Solidity: function totalNodes() view returns(uint256)
func (_NetworkStats *NetworkStatsSession) TotalNodes() (*big.Int, error) {
	return _NetworkStats.Contract.TotalNodes(&_NetworkStats.CallOpts)
}

// TotalNodes is a free data retrieval call binding the contract method 0x9592d424.
//
// Solidity: function totalNodes() view returns(uint256)
func (_NetworkStats *NetworkStatsCallerSession) TotalNodes() (*big.Int, error) {
	return _NetworkStats.Contract.TotalNodes(&_NetworkStats.CallOpts)
}

// TotalTasks is a free data retrieval call binding the contract method 0xd22c81e5.
//
// Solidity: function totalTasks() view returns(uint256)
func (_NetworkStats *NetworkStatsCaller) TotalTasks(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NetworkStats.contract.Call(opts, &out, "totalTasks")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalTasks is a free data retrieval call binding the contract method 0xd22c81e5.
//
// Solidity: function totalTasks() view returns(uint256)
func (_NetworkStats *NetworkStatsSession) TotalTasks() (*big.Int, error) {
	return _NetworkStats.Contract.TotalTasks(&_NetworkStats.CallOpts)
}

// TotalTasks is a free data retrieval call binding the contract method 0xd22c81e5.
//
// Solidity: function totalTasks() view returns(uint256)
func (_NetworkStats *NetworkStatsCallerSession) TotalTasks() (*big.Int, error) {
	return _NetworkStats.Contract.TotalTasks(&_NetworkStats.CallOpts)
}

// NodeAvailable is a paid mutator transaction binding the contract method 0x601d2a51.
//
// Solidity: function nodeAvailable() returns()
func (_NetworkStats *NetworkStatsTransactor) NodeAvailable(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NetworkStats.contract.Transact(opts, "nodeAvailable")
}

// NodeAvailable is a paid mutator transaction binding the contract method 0x601d2a51.
//
// Solidity: function nodeAvailable() returns()
func (_NetworkStats *NetworkStatsSession) NodeAvailable() (*types.Transaction, error) {
	return _NetworkStats.Contract.NodeAvailable(&_NetworkStats.TransactOpts)
}

// NodeAvailable is a paid mutator transaction binding the contract method 0x601d2a51.
//
// Solidity: function nodeAvailable() returns()
func (_NetworkStats *NetworkStatsTransactorSession) NodeAvailable() (*types.Transaction, error) {
	return _NetworkStats.Contract.NodeAvailable(&_NetworkStats.TransactOpts)
}

// NodeJoined is a paid mutator transaction binding the contract method 0xa52be2bd.
//
// Solidity: function nodeJoined(address nodeAddress, string gpuModel, uint256 vRAM) returns()
func (_NetworkStats *NetworkStatsTransactor) NodeJoined(opts *bind.TransactOpts, nodeAddress common.Address, gpuModel string, vRAM *big.Int) (*types.Transaction, error) {
	return _NetworkStats.contract.Transact(opts, "nodeJoined", nodeAddress, gpuModel, vRAM)
}

// NodeJoined is a paid mutator transaction binding the contract method 0xa52be2bd.
//
// Solidity: function nodeJoined(address nodeAddress, string gpuModel, uint256 vRAM) returns()
func (_NetworkStats *NetworkStatsSession) NodeJoined(nodeAddress common.Address, gpuModel string, vRAM *big.Int) (*types.Transaction, error) {
	return _NetworkStats.Contract.NodeJoined(&_NetworkStats.TransactOpts, nodeAddress, gpuModel, vRAM)
}

// NodeJoined is a paid mutator transaction binding the contract method 0xa52be2bd.
//
// Solidity: function nodeJoined(address nodeAddress, string gpuModel, uint256 vRAM) returns()
func (_NetworkStats *NetworkStatsTransactorSession) NodeJoined(nodeAddress common.Address, gpuModel string, vRAM *big.Int) (*types.Transaction, error) {
	return _NetworkStats.Contract.NodeJoined(&_NetworkStats.TransactOpts, nodeAddress, gpuModel, vRAM)
}

// NodeQuit is a paid mutator transaction binding the contract method 0x7a2af56e.
//
// Solidity: function nodeQuit() returns()
func (_NetworkStats *NetworkStatsTransactor) NodeQuit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NetworkStats.contract.Transact(opts, "nodeQuit")
}

// NodeQuit is a paid mutator transaction binding the contract method 0x7a2af56e.
//
// Solidity: function nodeQuit() returns()
func (_NetworkStats *NetworkStatsSession) NodeQuit() (*types.Transaction, error) {
	return _NetworkStats.Contract.NodeQuit(&_NetworkStats.TransactOpts)
}

// NodeQuit is a paid mutator transaction binding the contract method 0x7a2af56e.
//
// Solidity: function nodeQuit() returns()
func (_NetworkStats *NetworkStatsTransactorSession) NodeQuit() (*types.Transaction, error) {
	return _NetworkStats.Contract.NodeQuit(&_NetworkStats.TransactOpts)
}

// NodeTaskFinished is a paid mutator transaction binding the contract method 0xa5cc3442.
//
// Solidity: function nodeTaskFinished() returns()
func (_NetworkStats *NetworkStatsTransactor) NodeTaskFinished(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NetworkStats.contract.Transact(opts, "nodeTaskFinished")
}

// NodeTaskFinished is a paid mutator transaction binding the contract method 0xa5cc3442.
//
// Solidity: function nodeTaskFinished() returns()
func (_NetworkStats *NetworkStatsSession) NodeTaskFinished() (*types.Transaction, error) {
	return _NetworkStats.Contract.NodeTaskFinished(&_NetworkStats.TransactOpts)
}

// NodeTaskFinished is a paid mutator transaction binding the contract method 0xa5cc3442.
//
// Solidity: function nodeTaskFinished() returns()
func (_NetworkStats *NetworkStatsTransactorSession) NodeTaskFinished() (*types.Transaction, error) {
	return _NetworkStats.Contract.NodeTaskFinished(&_NetworkStats.TransactOpts)
}

// NodeTaskStarted is a paid mutator transaction binding the contract method 0x03659a66.
//
// Solidity: function nodeTaskStarted() returns()
func (_NetworkStats *NetworkStatsTransactor) NodeTaskStarted(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NetworkStats.contract.Transact(opts, "nodeTaskStarted")
}

// NodeTaskStarted is a paid mutator transaction binding the contract method 0x03659a66.
//
// Solidity: function nodeTaskStarted() returns()
func (_NetworkStats *NetworkStatsSession) NodeTaskStarted() (*types.Transaction, error) {
	return _NetworkStats.Contract.NodeTaskStarted(&_NetworkStats.TransactOpts)
}

// NodeTaskStarted is a paid mutator transaction binding the contract method 0x03659a66.
//
// Solidity: function nodeTaskStarted() returns()
func (_NetworkStats *NetworkStatsTransactorSession) NodeTaskStarted() (*types.Transaction, error) {
	return _NetworkStats.Contract.NodeTaskStarted(&_NetworkStats.TransactOpts)
}

// NodeUnavailable is a paid mutator transaction binding the contract method 0x9b1ee54d.
//
// Solidity: function nodeUnavailable() returns()
func (_NetworkStats *NetworkStatsTransactor) NodeUnavailable(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NetworkStats.contract.Transact(opts, "nodeUnavailable")
}

// NodeUnavailable is a paid mutator transaction binding the contract method 0x9b1ee54d.
//
// Solidity: function nodeUnavailable() returns()
func (_NetworkStats *NetworkStatsSession) NodeUnavailable() (*types.Transaction, error) {
	return _NetworkStats.Contract.NodeUnavailable(&_NetworkStats.TransactOpts)
}

// NodeUnavailable is a paid mutator transaction binding the contract method 0x9b1ee54d.
//
// Solidity: function nodeUnavailable() returns()
func (_NetworkStats *NetworkStatsTransactorSession) NodeUnavailable() (*types.Transaction, error) {
	return _NetworkStats.Contract.NodeUnavailable(&_NetworkStats.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NetworkStats *NetworkStatsTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NetworkStats.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NetworkStats *NetworkStatsSession) RenounceOwnership() (*types.Transaction, error) {
	return _NetworkStats.Contract.RenounceOwnership(&_NetworkStats.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NetworkStats *NetworkStatsTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _NetworkStats.Contract.RenounceOwnership(&_NetworkStats.TransactOpts)
}

// TaskCreated is a paid mutator transaction binding the contract method 0xd0e08cca.
//
// Solidity: function taskCreated() returns()
func (_NetworkStats *NetworkStatsTransactor) TaskCreated(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NetworkStats.contract.Transact(opts, "taskCreated")
}

// TaskCreated is a paid mutator transaction binding the contract method 0xd0e08cca.
//
// Solidity: function taskCreated() returns()
func (_NetworkStats *NetworkStatsSession) TaskCreated() (*types.Transaction, error) {
	return _NetworkStats.Contract.TaskCreated(&_NetworkStats.TransactOpts)
}

// TaskCreated is a paid mutator transaction binding the contract method 0xd0e08cca.
//
// Solidity: function taskCreated() returns()
func (_NetworkStats *NetworkStatsTransactorSession) TaskCreated() (*types.Transaction, error) {
	return _NetworkStats.Contract.TaskCreated(&_NetworkStats.TransactOpts)
}

// TaskDequeue is a paid mutator transaction binding the contract method 0xd513093c.
//
// Solidity: function taskDequeue() returns()
func (_NetworkStats *NetworkStatsTransactor) TaskDequeue(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NetworkStats.contract.Transact(opts, "taskDequeue")
}

// TaskDequeue is a paid mutator transaction binding the contract method 0xd513093c.
//
// Solidity: function taskDequeue() returns()
func (_NetworkStats *NetworkStatsSession) TaskDequeue() (*types.Transaction, error) {
	return _NetworkStats.Contract.TaskDequeue(&_NetworkStats.TransactOpts)
}

// TaskDequeue is a paid mutator transaction binding the contract method 0xd513093c.
//
// Solidity: function taskDequeue() returns()
func (_NetworkStats *NetworkStatsTransactorSession) TaskDequeue() (*types.Transaction, error) {
	return _NetworkStats.Contract.TaskDequeue(&_NetworkStats.TransactOpts)
}

// TaskEnqueue is a paid mutator transaction binding the contract method 0x7e13650a.
//
// Solidity: function taskEnqueue() returns()
func (_NetworkStats *NetworkStatsTransactor) TaskEnqueue(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NetworkStats.contract.Transact(opts, "taskEnqueue")
}

// TaskEnqueue is a paid mutator transaction binding the contract method 0x7e13650a.
//
// Solidity: function taskEnqueue() returns()
func (_NetworkStats *NetworkStatsSession) TaskEnqueue() (*types.Transaction, error) {
	return _NetworkStats.Contract.TaskEnqueue(&_NetworkStats.TransactOpts)
}

// TaskEnqueue is a paid mutator transaction binding the contract method 0x7e13650a.
//
// Solidity: function taskEnqueue() returns()
func (_NetworkStats *NetworkStatsTransactorSession) TaskEnqueue() (*types.Transaction, error) {
	return _NetworkStats.Contract.TaskEnqueue(&_NetworkStats.TransactOpts)
}

// TaskFinished is a paid mutator transaction binding the contract method 0xb1f7748f.
//
// Solidity: function taskFinished() returns()
func (_NetworkStats *NetworkStatsTransactor) TaskFinished(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NetworkStats.contract.Transact(opts, "taskFinished")
}

// TaskFinished is a paid mutator transaction binding the contract method 0xb1f7748f.
//
// Solidity: function taskFinished() returns()
func (_NetworkStats *NetworkStatsSession) TaskFinished() (*types.Transaction, error) {
	return _NetworkStats.Contract.TaskFinished(&_NetworkStats.TransactOpts)
}

// TaskFinished is a paid mutator transaction binding the contract method 0xb1f7748f.
//
// Solidity: function taskFinished() returns()
func (_NetworkStats *NetworkStatsTransactorSession) TaskFinished() (*types.Transaction, error) {
	return _NetworkStats.Contract.TaskFinished(&_NetworkStats.TransactOpts)
}

// TaskStarted is a paid mutator transaction binding the contract method 0x1e4a9b0c.
//
// Solidity: function taskStarted() returns()
func (_NetworkStats *NetworkStatsTransactor) TaskStarted(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NetworkStats.contract.Transact(opts, "taskStarted")
}

// TaskStarted is a paid mutator transaction binding the contract method 0x1e4a9b0c.
//
// Solidity: function taskStarted() returns()
func (_NetworkStats *NetworkStatsSession) TaskStarted() (*types.Transaction, error) {
	return _NetworkStats.Contract.TaskStarted(&_NetworkStats.TransactOpts)
}

// TaskStarted is a paid mutator transaction binding the contract method 0x1e4a9b0c.
//
// Solidity: function taskStarted() returns()
func (_NetworkStats *NetworkStatsTransactorSession) TaskStarted() (*types.Transaction, error) {
	return _NetworkStats.Contract.TaskStarted(&_NetworkStats.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NetworkStats *NetworkStatsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _NetworkStats.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NetworkStats *NetworkStatsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NetworkStats.Contract.TransferOwnership(&_NetworkStats.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NetworkStats *NetworkStatsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NetworkStats.Contract.TransferOwnership(&_NetworkStats.TransactOpts, newOwner)
}

// UpdateNodeContractAddress is a paid mutator transaction binding the contract method 0xc6c8b6f7.
//
// Solidity: function updateNodeContractAddress(address nodeContract) returns()
func (_NetworkStats *NetworkStatsTransactor) UpdateNodeContractAddress(opts *bind.TransactOpts, nodeContract common.Address) (*types.Transaction, error) {
	return _NetworkStats.contract.Transact(opts, "updateNodeContractAddress", nodeContract)
}

// UpdateNodeContractAddress is a paid mutator transaction binding the contract method 0xc6c8b6f7.
//
// Solidity: function updateNodeContractAddress(address nodeContract) returns()
func (_NetworkStats *NetworkStatsSession) UpdateNodeContractAddress(nodeContract common.Address) (*types.Transaction, error) {
	return _NetworkStats.Contract.UpdateNodeContractAddress(&_NetworkStats.TransactOpts, nodeContract)
}

// UpdateNodeContractAddress is a paid mutator transaction binding the contract method 0xc6c8b6f7.
//
// Solidity: function updateNodeContractAddress(address nodeContract) returns()
func (_NetworkStats *NetworkStatsTransactorSession) UpdateNodeContractAddress(nodeContract common.Address) (*types.Transaction, error) {
	return _NetworkStats.Contract.UpdateNodeContractAddress(&_NetworkStats.TransactOpts, nodeContract)
}

// UpdateTaskContractAddress is a paid mutator transaction binding the contract method 0x42145230.
//
// Solidity: function updateTaskContractAddress(address taskContract) returns()
func (_NetworkStats *NetworkStatsTransactor) UpdateTaskContractAddress(opts *bind.TransactOpts, taskContract common.Address) (*types.Transaction, error) {
	return _NetworkStats.contract.Transact(opts, "updateTaskContractAddress", taskContract)
}

// UpdateTaskContractAddress is a paid mutator transaction binding the contract method 0x42145230.
//
// Solidity: function updateTaskContractAddress(address taskContract) returns()
func (_NetworkStats *NetworkStatsSession) UpdateTaskContractAddress(taskContract common.Address) (*types.Transaction, error) {
	return _NetworkStats.Contract.UpdateTaskContractAddress(&_NetworkStats.TransactOpts, taskContract)
}

// UpdateTaskContractAddress is a paid mutator transaction binding the contract method 0x42145230.
//
// Solidity: function updateTaskContractAddress(address taskContract) returns()
func (_NetworkStats *NetworkStatsTransactorSession) UpdateTaskContractAddress(taskContract common.Address) (*types.Transaction, error) {
	return _NetworkStats.Contract.UpdateTaskContractAddress(&_NetworkStats.TransactOpts, taskContract)
}

// NetworkStatsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the NetworkStats contract.
type NetworkStatsOwnershipTransferredIterator struct {
	Event *NetworkStatsOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *NetworkStatsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NetworkStatsOwnershipTransferred)
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
		it.Event = new(NetworkStatsOwnershipTransferred)
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
func (it *NetworkStatsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NetworkStatsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NetworkStatsOwnershipTransferred represents a OwnershipTransferred event raised by the NetworkStats contract.
type NetworkStatsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NetworkStats *NetworkStatsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*NetworkStatsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NetworkStats.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NetworkStatsOwnershipTransferredIterator{contract: _NetworkStats.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NetworkStats *NetworkStatsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NetworkStatsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NetworkStats.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NetworkStatsOwnershipTransferred)
				if err := _NetworkStats.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NetworkStats *NetworkStatsFilterer) ParseOwnershipTransferred(log types.Log) (*NetworkStatsOwnershipTransferred, error) {
	event := new(NetworkStatsOwnershipTransferred)
	if err := _NetworkStats.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
