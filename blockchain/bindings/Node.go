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

// NodeGPUInfo is an auto generated low-level Go binding around an user-defined struct.
type NodeGPUInfo struct {
	Name string
	Vram *big.Int
}

// NodeNodeInfo is an auto generated low-level Go binding around an user-defined struct.
type NodeNodeInfo struct {
	Status *big.Int
	GpuID  [32]byte
	Gpu    NodeGPUInfo
	Score  *big.Int
}

// NodeMetaData contains all meta data concerning the Node contract.
var NodeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"tokenInstance\",\"type\":\"address\"},{\"internalType\":\"contractQOS\",\"name\":\"qosInstance\",\"type\":\"address\"},{\"internalType\":\"contractNetworkStats\",\"name\":\"netStatsInstance\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nodeAddress\",\"type\":\"address\"}],\"name\":\"NodeKickedOut\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nodeAddress\",\"type\":\"address\"}],\"name\":\"NodeSlashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nodeAddress\",\"type\":\"address\"}],\"name\":\"getNodeInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gpuID\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"vram\",\"type\":\"uint256\"}],\"internalType\":\"structNode.GPUInfo\",\"name\":\"gpu\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"score\",\"type\":\"uint256\"}],\"internalType\":\"structNode.NodeInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAvailableGPUs\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"vram\",\"type\":\"uint256\"}],\"internalType\":\"structNode.GPUInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAvailableNodes\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nodeAddress\",\"type\":\"address\"}],\"name\":\"getNodeStatus\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"gpuName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"gpuVram\",\"type\":\"uint256\"}],\"name\":\"join\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resume\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nodeAddress\",\"type\":\"address\"}],\"name\":\"slash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nodeAddress\",\"type\":\"address\"}],\"name\":\"startTask\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nodeAddress\",\"type\":\"address\"}],\"name\":\"finishTask\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"taskContract\",\"type\":\"address\"}],\"name\":\"updateTaskContractAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"vramLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"countLimit\",\"type\":\"uint256\"}],\"name\":\"filterGPUID\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"gpuID\",\"type\":\"bytes32\"}],\"name\":\"filterNodesByGPUID\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"root\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"k\",\"type\":\"uint256\"}],\"name\":\"selectNodesWithRoot\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// NodeABI is the input ABI used to generate the binding from.
// Deprecated: Use NodeMetaData.ABI instead.
var NodeABI = NodeMetaData.ABI

// Node is an auto generated Go binding around an Ethereum contract.
type Node struct {
	NodeCaller     // Read-only binding to the contract
	NodeTransactor // Write-only binding to the contract
	NodeFilterer   // Log filterer for contract events
}

// NodeCaller is an auto generated read-only Go binding around an Ethereum contract.
type NodeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NodeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NodeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NodeSession struct {
	Contract     *Node             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NodeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NodeCallerSession struct {
	Contract *NodeCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// NodeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NodeTransactorSession struct {
	Contract     *NodeTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NodeRaw is an auto generated low-level Go binding around an Ethereum contract.
type NodeRaw struct {
	Contract *Node // Generic contract binding to access the raw methods on
}

// NodeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NodeCallerRaw struct {
	Contract *NodeCaller // Generic read-only contract binding to access the raw methods on
}

// NodeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NodeTransactorRaw struct {
	Contract *NodeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNode creates a new instance of Node, bound to a specific deployed contract.
func NewNode(address common.Address, backend bind.ContractBackend) (*Node, error) {
	contract, err := bindNode(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Node{NodeCaller: NodeCaller{contract: contract}, NodeTransactor: NodeTransactor{contract: contract}, NodeFilterer: NodeFilterer{contract: contract}}, nil
}

// NewNodeCaller creates a new read-only instance of Node, bound to a specific deployed contract.
func NewNodeCaller(address common.Address, caller bind.ContractCaller) (*NodeCaller, error) {
	contract, err := bindNode(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NodeCaller{contract: contract}, nil
}

// NewNodeTransactor creates a new write-only instance of Node, bound to a specific deployed contract.
func NewNodeTransactor(address common.Address, transactor bind.ContractTransactor) (*NodeTransactor, error) {
	contract, err := bindNode(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NodeTransactor{contract: contract}, nil
}

// NewNodeFilterer creates a new log filterer instance of Node, bound to a specific deployed contract.
func NewNodeFilterer(address common.Address, filterer bind.ContractFilterer) (*NodeFilterer, error) {
	contract, err := bindNode(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NodeFilterer{contract: contract}, nil
}

// bindNode binds a generic wrapper to an already deployed contract.
func bindNode(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NodeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Node *NodeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Node.Contract.NodeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Node *NodeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Node.Contract.NodeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Node *NodeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Node.Contract.NodeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Node *NodeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Node.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Node *NodeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Node.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Node *NodeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Node.Contract.contract.Transact(opts, method, params...)
}

// FilterGPUID is a free data retrieval call binding the contract method 0x9f09870a.
//
// Solidity: function filterGPUID(uint256 vramLimit, uint256 countLimit) view returns(bytes32[], uint256[])
func (_Node *NodeCaller) FilterGPUID(opts *bind.CallOpts, vramLimit *big.Int, countLimit *big.Int) ([][32]byte, []*big.Int, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "filterGPUID", vramLimit, countLimit)

	if err != nil {
		return *new([][32]byte), *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)
	out1 := *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return out0, out1, err

}

// FilterGPUID is a free data retrieval call binding the contract method 0x9f09870a.
//
// Solidity: function filterGPUID(uint256 vramLimit, uint256 countLimit) view returns(bytes32[], uint256[])
func (_Node *NodeSession) FilterGPUID(vramLimit *big.Int, countLimit *big.Int) ([][32]byte, []*big.Int, error) {
	return _Node.Contract.FilterGPUID(&_Node.CallOpts, vramLimit, countLimit)
}

// FilterGPUID is a free data retrieval call binding the contract method 0x9f09870a.
//
// Solidity: function filterGPUID(uint256 vramLimit, uint256 countLimit) view returns(bytes32[], uint256[])
func (_Node *NodeCallerSession) FilterGPUID(vramLimit *big.Int, countLimit *big.Int) ([][32]byte, []*big.Int, error) {
	return _Node.Contract.FilterGPUID(&_Node.CallOpts, vramLimit, countLimit)
}

// FilterNodesByGPUID is a free data retrieval call binding the contract method 0x975ac9d9.
//
// Solidity: function filterNodesByGPUID(bytes32 gpuID) view returns(address[], uint256[])
func (_Node *NodeCaller) FilterNodesByGPUID(opts *bind.CallOpts, gpuID [32]byte) ([]common.Address, []*big.Int, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "filterNodesByGPUID", gpuID)

	if err != nil {
		return *new([]common.Address), *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return out0, out1, err

}

// FilterNodesByGPUID is a free data retrieval call binding the contract method 0x975ac9d9.
//
// Solidity: function filterNodesByGPUID(bytes32 gpuID) view returns(address[], uint256[])
func (_Node *NodeSession) FilterNodesByGPUID(gpuID [32]byte) ([]common.Address, []*big.Int, error) {
	return _Node.Contract.FilterNodesByGPUID(&_Node.CallOpts, gpuID)
}

// FilterNodesByGPUID is a free data retrieval call binding the contract method 0x975ac9d9.
//
// Solidity: function filterNodesByGPUID(bytes32 gpuID) view returns(address[], uint256[])
func (_Node *NodeCallerSession) FilterNodesByGPUID(gpuID [32]byte) ([]common.Address, []*big.Int, error) {
	return _Node.Contract.FilterNodesByGPUID(&_Node.CallOpts, gpuID)
}

// GetAvailableGPUs is a free data retrieval call binding the contract method 0x169eacef.
//
// Solidity: function getAvailableGPUs() view returns((string,uint256)[])
func (_Node *NodeCaller) GetAvailableGPUs(opts *bind.CallOpts) ([]NodeGPUInfo, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "getAvailableGPUs")

	if err != nil {
		return *new([]NodeGPUInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]NodeGPUInfo)).(*[]NodeGPUInfo)

	return out0, err

}

// GetAvailableGPUs is a free data retrieval call binding the contract method 0x169eacef.
//
// Solidity: function getAvailableGPUs() view returns((string,uint256)[])
func (_Node *NodeSession) GetAvailableGPUs() ([]NodeGPUInfo, error) {
	return _Node.Contract.GetAvailableGPUs(&_Node.CallOpts)
}

// GetAvailableGPUs is a free data retrieval call binding the contract method 0x169eacef.
//
// Solidity: function getAvailableGPUs() view returns((string,uint256)[])
func (_Node *NodeCallerSession) GetAvailableGPUs() ([]NodeGPUInfo, error) {
	return _Node.Contract.GetAvailableGPUs(&_Node.CallOpts)
}

// GetAvailableNodes is a free data retrieval call binding the contract method 0x436a014a.
//
// Solidity: function getAvailableNodes() view returns(address[])
func (_Node *NodeCaller) GetAvailableNodes(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "getAvailableNodes")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAvailableNodes is a free data retrieval call binding the contract method 0x436a014a.
//
// Solidity: function getAvailableNodes() view returns(address[])
func (_Node *NodeSession) GetAvailableNodes() ([]common.Address, error) {
	return _Node.Contract.GetAvailableNodes(&_Node.CallOpts)
}

// GetAvailableNodes is a free data retrieval call binding the contract method 0x436a014a.
//
// Solidity: function getAvailableNodes() view returns(address[])
func (_Node *NodeCallerSession) GetAvailableNodes() ([]common.Address, error) {
	return _Node.Contract.GetAvailableNodes(&_Node.CallOpts)
}

// GetNodeInfo is a free data retrieval call binding the contract method 0x582115fb.
//
// Solidity: function getNodeInfo(address nodeAddress) view returns((uint256,bytes32,(string,uint256),uint256))
func (_Node *NodeCaller) GetNodeInfo(opts *bind.CallOpts, nodeAddress common.Address) (NodeNodeInfo, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "getNodeInfo", nodeAddress)

	if err != nil {
		return *new(NodeNodeInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(NodeNodeInfo)).(*NodeNodeInfo)

	return out0, err

}

// GetNodeInfo is a free data retrieval call binding the contract method 0x582115fb.
//
// Solidity: function getNodeInfo(address nodeAddress) view returns((uint256,bytes32,(string,uint256),uint256))
func (_Node *NodeSession) GetNodeInfo(nodeAddress common.Address) (NodeNodeInfo, error) {
	return _Node.Contract.GetNodeInfo(&_Node.CallOpts, nodeAddress)
}

// GetNodeInfo is a free data retrieval call binding the contract method 0x582115fb.
//
// Solidity: function getNodeInfo(address nodeAddress) view returns((uint256,bytes32,(string,uint256),uint256))
func (_Node *NodeCallerSession) GetNodeInfo(nodeAddress common.Address) (NodeNodeInfo, error) {
	return _Node.Contract.GetNodeInfo(&_Node.CallOpts, nodeAddress)
}

// GetNodeStatus is a free data retrieval call binding the contract method 0xb65f8177.
//
// Solidity: function getNodeStatus(address nodeAddress) view returns(uint256)
func (_Node *NodeCaller) GetNodeStatus(opts *bind.CallOpts, nodeAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "getNodeStatus", nodeAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNodeStatus is a free data retrieval call binding the contract method 0xb65f8177.
//
// Solidity: function getNodeStatus(address nodeAddress) view returns(uint256)
func (_Node *NodeSession) GetNodeStatus(nodeAddress common.Address) (*big.Int, error) {
	return _Node.Contract.GetNodeStatus(&_Node.CallOpts, nodeAddress)
}

// GetNodeStatus is a free data retrieval call binding the contract method 0xb65f8177.
//
// Solidity: function getNodeStatus(address nodeAddress) view returns(uint256)
func (_Node *NodeCallerSession) GetNodeStatus(nodeAddress common.Address) (*big.Int, error) {
	return _Node.Contract.GetNodeStatus(&_Node.CallOpts, nodeAddress)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Node *NodeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Node *NodeSession) Owner() (common.Address, error) {
	return _Node.Contract.Owner(&_Node.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Node *NodeCallerSession) Owner() (common.Address, error) {
	return _Node.Contract.Owner(&_Node.CallOpts)
}

// SelectNodesWithRoot is a free data retrieval call binding the contract method 0x30f8ada8.
//
// Solidity: function selectNodesWithRoot(address root, uint256 k) view returns(address[])
func (_Node *NodeCaller) SelectNodesWithRoot(opts *bind.CallOpts, root common.Address, k *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "selectNodesWithRoot", root, k)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// SelectNodesWithRoot is a free data retrieval call binding the contract method 0x30f8ada8.
//
// Solidity: function selectNodesWithRoot(address root, uint256 k) view returns(address[])
func (_Node *NodeSession) SelectNodesWithRoot(root common.Address, k *big.Int) ([]common.Address, error) {
	return _Node.Contract.SelectNodesWithRoot(&_Node.CallOpts, root, k)
}

// SelectNodesWithRoot is a free data retrieval call binding the contract method 0x30f8ada8.
//
// Solidity: function selectNodesWithRoot(address root, uint256 k) view returns(address[])
func (_Node *NodeCallerSession) SelectNodesWithRoot(root common.Address, k *big.Int) ([]common.Address, error) {
	return _Node.Contract.SelectNodesWithRoot(&_Node.CallOpts, root, k)
}

// FinishTask is a paid mutator transaction binding the contract method 0x3fc0f48b.
//
// Solidity: function finishTask(address nodeAddress) returns()
func (_Node *NodeTransactor) FinishTask(opts *bind.TransactOpts, nodeAddress common.Address) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "finishTask", nodeAddress)
}

// FinishTask is a paid mutator transaction binding the contract method 0x3fc0f48b.
//
// Solidity: function finishTask(address nodeAddress) returns()
func (_Node *NodeSession) FinishTask(nodeAddress common.Address) (*types.Transaction, error) {
	return _Node.Contract.FinishTask(&_Node.TransactOpts, nodeAddress)
}

// FinishTask is a paid mutator transaction binding the contract method 0x3fc0f48b.
//
// Solidity: function finishTask(address nodeAddress) returns()
func (_Node *NodeTransactorSession) FinishTask(nodeAddress common.Address) (*types.Transaction, error) {
	return _Node.Contract.FinishTask(&_Node.TransactOpts, nodeAddress)
}

// Join is a paid mutator transaction binding the contract method 0x99b32360.
//
// Solidity: function join(string gpuName, uint256 gpuVram) returns()
func (_Node *NodeTransactor) Join(opts *bind.TransactOpts, gpuName string, gpuVram *big.Int) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "join", gpuName, gpuVram)
}

// Join is a paid mutator transaction binding the contract method 0x99b32360.
//
// Solidity: function join(string gpuName, uint256 gpuVram) returns()
func (_Node *NodeSession) Join(gpuName string, gpuVram *big.Int) (*types.Transaction, error) {
	return _Node.Contract.Join(&_Node.TransactOpts, gpuName, gpuVram)
}

// Join is a paid mutator transaction binding the contract method 0x99b32360.
//
// Solidity: function join(string gpuName, uint256 gpuVram) returns()
func (_Node *NodeTransactorSession) Join(gpuName string, gpuVram *big.Int) (*types.Transaction, error) {
	return _Node.Contract.Join(&_Node.TransactOpts, gpuName, gpuVram)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Node *NodeTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Node *NodeSession) Pause() (*types.Transaction, error) {
	return _Node.Contract.Pause(&_Node.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Node *NodeTransactorSession) Pause() (*types.Transaction, error) {
	return _Node.Contract.Pause(&_Node.TransactOpts)
}

// Quit is a paid mutator transaction binding the contract method 0xfc2b8cc3.
//
// Solidity: function quit() returns()
func (_Node *NodeTransactor) Quit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "quit")
}

// Quit is a paid mutator transaction binding the contract method 0xfc2b8cc3.
//
// Solidity: function quit() returns()
func (_Node *NodeSession) Quit() (*types.Transaction, error) {
	return _Node.Contract.Quit(&_Node.TransactOpts)
}

// Quit is a paid mutator transaction binding the contract method 0xfc2b8cc3.
//
// Solidity: function quit() returns()
func (_Node *NodeTransactorSession) Quit() (*types.Transaction, error) {
	return _Node.Contract.Quit(&_Node.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Node *NodeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Node *NodeSession) RenounceOwnership() (*types.Transaction, error) {
	return _Node.Contract.RenounceOwnership(&_Node.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Node *NodeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Node.Contract.RenounceOwnership(&_Node.TransactOpts)
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_Node *NodeTransactor) Resume(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "resume")
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_Node *NodeSession) Resume() (*types.Transaction, error) {
	return _Node.Contract.Resume(&_Node.TransactOpts)
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_Node *NodeTransactorSession) Resume() (*types.Transaction, error) {
	return _Node.Contract.Resume(&_Node.TransactOpts)
}

// Slash is a paid mutator transaction binding the contract method 0xc96be4cb.
//
// Solidity: function slash(address nodeAddress) returns()
func (_Node *NodeTransactor) Slash(opts *bind.TransactOpts, nodeAddress common.Address) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "slash", nodeAddress)
}

// Slash is a paid mutator transaction binding the contract method 0xc96be4cb.
//
// Solidity: function slash(address nodeAddress) returns()
func (_Node *NodeSession) Slash(nodeAddress common.Address) (*types.Transaction, error) {
	return _Node.Contract.Slash(&_Node.TransactOpts, nodeAddress)
}

// Slash is a paid mutator transaction binding the contract method 0xc96be4cb.
//
// Solidity: function slash(address nodeAddress) returns()
func (_Node *NodeTransactorSession) Slash(nodeAddress common.Address) (*types.Transaction, error) {
	return _Node.Contract.Slash(&_Node.TransactOpts, nodeAddress)
}

// StartTask is a paid mutator transaction binding the contract method 0x5f51c765.
//
// Solidity: function startTask(address nodeAddress) returns()
func (_Node *NodeTransactor) StartTask(opts *bind.TransactOpts, nodeAddress common.Address) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "startTask", nodeAddress)
}

// StartTask is a paid mutator transaction binding the contract method 0x5f51c765.
//
// Solidity: function startTask(address nodeAddress) returns()
func (_Node *NodeSession) StartTask(nodeAddress common.Address) (*types.Transaction, error) {
	return _Node.Contract.StartTask(&_Node.TransactOpts, nodeAddress)
}

// StartTask is a paid mutator transaction binding the contract method 0x5f51c765.
//
// Solidity: function startTask(address nodeAddress) returns()
func (_Node *NodeTransactorSession) StartTask(nodeAddress common.Address) (*types.Transaction, error) {
	return _Node.Contract.StartTask(&_Node.TransactOpts, nodeAddress)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Node *NodeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Node *NodeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Node.Contract.TransferOwnership(&_Node.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Node *NodeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Node.Contract.TransferOwnership(&_Node.TransactOpts, newOwner)
}

// UpdateTaskContractAddress is a paid mutator transaction binding the contract method 0x42145230.
//
// Solidity: function updateTaskContractAddress(address taskContract) returns()
func (_Node *NodeTransactor) UpdateTaskContractAddress(opts *bind.TransactOpts, taskContract common.Address) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "updateTaskContractAddress", taskContract)
}

// UpdateTaskContractAddress is a paid mutator transaction binding the contract method 0x42145230.
//
// Solidity: function updateTaskContractAddress(address taskContract) returns()
func (_Node *NodeSession) UpdateTaskContractAddress(taskContract common.Address) (*types.Transaction, error) {
	return _Node.Contract.UpdateTaskContractAddress(&_Node.TransactOpts, taskContract)
}

// UpdateTaskContractAddress is a paid mutator transaction binding the contract method 0x42145230.
//
// Solidity: function updateTaskContractAddress(address taskContract) returns()
func (_Node *NodeTransactorSession) UpdateTaskContractAddress(taskContract common.Address) (*types.Transaction, error) {
	return _Node.Contract.UpdateTaskContractAddress(&_Node.TransactOpts, taskContract)
}

// NodeNodeKickedOutIterator is returned from FilterNodeKickedOut and is used to iterate over the raw logs and unpacked data for NodeKickedOut events raised by the Node contract.
type NodeNodeKickedOutIterator struct {
	Event *NodeNodeKickedOut // Event containing the contract specifics and raw log

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
func (it *NodeNodeKickedOutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeNodeKickedOut)
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
		it.Event = new(NodeNodeKickedOut)
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
func (it *NodeNodeKickedOutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeNodeKickedOutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeNodeKickedOut represents a NodeKickedOut event raised by the Node contract.
type NodeNodeKickedOut struct {
	NodeAddress common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNodeKickedOut is a free log retrieval operation binding the contract event 0xb848d555de967481a57a6f519357b97fabdf06fbd2b63dece8f11876f3ab9323.
//
// Solidity: event NodeKickedOut(address nodeAddress)
func (_Node *NodeFilterer) FilterNodeKickedOut(opts *bind.FilterOpts) (*NodeNodeKickedOutIterator, error) {

	logs, sub, err := _Node.contract.FilterLogs(opts, "NodeKickedOut")
	if err != nil {
		return nil, err
	}
	return &NodeNodeKickedOutIterator{contract: _Node.contract, event: "NodeKickedOut", logs: logs, sub: sub}, nil
}

// WatchNodeKickedOut is a free log subscription operation binding the contract event 0xb848d555de967481a57a6f519357b97fabdf06fbd2b63dece8f11876f3ab9323.
//
// Solidity: event NodeKickedOut(address nodeAddress)
func (_Node *NodeFilterer) WatchNodeKickedOut(opts *bind.WatchOpts, sink chan<- *NodeNodeKickedOut) (event.Subscription, error) {

	logs, sub, err := _Node.contract.WatchLogs(opts, "NodeKickedOut")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeNodeKickedOut)
				if err := _Node.contract.UnpackLog(event, "NodeKickedOut", log); err != nil {
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

// ParseNodeKickedOut is a log parse operation binding the contract event 0xb848d555de967481a57a6f519357b97fabdf06fbd2b63dece8f11876f3ab9323.
//
// Solidity: event NodeKickedOut(address nodeAddress)
func (_Node *NodeFilterer) ParseNodeKickedOut(log types.Log) (*NodeNodeKickedOut, error) {
	event := new(NodeNodeKickedOut)
	if err := _Node.contract.UnpackLog(event, "NodeKickedOut", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeNodeSlashedIterator is returned from FilterNodeSlashed and is used to iterate over the raw logs and unpacked data for NodeSlashed events raised by the Node contract.
type NodeNodeSlashedIterator struct {
	Event *NodeNodeSlashed // Event containing the contract specifics and raw log

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
func (it *NodeNodeSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeNodeSlashed)
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
		it.Event = new(NodeNodeSlashed)
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
func (it *NodeNodeSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeNodeSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeNodeSlashed represents a NodeSlashed event raised by the Node contract.
type NodeNodeSlashed struct {
	NodeAddress common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNodeSlashed is a free log retrieval operation binding the contract event 0x29f3a9a9c7f6d4074ec8817742795e031d525ab8fe33b05ee339002580ef3a64.
//
// Solidity: event NodeSlashed(address nodeAddress)
func (_Node *NodeFilterer) FilterNodeSlashed(opts *bind.FilterOpts) (*NodeNodeSlashedIterator, error) {

	logs, sub, err := _Node.contract.FilterLogs(opts, "NodeSlashed")
	if err != nil {
		return nil, err
	}
	return &NodeNodeSlashedIterator{contract: _Node.contract, event: "NodeSlashed", logs: logs, sub: sub}, nil
}

// WatchNodeSlashed is a free log subscription operation binding the contract event 0x29f3a9a9c7f6d4074ec8817742795e031d525ab8fe33b05ee339002580ef3a64.
//
// Solidity: event NodeSlashed(address nodeAddress)
func (_Node *NodeFilterer) WatchNodeSlashed(opts *bind.WatchOpts, sink chan<- *NodeNodeSlashed) (event.Subscription, error) {

	logs, sub, err := _Node.contract.WatchLogs(opts, "NodeSlashed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeNodeSlashed)
				if err := _Node.contract.UnpackLog(event, "NodeSlashed", log); err != nil {
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

// ParseNodeSlashed is a log parse operation binding the contract event 0x29f3a9a9c7f6d4074ec8817742795e031d525ab8fe33b05ee339002580ef3a64.
//
// Solidity: event NodeSlashed(address nodeAddress)
func (_Node *NodeFilterer) ParseNodeSlashed(log types.Log) (*NodeNodeSlashed, error) {
	event := new(NodeNodeSlashed)
	if err := _Node.contract.UnpackLog(event, "NodeSlashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Node contract.
type NodeOwnershipTransferredIterator struct {
	Event *NodeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *NodeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeOwnershipTransferred)
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
		it.Event = new(NodeOwnershipTransferred)
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
func (it *NodeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeOwnershipTransferred represents a OwnershipTransferred event raised by the Node contract.
type NodeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Node *NodeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*NodeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Node.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NodeOwnershipTransferredIterator{contract: _Node.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Node *NodeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NodeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Node.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeOwnershipTransferred)
				if err := _Node.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Node *NodeFilterer) ParseOwnershipTransferred(log types.Log) (*NodeOwnershipTransferred, error) {
	event := new(NodeOwnershipTransferred)
	if err := _Node.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
