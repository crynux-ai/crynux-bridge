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

	prompt := "Self-portrait oil painting,a beautiful cyborg with golden hair,8k"
	seed := rand.Intn(100000000)
	taskArgs := fmt.Sprintf(`{"version":"2.0.0","base_model":{"name":"crynux-ai/stable-diffusion-xl-base-1.0"},"prompt":"%s","negative_prompt":"","task_config":{"num_images":1,"safety_checker":false,"seed":%d,"steps":40}}`, prompt, seed)
	taskType := models.TaskTypeSD
	var vramLimit uint64 = 10

	appConfig := config.GetConfig()
	taskFee := appConfig.Task.TaskFee

	task := models.InferenceTask{
		Client:    client,
		TaskArgs:  taskArgs,
		TaskType:  taskType,
		VramLimit: vramLimit,
		TaskFee:   taskFee,
		Cap:       1,
	}
	if err := config.GetDB().Create(&task).Error; err != nil {
		log.Errorf("AutoTask: auto create task failed: %v", err)
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
