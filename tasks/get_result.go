package tasks

import (
	log "github.com/sirupsen/logrus"
	"ig_server/config"
	"ig_server/models"
	"ig_server/relay"
	"strconv"
	"time"
)

func StartDownloadResultsWithTerminateChannel(ch <-chan int) {
	for {
		select {
		case stop := <-ch:
			if stop == 1 {
				return
			} else {
				processGetResultTasks()
			}
		default:
			processGetResultTasks()
		}
	}
}

func StartDownloadResults() {
	for {
		processGetResultTasks()
	}
}

func processGetResultTasks() {
	batchSize := 100
	interval := 1

	var tasks []models.InferenceTask

	tx := config.GetDB().Limit(batchSize).Order("id asc")

	if err := tx.Where(&models.InferenceTask{Status: models.InferenceTaskPendingResult}).Find(&tasks).Error; err != nil {
		log.Errorln("error while getting tasks from database")
		log.Errorln(err)
		time.Sleep(time.Duration(interval) * time.Second)
		return
	}

	for _, task := range tasks {

		log.Debugln("downloading task results for task id: " + strconv.FormatUint(task.TaskId, 10))

		if err := relay.DownloadTaskResult(&task); err != nil {

			log.Errorln("error while getting image from relay server")
			log.Errorln(err)
			log.Debugln("wait for a longer time")
			continue
		}

		task.Status = models.InferenceTaskSuccess

		if err := config.GetDB().Model(task).Select("Status", "TaskId").Updates(&task).Error; err != nil {
			log.Errorln(err)
			continue
		}
	}

	time.Sleep(time.Duration(interval) * time.Second)

}
