package tasks

import (
	"context"
	"crynux_bridge/blockchain"
	"crynux_bridge/config"
	"crynux_bridge/models"
	"errors"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func StartSyncBlockWithTerminateChannel(ch <-chan int) {

	syncedBlock, err := getSyncedBlock()

	if err != nil {
		log.Errorln("error getting synced block from the database")
		log.Fatal(err)
	}

	for {
		select {
		case stop := <-ch:
			if stop == 1 {
				return
			} else {
				processChannel(syncedBlock)
			}
		default:
			processChannel(syncedBlock)
		}
	}
}

func StartSyncBlock() {
	syncedBlock, err := getSyncedBlock()

	if err != nil {
		log.Errorln("error getting synced block from the database")
		log.Fatal(err)
	}

	for {
		processChannel(syncedBlock)
	}
}

func getSyncedBlock() (*models.SyncedBlock, error) {
	appConfig := config.GetConfig()
	syncedBlock := &models.SyncedBlock{}

	if err := config.GetDB().First(&syncedBlock).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			syncedBlock.BlockNumber = appConfig.Blockchain.StartBlockNum
		} else {
			return nil, err
		}
	}

	return syncedBlock, nil
}

func processChannel(syncedBlock *models.SyncedBlock) {

	interval := 1
	batchSize := uint64(500)

	client, err := blockchain.GetRpcClient()
	if err != nil {
		log.Errorln("error getting the eth rpc client")
		log.Errorln(err)
		time.Sleep(time.Duration(interval) * time.Second)
		return
	}

	latestBlockNum, err := client.BlockNumber(context.Background())
	if err != nil {
		log.Errorln("error getting the latest block number")
		log.Errorln(err)
		time.Sleep(time.Duration(interval) * time.Second)
		return
	}

	if latestBlockNum <= syncedBlock.BlockNumber {
		time.Sleep(time.Duration(interval) * time.Second)
		return
	}

	log.Debugln("new block received: " + strconv.FormatUint(latestBlockNum, 10))

	for start := syncedBlock.BlockNumber + 1; start <= latestBlockNum; start += batchSize {

		end := start + batchSize - 1

		if end > latestBlockNum {
			end = latestBlockNum
		}

		log.Debugln("processing blocks from " +
			strconv.FormatUint(start, 10) +
			" to " +
			strconv.FormatUint(end, 10) +
			" / " +
			strconv.FormatUint(latestBlockNum, 10))

		if err := processTaskSuccess(start, end); err != nil {
			log.Errorln(err)
			time.Sleep(time.Duration(interval) * time.Second)
			return
		}

		if err := processTaskAborted(start, end); err != nil {
			log.Errorln(err)
			time.Sleep(time.Duration(interval) * time.Second)
			return
		}

		oldNum := syncedBlock.BlockNumber
		syncedBlock.BlockNumber = end
		if err := config.GetDB().Save(syncedBlock).Error; err != nil {
			syncedBlock.BlockNumber = oldNum
			log.Errorln(err)
			time.Sleep(time.Duration(interval) * time.Second)
		}

		if end != latestBlockNum {
			time.Sleep(time.Duration(interval) * time.Second)
		}
	}

	time.Sleep(time.Duration(interval) * time.Second * 3)
}

func processTaskSuccess(startBlockNum, endBlockNum uint64) error {
	taskContractInstance, err := blockchain.GetTaskContractInstance()
	if err != nil {
		return err
	}

	taskSuccessEventIterator, err := taskContractInstance.FilterTaskSuccess(
		&bind.FilterOpts{
			Start:   startBlockNum,
			End:     &endBlockNum,
			Context: context.Background(),
		},
		nil,
		nil,
	)

	if err != nil {
		return err
	}

	for {
		if !taskSuccessEventIterator.Next() {
			log.Debugln("no more task success events.")
			break
		}

		taskSuccess := taskSuccessEventIterator.Event

		log.Debugln("received TaskSuccess: " + taskSuccess.TaskId.String())
		log.Debugln("result node: " + taskSuccess.ResultNode.Hex())

		task := &models.InferenceTask{
			TaskId: taskSuccess.TaskId.Uint64(),
			Status: models.InferenceTaskParamsUploaded,
		}

		if err := config.GetDB().Where(task).Select("Status", "ID").First(task).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// This task is not created by our server.
				// Just skip it.
				log.Debugln("Task not found in db. Might be created by other servers")
				continue
			} else {
				return err
			}
		}

		log.Debugln("Update task status for task of primary key: " + strconv.FormatUint(uint64(task.ID), 10))

		updates := &models.InferenceTask{
			ResultNode: taskSuccess.ResultNode.Hex(),
			Status:     models.InferenceTaskPendingResult,
		}

		tx := config.GetDB().Model(task)

		if err := tx.Select("ResultNode", "Status").Updates(updates).Error; err != nil {
			return err
		}
	}

	if err := taskSuccessEventIterator.Close(); err != nil {
		return err
	}

	return nil
}

func processTaskAborted(startBlockNum, endBlockNum uint64) error {
	taskContractInstance, err := blockchain.GetTaskContractInstance()
	if err != nil {
		log.Fatal(err)
	}

	taskAbortedEventIterator, err := taskContractInstance.FilterTaskAborted(
		&bind.FilterOpts{
			Start:   startBlockNum,
			End:     &endBlockNum,
			Context: context.Background(),
		},
		nil,
	)

	if err != nil {
		return err
	}

	for {
		if !taskAbortedEventIterator.Next() {
			log.Debugln("no more task aborted events.")
			break
		}
		taskAborted := taskAbortedEventIterator.Event

		log.Debugf("%s TaskAborted, reason: %s", taskAborted.TaskId.String(), taskAborted.Reason)

		task := &models.InferenceTask{
			TaskId: taskAborted.TaskId.Uint64(),
		}

		if err := config.GetDB().Where(task).Select("Status", "ID").First(task).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// This task is not created by our server.
				// Just skip it.
				log.Debugln("Task not found in db. Might be created by other servers")
				continue
			} else {
				return err
			}
		}

		if err := task.AbortWithReason("Task aborted on the Blockchain: "+taskAborted.Reason, config.GetDB()); err != nil {
			return err
		}
	}

	if err := taskAbortedEventIterator.Close(); err != nil {
		return err
	}

	return nil
}
