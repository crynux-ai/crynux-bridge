package tasks

import (
	"crynux_bridge/blockchain"
	"crynux_bridge/config"
	"crynux_bridge/models"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
)

func StartSendTaskOnChainWithTerminateChannel(ch <-chan int) {

	for {
		select {
		case stop := <-ch:
			if stop == 1 {
				return
			} else {
				processTasks()
			}
		default:
			processTasks()
		}
	}
}

func StartSendTaskOnChain() {
	for {
		processTasks()
	}
}

func processTasks() {

	batchSize := 100
	interval := 1

	var tasks []models.InferenceTask

	tx := config.GetDB().Order("id asc").Limit(batchSize)

	if err := tx.Where(map[string]interface{}{"Status": models.InferenceTaskPending}).Find(&tasks).Error; err != nil {
		log.Errorln("error while getting tasks from database")
		log.Errorln(err)
		time.Sleep(time.Duration(interval) * time.Second)
		return
	}

	for _, task := range tasks {

		log.Debugln("send task to the blockchain: " + strconv.FormatUint(uint64(task.ID), 10))

		txHash, err := blockchain.CreateTaskOnChain(&task)

		if err != nil {

			log.Errorln("error while sending task to the blockchain")
			log.Errorln(err)
			log.Errorln("abort the task for now")

			if err := task.AbortWithReason(err.Error(), config.GetDB()); err != nil {
				log.Errorln("error while updating task status to the db")
				log.Errorln(err)
			}

			continue
		}

		log.Debugln("tx hash: " + txHash)

		updateModel := &models.InferenceTask{
			TxHash: txHash,
			Status: models.InferenceTaskTransactionSent,
		}

		if err := config.GetDB().Model(&task).Select("TxHash", "Status").Updates(updateModel).Error; err != nil {
			log.Errorln("error while updating task status to the db")
			log.Errorln(err)
			continue
		}
	}

	time.Sleep(time.Duration(interval) * time.Second)
}
