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

// QOSMetaData contains all meta data concerning the QOS contract.
var QOSMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nodeAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"addTaskScore\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nodeAddress\",\"type\":\"address\"}],\"name\":\"finishTask\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nodeAddress\",\"type\":\"address\"}],\"name\":\"getCurrentTaskScore\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nodeAddress\",\"type\":\"address\"}],\"name\":\"getRecentTaskCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nodeAddress\",\"type\":\"address\"}],\"name\":\"getRecentTaskScore\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nodeAddress\",\"type\":\"address\"}],\"name\":\"getTaskCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nodeAddress\",\"type\":\"address\"}],\"name\":\"getTaskScore\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTaskScoreLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nodeAddress\",\"type\":\"address\"}],\"name\":\"kickout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nodeAddress\",\"type\":\"address\"}],\"name\":\"punish\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nodeAddress\",\"type\":\"address\"}],\"name\":\"shouldKickOut\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nodeAddress\",\"type\":\"address\"}],\"name\":\"startTask\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"}],\"name\":\"updateKickoutThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nodeContract\",\"type\":\"address\"}],\"name\":\"updateNodeContractAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"taskContract\",\"type\":\"address\"}],\"name\":\"updateTaskContractAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// QOSABI is the input ABI used to generate the binding from.
// Deprecated: Use QOSMetaData.ABI instead.
var QOSABI = QOSMetaData.ABI

// QOS is an auto generated Go binding around an Ethereum contract.
type QOS struct {
	QOSCaller     // Read-only binding to the contract
	QOSTransactor // Write-only binding to the contract
	QOSFilterer   // Log filterer for contract events
}

// QOSCaller is an auto generated read-only Go binding around an Ethereum contract.
type QOSCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QOSTransactor is an auto generated write-only Go binding around an Ethereum contract.
type QOSTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QOSFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type QOSFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QOSSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type QOSSession struct {
	Contract     *QOS              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// QOSCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type QOSCallerSession struct {
	Contract *QOSCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// QOSTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type QOSTransactorSession struct {
	Contract     *QOSTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// QOSRaw is an auto generated low-level Go binding around an Ethereum contract.
type QOSRaw struct {
	Contract *QOS // Generic contract binding to access the raw methods on
}

// QOSCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type QOSCallerRaw struct {
	Contract *QOSCaller // Generic read-only contract binding to access the raw methods on
}

// QOSTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type QOSTransactorRaw struct {
	Contract *QOSTransactor // Generic write-only contract binding to access the raw methods on
}

// NewQOS creates a new instance of QOS, bound to a specific deployed contract.
func NewQOS(address common.Address, backend bind.ContractBackend) (*QOS, error) {
	contract, err := bindQOS(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &QOS{QOSCaller: QOSCaller{contract: contract}, QOSTransactor: QOSTransactor{contract: contract}, QOSFilterer: QOSFilterer{contract: contract}}, nil
}

// NewQOSCaller creates a new read-only instance of QOS, bound to a specific deployed contract.
func NewQOSCaller(address common.Address, caller bind.ContractCaller) (*QOSCaller, error) {
	contract, err := bindQOS(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &QOSCaller{contract: contract}, nil
}

// NewQOSTransactor creates a new write-only instance of QOS, bound to a specific deployed contract.
func NewQOSTransactor(address common.Address, transactor bind.ContractTransactor) (*QOSTransactor, error) {
	contract, err := bindQOS(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &QOSTransactor{contract: contract}, nil
}

// NewQOSFilterer creates a new log filterer instance of QOS, bound to a specific deployed contract.
func NewQOSFilterer(address common.Address, filterer bind.ContractFilterer) (*QOSFilterer, error) {
	contract, err := bindQOS(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &QOSFilterer{contract: contract}, nil
}

// bindQOS binds a generic wrapper to an already deployed contract.
func bindQOS(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := QOSMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_QOS *QOSRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _QOS.Contract.QOSCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_QOS *QOSRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QOS.Contract.QOSTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_QOS *QOSRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _QOS.Contract.QOSTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_QOS *QOSCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _QOS.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_QOS *QOSTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QOS.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_QOS *QOSTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _QOS.Contract.contract.Transact(opts, method, params...)
}

// GetCurrentTaskScore is a free data retrieval call binding the contract method 0xb055757d.
//
// Solidity: function getCurrentTaskScore(address nodeAddress) view returns(uint256)
func (_QOS *QOSCaller) GetCurrentTaskScore(opts *bind.CallOpts, nodeAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _QOS.contract.Call(opts, &out, "getCurrentTaskScore", nodeAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentTaskScore is a free data retrieval call binding the contract method 0xb055757d.
//
// Solidity: function getCurrentTaskScore(address nodeAddress) view returns(uint256)
func (_QOS *QOSSession) GetCurrentTaskScore(nodeAddress common.Address) (*big.Int, error) {
	return _QOS.Contract.GetCurrentTaskScore(&_QOS.CallOpts, nodeAddress)
}

// GetCurrentTaskScore is a free data retrieval call binding the contract method 0xb055757d.
//
// Solidity: function getCurrentTaskScore(address nodeAddress) view returns(uint256)
func (_QOS *QOSCallerSession) GetCurrentTaskScore(nodeAddress common.Address) (*big.Int, error) {
	return _QOS.Contract.GetCurrentTaskScore(&_QOS.CallOpts, nodeAddress)
}

// GetRecentTaskCount is a free data retrieval call binding the contract method 0x1b0883a3.
//
// Solidity: function getRecentTaskCount(address nodeAddress) view returns(uint256)
func (_QOS *QOSCaller) GetRecentTaskCount(opts *bind.CallOpts, nodeAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _QOS.contract.Call(opts, &out, "getRecentTaskCount", nodeAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRecentTaskCount is a free data retrieval call binding the contract method 0x1b0883a3.
//
// Solidity: function getRecentTaskCount(address nodeAddress) view returns(uint256)
func (_QOS *QOSSession) GetRecentTaskCount(nodeAddress common.Address) (*big.Int, error) {
	return _QOS.Contract.GetRecentTaskCount(&_QOS.CallOpts, nodeAddress)
}

// GetRecentTaskCount is a free data retrieval call binding the contract method 0x1b0883a3.
//
// Solidity: function getRecentTaskCount(address nodeAddress) view returns(uint256)
func (_QOS *QOSCallerSession) GetRecentTaskCount(nodeAddress common.Address) (*big.Int, error) {
	return _QOS.Contract.GetRecentTaskCount(&_QOS.CallOpts, nodeAddress)
}

// GetRecentTaskScore is a free data retrieval call binding the contract method 0x886673a3.
//
// Solidity: function getRecentTaskScore(address nodeAddress) view returns(uint256)
func (_QOS *QOSCaller) GetRecentTaskScore(opts *bind.CallOpts, nodeAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _QOS.contract.Call(opts, &out, "getRecentTaskScore", nodeAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRecentTaskScore is a free data retrieval call binding the contract method 0x886673a3.
//
// Solidity: function getRecentTaskScore(address nodeAddress) view returns(uint256)
func (_QOS *QOSSession) GetRecentTaskScore(nodeAddress common.Address) (*big.Int, error) {
	return _QOS.Contract.GetRecentTaskScore(&_QOS.CallOpts, nodeAddress)
}

// GetRecentTaskScore is a free data retrieval call binding the contract method 0x886673a3.
//
// Solidity: function getRecentTaskScore(address nodeAddress) view returns(uint256)
func (_QOS *QOSCallerSession) GetRecentTaskScore(nodeAddress common.Address) (*big.Int, error) {
	return _QOS.Contract.GetRecentTaskScore(&_QOS.CallOpts, nodeAddress)
}

// GetTaskCount is a free data retrieval call binding the contract method 0xc64ecb8f.
//
// Solidity: function getTaskCount(address nodeAddress) view returns(uint256)
func (_QOS *QOSCaller) GetTaskCount(opts *bind.CallOpts, nodeAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _QOS.contract.Call(opts, &out, "getTaskCount", nodeAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTaskCount is a free data retrieval call binding the contract method 0xc64ecb8f.
//
// Solidity: function getTaskCount(address nodeAddress) view returns(uint256)
func (_QOS *QOSSession) GetTaskCount(nodeAddress common.Address) (*big.Int, error) {
	return _QOS.Contract.GetTaskCount(&_QOS.CallOpts, nodeAddress)
}

// GetTaskCount is a free data retrieval call binding the contract method 0xc64ecb8f.
//
// Solidity: function getTaskCount(address nodeAddress) view returns(uint256)
func (_QOS *QOSCallerSession) GetTaskCount(nodeAddress common.Address) (*big.Int, error) {
	return _QOS.Contract.GetTaskCount(&_QOS.CallOpts, nodeAddress)
}

// GetTaskScore is a free data retrieval call binding the contract method 0xa903e689.
//
// Solidity: function getTaskScore(address nodeAddress) view returns(uint256)
func (_QOS *QOSCaller) GetTaskScore(opts *bind.CallOpts, nodeAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _QOS.contract.Call(opts, &out, "getTaskScore", nodeAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTaskScore is a free data retrieval call binding the contract method 0xa903e689.
//
// Solidity: function getTaskScore(address nodeAddress) view returns(uint256)
func (_QOS *QOSSession) GetTaskScore(nodeAddress common.Address) (*big.Int, error) {
	return _QOS.Contract.GetTaskScore(&_QOS.CallOpts, nodeAddress)
}

// GetTaskScore is a free data retrieval call binding the contract method 0xa903e689.
//
// Solidity: function getTaskScore(address nodeAddress) view returns(uint256)
func (_QOS *QOSCallerSession) GetTaskScore(nodeAddress common.Address) (*big.Int, error) {
	return _QOS.Contract.GetTaskScore(&_QOS.CallOpts, nodeAddress)
}

// GetTaskScoreLimit is a free data retrieval call binding the contract method 0x46af6c7b.
//
// Solidity: function getTaskScoreLimit() view returns(uint256)
func (_QOS *QOSCaller) GetTaskScoreLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _QOS.contract.Call(opts, &out, "getTaskScoreLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTaskScoreLimit is a free data retrieval call binding the contract method 0x46af6c7b.
//
// Solidity: function getTaskScoreLimit() view returns(uint256)
func (_QOS *QOSSession) GetTaskScoreLimit() (*big.Int, error) {
	return _QOS.Contract.GetTaskScoreLimit(&_QOS.CallOpts)
}

// GetTaskScoreLimit is a free data retrieval call binding the contract method 0x46af6c7b.
//
// Solidity: function getTaskScoreLimit() view returns(uint256)
func (_QOS *QOSCallerSession) GetTaskScoreLimit() (*big.Int, error) {
	return _QOS.Contract.GetTaskScoreLimit(&_QOS.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_QOS *QOSCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _QOS.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_QOS *QOSSession) Owner() (common.Address, error) {
	return _QOS.Contract.Owner(&_QOS.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_QOS *QOSCallerSession) Owner() (common.Address, error) {
	return _QOS.Contract.Owner(&_QOS.CallOpts)
}

// ShouldKickOut is a free data retrieval call binding the contract method 0x739ecbc4.
//
// Solidity: function shouldKickOut(address nodeAddress) view returns(bool)
func (_QOS *QOSCaller) ShouldKickOut(opts *bind.CallOpts, nodeAddress common.Address) (bool, error) {
	var out []interface{}
	err := _QOS.contract.Call(opts, &out, "shouldKickOut", nodeAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ShouldKickOut is a free data retrieval call binding the contract method 0x739ecbc4.
//
// Solidity: function shouldKickOut(address nodeAddress) view returns(bool)
func (_QOS *QOSSession) ShouldKickOut(nodeAddress common.Address) (bool, error) {
	return _QOS.Contract.ShouldKickOut(&_QOS.CallOpts, nodeAddress)
}

// ShouldKickOut is a free data retrieval call binding the contract method 0x739ecbc4.
//
// Solidity: function shouldKickOut(address nodeAddress) view returns(bool)
func (_QOS *QOSCallerSession) ShouldKickOut(nodeAddress common.Address) (bool, error) {
	return _QOS.Contract.ShouldKickOut(&_QOS.CallOpts, nodeAddress)
}

// AddTaskScore is a paid mutator transaction binding the contract method 0x81d6d1a3.
//
// Solidity: function addTaskScore(address nodeAddress, uint256 i) returns()
func (_QOS *QOSTransactor) AddTaskScore(opts *bind.TransactOpts, nodeAddress common.Address, i *big.Int) (*types.Transaction, error) {
	return _QOS.contract.Transact(opts, "addTaskScore", nodeAddress, i)
}

// AddTaskScore is a paid mutator transaction binding the contract method 0x81d6d1a3.
//
// Solidity: function addTaskScore(address nodeAddress, uint256 i) returns()
func (_QOS *QOSSession) AddTaskScore(nodeAddress common.Address, i *big.Int) (*types.Transaction, error) {
	return _QOS.Contract.AddTaskScore(&_QOS.TransactOpts, nodeAddress, i)
}

// AddTaskScore is a paid mutator transaction binding the contract method 0x81d6d1a3.
//
// Solidity: function addTaskScore(address nodeAddress, uint256 i) returns()
func (_QOS *QOSTransactorSession) AddTaskScore(nodeAddress common.Address, i *big.Int) (*types.Transaction, error) {
	return _QOS.Contract.AddTaskScore(&_QOS.TransactOpts, nodeAddress, i)
}

// FinishTask is a paid mutator transaction binding the contract method 0x3fc0f48b.
//
// Solidity: function finishTask(address nodeAddress) returns()
func (_QOS *QOSTransactor) FinishTask(opts *bind.TransactOpts, nodeAddress common.Address) (*types.Transaction, error) {
	return _QOS.contract.Transact(opts, "finishTask", nodeAddress)
}

// FinishTask is a paid mutator transaction binding the contract method 0x3fc0f48b.
//
// Solidity: function finishTask(address nodeAddress) returns()
func (_QOS *QOSSession) FinishTask(nodeAddress common.Address) (*types.Transaction, error) {
	return _QOS.Contract.FinishTask(&_QOS.TransactOpts, nodeAddress)
}

// FinishTask is a paid mutator transaction binding the contract method 0x3fc0f48b.
//
// Solidity: function finishTask(address nodeAddress) returns()
func (_QOS *QOSTransactorSession) FinishTask(nodeAddress common.Address) (*types.Transaction, error) {
	return _QOS.Contract.FinishTask(&_QOS.TransactOpts, nodeAddress)
}

// Kickout is a paid mutator transaction binding the contract method 0xc05a0469.
//
// Solidity: function kickout(address nodeAddress) returns()
func (_QOS *QOSTransactor) Kickout(opts *bind.TransactOpts, nodeAddress common.Address) (*types.Transaction, error) {
	return _QOS.contract.Transact(opts, "kickout", nodeAddress)
}

// Kickout is a paid mutator transaction binding the contract method 0xc05a0469.
//
// Solidity: function kickout(address nodeAddress) returns()
func (_QOS *QOSSession) Kickout(nodeAddress common.Address) (*types.Transaction, error) {
	return _QOS.Contract.Kickout(&_QOS.TransactOpts, nodeAddress)
}

// Kickout is a paid mutator transaction binding the contract method 0xc05a0469.
//
// Solidity: function kickout(address nodeAddress) returns()
func (_QOS *QOSTransactorSession) Kickout(nodeAddress common.Address) (*types.Transaction, error) {
	return _QOS.Contract.Kickout(&_QOS.TransactOpts, nodeAddress)
}

// Punish is a paid mutator transaction binding the contract method 0xea7221a1.
//
// Solidity: function punish(address nodeAddress) returns()
func (_QOS *QOSTransactor) Punish(opts *bind.TransactOpts, nodeAddress common.Address) (*types.Transaction, error) {
	return _QOS.contract.Transact(opts, "punish", nodeAddress)
}

// Punish is a paid mutator transaction binding the contract method 0xea7221a1.
//
// Solidity: function punish(address nodeAddress) returns()
func (_QOS *QOSSession) Punish(nodeAddress common.Address) (*types.Transaction, error) {
	return _QOS.Contract.Punish(&_QOS.TransactOpts, nodeAddress)
}

// Punish is a paid mutator transaction binding the contract method 0xea7221a1.
//
// Solidity: function punish(address nodeAddress) returns()
func (_QOS *QOSTransactorSession) Punish(nodeAddress common.Address) (*types.Transaction, error) {
	return _QOS.Contract.Punish(&_QOS.TransactOpts, nodeAddress)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_QOS *QOSTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QOS.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_QOS *QOSSession) RenounceOwnership() (*types.Transaction, error) {
	return _QOS.Contract.RenounceOwnership(&_QOS.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_QOS *QOSTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _QOS.Contract.RenounceOwnership(&_QOS.TransactOpts)
}

// StartTask is a paid mutator transaction binding the contract method 0x5f51c765.
//
// Solidity: function startTask(address nodeAddress) returns()
func (_QOS *QOSTransactor) StartTask(opts *bind.TransactOpts, nodeAddress common.Address) (*types.Transaction, error) {
	return _QOS.contract.Transact(opts, "startTask", nodeAddress)
}

// StartTask is a paid mutator transaction binding the contract method 0x5f51c765.
//
// Solidity: function startTask(address nodeAddress) returns()
func (_QOS *QOSSession) StartTask(nodeAddress common.Address) (*types.Transaction, error) {
	return _QOS.Contract.StartTask(&_QOS.TransactOpts, nodeAddress)
}

// StartTask is a paid mutator transaction binding the contract method 0x5f51c765.
//
// Solidity: function startTask(address nodeAddress) returns()
func (_QOS *QOSTransactorSession) StartTask(nodeAddress common.Address) (*types.Transaction, error) {
	return _QOS.Contract.StartTask(&_QOS.TransactOpts, nodeAddress)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_QOS *QOSTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _QOS.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_QOS *QOSSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _QOS.Contract.TransferOwnership(&_QOS.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_QOS *QOSTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _QOS.Contract.TransferOwnership(&_QOS.TransactOpts, newOwner)
}

// UpdateKickoutThreshold is a paid mutator transaction binding the contract method 0xa5a4e94a.
//
// Solidity: function updateKickoutThreshold(uint256 threshold) returns()
func (_QOS *QOSTransactor) UpdateKickoutThreshold(opts *bind.TransactOpts, threshold *big.Int) (*types.Transaction, error) {
	return _QOS.contract.Transact(opts, "updateKickoutThreshold", threshold)
}

// UpdateKickoutThreshold is a paid mutator transaction binding the contract method 0xa5a4e94a.
//
// Solidity: function updateKickoutThreshold(uint256 threshold) returns()
func (_QOS *QOSSession) UpdateKickoutThreshold(threshold *big.Int) (*types.Transaction, error) {
	return _QOS.Contract.UpdateKickoutThreshold(&_QOS.TransactOpts, threshold)
}

// UpdateKickoutThreshold is a paid mutator transaction binding the contract method 0xa5a4e94a.
//
// Solidity: function updateKickoutThreshold(uint256 threshold) returns()
func (_QOS *QOSTransactorSession) UpdateKickoutThreshold(threshold *big.Int) (*types.Transaction, error) {
	return _QOS.Contract.UpdateKickoutThreshold(&_QOS.TransactOpts, threshold)
}

// UpdateNodeContractAddress is a paid mutator transaction binding the contract method 0xc6c8b6f7.
//
// Solidity: function updateNodeContractAddress(address nodeContract) returns()
func (_QOS *QOSTransactor) UpdateNodeContractAddress(opts *bind.TransactOpts, nodeContract common.Address) (*types.Transaction, error) {
	return _QOS.contract.Transact(opts, "updateNodeContractAddress", nodeContract)
}

// UpdateNodeContractAddress is a paid mutator transaction binding the contract method 0xc6c8b6f7.
//
// Solidity: function updateNodeContractAddress(address nodeContract) returns()
func (_QOS *QOSSession) UpdateNodeContractAddress(nodeContract common.Address) (*types.Transaction, error) {
	return _QOS.Contract.UpdateNodeContractAddress(&_QOS.TransactOpts, nodeContract)
}

// UpdateNodeContractAddress is a paid mutator transaction binding the contract method 0xc6c8b6f7.
//
// Solidity: function updateNodeContractAddress(address nodeContract) returns()
func (_QOS *QOSTransactorSession) UpdateNodeContractAddress(nodeContract common.Address) (*types.Transaction, error) {
	return _QOS.Contract.UpdateNodeContractAddress(&_QOS.TransactOpts, nodeContract)
}

// UpdateTaskContractAddress is a paid mutator transaction binding the contract method 0x42145230.
//
// Solidity: function updateTaskContractAddress(address taskContract) returns()
func (_QOS *QOSTransactor) UpdateTaskContractAddress(opts *bind.TransactOpts, taskContract common.Address) (*types.Transaction, error) {
	return _QOS.contract.Transact(opts, "updateTaskContractAddress", taskContract)
}

// UpdateTaskContractAddress is a paid mutator transaction binding the contract method 0x42145230.
//
// Solidity: function updateTaskContractAddress(address taskContract) returns()
func (_QOS *QOSSession) UpdateTaskContractAddress(taskContract common.Address) (*types.Transaction, error) {
	return _QOS.Contract.UpdateTaskContractAddress(&_QOS.TransactOpts, taskContract)
}

// UpdateTaskContractAddress is a paid mutator transaction binding the contract method 0x42145230.
//
// Solidity: function updateTaskContractAddress(address taskContract) returns()
func (_QOS *QOSTransactorSession) UpdateTaskContractAddress(taskContract common.Address) (*types.Transaction, error) {
	return _QOS.Contract.UpdateTaskContractAddress(&_QOS.TransactOpts, taskContract)
}

// QOSOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the QOS contract.
type QOSOwnershipTransferredIterator struct {
	Event *QOSOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *QOSOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QOSOwnershipTransferred)
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
		it.Event = new(QOSOwnershipTransferred)
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
func (it *QOSOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QOSOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QOSOwnershipTransferred represents a OwnershipTransferred event raised by the QOS contract.
type QOSOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_QOS *QOSFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*QOSOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _QOS.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &QOSOwnershipTransferredIterator{contract: _QOS.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_QOS *QOSFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *QOSOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _QOS.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QOSOwnershipTransferred)
				if err := _QOS.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_QOS *QOSFilterer) ParseOwnershipTransferred(log types.Log) (*QOSOwnershipTransferred, error) {
	event := new(QOSOwnershipTransferred)
	if err := _QOS.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
