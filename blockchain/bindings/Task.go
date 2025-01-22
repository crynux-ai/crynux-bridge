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

// VSSTaskTaskInfo is an auto generated low-level Go binding around an user-defined struct.
type VSSTaskTaskInfo struct {
	TaskType            uint8
	Creator             common.Address
	TaskIDCommitment    [32]byte
	SamplingSeed        [32]byte
	Nonce               [32]byte
	Sequence            *big.Int
	Status              uint8
	SelectedNode        common.Address
	Timeout             *big.Int
	Score               []byte
	TaskFee             *big.Int
	TaskSize            *big.Int
	ModelIDs            []string
	MinimumVRAM         *big.Int
	RequiredGPU         string
	RequiredGPUVRAM     *big.Int
	TaskVersion         [3]*big.Int
	AbortReason         uint8
	Error               uint8
	PaymentAddresses    []common.Address
	Payments            []*big.Int
	CreateTimestamp     *big.Int
	StartTimestamp      *big.Int
	ScoreReadyTimestamp *big.Int
}

// TaskMetaData contains all meta data concerning the Task contract.
var TaskMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractNode\",\"name\":\"nodeInstance\",\"type\":\"address\"},{\"internalType\":\"contractQOS\",\"name\":\"qosInstance\",\"type\":\"address\"},{\"internalType\":\"contractTaskQueue\",\"name\":\"taskQueueInstance\",\"type\":\"address\"},{\"internalType\":\"contractNetworkStats\",\"name\":\"networkStatsInstance\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nodeAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"modelID\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"enumVSSTask.TaskType\",\"name\":\"taskType\",\"type\":\"uint8\"}],\"name\":\"DownloadModel\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"taskIDCommitment\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"abortIssuer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumVSSTask.TaskStatus\",\"name\":\"lastStatus\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"enumVSSTask.TaskAbortReason\",\"name\":\"abortReason\",\"type\":\"uint8\"}],\"name\":\"TaskEndAborted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"taskIDCommitment\",\"type\":\"bytes32\"}],\"name\":\"TaskEndGroupRefund\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"taskIDCommitment\",\"type\":\"bytes32\"}],\"name\":\"TaskEndGroupSuccess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"taskIDCommitment\",\"type\":\"bytes32\"}],\"name\":\"TaskEndInvalidated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"taskIDCommitment\",\"type\":\"bytes32\"}],\"name\":\"TaskEndSuccess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"taskIDCommitment\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"selectedNode\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumVSSTask.TaskError\",\"name\":\"error\",\"type\":\"uint8\"}],\"name\":\"TaskErrorReported\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"taskIDCommitment\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"selectedNode\",\"type\":\"address\"}],\"name\":\"TaskParametersUploaded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"taskIDCommitment\",\"type\":\"bytes32\"}],\"name\":\"TaskQueued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"taskIDCommitment\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"selectedNode\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"taskScore\",\"type\":\"bytes\"}],\"name\":\"TaskScoreReady\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"taskIDCommitment\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"selectedNode\",\"type\":\"address\"}],\"name\":\"TaskStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"taskIDCommitment\",\"type\":\"bytes32\"}],\"name\":\"TaskValidated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"taskIDCommitment\",\"type\":\"bytes32\"},{\"internalType\":\"enumVSSTask.TaskAbortReason\",\"name\":\"abortReason\",\"type\":\"uint8\"}],\"name\":\"abortTask\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumVSSTask.TaskType\",\"name\":\"taskType\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"taskIDCommitment\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"nonce\",\"type\":\"bytes32\"},{\"internalType\":\"string[]\",\"name\":\"modelIDs\",\"type\":\"string[]\"},{\"internalType\":\"uint256\",\"name\":\"minimumVRAM\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"requiredGPU\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"requiredGPUVRAM\",\"type\":\"uint256\"},{\"internalType\":\"uint256[3]\",\"name\":\"taskVersion\",\"type\":\"uint256[3]\"},{\"internalType\":\"uint256\",\"name\":\"taskSize\",\"type\":\"uint256\"}],\"name\":\"createTask\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nodeAddress\",\"type\":\"address\"}],\"name\":\"getNodeTask\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"taskIDCommitment\",\"type\":\"bytes32\"}],\"name\":\"getTask\",\"outputs\":[{\"components\":[{\"internalType\":\"enumVSSTask.TaskType\",\"name\":\"taskType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"taskIDCommitment\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"samplingSeed\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"nonce\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"sequence\",\"type\":\"uint256\"},{\"internalType\":\"enumVSSTask.TaskStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"selectedNode\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"score\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"taskFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"taskSize\",\"type\":\"uint256\"},{\"internalType\":\"string[]\",\"name\":\"modelIDs\",\"type\":\"string[]\"},{\"internalType\":\"uint256\",\"name\":\"minimumVRAM\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"requiredGPU\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"requiredGPUVRAM\",\"type\":\"uint256\"},{\"internalType\":\"uint256[3]\",\"name\":\"taskVersion\",\"type\":\"uint256[3]\"},{\"internalType\":\"enumVSSTask.TaskAbortReason\",\"name\":\"abortReason\",\"type\":\"uint8\"},{\"internalType\":\"enumVSSTask.TaskError\",\"name\":\"error\",\"type\":\"uint8\"},{\"internalType\":\"address[]\",\"name\":\"paymentAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"payments\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"createTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"scoreReadyTimestamp\",\"type\":\"uint256\"}],\"internalType\":\"structVSSTask.TaskInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"root\",\"type\":\"address\"}],\"name\":\"nodeAvailableCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"taskIDCommitment\",\"type\":\"bytes32\"},{\"internalType\":\"enumVSSTask.TaskError\",\"name\":\"error\",\"type\":\"uint8\"}],\"name\":\"reportTaskError\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"taskIDCommitment\",\"type\":\"bytes32\"}],\"name\":\"reportTaskParametersUploaded\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"taskIDCommitment\",\"type\":\"bytes32\"}],\"name\":\"reportTaskResultUploaded\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setRelayAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"taskIDCommitment\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"taskScore\",\"type\":\"bytes\"}],\"name\":\"submitTaskScore\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"}],\"name\":\"updateDistanceThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"t\",\"type\":\"uint256\"}],\"name\":\"updateTimeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"taskIDCommitment\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"vrfProof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"}],\"name\":\"validateSingleTask\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"taskIDCommitment1\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"taskIDCommitment2\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"taskIDCommitment3\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"taskGUID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"vrfProof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"}],\"name\":\"validateTaskGroup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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
// Solidity: function getNodeTask(address nodeAddress) view returns(bytes32)
func (_Task *TaskCaller) GetNodeTask(opts *bind.CallOpts, nodeAddress common.Address) ([32]byte, error) {
	var out []interface{}
	err := _Task.contract.Call(opts, &out, "getNodeTask", nodeAddress)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetNodeTask is a free data retrieval call binding the contract method 0x9877ad45.
//
// Solidity: function getNodeTask(address nodeAddress) view returns(bytes32)
func (_Task *TaskSession) GetNodeTask(nodeAddress common.Address) ([32]byte, error) {
	return _Task.Contract.GetNodeTask(&_Task.CallOpts, nodeAddress)
}

// GetNodeTask is a free data retrieval call binding the contract method 0x9877ad45.
//
// Solidity: function getNodeTask(address nodeAddress) view returns(bytes32)
func (_Task *TaskCallerSession) GetNodeTask(nodeAddress common.Address) ([32]byte, error) {
	return _Task.Contract.GetNodeTask(&_Task.CallOpts, nodeAddress)
}

// GetTask is a free data retrieval call binding the contract method 0x15a29035.
//
// Solidity: function getTask(bytes32 taskIDCommitment) view returns((uint8,address,bytes32,bytes32,bytes32,uint256,uint8,address,uint256,bytes,uint256,uint256,string[],uint256,string,uint256,uint256[3],uint8,uint8,address[],uint256[],uint256,uint256,uint256))
func (_Task *TaskCaller) GetTask(opts *bind.CallOpts, taskIDCommitment [32]byte) (VSSTaskTaskInfo, error) {
	var out []interface{}
	err := _Task.contract.Call(opts, &out, "getTask", taskIDCommitment)

	if err != nil {
		return *new(VSSTaskTaskInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(VSSTaskTaskInfo)).(*VSSTaskTaskInfo)

	return out0, err

}

// GetTask is a free data retrieval call binding the contract method 0x15a29035.
//
// Solidity: function getTask(bytes32 taskIDCommitment) view returns((uint8,address,bytes32,bytes32,bytes32,uint256,uint8,address,uint256,bytes,uint256,uint256,string[],uint256,string,uint256,uint256[3],uint8,uint8,address[],uint256[],uint256,uint256,uint256))
func (_Task *TaskSession) GetTask(taskIDCommitment [32]byte) (VSSTaskTaskInfo, error) {
	return _Task.Contract.GetTask(&_Task.CallOpts, taskIDCommitment)
}

// GetTask is a free data retrieval call binding the contract method 0x15a29035.
//
// Solidity: function getTask(bytes32 taskIDCommitment) view returns((uint8,address,bytes32,bytes32,bytes32,uint256,uint8,address,uint256,bytes,uint256,uint256,string[],uint256,string,uint256,uint256[3],uint8,uint8,address[],uint256[],uint256,uint256,uint256))
func (_Task *TaskCallerSession) GetTask(taskIDCommitment [32]byte) (VSSTaskTaskInfo, error) {
	return _Task.Contract.GetTask(&_Task.CallOpts, taskIDCommitment)
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

// AbortTask is a paid mutator transaction binding the contract method 0xb03d2b11.
//
// Solidity: function abortTask(bytes32 taskIDCommitment, uint8 abortReason) returns()
func (_Task *TaskTransactor) AbortTask(opts *bind.TransactOpts, taskIDCommitment [32]byte, abortReason uint8) (*types.Transaction, error) {
	return _Task.contract.Transact(opts, "abortTask", taskIDCommitment, abortReason)
}

// AbortTask is a paid mutator transaction binding the contract method 0xb03d2b11.
//
// Solidity: function abortTask(bytes32 taskIDCommitment, uint8 abortReason) returns()
func (_Task *TaskSession) AbortTask(taskIDCommitment [32]byte, abortReason uint8) (*types.Transaction, error) {
	return _Task.Contract.AbortTask(&_Task.TransactOpts, taskIDCommitment, abortReason)
}

// AbortTask is a paid mutator transaction binding the contract method 0xb03d2b11.
//
// Solidity: function abortTask(bytes32 taskIDCommitment, uint8 abortReason) returns()
func (_Task *TaskTransactorSession) AbortTask(taskIDCommitment [32]byte, abortReason uint8) (*types.Transaction, error) {
	return _Task.Contract.AbortTask(&_Task.TransactOpts, taskIDCommitment, abortReason)
}

// CreateTask is a paid mutator transaction binding the contract method 0xe3697ca8.
//
// Solidity: function createTask(uint8 taskType, bytes32 taskIDCommitment, bytes32 nonce, string[] modelIDs, uint256 minimumVRAM, string requiredGPU, uint256 requiredGPUVRAM, uint256[3] taskVersion, uint256 taskSize) payable returns()
func (_Task *TaskTransactor) CreateTask(opts *bind.TransactOpts, taskType uint8, taskIDCommitment [32]byte, nonce [32]byte, modelIDs []string, minimumVRAM *big.Int, requiredGPU string, requiredGPUVRAM *big.Int, taskVersion [3]*big.Int, taskSize *big.Int) (*types.Transaction, error) {
	return _Task.contract.Transact(opts, "createTask", taskType, taskIDCommitment, nonce, modelIDs, minimumVRAM, requiredGPU, requiredGPUVRAM, taskVersion, taskSize)
}

// CreateTask is a paid mutator transaction binding the contract method 0xe3697ca8.
//
// Solidity: function createTask(uint8 taskType, bytes32 taskIDCommitment, bytes32 nonce, string[] modelIDs, uint256 minimumVRAM, string requiredGPU, uint256 requiredGPUVRAM, uint256[3] taskVersion, uint256 taskSize) payable returns()
func (_Task *TaskSession) CreateTask(taskType uint8, taskIDCommitment [32]byte, nonce [32]byte, modelIDs []string, minimumVRAM *big.Int, requiredGPU string, requiredGPUVRAM *big.Int, taskVersion [3]*big.Int, taskSize *big.Int) (*types.Transaction, error) {
	return _Task.Contract.CreateTask(&_Task.TransactOpts, taskType, taskIDCommitment, nonce, modelIDs, minimumVRAM, requiredGPU, requiredGPUVRAM, taskVersion, taskSize)
}

// CreateTask is a paid mutator transaction binding the contract method 0xe3697ca8.
//
// Solidity: function createTask(uint8 taskType, bytes32 taskIDCommitment, bytes32 nonce, string[] modelIDs, uint256 minimumVRAM, string requiredGPU, uint256 requiredGPUVRAM, uint256[3] taskVersion, uint256 taskSize) payable returns()
func (_Task *TaskTransactorSession) CreateTask(taskType uint8, taskIDCommitment [32]byte, nonce [32]byte, modelIDs []string, minimumVRAM *big.Int, requiredGPU string, requiredGPUVRAM *big.Int, taskVersion [3]*big.Int, taskSize *big.Int) (*types.Transaction, error) {
	return _Task.Contract.CreateTask(&_Task.TransactOpts, taskType, taskIDCommitment, nonce, modelIDs, minimumVRAM, requiredGPU, requiredGPUVRAM, taskVersion, taskSize)
}

// NodeAvailableCallback is a paid mutator transaction binding the contract method 0xd06ac297.
//
// Solidity: function nodeAvailableCallback(address root) returns()
func (_Task *TaskTransactor) NodeAvailableCallback(opts *bind.TransactOpts, root common.Address) (*types.Transaction, error) {
	return _Task.contract.Transact(opts, "nodeAvailableCallback", root)
}

// NodeAvailableCallback is a paid mutator transaction binding the contract method 0xd06ac297.
//
// Solidity: function nodeAvailableCallback(address root) returns()
func (_Task *TaskSession) NodeAvailableCallback(root common.Address) (*types.Transaction, error) {
	return _Task.Contract.NodeAvailableCallback(&_Task.TransactOpts, root)
}

// NodeAvailableCallback is a paid mutator transaction binding the contract method 0xd06ac297.
//
// Solidity: function nodeAvailableCallback(address root) returns()
func (_Task *TaskTransactorSession) NodeAvailableCallback(root common.Address) (*types.Transaction, error) {
	return _Task.Contract.NodeAvailableCallback(&_Task.TransactOpts, root)
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

// ReportTaskError is a paid mutator transaction binding the contract method 0xcb9c9475.
//
// Solidity: function reportTaskError(bytes32 taskIDCommitment, uint8 error) returns()
func (_Task *TaskTransactor) ReportTaskError(opts *bind.TransactOpts, taskIDCommitment [32]byte, error uint8) (*types.Transaction, error) {
	return _Task.contract.Transact(opts, "reportTaskError", taskIDCommitment, error)
}

// ReportTaskError is a paid mutator transaction binding the contract method 0xcb9c9475.
//
// Solidity: function reportTaskError(bytes32 taskIDCommitment, uint8 error) returns()
func (_Task *TaskSession) ReportTaskError(taskIDCommitment [32]byte, error uint8) (*types.Transaction, error) {
	return _Task.Contract.ReportTaskError(&_Task.TransactOpts, taskIDCommitment, error)
}

// ReportTaskError is a paid mutator transaction binding the contract method 0xcb9c9475.
//
// Solidity: function reportTaskError(bytes32 taskIDCommitment, uint8 error) returns()
func (_Task *TaskTransactorSession) ReportTaskError(taskIDCommitment [32]byte, error uint8) (*types.Transaction, error) {
	return _Task.Contract.ReportTaskError(&_Task.TransactOpts, taskIDCommitment, error)
}

// ReportTaskParametersUploaded is a paid mutator transaction binding the contract method 0xf048ef29.
//
// Solidity: function reportTaskParametersUploaded(bytes32 taskIDCommitment) returns()
func (_Task *TaskTransactor) ReportTaskParametersUploaded(opts *bind.TransactOpts, taskIDCommitment [32]byte) (*types.Transaction, error) {
	return _Task.contract.Transact(opts, "reportTaskParametersUploaded", taskIDCommitment)
}

// ReportTaskParametersUploaded is a paid mutator transaction binding the contract method 0xf048ef29.
//
// Solidity: function reportTaskParametersUploaded(bytes32 taskIDCommitment) returns()
func (_Task *TaskSession) ReportTaskParametersUploaded(taskIDCommitment [32]byte) (*types.Transaction, error) {
	return _Task.Contract.ReportTaskParametersUploaded(&_Task.TransactOpts, taskIDCommitment)
}

// ReportTaskParametersUploaded is a paid mutator transaction binding the contract method 0xf048ef29.
//
// Solidity: function reportTaskParametersUploaded(bytes32 taskIDCommitment) returns()
func (_Task *TaskTransactorSession) ReportTaskParametersUploaded(taskIDCommitment [32]byte) (*types.Transaction, error) {
	return _Task.Contract.ReportTaskParametersUploaded(&_Task.TransactOpts, taskIDCommitment)
}

// ReportTaskResultUploaded is a paid mutator transaction binding the contract method 0x44f76671.
//
// Solidity: function reportTaskResultUploaded(bytes32 taskIDCommitment) returns()
func (_Task *TaskTransactor) ReportTaskResultUploaded(opts *bind.TransactOpts, taskIDCommitment [32]byte) (*types.Transaction, error) {
	return _Task.contract.Transact(opts, "reportTaskResultUploaded", taskIDCommitment)
}

// ReportTaskResultUploaded is a paid mutator transaction binding the contract method 0x44f76671.
//
// Solidity: function reportTaskResultUploaded(bytes32 taskIDCommitment) returns()
func (_Task *TaskSession) ReportTaskResultUploaded(taskIDCommitment [32]byte) (*types.Transaction, error) {
	return _Task.Contract.ReportTaskResultUploaded(&_Task.TransactOpts, taskIDCommitment)
}

// ReportTaskResultUploaded is a paid mutator transaction binding the contract method 0x44f76671.
//
// Solidity: function reportTaskResultUploaded(bytes32 taskIDCommitment) returns()
func (_Task *TaskTransactorSession) ReportTaskResultUploaded(taskIDCommitment [32]byte) (*types.Transaction, error) {
	return _Task.Contract.ReportTaskResultUploaded(&_Task.TransactOpts, taskIDCommitment)
}

// SetRelayAddress is a paid mutator transaction binding the contract method 0xaeba8f3c.
//
// Solidity: function setRelayAddress(address addr) returns()
func (_Task *TaskTransactor) SetRelayAddress(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Task.contract.Transact(opts, "setRelayAddress", addr)
}

// SetRelayAddress is a paid mutator transaction binding the contract method 0xaeba8f3c.
//
// Solidity: function setRelayAddress(address addr) returns()
func (_Task *TaskSession) SetRelayAddress(addr common.Address) (*types.Transaction, error) {
	return _Task.Contract.SetRelayAddress(&_Task.TransactOpts, addr)
}

// SetRelayAddress is a paid mutator transaction binding the contract method 0xaeba8f3c.
//
// Solidity: function setRelayAddress(address addr) returns()
func (_Task *TaskTransactorSession) SetRelayAddress(addr common.Address) (*types.Transaction, error) {
	return _Task.Contract.SetRelayAddress(&_Task.TransactOpts, addr)
}

// SubmitTaskScore is a paid mutator transaction binding the contract method 0x35b6175e.
//
// Solidity: function submitTaskScore(bytes32 taskIDCommitment, bytes taskScore) returns()
func (_Task *TaskTransactor) SubmitTaskScore(opts *bind.TransactOpts, taskIDCommitment [32]byte, taskScore []byte) (*types.Transaction, error) {
	return _Task.contract.Transact(opts, "submitTaskScore", taskIDCommitment, taskScore)
}

// SubmitTaskScore is a paid mutator transaction binding the contract method 0x35b6175e.
//
// Solidity: function submitTaskScore(bytes32 taskIDCommitment, bytes taskScore) returns()
func (_Task *TaskSession) SubmitTaskScore(taskIDCommitment [32]byte, taskScore []byte) (*types.Transaction, error) {
	return _Task.Contract.SubmitTaskScore(&_Task.TransactOpts, taskIDCommitment, taskScore)
}

// SubmitTaskScore is a paid mutator transaction binding the contract method 0x35b6175e.
//
// Solidity: function submitTaskScore(bytes32 taskIDCommitment, bytes taskScore) returns()
func (_Task *TaskTransactorSession) SubmitTaskScore(taskIDCommitment [32]byte, taskScore []byte) (*types.Transaction, error) {
	return _Task.Contract.SubmitTaskScore(&_Task.TransactOpts, taskIDCommitment, taskScore)
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

// ValidateSingleTask is a paid mutator transaction binding the contract method 0x93c7c00b.
//
// Solidity: function validateSingleTask(bytes32 taskIDCommitment, bytes vrfProof, bytes publicKey) returns()
func (_Task *TaskTransactor) ValidateSingleTask(opts *bind.TransactOpts, taskIDCommitment [32]byte, vrfProof []byte, publicKey []byte) (*types.Transaction, error) {
	return _Task.contract.Transact(opts, "validateSingleTask", taskIDCommitment, vrfProof, publicKey)
}

// ValidateSingleTask is a paid mutator transaction binding the contract method 0x93c7c00b.
//
// Solidity: function validateSingleTask(bytes32 taskIDCommitment, bytes vrfProof, bytes publicKey) returns()
func (_Task *TaskSession) ValidateSingleTask(taskIDCommitment [32]byte, vrfProof []byte, publicKey []byte) (*types.Transaction, error) {
	return _Task.Contract.ValidateSingleTask(&_Task.TransactOpts, taskIDCommitment, vrfProof, publicKey)
}

// ValidateSingleTask is a paid mutator transaction binding the contract method 0x93c7c00b.
//
// Solidity: function validateSingleTask(bytes32 taskIDCommitment, bytes vrfProof, bytes publicKey) returns()
func (_Task *TaskTransactorSession) ValidateSingleTask(taskIDCommitment [32]byte, vrfProof []byte, publicKey []byte) (*types.Transaction, error) {
	return _Task.Contract.ValidateSingleTask(&_Task.TransactOpts, taskIDCommitment, vrfProof, publicKey)
}

// ValidateTaskGroup is a paid mutator transaction binding the contract method 0x127be3ca.
//
// Solidity: function validateTaskGroup(bytes32 taskIDCommitment1, bytes32 taskIDCommitment2, bytes32 taskIDCommitment3, bytes32 taskGUID, bytes vrfProof, bytes publicKey) returns()
func (_Task *TaskTransactor) ValidateTaskGroup(opts *bind.TransactOpts, taskIDCommitment1 [32]byte, taskIDCommitment2 [32]byte, taskIDCommitment3 [32]byte, taskGUID [32]byte, vrfProof []byte, publicKey []byte) (*types.Transaction, error) {
	return _Task.contract.Transact(opts, "validateTaskGroup", taskIDCommitment1, taskIDCommitment2, taskIDCommitment3, taskGUID, vrfProof, publicKey)
}

// ValidateTaskGroup is a paid mutator transaction binding the contract method 0x127be3ca.
//
// Solidity: function validateTaskGroup(bytes32 taskIDCommitment1, bytes32 taskIDCommitment2, bytes32 taskIDCommitment3, bytes32 taskGUID, bytes vrfProof, bytes publicKey) returns()
func (_Task *TaskSession) ValidateTaskGroup(taskIDCommitment1 [32]byte, taskIDCommitment2 [32]byte, taskIDCommitment3 [32]byte, taskGUID [32]byte, vrfProof []byte, publicKey []byte) (*types.Transaction, error) {
	return _Task.Contract.ValidateTaskGroup(&_Task.TransactOpts, taskIDCommitment1, taskIDCommitment2, taskIDCommitment3, taskGUID, vrfProof, publicKey)
}

// ValidateTaskGroup is a paid mutator transaction binding the contract method 0x127be3ca.
//
// Solidity: function validateTaskGroup(bytes32 taskIDCommitment1, bytes32 taskIDCommitment2, bytes32 taskIDCommitment3, bytes32 taskGUID, bytes vrfProof, bytes publicKey) returns()
func (_Task *TaskTransactorSession) ValidateTaskGroup(taskIDCommitment1 [32]byte, taskIDCommitment2 [32]byte, taskIDCommitment3 [32]byte, taskGUID [32]byte, vrfProof []byte, publicKey []byte) (*types.Transaction, error) {
	return _Task.Contract.ValidateTaskGroup(&_Task.TransactOpts, taskIDCommitment1, taskIDCommitment2, taskIDCommitment3, taskGUID, vrfProof, publicKey)
}

// TaskDownloadModelIterator is returned from FilterDownloadModel and is used to iterate over the raw logs and unpacked data for DownloadModel events raised by the Task contract.
type TaskDownloadModelIterator struct {
	Event *TaskDownloadModel // Event containing the contract specifics and raw log

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
func (it *TaskDownloadModelIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskDownloadModel)
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
		it.Event = new(TaskDownloadModel)
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
func (it *TaskDownloadModelIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskDownloadModelIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskDownloadModel represents a DownloadModel event raised by the Task contract.
type TaskDownloadModel struct {
	NodeAddress common.Address
	ModelID     string
	TaskType    uint8
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDownloadModel is a free log retrieval operation binding the contract event 0x2bb8f1e759285a56f96394fa1cc9f2583f0b954c8291d771cdef39177614d157.
//
// Solidity: event DownloadModel(address nodeAddress, string modelID, uint8 taskType)
func (_Task *TaskFilterer) FilterDownloadModel(opts *bind.FilterOpts) (*TaskDownloadModelIterator, error) {

	logs, sub, err := _Task.contract.FilterLogs(opts, "DownloadModel")
	if err != nil {
		return nil, err
	}
	return &TaskDownloadModelIterator{contract: _Task.contract, event: "DownloadModel", logs: logs, sub: sub}, nil
}

// WatchDownloadModel is a free log subscription operation binding the contract event 0x2bb8f1e759285a56f96394fa1cc9f2583f0b954c8291d771cdef39177614d157.
//
// Solidity: event DownloadModel(address nodeAddress, string modelID, uint8 taskType)
func (_Task *TaskFilterer) WatchDownloadModel(opts *bind.WatchOpts, sink chan<- *TaskDownloadModel) (event.Subscription, error) {

	logs, sub, err := _Task.contract.WatchLogs(opts, "DownloadModel")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskDownloadModel)
				if err := _Task.contract.UnpackLog(event, "DownloadModel", log); err != nil {
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

// ParseDownloadModel is a log parse operation binding the contract event 0x2bb8f1e759285a56f96394fa1cc9f2583f0b954c8291d771cdef39177614d157.
//
// Solidity: event DownloadModel(address nodeAddress, string modelID, uint8 taskType)
func (_Task *TaskFilterer) ParseDownloadModel(log types.Log) (*TaskDownloadModel, error) {
	event := new(TaskDownloadModel)
	if err := _Task.contract.UnpackLog(event, "DownloadModel", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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

// TaskTaskEndAbortedIterator is returned from FilterTaskEndAborted and is used to iterate over the raw logs and unpacked data for TaskEndAborted events raised by the Task contract.
type TaskTaskEndAbortedIterator struct {
	Event *TaskTaskEndAborted // Event containing the contract specifics and raw log

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
func (it *TaskTaskEndAbortedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskTaskEndAborted)
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
		it.Event = new(TaskTaskEndAborted)
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
func (it *TaskTaskEndAbortedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskTaskEndAbortedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskTaskEndAborted represents a TaskEndAborted event raised by the Task contract.
type TaskTaskEndAborted struct {
	TaskIDCommitment [32]byte
	AbortIssuer      common.Address
	LastStatus       uint8
	AbortReason      uint8
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterTaskEndAborted is a free log retrieval operation binding the contract event 0xe4161023d495c9541aba2a4e2be52416c6da996acb4a4481329f4ec310f33046.
//
// Solidity: event TaskEndAborted(bytes32 taskIDCommitment, address abortIssuer, uint8 lastStatus, uint8 abortReason)
func (_Task *TaskFilterer) FilterTaskEndAborted(opts *bind.FilterOpts) (*TaskTaskEndAbortedIterator, error) {

	logs, sub, err := _Task.contract.FilterLogs(opts, "TaskEndAborted")
	if err != nil {
		return nil, err
	}
	return &TaskTaskEndAbortedIterator{contract: _Task.contract, event: "TaskEndAborted", logs: logs, sub: sub}, nil
}

// WatchTaskEndAborted is a free log subscription operation binding the contract event 0xe4161023d495c9541aba2a4e2be52416c6da996acb4a4481329f4ec310f33046.
//
// Solidity: event TaskEndAborted(bytes32 taskIDCommitment, address abortIssuer, uint8 lastStatus, uint8 abortReason)
func (_Task *TaskFilterer) WatchTaskEndAborted(opts *bind.WatchOpts, sink chan<- *TaskTaskEndAborted) (event.Subscription, error) {

	logs, sub, err := _Task.contract.WatchLogs(opts, "TaskEndAborted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskTaskEndAborted)
				if err := _Task.contract.UnpackLog(event, "TaskEndAborted", log); err != nil {
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

// ParseTaskEndAborted is a log parse operation binding the contract event 0xe4161023d495c9541aba2a4e2be52416c6da996acb4a4481329f4ec310f33046.
//
// Solidity: event TaskEndAborted(bytes32 taskIDCommitment, address abortIssuer, uint8 lastStatus, uint8 abortReason)
func (_Task *TaskFilterer) ParseTaskEndAborted(log types.Log) (*TaskTaskEndAborted, error) {
	event := new(TaskTaskEndAborted)
	if err := _Task.contract.UnpackLog(event, "TaskEndAborted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaskTaskEndGroupRefundIterator is returned from FilterTaskEndGroupRefund and is used to iterate over the raw logs and unpacked data for TaskEndGroupRefund events raised by the Task contract.
type TaskTaskEndGroupRefundIterator struct {
	Event *TaskTaskEndGroupRefund // Event containing the contract specifics and raw log

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
func (it *TaskTaskEndGroupRefundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskTaskEndGroupRefund)
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
		it.Event = new(TaskTaskEndGroupRefund)
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
func (it *TaskTaskEndGroupRefundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskTaskEndGroupRefundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskTaskEndGroupRefund represents a TaskEndGroupRefund event raised by the Task contract.
type TaskTaskEndGroupRefund struct {
	TaskIDCommitment [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterTaskEndGroupRefund is a free log retrieval operation binding the contract event 0x013766fc0b6dd5ec82a37d86477d3f43021d4e1bd973cf17d16ddf174c6e1a2a.
//
// Solidity: event TaskEndGroupRefund(bytes32 taskIDCommitment)
func (_Task *TaskFilterer) FilterTaskEndGroupRefund(opts *bind.FilterOpts) (*TaskTaskEndGroupRefundIterator, error) {

	logs, sub, err := _Task.contract.FilterLogs(opts, "TaskEndGroupRefund")
	if err != nil {
		return nil, err
	}
	return &TaskTaskEndGroupRefundIterator{contract: _Task.contract, event: "TaskEndGroupRefund", logs: logs, sub: sub}, nil
}

// WatchTaskEndGroupRefund is a free log subscription operation binding the contract event 0x013766fc0b6dd5ec82a37d86477d3f43021d4e1bd973cf17d16ddf174c6e1a2a.
//
// Solidity: event TaskEndGroupRefund(bytes32 taskIDCommitment)
func (_Task *TaskFilterer) WatchTaskEndGroupRefund(opts *bind.WatchOpts, sink chan<- *TaskTaskEndGroupRefund) (event.Subscription, error) {

	logs, sub, err := _Task.contract.WatchLogs(opts, "TaskEndGroupRefund")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskTaskEndGroupRefund)
				if err := _Task.contract.UnpackLog(event, "TaskEndGroupRefund", log); err != nil {
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

// ParseTaskEndGroupRefund is a log parse operation binding the contract event 0x013766fc0b6dd5ec82a37d86477d3f43021d4e1bd973cf17d16ddf174c6e1a2a.
//
// Solidity: event TaskEndGroupRefund(bytes32 taskIDCommitment)
func (_Task *TaskFilterer) ParseTaskEndGroupRefund(log types.Log) (*TaskTaskEndGroupRefund, error) {
	event := new(TaskTaskEndGroupRefund)
	if err := _Task.contract.UnpackLog(event, "TaskEndGroupRefund", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaskTaskEndGroupSuccessIterator is returned from FilterTaskEndGroupSuccess and is used to iterate over the raw logs and unpacked data for TaskEndGroupSuccess events raised by the Task contract.
type TaskTaskEndGroupSuccessIterator struct {
	Event *TaskTaskEndGroupSuccess // Event containing the contract specifics and raw log

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
func (it *TaskTaskEndGroupSuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskTaskEndGroupSuccess)
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
		it.Event = new(TaskTaskEndGroupSuccess)
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
func (it *TaskTaskEndGroupSuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskTaskEndGroupSuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskTaskEndGroupSuccess represents a TaskEndGroupSuccess event raised by the Task contract.
type TaskTaskEndGroupSuccess struct {
	TaskIDCommitment [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterTaskEndGroupSuccess is a free log retrieval operation binding the contract event 0x34615f6df03b52f048165d16677ccd9416d392bdaf1a5540030abc3108b5134f.
//
// Solidity: event TaskEndGroupSuccess(bytes32 taskIDCommitment)
func (_Task *TaskFilterer) FilterTaskEndGroupSuccess(opts *bind.FilterOpts) (*TaskTaskEndGroupSuccessIterator, error) {

	logs, sub, err := _Task.contract.FilterLogs(opts, "TaskEndGroupSuccess")
	if err != nil {
		return nil, err
	}
	return &TaskTaskEndGroupSuccessIterator{contract: _Task.contract, event: "TaskEndGroupSuccess", logs: logs, sub: sub}, nil
}

// WatchTaskEndGroupSuccess is a free log subscription operation binding the contract event 0x34615f6df03b52f048165d16677ccd9416d392bdaf1a5540030abc3108b5134f.
//
// Solidity: event TaskEndGroupSuccess(bytes32 taskIDCommitment)
func (_Task *TaskFilterer) WatchTaskEndGroupSuccess(opts *bind.WatchOpts, sink chan<- *TaskTaskEndGroupSuccess) (event.Subscription, error) {

	logs, sub, err := _Task.contract.WatchLogs(opts, "TaskEndGroupSuccess")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskTaskEndGroupSuccess)
				if err := _Task.contract.UnpackLog(event, "TaskEndGroupSuccess", log); err != nil {
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

// ParseTaskEndGroupSuccess is a log parse operation binding the contract event 0x34615f6df03b52f048165d16677ccd9416d392bdaf1a5540030abc3108b5134f.
//
// Solidity: event TaskEndGroupSuccess(bytes32 taskIDCommitment)
func (_Task *TaskFilterer) ParseTaskEndGroupSuccess(log types.Log) (*TaskTaskEndGroupSuccess, error) {
	event := new(TaskTaskEndGroupSuccess)
	if err := _Task.contract.UnpackLog(event, "TaskEndGroupSuccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaskTaskEndInvalidatedIterator is returned from FilterTaskEndInvalidated and is used to iterate over the raw logs and unpacked data for TaskEndInvalidated events raised by the Task contract.
type TaskTaskEndInvalidatedIterator struct {
	Event *TaskTaskEndInvalidated // Event containing the contract specifics and raw log

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
func (it *TaskTaskEndInvalidatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskTaskEndInvalidated)
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
		it.Event = new(TaskTaskEndInvalidated)
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
func (it *TaskTaskEndInvalidatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskTaskEndInvalidatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskTaskEndInvalidated represents a TaskEndInvalidated event raised by the Task contract.
type TaskTaskEndInvalidated struct {
	TaskIDCommitment [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterTaskEndInvalidated is a free log retrieval operation binding the contract event 0x7dd089995216e4d8cb199db2119aaa97782d4b5014824e145e0b30bb7681b099.
//
// Solidity: event TaskEndInvalidated(bytes32 taskIDCommitment)
func (_Task *TaskFilterer) FilterTaskEndInvalidated(opts *bind.FilterOpts) (*TaskTaskEndInvalidatedIterator, error) {

	logs, sub, err := _Task.contract.FilterLogs(opts, "TaskEndInvalidated")
	if err != nil {
		return nil, err
	}
	return &TaskTaskEndInvalidatedIterator{contract: _Task.contract, event: "TaskEndInvalidated", logs: logs, sub: sub}, nil
}

// WatchTaskEndInvalidated is a free log subscription operation binding the contract event 0x7dd089995216e4d8cb199db2119aaa97782d4b5014824e145e0b30bb7681b099.
//
// Solidity: event TaskEndInvalidated(bytes32 taskIDCommitment)
func (_Task *TaskFilterer) WatchTaskEndInvalidated(opts *bind.WatchOpts, sink chan<- *TaskTaskEndInvalidated) (event.Subscription, error) {

	logs, sub, err := _Task.contract.WatchLogs(opts, "TaskEndInvalidated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskTaskEndInvalidated)
				if err := _Task.contract.UnpackLog(event, "TaskEndInvalidated", log); err != nil {
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

// ParseTaskEndInvalidated is a log parse operation binding the contract event 0x7dd089995216e4d8cb199db2119aaa97782d4b5014824e145e0b30bb7681b099.
//
// Solidity: event TaskEndInvalidated(bytes32 taskIDCommitment)
func (_Task *TaskFilterer) ParseTaskEndInvalidated(log types.Log) (*TaskTaskEndInvalidated, error) {
	event := new(TaskTaskEndInvalidated)
	if err := _Task.contract.UnpackLog(event, "TaskEndInvalidated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaskTaskEndSuccessIterator is returned from FilterTaskEndSuccess and is used to iterate over the raw logs and unpacked data for TaskEndSuccess events raised by the Task contract.
type TaskTaskEndSuccessIterator struct {
	Event *TaskTaskEndSuccess // Event containing the contract specifics and raw log

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
func (it *TaskTaskEndSuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskTaskEndSuccess)
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
		it.Event = new(TaskTaskEndSuccess)
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
func (it *TaskTaskEndSuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskTaskEndSuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskTaskEndSuccess represents a TaskEndSuccess event raised by the Task contract.
type TaskTaskEndSuccess struct {
	TaskIDCommitment [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterTaskEndSuccess is a free log retrieval operation binding the contract event 0xedbb681c299fed103abc0767f9969e4891c7f3e23b9fbb0d0bab0a2a3018fe65.
//
// Solidity: event TaskEndSuccess(bytes32 taskIDCommitment)
func (_Task *TaskFilterer) FilterTaskEndSuccess(opts *bind.FilterOpts) (*TaskTaskEndSuccessIterator, error) {

	logs, sub, err := _Task.contract.FilterLogs(opts, "TaskEndSuccess")
	if err != nil {
		return nil, err
	}
	return &TaskTaskEndSuccessIterator{contract: _Task.contract, event: "TaskEndSuccess", logs: logs, sub: sub}, nil
}

// WatchTaskEndSuccess is a free log subscription operation binding the contract event 0xedbb681c299fed103abc0767f9969e4891c7f3e23b9fbb0d0bab0a2a3018fe65.
//
// Solidity: event TaskEndSuccess(bytes32 taskIDCommitment)
func (_Task *TaskFilterer) WatchTaskEndSuccess(opts *bind.WatchOpts, sink chan<- *TaskTaskEndSuccess) (event.Subscription, error) {

	logs, sub, err := _Task.contract.WatchLogs(opts, "TaskEndSuccess")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskTaskEndSuccess)
				if err := _Task.contract.UnpackLog(event, "TaskEndSuccess", log); err != nil {
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

// ParseTaskEndSuccess is a log parse operation binding the contract event 0xedbb681c299fed103abc0767f9969e4891c7f3e23b9fbb0d0bab0a2a3018fe65.
//
// Solidity: event TaskEndSuccess(bytes32 taskIDCommitment)
func (_Task *TaskFilterer) ParseTaskEndSuccess(log types.Log) (*TaskTaskEndSuccess, error) {
	event := new(TaskTaskEndSuccess)
	if err := _Task.contract.UnpackLog(event, "TaskEndSuccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaskTaskErrorReportedIterator is returned from FilterTaskErrorReported and is used to iterate over the raw logs and unpacked data for TaskErrorReported events raised by the Task contract.
type TaskTaskErrorReportedIterator struct {
	Event *TaskTaskErrorReported // Event containing the contract specifics and raw log

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
func (it *TaskTaskErrorReportedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskTaskErrorReported)
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
		it.Event = new(TaskTaskErrorReported)
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
func (it *TaskTaskErrorReportedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskTaskErrorReportedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskTaskErrorReported represents a TaskErrorReported event raised by the Task contract.
type TaskTaskErrorReported struct {
	TaskIDCommitment [32]byte
	SelectedNode     common.Address
	Error            uint8
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterTaskErrorReported is a free log retrieval operation binding the contract event 0x3c1ceb9f05e479d4405f57845f8486cc0bc0045bb1bc9a185b82d19866eafa1c.
//
// Solidity: event TaskErrorReported(bytes32 taskIDCommitment, address selectedNode, uint8 error)
func (_Task *TaskFilterer) FilterTaskErrorReported(opts *bind.FilterOpts) (*TaskTaskErrorReportedIterator, error) {

	logs, sub, err := _Task.contract.FilterLogs(opts, "TaskErrorReported")
	if err != nil {
		return nil, err
	}
	return &TaskTaskErrorReportedIterator{contract: _Task.contract, event: "TaskErrorReported", logs: logs, sub: sub}, nil
}

// WatchTaskErrorReported is a free log subscription operation binding the contract event 0x3c1ceb9f05e479d4405f57845f8486cc0bc0045bb1bc9a185b82d19866eafa1c.
//
// Solidity: event TaskErrorReported(bytes32 taskIDCommitment, address selectedNode, uint8 error)
func (_Task *TaskFilterer) WatchTaskErrorReported(opts *bind.WatchOpts, sink chan<- *TaskTaskErrorReported) (event.Subscription, error) {

	logs, sub, err := _Task.contract.WatchLogs(opts, "TaskErrorReported")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskTaskErrorReported)
				if err := _Task.contract.UnpackLog(event, "TaskErrorReported", log); err != nil {
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

// ParseTaskErrorReported is a log parse operation binding the contract event 0x3c1ceb9f05e479d4405f57845f8486cc0bc0045bb1bc9a185b82d19866eafa1c.
//
// Solidity: event TaskErrorReported(bytes32 taskIDCommitment, address selectedNode, uint8 error)
func (_Task *TaskFilterer) ParseTaskErrorReported(log types.Log) (*TaskTaskErrorReported, error) {
	event := new(TaskTaskErrorReported)
	if err := _Task.contract.UnpackLog(event, "TaskErrorReported", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaskTaskParametersUploadedIterator is returned from FilterTaskParametersUploaded and is used to iterate over the raw logs and unpacked data for TaskParametersUploaded events raised by the Task contract.
type TaskTaskParametersUploadedIterator struct {
	Event *TaskTaskParametersUploaded // Event containing the contract specifics and raw log

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
func (it *TaskTaskParametersUploadedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskTaskParametersUploaded)
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
		it.Event = new(TaskTaskParametersUploaded)
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
func (it *TaskTaskParametersUploadedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskTaskParametersUploadedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskTaskParametersUploaded represents a TaskParametersUploaded event raised by the Task contract.
type TaskTaskParametersUploaded struct {
	TaskIDCommitment [32]byte
	SelectedNode     common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterTaskParametersUploaded is a free log retrieval operation binding the contract event 0x414ae8bfed908fc770ab249e9e7e13e2c2c023995d7a718453702fbd7f5d46a3.
//
// Solidity: event TaskParametersUploaded(bytes32 taskIDCommitment, address selectedNode)
func (_Task *TaskFilterer) FilterTaskParametersUploaded(opts *bind.FilterOpts) (*TaskTaskParametersUploadedIterator, error) {

	logs, sub, err := _Task.contract.FilterLogs(opts, "TaskParametersUploaded")
	if err != nil {
		return nil, err
	}
	return &TaskTaskParametersUploadedIterator{contract: _Task.contract, event: "TaskParametersUploaded", logs: logs, sub: sub}, nil
}

// WatchTaskParametersUploaded is a free log subscription operation binding the contract event 0x414ae8bfed908fc770ab249e9e7e13e2c2c023995d7a718453702fbd7f5d46a3.
//
// Solidity: event TaskParametersUploaded(bytes32 taskIDCommitment, address selectedNode)
func (_Task *TaskFilterer) WatchTaskParametersUploaded(opts *bind.WatchOpts, sink chan<- *TaskTaskParametersUploaded) (event.Subscription, error) {

	logs, sub, err := _Task.contract.WatchLogs(opts, "TaskParametersUploaded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskTaskParametersUploaded)
				if err := _Task.contract.UnpackLog(event, "TaskParametersUploaded", log); err != nil {
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

// ParseTaskParametersUploaded is a log parse operation binding the contract event 0x414ae8bfed908fc770ab249e9e7e13e2c2c023995d7a718453702fbd7f5d46a3.
//
// Solidity: event TaskParametersUploaded(bytes32 taskIDCommitment, address selectedNode)
func (_Task *TaskFilterer) ParseTaskParametersUploaded(log types.Log) (*TaskTaskParametersUploaded, error) {
	event := new(TaskTaskParametersUploaded)
	if err := _Task.contract.UnpackLog(event, "TaskParametersUploaded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaskTaskQueuedIterator is returned from FilterTaskQueued and is used to iterate over the raw logs and unpacked data for TaskQueued events raised by the Task contract.
type TaskTaskQueuedIterator struct {
	Event *TaskTaskQueued // Event containing the contract specifics and raw log

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
func (it *TaskTaskQueuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskTaskQueued)
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
		it.Event = new(TaskTaskQueued)
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
func (it *TaskTaskQueuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskTaskQueuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskTaskQueued represents a TaskQueued event raised by the Task contract.
type TaskTaskQueued struct {
	TaskIDCommitment [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterTaskQueued is a free log retrieval operation binding the contract event 0x24d242f18ee31f8680a00041d1cc841752e13001a3a9a67f5c7562ba5fe67b2b.
//
// Solidity: event TaskQueued(bytes32 taskIDCommitment)
func (_Task *TaskFilterer) FilterTaskQueued(opts *bind.FilterOpts) (*TaskTaskQueuedIterator, error) {

	logs, sub, err := _Task.contract.FilterLogs(opts, "TaskQueued")
	if err != nil {
		return nil, err
	}
	return &TaskTaskQueuedIterator{contract: _Task.contract, event: "TaskQueued", logs: logs, sub: sub}, nil
}

// WatchTaskQueued is a free log subscription operation binding the contract event 0x24d242f18ee31f8680a00041d1cc841752e13001a3a9a67f5c7562ba5fe67b2b.
//
// Solidity: event TaskQueued(bytes32 taskIDCommitment)
func (_Task *TaskFilterer) WatchTaskQueued(opts *bind.WatchOpts, sink chan<- *TaskTaskQueued) (event.Subscription, error) {

	logs, sub, err := _Task.contract.WatchLogs(opts, "TaskQueued")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskTaskQueued)
				if err := _Task.contract.UnpackLog(event, "TaskQueued", log); err != nil {
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

// ParseTaskQueued is a log parse operation binding the contract event 0x24d242f18ee31f8680a00041d1cc841752e13001a3a9a67f5c7562ba5fe67b2b.
//
// Solidity: event TaskQueued(bytes32 taskIDCommitment)
func (_Task *TaskFilterer) ParseTaskQueued(log types.Log) (*TaskTaskQueued, error) {
	event := new(TaskTaskQueued)
	if err := _Task.contract.UnpackLog(event, "TaskQueued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaskTaskScoreReadyIterator is returned from FilterTaskScoreReady and is used to iterate over the raw logs and unpacked data for TaskScoreReady events raised by the Task contract.
type TaskTaskScoreReadyIterator struct {
	Event *TaskTaskScoreReady // Event containing the contract specifics and raw log

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
func (it *TaskTaskScoreReadyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskTaskScoreReady)
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
		it.Event = new(TaskTaskScoreReady)
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
func (it *TaskTaskScoreReadyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskTaskScoreReadyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskTaskScoreReady represents a TaskScoreReady event raised by the Task contract.
type TaskTaskScoreReady struct {
	TaskIDCommitment [32]byte
	SelectedNode     common.Address
	TaskScore        []byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterTaskScoreReady is a free log retrieval operation binding the contract event 0x9aeb9b7daec64fcb1c8ca252f759fafa84039670d35f6265841a690646707725.
//
// Solidity: event TaskScoreReady(bytes32 taskIDCommitment, address selectedNode, bytes taskScore)
func (_Task *TaskFilterer) FilterTaskScoreReady(opts *bind.FilterOpts) (*TaskTaskScoreReadyIterator, error) {

	logs, sub, err := _Task.contract.FilterLogs(opts, "TaskScoreReady")
	if err != nil {
		return nil, err
	}
	return &TaskTaskScoreReadyIterator{contract: _Task.contract, event: "TaskScoreReady", logs: logs, sub: sub}, nil
}

// WatchTaskScoreReady is a free log subscription operation binding the contract event 0x9aeb9b7daec64fcb1c8ca252f759fafa84039670d35f6265841a690646707725.
//
// Solidity: event TaskScoreReady(bytes32 taskIDCommitment, address selectedNode, bytes taskScore)
func (_Task *TaskFilterer) WatchTaskScoreReady(opts *bind.WatchOpts, sink chan<- *TaskTaskScoreReady) (event.Subscription, error) {

	logs, sub, err := _Task.contract.WatchLogs(opts, "TaskScoreReady")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskTaskScoreReady)
				if err := _Task.contract.UnpackLog(event, "TaskScoreReady", log); err != nil {
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

// ParseTaskScoreReady is a log parse operation binding the contract event 0x9aeb9b7daec64fcb1c8ca252f759fafa84039670d35f6265841a690646707725.
//
// Solidity: event TaskScoreReady(bytes32 taskIDCommitment, address selectedNode, bytes taskScore)
func (_Task *TaskFilterer) ParseTaskScoreReady(log types.Log) (*TaskTaskScoreReady, error) {
	event := new(TaskTaskScoreReady)
	if err := _Task.contract.UnpackLog(event, "TaskScoreReady", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaskTaskStartedIterator is returned from FilterTaskStarted and is used to iterate over the raw logs and unpacked data for TaskStarted events raised by the Task contract.
type TaskTaskStartedIterator struct {
	Event *TaskTaskStarted // Event containing the contract specifics and raw log

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
func (it *TaskTaskStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskTaskStarted)
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
		it.Event = new(TaskTaskStarted)
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
func (it *TaskTaskStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskTaskStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskTaskStarted represents a TaskStarted event raised by the Task contract.
type TaskTaskStarted struct {
	TaskIDCommitment [32]byte
	SelectedNode     common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterTaskStarted is a free log retrieval operation binding the contract event 0x9bd27d3fcd577a82b21e57704c0d97fad539157111679935b4584f4e1abe37dc.
//
// Solidity: event TaskStarted(bytes32 taskIDCommitment, address selectedNode)
func (_Task *TaskFilterer) FilterTaskStarted(opts *bind.FilterOpts) (*TaskTaskStartedIterator, error) {

	logs, sub, err := _Task.contract.FilterLogs(opts, "TaskStarted")
	if err != nil {
		return nil, err
	}
	return &TaskTaskStartedIterator{contract: _Task.contract, event: "TaskStarted", logs: logs, sub: sub}, nil
}

// WatchTaskStarted is a free log subscription operation binding the contract event 0x9bd27d3fcd577a82b21e57704c0d97fad539157111679935b4584f4e1abe37dc.
//
// Solidity: event TaskStarted(bytes32 taskIDCommitment, address selectedNode)
func (_Task *TaskFilterer) WatchTaskStarted(opts *bind.WatchOpts, sink chan<- *TaskTaskStarted) (event.Subscription, error) {

	logs, sub, err := _Task.contract.WatchLogs(opts, "TaskStarted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskTaskStarted)
				if err := _Task.contract.UnpackLog(event, "TaskStarted", log); err != nil {
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

// ParseTaskStarted is a log parse operation binding the contract event 0x9bd27d3fcd577a82b21e57704c0d97fad539157111679935b4584f4e1abe37dc.
//
// Solidity: event TaskStarted(bytes32 taskIDCommitment, address selectedNode)
func (_Task *TaskFilterer) ParseTaskStarted(log types.Log) (*TaskTaskStarted, error) {
	event := new(TaskTaskStarted)
	if err := _Task.contract.UnpackLog(event, "TaskStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaskTaskValidatedIterator is returned from FilterTaskValidated and is used to iterate over the raw logs and unpacked data for TaskValidated events raised by the Task contract.
type TaskTaskValidatedIterator struct {
	Event *TaskTaskValidated // Event containing the contract specifics and raw log

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
func (it *TaskTaskValidatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskTaskValidated)
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
		it.Event = new(TaskTaskValidated)
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
func (it *TaskTaskValidatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskTaskValidatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskTaskValidated represents a TaskValidated event raised by the Task contract.
type TaskTaskValidated struct {
	TaskIDCommitment [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterTaskValidated is a free log retrieval operation binding the contract event 0xb86e6a82586576bc010ad08619492657c958d11f13926ea7cd171c66d20bad4d.
//
// Solidity: event TaskValidated(bytes32 taskIDCommitment)
func (_Task *TaskFilterer) FilterTaskValidated(opts *bind.FilterOpts) (*TaskTaskValidatedIterator, error) {

	logs, sub, err := _Task.contract.FilterLogs(opts, "TaskValidated")
	if err != nil {
		return nil, err
	}
	return &TaskTaskValidatedIterator{contract: _Task.contract, event: "TaskValidated", logs: logs, sub: sub}, nil
}

// WatchTaskValidated is a free log subscription operation binding the contract event 0xb86e6a82586576bc010ad08619492657c958d11f13926ea7cd171c66d20bad4d.
//
// Solidity: event TaskValidated(bytes32 taskIDCommitment)
func (_Task *TaskFilterer) WatchTaskValidated(opts *bind.WatchOpts, sink chan<- *TaskTaskValidated) (event.Subscription, error) {

	logs, sub, err := _Task.contract.WatchLogs(opts, "TaskValidated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskTaskValidated)
				if err := _Task.contract.UnpackLog(event, "TaskValidated", log); err != nil {
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

// ParseTaskValidated is a log parse operation binding the contract event 0xb86e6a82586576bc010ad08619492657c958d11f13926ea7cd171c66d20bad4d.
//
// Solidity: event TaskValidated(bytes32 taskIDCommitment)
func (_Task *TaskFilterer) ParseTaskValidated(log types.Log) (*TaskTaskValidated, error) {
	event := new(TaskTaskValidated)
	if err := _Task.contract.UnpackLog(event, "TaskValidated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
