package tasks

import (
	"crynux_bridge/config"
	"crynux_bridge/models"
	"fmt"
	"math/rand"
	"time"

	log "github.com/sirupsen/logrus"
)

func autoCreateTask() {
	clientID := "auto-task"
	client := models.Client{ClientId: clientID}

	clientTask := models.ClientTask{Client: client}
	if err := config.GetDB().Create(&clientTask).Error; err != nil {
		log.Errorf("AutoTask: auto create task failed: %v", err)
	}

	prompt := "Self-portrait oil painting,a beautiful cyborg with golden hair,8k"
	seed := rand.Intn(100000000)
	taskArgs := fmt.Sprintf(`{"version":"2.0.0","base_model":{"name":"crynux-ai/sdxl-turbo"},"prompt":"%s","negative_prompt":"","scheduler":{"method":"EulerAncestralDiscreteScheduler","args":{"timestep_spacing":"trailing"}},"task_config":{"num_images":1,"seed":%d,"steps":1,"cfg":0}}`, prompt, seed)
	taskType := models.TaskTypeSD
	var vramLimit uint64 = 10

	appConfig := config.GetConfig()
	taskFee := appConfig.Task.TaskFee

	for i := 0; i < appConfig.Task.RepeatNum; i++ {
		task := models.InferenceTask{
			Client:     client,
			ClientTask: clientTask,
			TaskArgs:   taskArgs,
			TaskType:   taskType,
			VramLimit:  vramLimit,
			TaskFee:    taskFee,
			Cap:        1,
		}
		if err := config.GetDB().Create(&task).Error; err != nil {
			log.Errorf("AutoTask: auto create task failed: %v", err)
		}
	}
	log.Info("AutoTask: auto create task success")
}

func StartAutoCreateTask() {
	for {
		autoCreateTask()
		time.Sleep(5 * time.Minute)
	}
}

func StartAutoCreateTaskWithTerminalChannel(ch <-chan int) {
	for {
		select {
		case stop := <-ch:
			if stop == 1 {
				return
			} else {
				autoCreateTask()
			}
		default:
			autoCreateTask()
		}
		time.Sleep(5 * time.Minute)
	}
}
