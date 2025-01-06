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

// TaskQueueMetaData contains all meta data concerning the TaskQueue contract.
var TaskQueueMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getCheapestTask\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSizeLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"taskIDCommitment\",\"type\":\"bytes32\"}],\"name\":\"include\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"gpuName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"gpuVRAM\",\"type\":\"uint256\"},{\"internalType\":\"uint256[3]\",\"name\":\"version\",\"type\":\"uint256[3]\"},{\"internalType\":\"string[]\",\"name\":\"lastModelIDs\",\"type\":\"string[]\"}],\"name\":\"popTask\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"taskIDCommitment\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"taskFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"taskSize\",\"type\":\"uint256\"},{\"internalType\":\"string[]\",\"name\":\"modelIDs\",\"type\":\"string[]\"},{\"internalType\":\"uint256\",\"name\":\"minimumVRAM\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"requiredGPU\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"requiredGPUVRAM\",\"type\":\"uint256\"},{\"internalType\":\"uint256[3]\",\"name\":\"taskVersion\",\"type\":\"uint256[3]\"}],\"name\":\"pushTask\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"taskIDCommitment\",\"type\":\"bytes32\"}],\"name\":\"removeTask\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"size\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"updateSizeLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"taskContract\",\"type\":\"address\"}],\"name\":\"updateTaskContractAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TaskQueueABI is the input ABI used to generate the binding from.
// Deprecated: Use TaskQueueMetaData.ABI instead.
var TaskQueueABI = TaskQueueMetaData.ABI

// TaskQueue is an auto generated Go binding around an Ethereum contract.
type TaskQueue struct {
	TaskQueueCaller     // Read-only binding to the contract
	TaskQueueTransactor // Write-only binding to the contract
	TaskQueueFilterer   // Log filterer for contract events
}

// TaskQueueCaller is an auto generated read-only Go binding around an Ethereum contract.
type TaskQueueCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaskQueueTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TaskQueueTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaskQueueFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TaskQueueFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaskQueueSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TaskQueueSession struct {
	Contract     *TaskQueue        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TaskQueueCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TaskQueueCallerSession struct {
	Contract *TaskQueueCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// TaskQueueTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TaskQueueTransactorSession struct {
	Contract     *TaskQueueTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// TaskQueueRaw is an auto generated low-level Go binding around an Ethereum contract.
type TaskQueueRaw struct {
	Contract *TaskQueue // Generic contract binding to access the raw methods on
}

// TaskQueueCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TaskQueueCallerRaw struct {
	Contract *TaskQueueCaller // Generic read-only contract binding to access the raw methods on
}

// TaskQueueTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TaskQueueTransactorRaw struct {
	Contract *TaskQueueTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTaskQueue creates a new instance of TaskQueue, bound to a specific deployed contract.
func NewTaskQueue(address common.Address, backend bind.ContractBackend) (*TaskQueue, error) {
	contract, err := bindTaskQueue(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TaskQueue{TaskQueueCaller: TaskQueueCaller{contract: contract}, TaskQueueTransactor: TaskQueueTransactor{contract: contract}, TaskQueueFilterer: TaskQueueFilterer{contract: contract}}, nil
}

// NewTaskQueueCaller creates a new read-only instance of TaskQueue, bound to a specific deployed contract.
func NewTaskQueueCaller(address common.Address, caller bind.ContractCaller) (*TaskQueueCaller, error) {
	contract, err := bindTaskQueue(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TaskQueueCaller{contract: contract}, nil
}

// NewTaskQueueTransactor creates a new write-only instance of TaskQueue, bound to a specific deployed contract.
func NewTaskQueueTransactor(address common.Address, transactor bind.ContractTransactor) (*TaskQueueTransactor, error) {
	contract, err := bindTaskQueue(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TaskQueueTransactor{contract: contract}, nil
}

// NewTaskQueueFilterer creates a new log filterer instance of TaskQueue, bound to a specific deployed contract.
func NewTaskQueueFilterer(address common.Address, filterer bind.ContractFilterer) (*TaskQueueFilterer, error) {
	contract, err := bindTaskQueue(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TaskQueueFilterer{contract: contract}, nil
}

// bindTaskQueue binds a generic wrapper to an already deployed contract.
func bindTaskQueue(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TaskQueueMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TaskQueue *TaskQueueRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TaskQueue.Contract.TaskQueueCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TaskQueue *TaskQueueRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaskQueue.Contract.TaskQueueTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TaskQueue *TaskQueueRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TaskQueue.Contract.TaskQueueTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TaskQueue *TaskQueueCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TaskQueue.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TaskQueue *TaskQueueTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaskQueue.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TaskQueue *TaskQueueTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TaskQueue.Contract.contract.Transact(opts, method, params...)
}

// GetCheapestTask is a free data retrieval call binding the contract method 0x56b58d34.
//
// Solidity: function getCheapestTask() view returns(bytes32)
func (_TaskQueue *TaskQueueCaller) GetCheapestTask(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TaskQueue.contract.Call(opts, &out, "getCheapestTask")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetCheapestTask is a free data retrieval call binding the contract method 0x56b58d34.
//
// Solidity: function getCheapestTask() view returns(bytes32)
func (_TaskQueue *TaskQueueSession) GetCheapestTask() ([32]byte, error) {
	return _TaskQueue.Contract.GetCheapestTask(&_TaskQueue.CallOpts)
}

// GetCheapestTask is a free data retrieval call binding the contract method 0x56b58d34.
//
// Solidity: function getCheapestTask() view returns(bytes32)
func (_TaskQueue *TaskQueueCallerSession) GetCheapestTask() ([32]byte, error) {
	return _TaskQueue.Contract.GetCheapestTask(&_TaskQueue.CallOpts)
}

// GetSizeLimit is a free data retrieval call binding the contract method 0xf00c6741.
//
// Solidity: function getSizeLimit() view returns(uint256)
func (_TaskQueue *TaskQueueCaller) GetSizeLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TaskQueue.contract.Call(opts, &out, "getSizeLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSizeLimit is a free data retrieval call binding the contract method 0xf00c6741.
//
// Solidity: function getSizeLimit() view returns(uint256)
func (_TaskQueue *TaskQueueSession) GetSizeLimit() (*big.Int, error) {
	return _TaskQueue.Contract.GetSizeLimit(&_TaskQueue.CallOpts)
}

// GetSizeLimit is a free data retrieval call binding the contract method 0xf00c6741.
//
// Solidity: function getSizeLimit() view returns(uint256)
func (_TaskQueue *TaskQueueCallerSession) GetSizeLimit() (*big.Int, error) {
	return _TaskQueue.Contract.GetSizeLimit(&_TaskQueue.CallOpts)
}

// Include is a free data retrieval call binding the contract method 0x87abc5a4.
//
// Solidity: function include(bytes32 taskIDCommitment) view returns(bool)
func (_TaskQueue *TaskQueueCaller) Include(opts *bind.CallOpts, taskIDCommitment [32]byte) (bool, error) {
	var out []interface{}
	err := _TaskQueue.contract.Call(opts, &out, "include", taskIDCommitment)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Include is a free data retrieval call binding the contract method 0x87abc5a4.
//
// Solidity: function include(bytes32 taskIDCommitment) view returns(bool)
func (_TaskQueue *TaskQueueSession) Include(taskIDCommitment [32]byte) (bool, error) {
	return _TaskQueue.Contract.Include(&_TaskQueue.CallOpts, taskIDCommitment)
}

// Include is a free data retrieval call binding the contract method 0x87abc5a4.
//
// Solidity: function include(bytes32 taskIDCommitment) view returns(bool)
func (_TaskQueue *TaskQueueCallerSession) Include(taskIDCommitment [32]byte) (bool, error) {
	return _TaskQueue.Contract.Include(&_TaskQueue.CallOpts, taskIDCommitment)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TaskQueue *TaskQueueCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TaskQueue.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TaskQueue *TaskQueueSession) Owner() (common.Address, error) {
	return _TaskQueue.Contract.Owner(&_TaskQueue.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TaskQueue *TaskQueueCallerSession) Owner() (common.Address, error) {
	return _TaskQueue.Contract.Owner(&_TaskQueue.CallOpts)
}

// Size is a free data retrieval call binding the contract method 0x949d225d.
//
// Solidity: function size() view returns(uint256)
func (_TaskQueue *TaskQueueCaller) Size(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TaskQueue.contract.Call(opts, &out, "size")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Size is a free data retrieval call binding the contract method 0x949d225d.
//
// Solidity: function size() view returns(uint256)
func (_TaskQueue *TaskQueueSession) Size() (*big.Int, error) {
	return _TaskQueue.Contract.Size(&_TaskQueue.CallOpts)
}

// Size is a free data retrieval call binding the contract method 0x949d225d.
//
// Solidity: function size() view returns(uint256)
func (_TaskQueue *TaskQueueCallerSession) Size() (*big.Int, error) {
	return _TaskQueue.Contract.Size(&_TaskQueue.CallOpts)
}

// PopTask is a paid mutator transaction binding the contract method 0x848001f4.
//
// Solidity: function popTask(string gpuName, uint256 gpuVRAM, uint256[3] version, string[] lastModelIDs) returns(bytes32)
func (_TaskQueue *TaskQueueTransactor) PopTask(opts *bind.TransactOpts, gpuName string, gpuVRAM *big.Int, version [3]*big.Int, lastModelIDs []string) (*types.Transaction, error) {
	return _TaskQueue.contract.Transact(opts, "popTask", gpuName, gpuVRAM, version, lastModelIDs)
}

// PopTask is a paid mutator transaction binding the contract method 0x848001f4.
//
// Solidity: function popTask(string gpuName, uint256 gpuVRAM, uint256[3] version, string[] lastModelIDs) returns(bytes32)
func (_TaskQueue *TaskQueueSession) PopTask(gpuName string, gpuVRAM *big.Int, version [3]*big.Int, lastModelIDs []string) (*types.Transaction, error) {
	return _TaskQueue.Contract.PopTask(&_TaskQueue.TransactOpts, gpuName, gpuVRAM, version, lastModelIDs)
}

// PopTask is a paid mutator transaction binding the contract method 0x848001f4.
//
// Solidity: function popTask(string gpuName, uint256 gpuVRAM, uint256[3] version, string[] lastModelIDs) returns(bytes32)
func (_TaskQueue *TaskQueueTransactorSession) PopTask(gpuName string, gpuVRAM *big.Int, version [3]*big.Int, lastModelIDs []string) (*types.Transaction, error) {
	return _TaskQueue.Contract.PopTask(&_TaskQueue.TransactOpts, gpuName, gpuVRAM, version, lastModelIDs)
}

// PushTask is a paid mutator transaction binding the contract method 0xc826ae15.
//
// Solidity: function pushTask(bytes32 taskIDCommitment, uint256 taskFee, uint256 taskSize, string[] modelIDs, uint256 minimumVRAM, string requiredGPU, uint256 requiredGPUVRAM, uint256[3] taskVersion) returns()
func (_TaskQueue *TaskQueueTransactor) PushTask(opts *bind.TransactOpts, taskIDCommitment [32]byte, taskFee *big.Int, taskSize *big.Int, modelIDs []string, minimumVRAM *big.Int, requiredGPU string, requiredGPUVRAM *big.Int, taskVersion [3]*big.Int) (*types.Transaction, error) {
	return _TaskQueue.contract.Transact(opts, "pushTask", taskIDCommitment, taskFee, taskSize, modelIDs, minimumVRAM, requiredGPU, requiredGPUVRAM, taskVersion)
}

// PushTask is a paid mutator transaction binding the contract method 0xc826ae15.
//
// Solidity: function pushTask(bytes32 taskIDCommitment, uint256 taskFee, uint256 taskSize, string[] modelIDs, uint256 minimumVRAM, string requiredGPU, uint256 requiredGPUVRAM, uint256[3] taskVersion) returns()
func (_TaskQueue *TaskQueueSession) PushTask(taskIDCommitment [32]byte, taskFee *big.Int, taskSize *big.Int, modelIDs []string, minimumVRAM *big.Int, requiredGPU string, requiredGPUVRAM *big.Int, taskVersion [3]*big.Int) (*types.Transaction, error) {
	return _TaskQueue.Contract.PushTask(&_TaskQueue.TransactOpts, taskIDCommitment, taskFee, taskSize, modelIDs, minimumVRAM, requiredGPU, requiredGPUVRAM, taskVersion)
}

// PushTask is a paid mutator transaction binding the contract method 0xc826ae15.
//
// Solidity: function pushTask(bytes32 taskIDCommitment, uint256 taskFee, uint256 taskSize, string[] modelIDs, uint256 minimumVRAM, string requiredGPU, uint256 requiredGPUVRAM, uint256[3] taskVersion) returns()
func (_TaskQueue *TaskQueueTransactorSession) PushTask(taskIDCommitment [32]byte, taskFee *big.Int, taskSize *big.Int, modelIDs []string, minimumVRAM *big.Int, requiredGPU string, requiredGPUVRAM *big.Int, taskVersion [3]*big.Int) (*types.Transaction, error) {
	return _TaskQueue.Contract.PushTask(&_TaskQueue.TransactOpts, taskIDCommitment, taskFee, taskSize, modelIDs, minimumVRAM, requiredGPU, requiredGPUVRAM, taskVersion)
}

// RemoveTask is a paid mutator transaction binding the contract method 0x5035b622.
//
// Solidity: function removeTask(bytes32 taskIDCommitment) returns()
func (_TaskQueue *TaskQueueTransactor) RemoveTask(opts *bind.TransactOpts, taskIDCommitment [32]byte) (*types.Transaction, error) {
	return _TaskQueue.contract.Transact(opts, "removeTask", taskIDCommitment)
}

// RemoveTask is a paid mutator transaction binding the contract method 0x5035b622.
//
// Solidity: function removeTask(bytes32 taskIDCommitment) returns()
func (_TaskQueue *TaskQueueSession) RemoveTask(taskIDCommitment [32]byte) (*types.Transaction, error) {
	return _TaskQueue.Contract.RemoveTask(&_TaskQueue.TransactOpts, taskIDCommitment)
}

// RemoveTask is a paid mutator transaction binding the contract method 0x5035b622.
//
// Solidity: function removeTask(bytes32 taskIDCommitment) returns()
func (_TaskQueue *TaskQueueTransactorSession) RemoveTask(taskIDCommitment [32]byte) (*types.Transaction, error) {
	return _TaskQueue.Contract.RemoveTask(&_TaskQueue.TransactOpts, taskIDCommitment)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TaskQueue *TaskQueueTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaskQueue.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TaskQueue *TaskQueueSession) RenounceOwnership() (*types.Transaction, error) {
	return _TaskQueue.Contract.RenounceOwnership(&_TaskQueue.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TaskQueue *TaskQueueTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TaskQueue.Contract.RenounceOwnership(&_TaskQueue.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TaskQueue *TaskQueueTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TaskQueue.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TaskQueue *TaskQueueSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TaskQueue.Contract.TransferOwnership(&_TaskQueue.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TaskQueue *TaskQueueTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TaskQueue.Contract.TransferOwnership(&_TaskQueue.TransactOpts, newOwner)
}

// UpdateSizeLimit is a paid mutator transaction binding the contract method 0xfcdc8b8e.
//
// Solidity: function updateSizeLimit(uint256 limit) returns()
func (_TaskQueue *TaskQueueTransactor) UpdateSizeLimit(opts *bind.TransactOpts, limit *big.Int) (*types.Transaction, error) {
	return _TaskQueue.contract.Transact(opts, "updateSizeLimit", limit)
}

// UpdateSizeLimit is a paid mutator transaction binding the contract method 0xfcdc8b8e.
//
// Solidity: function updateSizeLimit(uint256 limit) returns()
func (_TaskQueue *TaskQueueSession) UpdateSizeLimit(limit *big.Int) (*types.Transaction, error) {
	return _TaskQueue.Contract.UpdateSizeLimit(&_TaskQueue.TransactOpts, limit)
}

// UpdateSizeLimit is a paid mutator transaction binding the contract method 0xfcdc8b8e.
//
// Solidity: function updateSizeLimit(uint256 limit) returns()
func (_TaskQueue *TaskQueueTransactorSession) UpdateSizeLimit(limit *big.Int) (*types.Transaction, error) {
	return _TaskQueue.Contract.UpdateSizeLimit(&_TaskQueue.TransactOpts, limit)
}

// UpdateTaskContractAddress is a paid mutator transaction binding the contract method 0x42145230.
//
// Solidity: function updateTaskContractAddress(address taskContract) returns()
func (_TaskQueue *TaskQueueTransactor) UpdateTaskContractAddress(opts *bind.TransactOpts, taskContract common.Address) (*types.Transaction, error) {
	return _TaskQueue.contract.Transact(opts, "updateTaskContractAddress", taskContract)
}

// UpdateTaskContractAddress is a paid mutator transaction binding the contract method 0x42145230.
//
// Solidity: function updateTaskContractAddress(address taskContract) returns()
func (_TaskQueue *TaskQueueSession) UpdateTaskContractAddress(taskContract common.Address) (*types.Transaction, error) {
	return _TaskQueue.Contract.UpdateTaskContractAddress(&_TaskQueue.TransactOpts, taskContract)
}

// UpdateTaskContractAddress is a paid mutator transaction binding the contract method 0x42145230.
//
// Solidity: function updateTaskContractAddress(address taskContract) returns()
func (_TaskQueue *TaskQueueTransactorSession) UpdateTaskContractAddress(taskContract common.Address) (*types.Transaction, error) {
	return _TaskQueue.Contract.UpdateTaskContractAddress(&_TaskQueue.TransactOpts, taskContract)
}

// TaskQueueOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TaskQueue contract.
type TaskQueueOwnershipTransferredIterator struct {
	Event *TaskQueueOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TaskQueueOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskQueueOwnershipTransferred)
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
		it.Event = new(TaskQueueOwnershipTransferred)
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
func (it *TaskQueueOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskQueueOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskQueueOwnershipTransferred represents a OwnershipTransferred event raised by the TaskQueue contract.
type TaskQueueOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TaskQueue *TaskQueueFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TaskQueueOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TaskQueue.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TaskQueueOwnershipTransferredIterator{contract: _TaskQueue.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TaskQueue *TaskQueueFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TaskQueueOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TaskQueue.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskQueueOwnershipTransferred)
				if err := _TaskQueue.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_TaskQueue *TaskQueueFilterer) ParseOwnershipTransferred(log types.Log) (*TaskQueueOwnershipTransferred, error) {
	event := new(TaskQueueOwnershipTransferred)
	if err := _TaskQueue.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
