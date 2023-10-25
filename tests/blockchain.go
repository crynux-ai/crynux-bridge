package tests

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"ig_server/blockchain"
	"ig_server/config"
	"ig_server/models"
	"math/big"
)

func PrepareAccounts() (addresses []string, privateKeys []string, err error) {

	// Create 4 accounts
	// TaskCreator, Node1, Node2, Node3
	// Transfer tokens to these accounts
	tokenInstance, err := blockchain.GetCrynuxTokenContractInstance()
	if err != nil {
		log.Errorln("error get token contract instance")
		log.Errorln(err)
		return nil, nil, err
	}

	client, err := blockchain.GetRpcClient()
	if err != nil {
		log.Errorln("error connect to the websocket endpoint")
		log.Errorln(err)
		return nil, nil, err
	}

	rootAddress := common.HexToAddress(config.GetConfig().Test.RootAddress)
	rootPrivateKey := config.GetConfig().Test.RootPrivateKey

	for i := 0; i <= 3; i++ {
		privateKey, err := crypto.GenerateKey()
		if err != nil {
			return nil, nil, err
		}
		privateKeyBytes := crypto.FromECDSA(privateKey)
		privateKeyStr := hexutil.Encode(privateKeyBytes)[2:]
		privateKeys = append(privateKeys, privateKeyStr)

		publicKey := privateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			return nil, nil, errors.New("error casting public key to ECDSA")
		}

		address := crypto.PubkeyToAddress(*publicKeyECDSA)
		addresses = append(addresses, address.Hex())

		// Transfer ether to these accounts
		auth, err := blockchain.GetAuth(client, rootAddress, rootPrivateKey)
		if err != nil {
			return nil, nil, err
		}

		amount := new(big.Int).Mul(big.NewInt(1000), big.NewInt(params.GWei))
		tx, err := blockchain.SendETH(rootAddress, address, amount, rootPrivateKey)

		_, err = bind.WaitMined(context.Background(), client, tx)
		if err != nil {
			return nil, nil, err
		}

		// Transfer CNX to these accounts
		auth, err = blockchain.GetAuth(client, rootAddress, rootPrivateKey)
		if err != nil {
			return nil, nil, err
		}

		tx, err = tokenInstance.Transfer(auth, address, new(big.Int).Mul(big.NewInt(1000), big.NewInt(params.Ether)))
		if err != nil {
			return nil, nil, err
		}

		_, err = bind.WaitMined(context.Background(), client, tx)
		if err != nil {
			return nil, nil, err
		}
	}

	appConfig := config.GetConfig()
	appConfig.Blockchain.Account.Address = addresses[0]
	appConfig.Blockchain.Account.PrivateKey = privateKeys[0]

	return addresses, privateKeys, nil
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

		tx, err = nodeInstance.Join(auth)
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

		_, err = bind.WaitMined(context.Background(), client, tx)
		if err != nil {
			return err
		}
	}

	log.Debugln("all task commitments submitted")

	for i := 1; i <= 3; i++ {
		address := common.HexToAddress(addresses[i])
		auth, err := blockchain.GetAuth(client, address, privateKeys[i])
		if err != nil {
			return err
		}

		tx, err := taskInstance.DiscloseTaskResult(auth, taskId, big.NewInt(int64(rounds[addresses[i]])), results[i-1])

		_, err = bind.WaitMined(context.Background(), client, tx)
		if err != nil {
			return err
		}
	}

	log.Debugln("all task results disclosed")

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
