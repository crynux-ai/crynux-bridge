package blockchain

import (
	"context"
	"crynux_bridge/blockchain/bindings"
	"crynux_bridge/config"
	"crynux_bridge/models"
	"crypto/sha256"
	"encoding/binary"
	"errors"
	"image"
	"image/png"
	"io"
	"math/big"
	"math/rand"
	"strconv"
	"time"

	"github.com/corona10/goimagehash"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	log "github.com/sirupsen/logrus"
)

var taskContractInstance *bindings.Task
var crynuxTokenContractInstance *bindings.CrynuxToken
var nodeContractInstance *bindings.Node

func CreateTaskOnChain(task *models.InferenceTask) (string, error) {

	appConfig := config.GetConfig()

	taskHash, err := task.GetTaskHash()
	if err != nil {
		return "", err
	}

	dataHash := &[32]byte{
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
	}

	taskContractAddress := common.HexToAddress(appConfig.Blockchain.Contracts.Task)
	accountAddress := common.HexToAddress(appConfig.Blockchain.Account.Address)
	accountPrivateKey := appConfig.Blockchain.Account.PrivateKey

	client, err := GetRpcClient()
	if err != nil {
		return "", err
	}

	instance, err := bindings.NewTask(taskContractAddress, client)
	if err != nil {
		return "", err
	}

	auth, err := GetAuth(client, accountAddress, accountPrivateKey)
	if err != nil {
		return "", err
	}

	log.Debugln("create task tx: TaskHash " + common.Bytes2Hex(taskHash[:]))
	log.Debugln("create task tx: DataHash " + common.Bytes2Hex(dataHash[:]))

	tx, err := instance.CreateTask(auth, big.NewInt(int64(task.TaskType)), *taskHash, *dataHash, big.NewInt(int64(task.VramLimit)))
	if err != nil {
		return "", err
	}

	return tx.Hash().Hex(), nil
}

func GetTaskCreationResult(txHash string) (*big.Int, error) {

	client, err := GetRpcClient()
	if err != nil {
		return nil, err
	}

	ctx, cancelFn := context.WithTimeout(context.Background(), time.Duration(3)*time.Second)
	defer cancelFn()

	receipt, err := client.TransactionReceipt(ctx, common.HexToHash(txHash))
	if err != nil {

		if errors.Is(err, ethereum.NotFound) {
			// Transaction pending
			return nil, nil
		}

		log.Errorln("error getting tx receipt for: " + txHash)
		return nil, err
	}

	if receipt.Status == 0 {
		// Transaction failed
		// Get reason
		reason, err := GetErrorMessageForTxHash(receipt.TxHash, receipt.BlockNumber)

		if err != nil {
			log.Errorln("error getting error message for: " + txHash)
			return nil, err
		}

		return nil, errors.New(reason)
	}

	// Transaction success
	// Extract taskId from the logs
	taskContractInstance, err := GetTaskContractInstance()
	if err != nil {
		log.Errorln("error get task contract instance: " + receipt.TxHash.Hex())
		return nil, err
	}

	// There are 5 events emitted from the CreateTask method
	// Approval, Transfer, TaskCreated x 3
	if len(receipt.Logs) != 5 {
		log.Errorln(receipt.Logs)
		return nil, errors.New("wrong event logs number:" + strconv.Itoa(len(receipt.Logs)))
	}

	taskCreatedEvent, err := taskContractInstance.ParseTaskCreated(*receipt.Logs[2])
	if err != nil {
		log.Errorln("error parse task created event: " + receipt.TxHash.Hex())
		return nil, err
	}

	taskId := taskCreatedEvent.TaskId

	return taskId, nil
}

func GetTaskResultCommitment(result []byte) (commitment [32]byte, nonce [32]byte) {
	nonceStr := strconv.Itoa(rand.Int())
	nonceHash := crypto.Keccak256Hash([]byte(nonceStr))
	commitmentHash := crypto.Keccak256Hash(result, nonceHash.Bytes())
	copy(commitment[:], commitmentHash.Bytes())
	copy(nonce[:], nonceHash.Bytes())
	return commitment, nonce
}

func GetPHashForImage(image image.Image) ([]byte, error) {
	pHash, err := goimagehash.PerceptionHash(image)
	if err != nil {
		return nil, err
	}

	bs := make([]byte, pHash.Bits()/8)
	binary.BigEndian.PutUint64(bs, pHash.GetHash())
	return bs, nil
}

func GetPHashForImageReader(reader io.Reader) ([]byte, error) {
	img, err := png.Decode(reader)
	if err != nil {
		return nil, err
	}
	return GetPHashForImage(img)
}


func GetHashForGPTResponse(resp string) []byte {
	h := sha256.Sum256([]byte(resp))
	return h[:]
}

func ApproveAllBalanceForTaskCreator() error {

	// Check ETH balance
	client, err := GetRpcClient()
	if err != nil {
		return err
	}

	appAddress := common.HexToAddress(config.GetConfig().Blockchain.Account.Address)

	log.Infoln("Approve all balance for the application account: " + config.GetConfig().Blockchain.Account.Address)

	currentETHBalance, err := client.BalanceAt(
		context.Background(),
		appAddress,
		nil,
	)
	if err != nil {
		return err
	}

	currentETHBalanceInEther := new(big.Int).Div(currentETHBalance, big.NewInt(params.Ether))
	log.Infoln("ETH balance for the application account: " + currentETHBalanceInEther.String())

	ethThreshold := new(big.Int).Mul(big.NewInt(10000000), big.NewInt(params.GWei))

	if currentETHBalance.Cmp(ethThreshold) != 1 {
		return errors.New("not enough ETH left")
	}

	// Check Crynux token balance
	cnxTokenInstance, err := GetCrynuxTokenContractInstance()
	if err != nil {
		return err
	}

	cnxBalance, err := cnxTokenInstance.BalanceOf(
		&bind.CallOpts{
			Pending: false,
			Context: context.Background(),
		},
		appAddress,
	)
	if err != nil {
		return err
	}

	cnxBalanceInEther := new(big.Int).Div(cnxBalance, big.NewInt(params.Ether))
	log.Infoln("Crynux token balance for the application account: " + cnxBalanceInEther.String())

	cnxThreshold := new(big.Int).Mul(big.NewInt(100), big.NewInt(params.Ether))

	if cnxBalance.Cmp(cnxThreshold) != 1 {
		return errors.New("not enough Crynux token left")
	}

	// Check Crynux token allowance
	taskContractAddress := common.HexToAddress(config.GetConfig().Blockchain.Contracts.Task)

	cnxAllowance, err := cnxTokenInstance.Allowance(
		&bind.CallOpts{
			Pending: false,
			Context: context.Background(),
		},
		appAddress,
		taskContractAddress,
	)

	if err != nil {
		return err
	}

	cnxAllowanceInEther := new(big.Int).Div(cnxAllowance, big.NewInt(params.Ether))
	log.Infoln("Crynux token allowance to the task contract for the application account: " + cnxAllowanceInEther.String())

	if cnxBalance.Cmp(cnxAllowance) == 1 {

		log.Infoln("Approve more tokens for the task contract")

		// Let's approve the remaining balance
		auth, err := GetAuth(
			client,
			appAddress,
			config.GetConfig().Blockchain.Account.PrivateKey)

		if err != nil {
			return err
		}

		tx, err := cnxTokenInstance.Approve(
			auth,
			taskContractAddress,
			cnxBalance,
		)

		if err != nil {
			return err
		}

		_, err = bind.WaitMined(context.Background(), client, tx)
		if err != nil {
			return err
		}
	}

	return nil
}
