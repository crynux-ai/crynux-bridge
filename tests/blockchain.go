package tests

import (
	"context"
	"crynux_bridge/blockchain"
	"crynux_bridge/config"
	"crynux_bridge/models"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

const (
	GPUName string = "NVIDIA GeForce GTX 1070 Ti"
	GPUVram int    = 8
)

func PrepareAccounts() (addresses []string, privateKeys []string, err error) {
	// return account created by ganache cli
	addresses = []string{
		"0xe563e647c53ad9d5d28Da50B4e6cc48594117CF1",
		"0x577887519278199ce8F8D80bAcc70fc32b48daD4",
		"0x9229d36c82E4e1d03B086C27d704741D0c78321e",
		"0xEa1A669fd6A705d28239011A074adB3Cfd6cd82B",
	}
	privateKeys = []string{
		"bde1aedbe693f34c3c1502e40fe17b18b7f71757523db93a3038a8cadfe43d4d",
		"a627246a109551432ac5db6535566af34fdddfaa11df17b8afd53eb987e209a2",
		"b171f296622b98cbdc08dcdcb0696f738c3a22d9d367c657117cd3c8d0b71d42",
		"8fb2fc9862b93b5b75cda8202f583711201e4cba5459eefe442b8c5dcc4bdab9",
	}
	err = nil
	return
}

func PrepareNetwork(addresses []string, privateKeys []string) error {

	// Approve the staking tokens to the node contract
	// Join 3 nodes to the network

	tokenInstance, err := blockchain.GetCrynuxTokenContractInstance()
	if err != nil {
		log.Errorln("error get token contract instance")
		log.Errorln(err)
		return err
	}

	nodeInstance, err := blockchain.GetNodeContractInstance()
	if err != nil {
		log.Errorln("error get node contract instance")
		log.Errorln(err)
		return err
	}

	client, err := blockchain.GetRpcClient()
	if err != nil {
		log.Errorln("error connect to the websocket endpoint")
		log.Errorln(err)
		return err
	}

	nodeContractAddress := common.HexToAddress(config.GetConfig().Blockchain.Contracts.Node)
	for i := 1; i <= 3; i++ {
		address := common.HexToAddress(addresses[i])
		auth, err := blockchain.GetAuth(client, address, privateKeys[i])
		if err != nil {
			return err
		}

		tx, err := tokenInstance.Approve(
			auth,
			nodeContractAddress,
			new(big.Int).Mul(big.NewInt(500), big.NewInt(params.Ether)))

		if err != nil {
			return err
		}

		_, err = bind.WaitMined(context.Background(), client, tx)
		if err != nil {
			return err
		}

		auth, err = blockchain.GetAuth(client, address, privateKeys[i])
		if err != nil {
			return err
		}

		tx, err = nodeInstance.Join(auth, GPUName, big.NewInt(int64(GPUVram)))
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

func PrepareTaskCreatorAccount(address string, privateKey string) error {
	client, err := blockchain.GetRpcClient()
	if err != nil {
		log.Errorln("error connect to the websocket endpoint")
		log.Errorln(err)
		return err
	}

	tokenInstance, err := blockchain.GetCrynuxTokenContractInstance()
	if err != nil {
		log.Errorln("error get token contract instance")
		log.Errorln(err)
		return err
	}

	// Approve some tokens to the task contract for the task creator account
	auth, err := blockchain.GetAuth(
		client,
		common.HexToAddress(address),
		privateKey)

	if err != nil {
		return err
	}

	taskContractAddress := common.HexToAddress(config.GetConfig().Blockchain.Contracts.Task)
	tx, err := tokenInstance.Approve(
		auth,
		taskContractAddress,
		new(big.Int).Mul(big.NewInt(1000), big.NewInt(params.Ether)))

	if err != nil {
		return err
	}

	_, err = bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		return err
	}

	// Display current approved amount
	opts := &bind.CallOpts{
		Pending: false,
		Context: context.Background(),
	}

	allowance, err := tokenInstance.Allowance(opts, common.HexToAddress(address), taskContractAddress)
	log.Debugln("allowance for task creator: " + allowance.String())
	balance, err := tokenInstance.BalanceOf(opts, common.HexToAddress(address))
	log.Debugln("balance of task creator: " + balance.String())

	return nil
}

func ClearNetwork(addresses []string, privateKeys []string) error {
	// Quit the network for 3 nodes
	nodeInstance, err := blockchain.GetNodeContractInstance()
	if err != nil {
		log.Errorln("error get node contract instance")
		log.Errorln(err)
		return err
	}

	client, err := blockchain.GetRpcClient()
	if err != nil {
		log.Errorln("error connect to the websocket endpoint")
		log.Errorln(err)
		return err
	}

	for i := 1; i <= 3; i++ {
		address := common.HexToAddress(addresses[i])
		auth, err := blockchain.GetAuth(client, address, privateKeys[i])
		if err != nil {
			return err
		}

		tx, err := nodeInstance.Quit(auth)

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

func SuccessTaskOnChain(taskId *big.Int, addresses []string, privateKeys []string) error {
	results := [3][]byte{
		[]byte("123456789"),
		[]byte("123456789"),
		[]byte("123456789"),
	}
	return SubmitAndDiscloseResults(taskId, addresses, privateKeys, results)
}

func AbortTaskOnChain(taskId *big.Int, addresses []string, privateKeys []string) error {
	results := [3][]byte{
		[]byte("123456789"),
		[]byte("555555555"),
		[]byte("987654321"),
	}
	return SubmitAndDiscloseResults(taskId, addresses, privateKeys, results)
}

func SubmitAndDiscloseResults(taskId *big.Int, addresses []string, privateKeys []string, results [3][]byte) error {
	taskInstance, err := blockchain.GetTaskContractInstance()
	if err != nil {
		return err
	}

	client, err := blockchain.GetRpcClient()
	if err != nil {
		return err
	}

	opts := &bind.CallOpts{
		Pending: false,
		Context: context.Background(),
	}

	log.Debugln("get task info for id: " + taskId.String())

	taskInfo, err := taskInstance.GetTask(opts, taskId)
	if err != nil {
		return err
	}

	if len(taskInfo.SelectedNodes) == 0 {
		// Task failed before successful creation
		log.Debugln("task failed before successful creation")
		return nil
	}

	rounds := make(map[string]int)

	for i := 0; i < 3; i++ {
		selectedAddress := taskInfo.SelectedNodes[i].Hex()
		rounds[selectedAddress] = i
	}

	for i := 1; i <= 3; i++ {
		address := common.HexToAddress(addresses[i])
		auth, err := blockchain.GetAuth(client, address, privateKeys[i])
		if err != nil {
			return err
		}

		commitment, nonce := blockchain.GetTaskResultCommitment(results[i-1])

		tx, err := taskInstance.SubmitTaskResultCommitment(
			auth, taskId, big.NewInt(int64(rounds[addresses[i]])),
			commitment, nonce)
		if err != nil {
			return err
		}

		_, err = bind.WaitMined(context.Background(), client, tx)
		if err != nil {
			return err
		}
	}

	log.Debugln("all task commitments submitted")

	var discloseBlockNum uint64 = 0

	for i := 1; i <= 3; i++ {
		address := common.HexToAddress(addresses[i])
		auth, err := blockchain.GetAuth(client, address, privateKeys[i])
		if err != nil {
			return err
		}

		tx, err := taskInstance.DiscloseTaskResult(auth, taskId, big.NewInt(int64(rounds[addresses[i]])), results[i-1])
		if err != nil {
			return err
		}

		receipt, err := bind.WaitMined(context.Background(), client, tx)
		if err != nil {
			return err
		}

		if discloseBlockNum == 0 {
			discloseBlockNum = receipt.BlockNumber.Uint64()
		}
	}

	log.Debugln("all task results disclosed")

	taskSuccessIterator, err := taskInstance.FilterTaskSuccess(&bind.FilterOpts{
		Start: discloseBlockNum,
	}, nil)
	if err != nil {
		return err
	}

	var resultNode string
	for taskSuccessIterator.Next() {
		taskSuccess := taskSuccessIterator.Event
		resultNode = taskSuccess.ResultNode.Hex()
		break
	}
	if err := taskSuccessIterator.Close(); err != nil {
		return err
	}

	for i := 1; i <= 3; i++ {
		if addresses[i] == resultNode {
			address := common.HexToAddress(addresses[i])
			auth, err := blockchain.GetAuth(client, address, privateKeys[i])
			if err != nil {
				return err
			}

			tx, err := taskInstance.ReportResultsUploaded(auth, taskId, big.NewInt(int64(rounds[resultNode])))
			if err != nil {
				return err
			}

			_, err = bind.WaitMined(context.Background(), client, tx)
			if err != nil {
				return err
			}
		}
	}
	log.Debugln("report results uploaded")

	return nil
}

func SyncToLatestBlock() error {
	client, err := blockchain.GetRpcClient()
	if err != nil {
		return err
	}

	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return err
	}

	syncedBlock := &models.SyncedBlock{}
	if err := config.GetDB().Where(syncedBlock).First(syncedBlock).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	}

	syncedBlock.BlockNumber = header.Number.Uint64()

	if err := config.GetDB().Save(syncedBlock).Error; err != nil {
		return err
	}

	return nil
}
