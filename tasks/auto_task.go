package tasks

import (
	"context"
	"crynux_bridge/blockchain"
	"crynux_bridge/config"
	"crynux_bridge/models"
	crand "crypto/rand"
	"fmt"
	"math/rand"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	log "github.com/sirupsen/logrus"
)

func autoCreateTask(ctx context.Context) error {
	clientID := "auto-task"
	client := models.Client{ClientId: clientID}

	clientTask := models.ClientTask{Client: client}
	if err := func() error {
		dbCtx, cancel := context.WithTimeout(ctx, time.Second)
		defer cancel()
		return config.GetDB().WithContext(dbCtx).Create(&clientTask).Error
	}(); err != nil {
		log.Errorf("AutoTask: auto create task failed: %v", err)
		return err
	}

	var taskArgs string
	var minVram uint64
	r := rand.Float64()
	if r < 0.5 {
		prompt := "Self-portrait oil painting,a beautiful cyborg with golden hair,8k"
		seed := rand.Intn(100000000)
		taskArgs = fmt.Sprintf(`{"version":"2.5.0","base_model":{"name":"crynux-ai/sdxl-turbo", "variant": "fp16"},"prompt":"%s","negative_prompt":"","scheduler":{"method":"EulerAncestralDiscreteScheduler","args":{"timestep_spacing":"trailing"}},"task_config":{"num_images":1,"seed":%d,"steps":1,"cfg":0}}`, prompt, seed)
		minVram = 14
	} else {
		prompt := "best quality, ultra high res, photorealistic++++, 1girl, off-shoulder sweater, smiling, faded ash gray messy bun hair+, border light, depth of field, looking at viewer, closeup"
		negativePrompt := "paintings, sketches, worst quality+++++, low quality+++++, normal quality+++++, lowres, normal quality, monochrome++, grayscale++, skin spots, acnes, skin blemishes, age spot, glans"
		seed := rand.Intn(100000000)
		taskArgs = fmt.Sprintf(`{"version":"2.5.0","base_model":{"name":"crynux-ai/stable-diffusion-v1-5", "variant": "fp16"},"prompt":"%s","negative_prompt":"%s","task_config":{"num_images":1,"seed":%d,"steps":25,"cfg":0,"safety_checker":false}}`, prompt, negativePrompt, seed)
		minVram = 4
	}
	taskType := models.TaskTypeSD
	taskModelIDs, err := models.GetTaskConfigModelIDs(taskArgs, taskType)
	if err != nil {
		log.Errorf("AutoTask: cannot get model ids from task args: %v", err)
		return err
	}

	appConfig := config.GetConfig()
	taskFee := appConfig.Task.TaskFee

	taskIDBytes := make([]byte, 32)
	crand.Read(taskIDBytes)
	taskID := hexutil.Encode(taskIDBytes)

	task := &models.InferenceTask{
		Client:       client,
		ClientTask:   clientTask,
		TaskArgs:     taskArgs,
		TaskType:     taskType,
		TaskModelIDs: taskModelIDs,
		TaskVersion:  "2.5.0",
		MinVram:      minVram,
		TaskFee:      taskFee,
		TaskSize:     1,
		TaskID:       taskID,
	}
	if err := task.Save(ctx, config.GetDB()); err != nil {
		log.Errorf("AutoTask: cannot save task %v", err)
		return err
	}

	log.Infof("AutoTask: auto create task %s", taskID)
	return nil
}

func autoCreateTasks(ctx context.Context) {
	appConfig := config.GetConfig()

	for {
		batchSize := 10
		cnt, err := blockchain.GetQueuedTasks(ctx)
		if err != nil {
			log.Errorf("AutoTask: cannot get queued tasks %v", err)
			continue
		}
		queuedTasks := cnt.Uint64()
		log.Infof("AutoTask: queued tasks count: %d", queuedTasks)
		if queuedTasks > appConfig.Task.QueuedTasksLimit {
			batchSize = 1
			time.Sleep(30 * time.Second)
		}
		for i := 0; i < batchSize; i++ {
			go autoCreateTask(ctx)
		}
	}
}

func AutoCreateTasks(ctx context.Context) {
	ctx1, cancel := context.WithCancel(ctx)
	defer cancel()

	go autoCreateTask(ctx1)
	<- ctx1.Done()
	err := ctx1.Err()
	log.Errorf("AutoTask: timeout %v, finish", err)
}