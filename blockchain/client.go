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
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
)

var ethRpcClient *ethclient.Client

var chainID *big.Int
var gasPrice *big.Int

var taskContractInstance *bindings.Task
var nodeContractInstance *bindings.Node
var netstatsContractInstance *bindings.NetworkStats

func GetRpcClient() *ethclient.Client {
	if ethRpcClient == nil {
		log.Panicln("eth rpc client is nil")
	}
	return ethRpcClient
}

func GetTaskContractInstance() *bindings.Task {
	if taskContractInstance == nil {
		log.Panicln("task contract instance is nil")
	}
	return taskContractInstance
}

func GetNodeContractInstance() *bindings.Node {
	if nodeContractInstance == nil {
		log.Panicln("node contract instance is nil")
	}
	return nodeContractInstance
}

func GetNetstatsContractInstance() *bindings.NetworkStats {
	if netstatsContractInstance == nil {
		log.Panicln("netstats contract instance is nil")
	}
	return netstatsContractInstance
}

func getGasPrice() *big.Int {
	if gasPrice == nil {
		log.Panicln("gas price is nil")
	}
	return gasPrice
}

func getChainID() *big.Int {
	if chainID == nil {
		log.Panicln("chain id is nil")
	}
	return chainID
}

func Init(ctx context.Context) error {
	appConfig := config.GetConfig()
	if err := initEthRpcClient(appConfig.Blockchain.RpcEndpoint); err != nil {
		return err
	}
	if err := initTaskContractInstance(appConfig.Blockchain.Contracts.Task); err != nil {
		return err
	}
	if err := initNodeContractAddress(appConfig.Blockchain.Contracts.Node); err != nil {
		return err
	}
	if err := initNetstatsContractAddress(appConfig.Blockchain.Contracts.Netstats); err != nil {
		return err
	}
	if err := initChainID(ctx, appConfig.Blockchain.ChainID); err != nil {
		return err
	}
	if err := initSuggestGasPrice(ctx, appConfig.Blockchain.GasPrice); err != nil {
		return err
	}
	return nil
}

func initEthRpcClient(endpoint string) error {
	client, err := ethclient.Dial(endpoint)
	if err != nil {
		return err
	}
	ethRpcClient = client
	return nil
}

func initTaskContractInstance(taskContractAddress string) error {
	client := GetRpcClient()
	taskInstance, err := bindings.NewTask(common.HexToAddress(taskContractAddress), client)
	if err != nil {
		return err
	}
	taskContractInstance = taskInstance
	return nil
}

func initNodeContractAddress(nodeContractAddress string) error {
	client := GetRpcClient()
	nodeInstance, err := bindings.NewNode(common.HexToAddress(nodeContractAddress), client)
	if err != nil {
		return err
	}
	nodeContractInstance = nodeInstance
	return nil
}

func initNetstatsContractAddress(netstatsContractAddress string) error {
	client := GetRpcClient()
	netstatsInstance, err := bindings.NewNetworkStats(common.HexToAddress(netstatsContractAddress), client)
	if err != nil {
		return err
	}
	netstatsContractInstance = netstatsInstance
	return nil
}

func initSuggestGasPrice(ctx context.Context, gasPriceNum uint64) error {
	if gasPriceNum > 0 {
		gasPrice = big.NewInt(0).SetUint64(gasPriceNum)
	} else {
		callCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
		defer cancel()
		client := GetRpcClient()
		p, err := client.SuggestGasPrice(callCtx)
		if err != nil {
			return err
		}
		log.Debugln("Estimated gas price from blockchain: " + p.String())
		gasPrice = p
	}
	return nil
}

func initChainID(ctx context.Context, chainIDNum uint64) error {
	if chainIDNum > 0 {
		chainID = big.NewInt(0).SetUint64(chainIDNum)
	} else {
		callCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
		defer cancel()
		client := GetRpcClient()
		id, err := client.ChainID(callCtx)
		if err != nil {
			return err
		}
		chainID = id
	}
	return nil
}

func GetAuth(ctx context.Context, address common.Address, privateKeyStr string) (*bind.TransactOpts, error) {
	appConfig := config.GetConfig()

	var err error
	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		return nil, err
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, getChainID())
	if err != nil {
		return nil, err
	}

	log.Debugln("Set gas limit to:" + strconv.FormatUint(appConfig.Blockchain.GasLimit, 10))

	auth.Value = big.NewInt(0)
	auth.GasLimit = appConfig.Blockchain.GasLimit
	auth.GasPrice = getGasPrice()

	return auth, nil
}

func WaitTxReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	deadline, hasDeadline := ctx.Deadline()
	client := GetRpcClient()

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
	client := GetRpcClient()
	gasLimit := config.GetConfig().Blockchain.GasLimit

	txMutex.Lock()
	defer txMutex.Unlock()
	nonce, err := getNonce(ctx, from)
	if err != nil {
		return nil, err
	}

	tx := types.NewTransaction(nonce, to, amount, gasLimit, getGasPrice(), nil)

	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		return nil, err
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(getChainID()), privateKey)
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

	client := GetRpcClient()
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

	errMsg, err := unpackError(res)
	if err != nil {
		errMsg = "Unknown tx error" + hexutil.Encode(res)
	}
	return errMsg, err
}

var (
	errorSig     = []byte{0x08, 0xc3, 0x79, 0xa0} // Keccak256("Error(string)")[:4]
	abiString, _ = abi.NewType("string", "", nil)
)

func unpackError(result []byte) (string, error) {
	if len(result) < 4 {
		return "", errors.New("tx result length too short")
	}
	if !bytes.Equal(result[:4], errorSig) {
		return "", errors.New("tx result not of type Error(string)")
	}

	vs, err := abi.Arguments{{Type: abiString}}.UnpackValues(result[4:])
	if err != nil {
		return "", err
	}

	return vs[0].(string), nil
}
