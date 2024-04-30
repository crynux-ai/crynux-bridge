package blockchain

import (
	"bytes"
	"context"
	"crynux_bridge/blockchain/bindings"
	"crynux_bridge/config"
	"errors"
	"math/big"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
)

var ethWSClient *ethclient.Client

func GetRpcClient() (*ethclient.Client, error) {

	if ethWSClient == nil {
		appConfig := config.GetConfig()
		client, err := ethclient.Dial(appConfig.Blockchain.RpcEndpoint)

		if err != nil {
			return nil, err
		}

		ethWSClient = client
	}

	return ethWSClient, nil
}

func GetAuth(client *ethclient.Client, address common.Address, privateKeyStr string) (*bind.TransactOpts, error) {

	appConfig := config.GetConfig()

	nonce, err := client.PendingNonceAt(context.Background(), address)
	if err != nil {
		return nil, err
	}

	log.Debugln("Nonce from blockchain: " + strconv.FormatUint(nonce, 10))

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	log.Debugln("Estimated gas price from blockchain: " + gasPrice.String())

	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		return nil, err
	}

	chainId, err := client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}

	log.Debugln("Chain ID from the blockchain: " + chainId.String())

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		return nil, err
	}

	log.Debugln("Set gas limit to:" + strconv.FormatUint(appConfig.Blockchain.GasLimit, 10))

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = appConfig.Blockchain.GasLimit
	auth.GasPrice = gasPrice

	return auth, nil
}

func SendETH(from common.Address, to common.Address, amount *big.Int, privateKeyStr string) (*types.Transaction, error) {

	client, err := GetRpcClient()
	if err != nil {
		return nil, err
	}

	nonce, err := client.PendingNonceAt(context.Background(), from)
	if err != nil {
		return nil, err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	gasLimit := config.GetConfig().Blockchain.GasLimit
	tx := types.NewTransaction(nonce, to, amount, gasLimit, gasPrice, nil)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return nil, err
	}

	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		return nil, err
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return nil, err
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return nil, err
	}

	return signedTx, nil
}

func GetErrorMessageForTxHash(txHash common.Hash, blockNumber *big.Int) (string, error) {

	client, err := GetRpcClient()
	if err != nil {
		return "", err
	}

	tx, _, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		return "", err
	}

	msg := ethereum.CallMsg{
		From:     common.HexToAddress(config.GetConfig().Blockchain.Account.Address),
		To:       tx.To(),
		Gas:      tx.Gas(),
		GasPrice: tx.GasPrice(),
		Value:    tx.Value(),
		Data:     tx.Data(),
	}

	ctx, cancelFn := context.WithTimeout(context.Background(), time.Duration(3)*time.Second)
	defer cancelFn()

	res, err := client.CallContract(ctx, msg, blockNumber)
	if err != nil {
		return err.Error(), nil
	}

	return unpackError(res)
}

var (
	errorSig     = []byte{0x08, 0xc3, 0x79, 0xa0} // Keccak256("Error(string)")[:4]
	abiString, _ = abi.NewType("string", "", nil)
)

func unpackError(result []byte) (string, error) {
	if !bytes.Equal(result[:4], errorSig) {
		return "", errors.New("TX result not of type Error(string)")
	}

	vs, err := abi.Arguments{{Type: abiString}}.UnpackValues(result[4:])
	if err != nil {
		return "", err
	}

	return vs[0].(string), nil
}

var taskContractInstance *bindings.Task
var nodeContractInstance *bindings.Node
var netstatsContractInstance *bindings.NetworkStats

func GetTaskContractInstance() (*bindings.Task, error) {

	if taskContractInstance == nil {
		appConfig := config.GetConfig()
		taskContractAddress := common.HexToAddress(appConfig.Blockchain.Contracts.Task)

		client, err := GetRpcClient()
		if err != nil {
			return nil, err
		}

		instance, err := bindings.NewTask(taskContractAddress, client)

		if err != nil {
			return nil, err
		}

		taskContractInstance = instance
	}

	return taskContractInstance, nil
}

func GetNodeContractInstance() (*bindings.Node, error) {
	if nodeContractInstance == nil {
		appConfig := config.GetConfig()
		nodeContractAddress := common.HexToAddress(appConfig.Blockchain.Contracts.Node)

		client, err := GetRpcClient()
		if err != nil {
			return nil, err
		}

		instance, err := bindings.NewNode(nodeContractAddress, client)

		if err != nil {
			return nil, err
		}

		nodeContractInstance = instance
	}

	return nodeContractInstance, nil
}

func GetNetstatsContractInstance() (*bindings.NetworkStats, error) {
	if netstatsContractInstance == nil {
		appConfig := config.GetConfig()
		address := common.HexToAddress(appConfig.Blockchain.Contracts.Netstats)

		client, err := GetRpcClient()
		if err != nil {
			return nil, err
		}

		instance, err := bindings.NewNetworkStats(address, client)

		if err != nil {
			return nil, err
		}

		netstatsContractInstance = instance
	}

	return netstatsContractInstance, nil
}
