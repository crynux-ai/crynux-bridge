package tasks

import (
	"context"
	"crynux_bridge/config"
	"crynux_bridge/models"
	"crynux_bridge/relay"
	crand "crypto/rand"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func generateRandomTask(client models.Client) *models.InferenceTask {
	clientTask := models.ClientTask{Client: client}

	var taskArgs string
	var minVram uint64 = 0
	var requiredGPU string = ""
	var requiredGPUVram uint64 = 0
	var taskType models.ChainTaskType
	r := rand.Float64()
	if r < 0.25 {
		prompt := "Self-portrait oil painting,a beautiful cyborg with golden hair,8k"
		seed := rand.Intn(100000000)
		taskArgs = fmt.Sprintf(`{"base_model":{"name":"crynux-ai/sdxl-turbo", "variant": "fp16"},"prompt":"%s","negative_prompt":"","scheduler":{"method":"EulerAncestralDiscreteScheduler","args":{"timestep_spacing":"trailing"}},"task_config":{"num_images":1,"seed":%d,"steps":1,"cfg":0}}`, prompt, seed)
		minVram = 14
		taskType = models.TaskTypeSD
	} else if r < 0.5 {
		prompt := "best quality, ultra high res, photorealistic++++, 1girl, off-shoulder sweater, smiling, faded ash gray messy bun hair+, border light, depth of field, looking at viewer, closeup"
		negativePrompt := "paintings, sketches, worst quality+++++, low quality+++++, normal quality+++++, lowres, normal quality, monochrome++, grayscale++, skin spots, acnes, skin blemishes, age spot, glans"
		seed := rand.Intn(100000000)
		taskArgs = fmt.Sprintf(`{"base_model":{"name":"crynux-ai/stable-diffusion-v1-5", "variant": "fp16"},"prompt":"%s","negative_prompt":"%s","task_config":{"num_images":1,"seed":%d,"steps":25,"cfg":0,"safety_checker":false}}`, prompt, negativePrompt, seed)
		minVram = 4
		taskType = models.TaskTypeSD
	} else if r < 0.75 {
		seed := rand.Intn(100000000)
		taskArgs = fmt.Sprintf(`{"model":"Qwen/Qwen2.5-7B","messages":[{"role":"user","content":"I want to create a chat bot. Any suggestions?"}],"tools":null,"generation_config":{"max_new_tokens":250,"do_sample":true,"temperature":0.8,"repetition_penalty":1.1},"seed":%d,"dtype":"bfloat16","quantize_bits":4}`, seed)
		requiredGPU = "NVIDIA GeForce RTX 4060"
		requiredGPUVram = 8
		taskType = models.TaskTypeLLM
	} else {
		seed := rand.Intn(100000000)
		taskArgs = fmt.Sprintf(`{"model":"Qwen/Qwen2.5-7B","messages":[{"role":"user","content":"I want to create a chat bot. Any suggestions?"}],"tools":null,"generation_config":{"max_new_tokens":250,"do_sample":true,"temperature":0.8,"repetition_penalty":1.1},"seed":%d,"dtype":"bfloat16"}`, seed)
		requiredGPU = "NVIDIA GeForce RTX 4090"
		requiredGPUVram = 24
		taskType = models.TaskTypeLLM
	}
	taskModelIDs, _ := models.GetTaskConfigModelIDs(taskArgs, taskType)

	appConfig := config.GetConfig()
	taskFee := appConfig.Task.TaskFee

	taskIDBytes := make([]byte, 32)
	crand.Read(taskIDBytes)
	taskID := hexutil.Encode(taskIDBytes)

	task := &models.InferenceTask{
		Client:          client,
		ClientTask:      clientTask,
		TaskArgs:        taskArgs,
		TaskType:        taskType,
		TaskModelIDs:    taskModelIDs,
		TaskVersion:     "2.5.0",
		MinVram:         minVram,
		RequiredGPU:     requiredGPU,
		RequiredGPUVram: requiredGPUVram,
		TaskFee:         taskFee,
		TaskSize:        1,
		TaskID:          taskID,
	}
	return task
}

func getPendingAutoTasksCount(ctx context.Context, client models.Client) (uint64, error) {
	dbCtx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	task := &models.InferenceTask{
		Client: client,
	}
	var count int64
	if err := config.GetDB().WithContext(dbCtx).Model(&task).Where(&task).Where("(status = ? OR status = ?)", models.InferenceTaskPending, models.InferenceTaskStarted).Count(&count).Error; err != nil {
		return 0, err
	}
	return uint64(count), nil
}

func autoCreateTasks(ctx context.Context) error {
	appConfig := config.GetConfig()

	clientID := "auto-task"
	client := models.Client{ClientId: clientID}

	if err := func() error {
		dbCtx, cancel := context.WithTimeout(ctx, time.Second)
		defer cancel()
		err := config.GetDB().WithContext(dbCtx).Model(&client).Where(&client).First(&client).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return config.GetDB().WithContext(dbCtx).Create(&client).Error
			}
			return err
		}
		return nil
	}(); err != nil {
		log.Errorf("AutoTask: create client failed: %v", err)
		return err
	}

	for {
		batchSize := int(appConfig.Task.AutoTasksBatchSize)
		tasks := make([]*models.InferenceTask, batchSize)
		cnt, err := getPendingAutoTasksCount(ctx, client)
		if err != nil {
			log.Errorf("AutoTask: cannot get pending auto tasks count %v", err)
			time.Sleep(2 * time.Second)
			continue
		}
		log.Infof("AutoTask: pending auto tasks count: %d", cnt)
		if cnt > appConfig.Task.PendingAutoTasksLimit {
			time.Sleep(30 * time.Second)
			continue
		}
		queuedTasks, err := relay.GetQueuedTasks(ctx)
		if err != nil {
			log.Errorf("AutoTask: cannot get queued tasks count %v", err)
			time.Sleep(2 * time.Second)
			continue
		}
		log.Infof("AutoTask: queued task count %d", queuedTasks)
		if uint64(queuedTasks) > appConfig.Task.PendingAutoTasksLimit {
			time.Sleep(30 * time.Second)
			continue
		}

		for i := 0; i < batchSize; i++ {
			task := generateRandomTask(client)
			tasks[i] = task
		}
		if err := models.SaveTasks(ctx, config.GetDB(), tasks); err != nil {
			log.Errorf("AutoTask: cannot save auto tasks: %v", err)
			return err
		}
		time.Sleep(5 * time.Second)
	}
}

func AutoCreateTasks(ctx context.Context) {
	ctx1, cancel := context.WithCancel(ctx)
	defer cancel()

	go func() {
		for {
			err := autoCreateTasks(ctx1)
			if err != nil {
				log.Errorf("AutoTask: auto create tasks error: %v", err)
				time.Sleep(5 * time.Second)
			}
		}
	}()
	<-ctx1.Done()
	err := ctx1.Err()
	log.Errorf("AutoTask: timeout %v, finish", err)
}
