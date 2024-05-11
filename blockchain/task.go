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
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	log "github.com/sirupsen/logrus"
)


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

	taskFee := new(big.Int).Mul(big.NewInt(int64(task.TaskFee)), big.NewInt(params.Ether))
	cap := big.NewInt(int64(task.Cap))
	auth.Value = taskFee
	tx, err := instance.CreateTask(auth, big.NewInt(int64(task.TaskType)), *taskHash, *dataHash, big.NewInt(int64(task.VramLimit)), cap)
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

	// There are 6 events emitted from the CreateTask method
	// Approval, Transfer, TaskPending, TaskCreated x 3
	var taskId *big.Int = nil

	for _, eventLog := range receipt.Logs {
		taskPendingEvent, err := taskContractInstance.ParseTaskPending(*eventLog)
		if err != nil {
			errS := err.Error()
			if errS == "no event signature" || errS == "event signature mismatch" {
				continue
			}
			log.Errorln("error parse task pending event: " + receipt.TxHash.Hex())
			return nil, err
		}
		taskId = taskPendingEvent.TaskId
	}

	if taskId == nil {
		log.Errorln("task pending event not found: " + receipt.TxHash.Hex())
		return nil, errors.New("task pending event not found: " + receipt.TxHash.Hex())
	}

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

	ethThreshold := new(big.Int).Mul(big.NewInt(500), big.NewInt(params.Ether))

	if currentETHBalance.Cmp(ethThreshold) != 1 {
		return errors.New("not enough ETH left")
	}

	return nil
}
