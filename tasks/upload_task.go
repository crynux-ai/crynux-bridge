package tasks

import (
	"crynux_bridge/config"
	"crynux_bridge/models"
	"crynux_bridge/relay"
	"regexp"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
)

func StartUploadTaskParamsWithTerminateChannel(ch <-chan int) {

	for {
		select {
		case stop := <-ch:
			if stop == 1 {
				return
			} else {
				processUploadTasks()
			}
		default:
			processUploadTasks()
		}
	}
}

func StartUploadTaskParams() {
	for {
		processUploadTasks()
	}
}

func processUploadTasks() {

	batchSize := 100
	interval := 1

	var tasks []models.InferenceTask

	tx := config.GetDB().Limit(batchSize).Order("id asc")

	if err := tx.Where(&models.InferenceTask{Status: models.InferenceTaskBlockchainConfirmed}).Find(&tasks).Error; err != nil {
		log.Errorln("error while getting tasks from database")
		log.Errorln(err)
		time.Sleep(time.Duration(interval) * time.Second)
		return
	}

	for _, task := range tasks {
		if err := relay.UploadTask(&task); err != nil {

			match, errM := regexp.MatchString("Task not found on the Blockchain", err.Error())

			if errM != nil {
				log.Errorln("Error parsing error message from relay server")
				log.Errorln("Try again later...")
				continue
			}

			if match {
				log.Debugln("Relay server reported task not confirmed on the Blockchain")
				log.Debugln("Still need to wait")
				continue
			}

			match, errM = regexp.MatchString("Task already uploaded", err.Error())
			if errM != nil {
				log.Errorln("Error parsing error message from relay server")
				log.Errorln("Try again later...")
				continue
			}

			if match {
				log.Debugln("Task params already uploaded.")
				log.Debugln("Might be due to the last timeout")
				log.Debugln("Proceed to the next step")
				task.Status = models.InferenceTaskParamsUploaded
			} else {
				log.Errorln("Error while uploading task params to the relay server")
				log.Errorln(err)
				log.Errorln("Try again later...")
				continue
			}

		} else {
			task.Status = models.InferenceTaskParamsUploaded
		}

		if err := config.GetDB().Model(task).Update("Status", task.Status).Error; err != nil {
			log.Errorln(err)
			continue
		}

		log.Debugln("Task uploaded to the relay server:" + strconv.FormatUint(task.TaskID, 10))
	}

	time.Sleep(time.Duration(interval) * time.Second)
}
