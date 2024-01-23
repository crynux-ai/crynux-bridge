package tasks

import (
	"crynux_bridge/blockchain"
	"crynux_bridge/config"
	"crynux_bridge/models"
	"regexp"
	"time"

	log "github.com/sirupsen/logrus"
)

func StartGetTaskCreationResultWithTerminateChannel(ch <-chan int) {
	for {
		select {
		case stop := <-ch:
			if stop == 1 {
				return
			} else {
				processGetTaskResults()
			}
		default:
			processGetTaskResults()
		}
	}
}

func StartGetTaskCreationResult() {
	for {
		processGetTaskResults()
	}
}

func processGetTaskResults() {
	batchSize := 100
	interval := 1

	var tasks []models.InferenceTask

	tx := config.GetDB().Limit(batchSize).Order("id asc")

	if err := tx.Where(&models.InferenceTask{Status: models.InferenceTaskTransactionSent}).Omit("TaskArgs").Find(&tasks).Error; err != nil {
		log.Errorln("error while getting tasks from database")
		log.Errorln(err)
		time.Sleep(time.Duration(interval) * time.Second)
		return
	}

	for _, task := range tasks {

		taskId, err := blockchain.GetTaskCreationResult(task.TxHash)

		if err != nil {

			errMsg := err.Error()

			match, errM := regexp.MatchString("VM Exception while processing transaction: revert", errMsg)

			if errM != nil {
				log.Errorln("error parsing error message from relay server")
				continue
			}

			if match {
				log.Infoln("Transaction has been reverted")
				log.Infoln(errMsg)
				log.Infoln("Task should be aborted")

				if err := task.AbortWithReason(errMsg, config.GetDB()); err != nil {
					log.Errorln("error updating task status")
					log.Errorln(err)
				}

				continue
			}

			log.Errorln("Error while getting task creation result from the blockchain")
			log.Errorln(err)
			log.Errorln("Try again later...")

			continue
		}

		if taskId == nil {
			// Not ready yet
			log.Debugln("transaction not confirmed: " + task.TxHash)
			continue
		}

		log.Debugln("transaction confirmed: " + task.TxHash)
		log.Debugln("created task id onchain: " + taskId.String())

		if err := updateTaskIdAndStatus(&task, taskId.Uint64(), models.InferenceTaskBlockchainConfirmed); err != nil {
			log.Errorln("error updating task status")
			log.Errorln(err)
		}
	}

	time.Sleep(time.Duration(interval) * time.Second)
}

func updateTaskIdAndStatus(task *models.InferenceTask, taskId uint64, status models.TaskStatus) error {

	task.TaskId = taskId
	task.Status = status

	tx := config.GetDB().Model(task).Select("Status", "TaskId")
	return tx.Updates(&task).Error
}
