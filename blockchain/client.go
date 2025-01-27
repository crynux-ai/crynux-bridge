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

var ethRpcClient *ethclient.Client

var chainID *big.Int
var gasPrice *big.Int

func GetRpcClient() (*ethclient.Client, error) {
	if ethRpcClient == nil {
		appConfig := config.GetConfig()
		client, err := ethclient.Dial(appConfig.Blockchain.RpcEndpoint)

		if err != nil {
			return nil, err
		}

		ethRpcClient = client
	}

	return ethRpcClient, nil
}

func getSuggestGasPrice(ctx context.Context) (*big.Int, error) {
	if gasPrice == nil {
		client, err := GetRpcClient()
		if err != nil {
			return nil, nil
		}

		if err := getLimiter().Wait(ctx); err != nil {
			return nil, nil
		}
		callCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
		defer cancel()
		p, err := client.SuggestGasPrice(callCtx)
		if err != nil {
			return nil, err
		}
		log.Debugln("Estimated gas price from blockchain: " + p.String())
		gasPrice = p
	}
	return gasPrice, nil
}

func getChainID(ctx context.Context) (*big.Int, error) {
	if chainID == nil {
		client, err := GetRpcClient()
		if err != nil {
			return nil, nil
		}

		if err := getLimiter().Wait(ctx); err != nil {
			return nil, nil
		}
		callCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
		defer cancel()
		id, err := client.ChainID(callCtx)
		if err != nil {
			return nil, err
		}
		chainID = id
	}
	return chainID, nil
}

func GetAuth(ctx context.Context, address common.Address, privateKeyStr string) (*bind.TransactOpts, error) {

	appConfig := config.GetConfig()

	var err error
	gasPrice := big.NewInt(0)
	if appConfig.Blockchain.GasPrice > 0 {
		gasPrice.SetUint64(appConfig.Blockchain.GasPrice)
	} else {
		gasPrice, err = getSuggestGasPrice(ctx)
		if err != nil {
			return nil, err
		}
	}

	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		return nil, err
	}

	chainID := big.NewInt(0)
	if appConfig.Blockchain.ChainID > 0 {
		chainID.SetUint64(appConfig.Blockchain.ChainID)
	} else {
		chainID, err = getChainID(ctx)
		if err != nil {
			return nil, err
		}
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, err
	}

	log.Debugln("Set gas limit to:" + strconv.FormatUint(appConfig.Blockchain.GasLimit, 10))

	auth.Value = big.NewInt(0)
	auth.GasLimit = appConfig.Blockchain.GasLimit
	auth.GasPrice = gasPrice

	return auth, nil
}

func WaitTxReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	deadline, hasDeadline := ctx.Deadline()
	client, err := GetRpcClient()
	if err != nil {
		return nil, err
	}

	for {
		r, err := func() (*types.Receipt, error) {
			callCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
			defer cancel()
			return client.TransactionReceipt(callCtx, txHash)
		}()
		if err == ethereum.NotFound {
			time.Sleep(time.Second)
			continue
		}
		if hasDeadline && time.Now().Compare(deadline) >= 0 && err == context.DeadlineExceeded {
			log.Errorf("wait receipt of tx %s timeout", txHash.Hex())
			return nil, err
		}
		if err != nil {
			return nil, err
		}
		return r, nil
	}
}

func SendETH(ctx context.Context, from common.Address, to common.Address, amount *big.Int, privateKeyStr string) (*types.Transaction, error) {
	appConfig := config.GetConfig()

	client, err := GetRpcClient()
	if err != nil {
		return nil, err
	}

	gasPrice := big.NewInt(0)
	if appConfig.Blockchain.GasPrice > 0 {
		gasPrice.SetUint64(appConfig.Blockchain.GasPrice)
	} else {
		gasPrice, err = getSuggestGasPrice(ctx)
		if err != nil {
			return nil, err
		}
	}

	chainID := big.NewInt(0)
	if appConfig.Blockchain.ChainID > 0 {
		chainID.SetUint64(appConfig.Blockchain.ChainID)
	} else {
		chainID, err = getChainID(ctx)
		if err != nil {
			return nil, err
		}
	}

	gasLimit := config.GetConfig().Blockchain.GasLimit

	txMutex.Lock()
	defer txMutex.Unlock()
	nonce, err := getNonce(ctx, from)
	if err != nil {
		return nil, err
	}

	tx := types.NewTransaction(nonce, to, amount, gasLimit, gasPrice, nil)

	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		return nil, err
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return nil, err
	}

	if err := getLimiter().Wait(ctx); err != nil {
		return nil, err
	}

	callCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	err = client.SendTransaction(callCtx, signedTx)
	if err != nil {
		err = processSendingTxError(err)
		return nil, err
	}

	addNonce(nonce)
	return signedTx, nil
}

func GetErrorMessageFromReceipt(ctx context.Context, receipt *types.Receipt) (string, error) {

	client, err := GetRpcClient()
	if err != nil {
		return "", err
	}

	if err := getLimiter().Wait(ctx); err != nil {
		return "", err
	}

	ctx1, cancel1 := context.WithTimeout(ctx, 30*time.Second)
	defer cancel1()
	tx, _, err := client.TransactionByHash(ctx1, receipt.TxHash)
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

	if err := getLimiter().Wait(ctx); err != nil {
		return "", err
	}

	blockNumber := big.NewInt(0).Sub(receipt.BlockNumber, big.NewInt(1))

	ctx2, cancel2 := context.WithTimeout(ctx, 30*time.Second)
	defer cancel2()

	res, err := client.CallContract(ctx2, msg, blockNumber)
	if err != nil {
		return "", err
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
