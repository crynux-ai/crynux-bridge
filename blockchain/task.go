package blockchain

import (
	"context"
	"crynux_bridge/blockchain/bindings"
	"crynux_bridge/config"
	"crynux_bridge/models"
	"crynux_bridge/utils"
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/binary"
	"errors"
	"image"
	"image/png"
	"io"
	"math/big"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/corona10/goimagehash"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	log "github.com/sirupsen/logrus"
)

func GetTaskByCommitment(ctx context.Context, taskIDCommitment [32]byte) (*bindings.VSSTaskTaskInfo, error) {
	taskInstance, err := GetTaskContractInstance()
	if err != nil {
		return nil, err
	}

	callCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	opts := &bind.CallOpts{
		Pending: false,
		Context: callCtx,
	}

	if err := getLimiter().Wait(callCtx); err != nil {
		return nil, err
	}
	taskInfo, err := taskInstance.GetTask(opts, taskIDCommitment)
	if err != nil {
		return nil, err
	}

	return &taskInfo, nil
}

func CreateTaskOnChain(ctx context.Context, task *models.InferenceTask) (string, error) {
	taskInstance, err := GetTaskContractInstance()
	if err != nil {
		return "", err
	}

	appConfig := config.GetConfig()
	address := common.HexToAddress(appConfig.Blockchain.Account.Address)
	privkey := appConfig.Blockchain.Account.PrivateKey

	auth, err := GetAuth(ctx, address, privkey)
	if err != nil {
		return "", err
	}

	callCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	if err := getLimiter().Wait(callCtx); err != nil {
		return "", err
	}
	auth.Context = callCtx

	taskIDCommitment, _ := utils.HexStrToBytes32(task.TaskIDCommitment)
	nonce, _ := utils.HexStrToBytes32(task.Nonce)

	versionArr := strings.SplitN(task.TaskVersion, ".", 3)
	if len(versionArr) != 3 {
		return "", errors.New("Task version invalid")
	}
	var taskVersion [3]*big.Int
	for i := 0; i < 3; i++ {
		versionStr := versionArr[i]
		version, err := strconv.Atoi(versionStr)
		if err != nil {
			return "", errors.New("Task version invalid")
		}
		taskVersion[i] = big.NewInt(int64(version))
	}

	tx, err := taskInstance.CreateTask(
		auth,
		uint8(task.TaskType),
		*taskIDCommitment,
		*nonce,
		task.TaskModelIDs,
		big.NewInt(int64(task.MinVram)),
		task.RequiredGPU,
		big.NewInt(int64(task.RequiredGPUVram)),
		taskVersion,
		big.NewInt(int64(task.TaskSize)),
	)
	if err != nil {
		return "", err
	}

	return tx.Hash().Hex(), nil
}

func ValidateSingleTask(ctx context.Context, task *models.InferenceTask) (string, error) {
	taskInstance, err := GetTaskContractInstance()
	if err != nil {
		return "", err
	}

	appConfig := config.GetConfig()
	address := common.HexToAddress(appConfig.Blockchain.Account.Address)
	privkey := appConfig.Blockchain.Account.PrivateKey

	auth, err := GetAuth(ctx, address, privkey)
	if err != nil {
		return "", err
	}

	callCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	if err := getLimiter().Wait(callCtx); err != nil {
		return "", err
	}
	auth.Context = callCtx

	taskIDCommitment, _ := utils.HexStrToBytes32(task.TaskIDCommitment)
	vrfProof, _ := hexutil.Decode(task.VRFProof)
	privateKey, err := crypto.HexToECDSA(privkey)
	if err != nil {
		return "", err
	}

	publicKey := privateKey.Public()

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", errors.New("error casting public key to ECDSA")
	}
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	if len(publicKeyBytes) != 65 {
		return "", errors.New("umcompressed public key bytes length is not 65")
	}
	publicKeyBytes = publicKeyBytes[1:]

	tx, err := taskInstance.ValidateSingleTask(auth, *taskIDCommitment, vrfProof, publicKeyBytes)
	if err != nil {
		return "", err
	}
	return tx.Hash().Hex(), nil
}

func ValidateTaskGroup(ctx context.Context, task1, task2, task3 *models.InferenceTask) (string, error) {
	taskInstance, err := GetTaskContractInstance()
	if err != nil {
		return "", err
	}

	appConfig := config.GetConfig()
	address := common.HexToAddress(appConfig.Blockchain.Account.Address)
	privkey := appConfig.Blockchain.Account.PrivateKey

	auth, err := GetAuth(ctx, address, privkey)
	if err != nil {
		return "", err
	}

	callCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	if err := getLimiter().Wait(callCtx); err != nil {
		return "", err
	}
	auth.Context = callCtx

	if !(task1.TaskID == task2.TaskID && task1.TaskID == task3.TaskID) {
		return "", errors.New("taskID of tasks in group is not the same")
	}
	if !(task1.Sequence < task2.Sequence && task2.Sequence < task3.Sequence) {
		return "", errors.New("task order of tasks in group is incorrect")
	}

	taskIDCommitment1, _ := utils.HexStrToBytes32(task1.TaskIDCommitment)
	taskIDCommitment2, _ := utils.HexStrToBytes32(task2.TaskIDCommitment)
	taskIDCommitment3, _ := utils.HexStrToBytes32(task3.TaskIDCommitment)
	taskID, _ := utils.HexStrToBytes32(task1.TaskID)
	vrfProof, _ := hexutil.Decode(task1.VRFProof)
	privateKey, err := crypto.HexToECDSA(privkey)
	if err != nil {
		return "", err
	}

	publicKey := privateKey.Public()

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", errors.New("error casting public key to ECDSA")
	}
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	if len(publicKeyBytes) != 65 {
		return "", errors.New("umcompressed public key bytes length is not 65")
	}
	publicKeyBytes = publicKeyBytes[1:]
	tx, err := taskInstance.ValidateTaskGroup(auth, *taskIDCommitment1, *taskIDCommitment2, *taskIDCommitment3, *taskID, vrfProof, publicKeyBytes)
	if err != nil {
		return "", err
	}
	return tx.Hash().Hex(), nil

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
