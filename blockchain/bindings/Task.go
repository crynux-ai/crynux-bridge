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

// TaskInfo is an auto generated low-level Go binding around an user-defined struct.
type TaskInfo struct {
	Id                    *big.Int
	TaskType              *big.Int
	Creator               common.Address
	TaskHash              [32]byte
	DataHash              [32]byte
	VramLimit             *big.Int
	IsSuccess             bool
	SelectedNodes         []common.Address
	Commitments           [][32]byte
	Nonces                [][32]byte
	Results               [][]byte
	ResultDisclosedRounds []*big.Int
	ResultNode            common.Address
	Aborted               bool
	Timeout               *big.Int
	Balance               *big.Int
}

// TaskMetaData contains all meta data concerning the Task contract.
var TaskMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractNode\",\"name\":\"nodeInstance\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenInstance\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"taskId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"TaskAborted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"taskId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"taskType\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"selectedNode\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"taskHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"}],\"name\":\"TaskCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"taskId\",\"type\":\"uint256\"}],\"name\":\"TaskResultCommitmentsReady\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"taskId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"result\",\"type\":\"bytes\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"resultNode\",\"type\":\"address\"}],\"name\":\"TaskSuccess\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"taskType\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"taskHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"vramLimit\",\"type\":\"uint256\"}],\"name\":\"createTask\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"taskId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"nonce\",\"type\":\"bytes32\"}],\"name\":\"submitTaskResultCommitment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"taskId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"result\",\"type\":\"bytes\"}],\"name\":\"discloseTaskResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"taskId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"}],\"name\":\"reportResultsUploaded\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"taskId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"}],\"name\":\"reportTaskError\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"taskId\",\"type\":\"uint256\"}],\"name\":\"cancelTask\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"updateTaskFeePerNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"}],\"name\":\"updateDistanceThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"t\",\"type\":\"uint256\"}],\"name\":\"updateTimeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"taskId\",\"type\":\"uint256\"}],\"name\":\"getTask\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"taskType\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"taskHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"vramLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isSuccess\",\"type\":\"bool\"},{\"internalType\":\"address[]\",\"name\":\"selectedNodes\",\"type\":\"address[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"commitments\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"nonces\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256[]\",\"name\":\"resultDisclosedRounds\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"resultNode\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"aborted\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"internalType\":\"structTask.TaskInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nodeAddress\",\"type\":\"address\"}],\"name\":\"getNodeTask\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalTasks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSuccessTasks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAbortedTasks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// TaskABI is the input ABI used to generate the binding from.
// Deprecated: Use TaskMetaData.ABI instead.
var TaskABI = TaskMetaData.ABI

// Task is an auto generated Go binding around an Ethereum contract.
type Task struct {
	TaskCaller     // Read-only binding to the contract
	TaskTransactor // Write-only binding to the contract
	TaskFilterer   // Log filterer for contract events
}

// TaskCaller is an auto generated read-only Go binding around an Ethereum contract.
type TaskCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaskTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TaskTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaskFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TaskFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaskSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TaskSession struct {
	Contract     *Task             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TaskCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TaskCallerSession struct {
	Contract *TaskCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TaskTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TaskTransactorSession struct {
	Contract     *TaskTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TaskRaw is an auto generated low-level Go binding around an Ethereum contract.
type TaskRaw struct {
	Contract *Task // Generic contract binding to access the raw methods on
}

// TaskCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TaskCallerRaw struct {
	Contract *TaskCaller // Generic read-only contract binding to access the raw methods on
}

// TaskTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TaskTransactorRaw struct {
	Contract *TaskTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTask creates a new instance of Task, bound to a specific deployed contract.
func NewTask(address common.Address, backend bind.ContractBackend) (*Task, error) {
	contract, err := bindTask(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Task{TaskCaller: TaskCaller{contract: contract}, TaskTransactor: TaskTransactor{contract: contract}, TaskFilterer: TaskFilterer{contract: contract}}, nil
}

// NewTaskCaller creates a new read-only instance of Task, bound to a specific deployed contract.
func NewTaskCaller(address common.Address, caller bind.ContractCaller) (*TaskCaller, error) {
	contract, err := bindTask(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TaskCaller{contract: contract}, nil
}

// NewTaskTransactor creates a new write-only instance of Task, bound to a specific deployed contract.
func NewTaskTransactor(address common.Address, transactor bind.ContractTransactor) (*TaskTransactor, error) {
	contract, err := bindTask(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TaskTransactor{contract: contract}, nil
}

// NewTaskFilterer creates a new log filterer instance of Task, bound to a specific deployed contract.
func NewTaskFilterer(address common.Address, filterer bind.ContractFilterer) (*TaskFilterer, error) {
	contract, err := bindTask(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TaskFilterer{contract: contract}, nil
}

// bindTask binds a generic wrapper to an already deployed contract.
func bindTask(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TaskMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Task *TaskRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Task.Contract.TaskCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Task *TaskRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Task.Contract.TaskTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Task *TaskRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Task.Contract.TaskTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Task *TaskCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Task.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Task *TaskTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Task.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Task *TaskTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Task.Contract.contract.Transact(opts, method, params...)
}

// GetNodeTask is a free data retrieval call binding the contract method 0x9877ad45.
//
// Solidity: function getNodeTask(address nodeAddress) view returns(uint256)
func (_Task *TaskCaller) GetNodeTask(opts *bind.CallOpts, nodeAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Task.contract.Call(opts, &out, "getNodeTask", nodeAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNodeTask is a free data retrieval call binding the contract method 0x9877ad45.
//
// Solidity: function getNodeTask(address nodeAddress) view returns(uint256)
func (_Task *TaskSession) GetNodeTask(nodeAddress common.Address) (*big.Int, error) {
	return _Task.Contract.GetNodeTask(&_Task.CallOpts, nodeAddress)
}

// GetNodeTask is a free data retrieval call binding the contract method 0x9877ad45.
//
// Solidity: function getNodeTask(address nodeAddress) view returns(uint256)
func (_Task *TaskCallerSession) GetNodeTask(nodeAddress common.Address) (*big.Int, error) {
	return _Task.Contract.GetNodeTask(&_Task.CallOpts, nodeAddress)
}

// GetTask is a free data retrieval call binding the contract method 0x1d65e77e.
//
// Solidity: function getTask(uint256 taskId) view returns((uint256,uint256,address,bytes32,bytes32,uint256,bool,address[],bytes32[],bytes32[],bytes[],uint256[],address,bool,uint256,uint256))
func (_Task *TaskCaller) GetTask(opts *bind.CallOpts, taskId *big.Int) (TaskInfo, error) {
	var out []interface{}
	err := _Task.contract.Call(opts, &out, "getTask", taskId)

	if err != nil {
		return *new(TaskInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(TaskInfo)).(*TaskInfo)

	return out0, err

}

// GetTask is a free data retrieval call binding the contract method 0x1d65e77e.
//
// Solidity: function getTask(uint256 taskId) view returns((uint256,uint256,address,bytes32,bytes32,uint256,bool,address[],bytes32[],bytes32[],bytes[],uint256[],address,bool,uint256,uint256))
func (_Task *TaskSession) GetTask(taskId *big.Int) (TaskInfo, error) {
	return _Task.Contract.GetTask(&_Task.CallOpts, taskId)
}

// GetTask is a free data retrieval call binding the contract method 0x1d65e77e.
//
// Solidity: function getTask(uint256 taskId) view returns((uint256,uint256,address,bytes32,bytes32,uint256,bool,address[],bytes32[],bytes32[],bytes[],uint256[],address,bool,uint256,uint256))
func (_Task *TaskCallerSession) GetTask(taskId *big.Int) (TaskInfo, error) {
	return _Task.Contract.GetTask(&_Task.CallOpts, taskId)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Task *TaskCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Task.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Task *TaskSession) Owner() (common.Address, error) {
	return _Task.Contract.Owner(&_Task.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Task *TaskCallerSession) Owner() (common.Address, error) {
	return _Task.Contract.Owner(&_Task.CallOpts)
}

// TotalAbortedTasks is a free data retrieval call binding the contract method 0x2ff6d2ca.
//
// Solidity: function totalAbortedTasks() view returns(uint256)
func (_Task *TaskCaller) TotalAbortedTasks(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Task.contract.Call(opts, &out, "totalAbortedTasks")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAbortedTasks is a free data retrieval call binding the contract method 0x2ff6d2ca.
//
// Solidity: function totalAbortedTasks() view returns(uint256)
func (_Task *TaskSession) TotalAbortedTasks() (*big.Int, error) {
	return _Task.Contract.TotalAbortedTasks(&_Task.CallOpts)
}

// TotalAbortedTasks is a free data retrieval call binding the contract method 0x2ff6d2ca.
//
// Solidity: function totalAbortedTasks() view returns(uint256)
func (_Task *TaskCallerSession) TotalAbortedTasks() (*big.Int, error) {
	return _Task.Contract.TotalAbortedTasks(&_Task.CallOpts)
}

// TotalSuccessTasks is a free data retrieval call binding the contract method 0x775820f8.
//
// Solidity: function totalSuccessTasks() view returns(uint256)
func (_Task *TaskCaller) TotalSuccessTasks(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Task.contract.Call(opts, &out, "totalSuccessTasks")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSuccessTasks is a free data retrieval call binding the contract method 0x775820f8.
//
// Solidity: function totalSuccessTasks() view returns(uint256)
func (_Task *TaskSession) TotalSuccessTasks() (*big.Int, error) {
	return _Task.Contract.TotalSuccessTasks(&_Task.CallOpts)
}

// TotalSuccessTasks is a free data retrieval call binding the contract method 0x775820f8.
//
// Solidity: function totalSuccessTasks() view returns(uint256)
func (_Task *TaskCallerSession) TotalSuccessTasks() (*big.Int, error) {
	return _Task.Contract.TotalSuccessTasks(&_Task.CallOpts)
}

// TotalTasks is a free data retrieval call binding the contract method 0xd22c81e5.
//
// Solidity: function totalTasks() view returns(uint256)
func (_Task *TaskCaller) TotalTasks(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Task.contract.Call(opts, &out, "totalTasks")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalTasks is a free data retrieval call binding the contract method 0xd22c81e5.
//
// Solidity: function totalTasks() view returns(uint256)
func (_Task *TaskSession) TotalTasks() (*big.Int, error) {
	return _Task.Contract.TotalTasks(&_Task.CallOpts)
}

// TotalTasks is a free data retrieval call binding the contract method 0xd22c81e5.
//
// Solidity: function totalTasks() view returns(uint256)
func (_Task *TaskCallerSession) TotalTasks() (*big.Int, error) {
	return _Task.Contract.TotalTasks(&_Task.CallOpts)
}

// CancelTask is a paid mutator transaction binding the contract method 0x7eec20a8.
//
// Solidity: function cancelTask(uint256 taskId) returns()
func (_Task *TaskTransactor) CancelTask(opts *bind.TransactOpts, taskId *big.Int) (*types.Transaction, error) {
	return _Task.contract.Transact(opts, "cancelTask", taskId)
}

// CancelTask is a paid mutator transaction binding the contract method 0x7eec20a8.
//
// Solidity: function cancelTask(uint256 taskId) returns()
func (_Task *TaskSession) CancelTask(taskId *big.Int) (*types.Transaction, error) {
	return _Task.Contract.CancelTask(&_Task.TransactOpts, taskId)
}

// CancelTask is a paid mutator transaction binding the contract method 0x7eec20a8.
//
// Solidity: function cancelTask(uint256 taskId) returns()
func (_Task *TaskTransactorSession) CancelTask(taskId *big.Int) (*types.Transaction, error) {
	return _Task.Contract.CancelTask(&_Task.TransactOpts, taskId)
}

// CreateTask is a paid mutator transaction binding the contract method 0x6ab702ed.
//
// Solidity: function createTask(uint256 taskType, bytes32 taskHash, bytes32 dataHash, uint256 vramLimit) returns()
func (_Task *TaskTransactor) CreateTask(opts *bind.TransactOpts, taskType *big.Int, taskHash [32]byte, dataHash [32]byte, vramLimit *big.Int) (*types.Transaction, error) {
	return _Task.contract.Transact(opts, "createTask", taskType, taskHash, dataHash, vramLimit)
}

// CreateTask is a paid mutator transaction binding the contract method 0x6ab702ed.
//
// Solidity: function createTask(uint256 taskType, bytes32 taskHash, bytes32 dataHash, uint256 vramLimit) returns()
func (_Task *TaskSession) CreateTask(taskType *big.Int, taskHash [32]byte, dataHash [32]byte, vramLimit *big.Int) (*types.Transaction, error) {
	return _Task.Contract.CreateTask(&_Task.TransactOpts, taskType, taskHash, dataHash, vramLimit)
}

// CreateTask is a paid mutator transaction binding the contract method 0x6ab702ed.
//
// Solidity: function createTask(uint256 taskType, bytes32 taskHash, bytes32 dataHash, uint256 vramLimit) returns()
func (_Task *TaskTransactorSession) CreateTask(taskType *big.Int, taskHash [32]byte, dataHash [32]byte, vramLimit *big.Int) (*types.Transaction, error) {
	return _Task.Contract.CreateTask(&_Task.TransactOpts, taskType, taskHash, dataHash, vramLimit)
}

// DiscloseTaskResult is a paid mutator transaction binding the contract method 0x63be8c33.
//
// Solidity: function discloseTaskResult(uint256 taskId, uint256 round, bytes result) returns()
func (_Task *TaskTransactor) DiscloseTaskResult(opts *bind.TransactOpts, taskId *big.Int, round *big.Int, result []byte) (*types.Transaction, error) {
	return _Task.contract.Transact(opts, "discloseTaskResult", taskId, round, result)
}

// DiscloseTaskResult is a paid mutator transaction binding the contract method 0x63be8c33.
//
// Solidity: function discloseTaskResult(uint256 taskId, uint256 round, bytes result) returns()
func (_Task *TaskSession) DiscloseTaskResult(taskId *big.Int, round *big.Int, result []byte) (*types.Transaction, error) {
	return _Task.Contract.DiscloseTaskResult(&_Task.TransactOpts, taskId, round, result)
}

// DiscloseTaskResult is a paid mutator transaction binding the contract method 0x63be8c33.
//
// Solidity: function discloseTaskResult(uint256 taskId, uint256 round, bytes result) returns()
func (_Task *TaskTransactorSession) DiscloseTaskResult(taskId *big.Int, round *big.Int, result []byte) (*types.Transaction, error) {
	return _Task.Contract.DiscloseTaskResult(&_Task.TransactOpts, taskId, round, result)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Task *TaskTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Task.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Task *TaskSession) RenounceOwnership() (*types.Transaction, error) {
	return _Task.Contract.RenounceOwnership(&_Task.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Task *TaskTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Task.Contract.RenounceOwnership(&_Task.TransactOpts)
}

// ReportResultsUploaded is a paid mutator transaction binding the contract method 0x95b1d89b.
//
// Solidity: function reportResultsUploaded(uint256 taskId, uint256 round) returns()
func (_Task *TaskTransactor) ReportResultsUploaded(opts *bind.TransactOpts, taskId *big.Int, round *big.Int) (*types.Transaction, error) {
	return _Task.contract.Transact(opts, "reportResultsUploaded", taskId, round)
}

// ReportResultsUploaded is a paid mutator transaction binding the contract method 0x95b1d89b.
//
// Solidity: function reportResultsUploaded(uint256 taskId, uint256 round) returns()
func (_Task *TaskSession) ReportResultsUploaded(taskId *big.Int, round *big.Int) (*types.Transaction, error) {
	return _Task.Contract.ReportResultsUploaded(&_Task.TransactOpts, taskId, round)
}

// ReportResultsUploaded is a paid mutator transaction binding the contract method 0x95b1d89b.
//
// Solidity: function reportResultsUploaded(uint256 taskId, uint256 round) returns()
func (_Task *TaskTransactorSession) ReportResultsUploaded(taskId *big.Int, round *big.Int) (*types.Transaction, error) {
	return _Task.Contract.ReportResultsUploaded(&_Task.TransactOpts, taskId, round)
}

// ReportTaskError is a paid mutator transaction binding the contract method 0x695b1c8f.
//
// Solidity: function reportTaskError(uint256 taskId, uint256 round) returns()
func (_Task *TaskTransactor) ReportTaskError(opts *bind.TransactOpts, taskId *big.Int, round *big.Int) (*types.Transaction, error) {
	return _Task.contract.Transact(opts, "reportTaskError", taskId, round)
}

// ReportTaskError is a paid mutator transaction binding the contract method 0x695b1c8f.
//
// Solidity: function reportTaskError(uint256 taskId, uint256 round) returns()
func (_Task *TaskSession) ReportTaskError(taskId *big.Int, round *big.Int) (*types.Transaction, error) {
	return _Task.Contract.ReportTaskError(&_Task.TransactOpts, taskId, round)
}

// ReportTaskError is a paid mutator transaction binding the contract method 0x695b1c8f.
//
// Solidity: function reportTaskError(uint256 taskId, uint256 round) returns()
func (_Task *TaskTransactorSession) ReportTaskError(taskId *big.Int, round *big.Int) (*types.Transaction, error) {
	return _Task.Contract.ReportTaskError(&_Task.TransactOpts, taskId, round)
}

// SubmitTaskResultCommitment is a paid mutator transaction binding the contract method 0x47f90738.
//
// Solidity: function submitTaskResultCommitment(uint256 taskId, uint256 round, bytes32 commitment, bytes32 nonce) returns()
func (_Task *TaskTransactor) SubmitTaskResultCommitment(opts *bind.TransactOpts, taskId *big.Int, round *big.Int, commitment [32]byte, nonce [32]byte) (*types.Transaction, error) {
	return _Task.contract.Transact(opts, "submitTaskResultCommitment", taskId, round, commitment, nonce)
}

// SubmitTaskResultCommitment is a paid mutator transaction binding the contract method 0x47f90738.
//
// Solidity: function submitTaskResultCommitment(uint256 taskId, uint256 round, bytes32 commitment, bytes32 nonce) returns()
func (_Task *TaskSession) SubmitTaskResultCommitment(taskId *big.Int, round *big.Int, commitment [32]byte, nonce [32]byte) (*types.Transaction, error) {
	return _Task.Contract.SubmitTaskResultCommitment(&_Task.TransactOpts, taskId, round, commitment, nonce)
}

// SubmitTaskResultCommitment is a paid mutator transaction binding the contract method 0x47f90738.
//
// Solidity: function submitTaskResultCommitment(uint256 taskId, uint256 round, bytes32 commitment, bytes32 nonce) returns()
func (_Task *TaskTransactorSession) SubmitTaskResultCommitment(taskId *big.Int, round *big.Int, commitment [32]byte, nonce [32]byte) (*types.Transaction, error) {
	return _Task.Contract.SubmitTaskResultCommitment(&_Task.TransactOpts, taskId, round, commitment, nonce)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Task *TaskTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Task.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Task *TaskSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Task.Contract.TransferOwnership(&_Task.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Task *TaskTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Task.Contract.TransferOwnership(&_Task.TransactOpts, newOwner)
}

// UpdateDistanceThreshold is a paid mutator transaction binding the contract method 0x9244e462.
//
// Solidity: function updateDistanceThreshold(uint256 threshold) returns()
func (_Task *TaskTransactor) UpdateDistanceThreshold(opts *bind.TransactOpts, threshold *big.Int) (*types.Transaction, error) {
	return _Task.contract.Transact(opts, "updateDistanceThreshold", threshold)
}

// UpdateDistanceThreshold is a paid mutator transaction binding the contract method 0x9244e462.
//
// Solidity: function updateDistanceThreshold(uint256 threshold) returns()
func (_Task *TaskSession) UpdateDistanceThreshold(threshold *big.Int) (*types.Transaction, error) {
	return _Task.Contract.UpdateDistanceThreshold(&_Task.TransactOpts, threshold)
}

// UpdateDistanceThreshold is a paid mutator transaction binding the contract method 0x9244e462.
//
// Solidity: function updateDistanceThreshold(uint256 threshold) returns()
func (_Task *TaskTransactorSession) UpdateDistanceThreshold(threshold *big.Int) (*types.Transaction, error) {
	return _Task.Contract.UpdateDistanceThreshold(&_Task.TransactOpts, threshold)
}

// UpdateTaskFeePerNode is a paid mutator transaction binding the contract method 0x96d54e34.
//
// Solidity: function updateTaskFeePerNode(uint256 fee) returns()
func (_Task *TaskTransactor) UpdateTaskFeePerNode(opts *bind.TransactOpts, fee *big.Int) (*types.Transaction, error) {
	return _Task.contract.Transact(opts, "updateTaskFeePerNode", fee)
}

// UpdateTaskFeePerNode is a paid mutator transaction binding the contract method 0x96d54e34.
//
// Solidity: function updateTaskFeePerNode(uint256 fee) returns()
func (_Task *TaskSession) UpdateTaskFeePerNode(fee *big.Int) (*types.Transaction, error) {
	return _Task.Contract.UpdateTaskFeePerNode(&_Task.TransactOpts, fee)
}

// UpdateTaskFeePerNode is a paid mutator transaction binding the contract method 0x96d54e34.
//
// Solidity: function updateTaskFeePerNode(uint256 fee) returns()
func (_Task *TaskTransactorSession) UpdateTaskFeePerNode(fee *big.Int) (*types.Transaction, error) {
	return _Task.Contract.UpdateTaskFeePerNode(&_Task.TransactOpts, fee)
}

// UpdateTimeout is a paid mutator transaction binding the contract method 0xa330214e.
//
// Solidity: function updateTimeout(uint256 t) returns()
func (_Task *TaskTransactor) UpdateTimeout(opts *bind.TransactOpts, t *big.Int) (*types.Transaction, error) {
	return _Task.contract.Transact(opts, "updateTimeout", t)
}

// UpdateTimeout is a paid mutator transaction binding the contract method 0xa330214e.
//
// Solidity: function updateTimeout(uint256 t) returns()
func (_Task *TaskSession) UpdateTimeout(t *big.Int) (*types.Transaction, error) {
	return _Task.Contract.UpdateTimeout(&_Task.TransactOpts, t)
}

// UpdateTimeout is a paid mutator transaction binding the contract method 0xa330214e.
//
// Solidity: function updateTimeout(uint256 t) returns()
func (_Task *TaskTransactorSession) UpdateTimeout(t *big.Int) (*types.Transaction, error) {
	return _Task.Contract.UpdateTimeout(&_Task.TransactOpts, t)
}

// TaskOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Task contract.
type TaskOwnershipTransferredIterator struct {
	Event *TaskOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TaskOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskOwnershipTransferred)
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
		it.Event = new(TaskOwnershipTransferred)
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
func (it *TaskOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskOwnershipTransferred represents a OwnershipTransferred event raised by the Task contract.
type TaskOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Task *TaskFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TaskOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Task.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TaskOwnershipTransferredIterator{contract: _Task.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Task *TaskFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TaskOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Task.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskOwnershipTransferred)
				if err := _Task.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Task *TaskFilterer) ParseOwnershipTransferred(log types.Log) (*TaskOwnershipTransferred, error) {
	event := new(TaskOwnershipTransferred)
	if err := _Task.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaskTaskAbortedIterator is returned from FilterTaskAborted and is used to iterate over the raw logs and unpacked data for TaskAborted events raised by the Task contract.
type TaskTaskAbortedIterator struct {
	Event *TaskTaskAborted // Event containing the contract specifics and raw log

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
func (it *TaskTaskAbortedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskTaskAborted)
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
		it.Event = new(TaskTaskAborted)
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
func (it *TaskTaskAbortedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskTaskAbortedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskTaskAborted represents a TaskAborted event raised by the Task contract.
type TaskTaskAborted struct {
	TaskId *big.Int
	Reason string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTaskAborted is a free log retrieval operation binding the contract event 0x346216bb718ede47b40f85e7caf58cd997eebda879d146a8b49091c106c851d9.
//
// Solidity: event TaskAborted(uint256 taskId, string reason)
func (_Task *TaskFilterer) FilterTaskAborted(opts *bind.FilterOpts) (*TaskTaskAbortedIterator, error) {

	logs, sub, err := _Task.contract.FilterLogs(opts, "TaskAborted")
	if err != nil {
		return nil, err
	}
	return &TaskTaskAbortedIterator{contract: _Task.contract, event: "TaskAborted", logs: logs, sub: sub}, nil
}

// WatchTaskAborted is a free log subscription operation binding the contract event 0x346216bb718ede47b40f85e7caf58cd997eebda879d146a8b49091c106c851d9.
//
// Solidity: event TaskAborted(uint256 taskId, string reason)
func (_Task *TaskFilterer) WatchTaskAborted(opts *bind.WatchOpts, sink chan<- *TaskTaskAborted) (event.Subscription, error) {

	logs, sub, err := _Task.contract.WatchLogs(opts, "TaskAborted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskTaskAborted)
				if err := _Task.contract.UnpackLog(event, "TaskAborted", log); err != nil {
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

// ParseTaskAborted is a log parse operation binding the contract event 0x346216bb718ede47b40f85e7caf58cd997eebda879d146a8b49091c106c851d9.
//
// Solidity: event TaskAborted(uint256 taskId, string reason)
func (_Task *TaskFilterer) ParseTaskAborted(log types.Log) (*TaskTaskAborted, error) {
	event := new(TaskTaskAborted)
	if err := _Task.contract.UnpackLog(event, "TaskAborted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaskTaskCreatedIterator is returned from FilterTaskCreated and is used to iterate over the raw logs and unpacked data for TaskCreated events raised by the Task contract.
type TaskTaskCreatedIterator struct {
	Event *TaskTaskCreated // Event containing the contract specifics and raw log

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
func (it *TaskTaskCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskTaskCreated)
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
		it.Event = new(TaskTaskCreated)
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
func (it *TaskTaskCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskTaskCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskTaskCreated represents a TaskCreated event raised by the Task contract.
type TaskTaskCreated struct {
	TaskId       *big.Int
	TaskType     *big.Int
	Creator      common.Address
	SelectedNode common.Address
	TaskHash     [32]byte
	DataHash     [32]byte
	Round        *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTaskCreated is a free log retrieval operation binding the contract event 0xbb7c6b96a3c44e4d2f31e6c30253c53b53f88b24017a672bc1781402f646a8ad.
//
// Solidity: event TaskCreated(uint256 taskId, uint256 taskType, address indexed creator, address indexed selectedNode, bytes32 taskHash, bytes32 dataHash, uint256 round)
func (_Task *TaskFilterer) FilterTaskCreated(opts *bind.FilterOpts, creator []common.Address, selectedNode []common.Address) (*TaskTaskCreatedIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var selectedNodeRule []interface{}
	for _, selectedNodeItem := range selectedNode {
		selectedNodeRule = append(selectedNodeRule, selectedNodeItem)
	}

	logs, sub, err := _Task.contract.FilterLogs(opts, "TaskCreated", creatorRule, selectedNodeRule)
	if err != nil {
		return nil, err
	}
	return &TaskTaskCreatedIterator{contract: _Task.contract, event: "TaskCreated", logs: logs, sub: sub}, nil
}

// WatchTaskCreated is a free log subscription operation binding the contract event 0xbb7c6b96a3c44e4d2f31e6c30253c53b53f88b24017a672bc1781402f646a8ad.
//
// Solidity: event TaskCreated(uint256 taskId, uint256 taskType, address indexed creator, address indexed selectedNode, bytes32 taskHash, bytes32 dataHash, uint256 round)
func (_Task *TaskFilterer) WatchTaskCreated(opts *bind.WatchOpts, sink chan<- *TaskTaskCreated, creator []common.Address, selectedNode []common.Address) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var selectedNodeRule []interface{}
	for _, selectedNodeItem := range selectedNode {
		selectedNodeRule = append(selectedNodeRule, selectedNodeItem)
	}

	logs, sub, err := _Task.contract.WatchLogs(opts, "TaskCreated", creatorRule, selectedNodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskTaskCreated)
				if err := _Task.contract.UnpackLog(event, "TaskCreated", log); err != nil {
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

// ParseTaskCreated is a log parse operation binding the contract event 0xbb7c6b96a3c44e4d2f31e6c30253c53b53f88b24017a672bc1781402f646a8ad.
//
// Solidity: event TaskCreated(uint256 taskId, uint256 taskType, address indexed creator, address indexed selectedNode, bytes32 taskHash, bytes32 dataHash, uint256 round)
func (_Task *TaskFilterer) ParseTaskCreated(log types.Log) (*TaskTaskCreated, error) {
	event := new(TaskTaskCreated)
	if err := _Task.contract.UnpackLog(event, "TaskCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaskTaskResultCommitmentsReadyIterator is returned from FilterTaskResultCommitmentsReady and is used to iterate over the raw logs and unpacked data for TaskResultCommitmentsReady events raised by the Task contract.
type TaskTaskResultCommitmentsReadyIterator struct {
	Event *TaskTaskResultCommitmentsReady // Event containing the contract specifics and raw log

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
func (it *TaskTaskResultCommitmentsReadyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskTaskResultCommitmentsReady)
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
		it.Event = new(TaskTaskResultCommitmentsReady)
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
func (it *TaskTaskResultCommitmentsReadyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskTaskResultCommitmentsReadyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskTaskResultCommitmentsReady represents a TaskResultCommitmentsReady event raised by the Task contract.
type TaskTaskResultCommitmentsReady struct {
	TaskId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTaskResultCommitmentsReady is a free log retrieval operation binding the contract event 0xb16812d8924e5125b1e331ca0097225a801aaa45b056a9ab12ab6ba658e6c9e5.
//
// Solidity: event TaskResultCommitmentsReady(uint256 taskId)
func (_Task *TaskFilterer) FilterTaskResultCommitmentsReady(opts *bind.FilterOpts) (*TaskTaskResultCommitmentsReadyIterator, error) {

	logs, sub, err := _Task.contract.FilterLogs(opts, "TaskResultCommitmentsReady")
	if err != nil {
		return nil, err
	}
	return &TaskTaskResultCommitmentsReadyIterator{contract: _Task.contract, event: "TaskResultCommitmentsReady", logs: logs, sub: sub}, nil
}

// WatchTaskResultCommitmentsReady is a free log subscription operation binding the contract event 0xb16812d8924e5125b1e331ca0097225a801aaa45b056a9ab12ab6ba658e6c9e5.
//
// Solidity: event TaskResultCommitmentsReady(uint256 taskId)
func (_Task *TaskFilterer) WatchTaskResultCommitmentsReady(opts *bind.WatchOpts, sink chan<- *TaskTaskResultCommitmentsReady) (event.Subscription, error) {

	logs, sub, err := _Task.contract.WatchLogs(opts, "TaskResultCommitmentsReady")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskTaskResultCommitmentsReady)
				if err := _Task.contract.UnpackLog(event, "TaskResultCommitmentsReady", log); err != nil {
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

// ParseTaskResultCommitmentsReady is a log parse operation binding the contract event 0xb16812d8924e5125b1e331ca0097225a801aaa45b056a9ab12ab6ba658e6c9e5.
//
// Solidity: event TaskResultCommitmentsReady(uint256 taskId)
func (_Task *TaskFilterer) ParseTaskResultCommitmentsReady(log types.Log) (*TaskTaskResultCommitmentsReady, error) {
	event := new(TaskTaskResultCommitmentsReady)
	if err := _Task.contract.UnpackLog(event, "TaskResultCommitmentsReady", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaskTaskSuccessIterator is returned from FilterTaskSuccess and is used to iterate over the raw logs and unpacked data for TaskSuccess events raised by the Task contract.
type TaskTaskSuccessIterator struct {
	Event *TaskTaskSuccess // Event containing the contract specifics and raw log

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
func (it *TaskTaskSuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskTaskSuccess)
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
		it.Event = new(TaskTaskSuccess)
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
func (it *TaskTaskSuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskTaskSuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskTaskSuccess represents a TaskSuccess event raised by the Task contract.
type TaskTaskSuccess struct {
	TaskId     *big.Int
	Result     []byte
	ResultNode common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTaskSuccess is a free log retrieval operation binding the contract event 0x3cd644a756687c7dd87bbe74e2fd65bde569a3335e26f68f8fa334b2849673f7.
//
// Solidity: event TaskSuccess(uint256 taskId, bytes result, address indexed resultNode)
func (_Task *TaskFilterer) FilterTaskSuccess(opts *bind.FilterOpts, resultNode []common.Address) (*TaskTaskSuccessIterator, error) {

	var resultNodeRule []interface{}
	for _, resultNodeItem := range resultNode {
		resultNodeRule = append(resultNodeRule, resultNodeItem)
	}

	logs, sub, err := _Task.contract.FilterLogs(opts, "TaskSuccess", resultNodeRule)
	if err != nil {
		return nil, err
	}
	return &TaskTaskSuccessIterator{contract: _Task.contract, event: "TaskSuccess", logs: logs, sub: sub}, nil
}

// WatchTaskSuccess is a free log subscription operation binding the contract event 0x3cd644a756687c7dd87bbe74e2fd65bde569a3335e26f68f8fa334b2849673f7.
//
// Solidity: event TaskSuccess(uint256 taskId, bytes result, address indexed resultNode)
func (_Task *TaskFilterer) WatchTaskSuccess(opts *bind.WatchOpts, sink chan<- *TaskTaskSuccess, resultNode []common.Address) (event.Subscription, error) {

	var resultNodeRule []interface{}
	for _, resultNodeItem := range resultNode {
		resultNodeRule = append(resultNodeRule, resultNodeItem)
	}

	logs, sub, err := _Task.contract.WatchLogs(opts, "TaskSuccess", resultNodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskTaskSuccess)
				if err := _Task.contract.UnpackLog(event, "TaskSuccess", log); err != nil {
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

// ParseTaskSuccess is a log parse operation binding the contract event 0x3cd644a756687c7dd87bbe74e2fd65bde569a3335e26f68f8fa334b2849673f7.
//
// Solidity: event TaskSuccess(uint256 taskId, bytes result, address indexed resultNode)
func (_Task *TaskFilterer) ParseTaskSuccess(log types.Log) (*TaskTaskSuccess, error) {
	event := new(TaskTaskSuccess)
	if err := _Task.contract.UnpackLog(event, "TaskSuccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
