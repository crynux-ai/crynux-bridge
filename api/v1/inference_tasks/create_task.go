package inference_tasks

import (
	"context"
	"crynux_bridge/api/v1/response"
	"crynux_bridge/api/v1/tools"
	"crynux_bridge/config"
	"crynux_bridge/models"
	"crypto/rand"
	"errors"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"gorm.io/gorm"
)

type TaskInput struct {
	ClientID        string                `json:"client_id" description:"Client id" validate:"required"`
	TaskArgs        string                `json:"task_args" description:"Task args" validate:"required"`
	TaskType        *models.ChainTaskType `json:"task_type" description:"Task type. 0 - SD task, 1 - LLM task, 2 - SD Finetune task" validate:"required"`
	TaskVersion     *string               `json:"task_version,omitempty" description:"Task version. Default is 2.5.0" validate:"omitempty"`
	MinVram         *uint64               `json:"min_vram,omitempty" description:"Task minimal vram requirement" validate:"omitempty"`
	RequiredGPU     string                `json:"required_gpu,omitempty" description:"Task required GPU name" validate:"omitempty"`
	RequiredGPUVram uint64                `json:"required_gpu_vram,omitempty" description:"Task required GPU Vram" validate:"omitempty"`
	RepeatNum       *int                  `json:"repeat_num,omitempty" description:"Task repeat number" validate:"omitempty"`
}

type TaskResponse struct {
	response.Response
	Data *models.ClientTask `json:"data"`
}

func getDefaultMinVram(taskType models.ChainTaskType, taskArgs string) (uint64, error) {
	if taskType == models.TaskTypeSD {
		baseModel, err := models.GetSDTaskConfigBaseModel(taskArgs)
		if err != nil {
			return 0, err
		}
		if baseModel == "crynux-ai/stable-diffusion-v1-5" {
			return 8, nil
		} else if baseModel == "crynux-ai/sdxl-turbo" || baseModel == "crynux-ai/stable-diffusion-xl-base-1.0" {
			return 14, nil
		} else {
			return 10, nil
		}
	} else {
		return 8, nil
	}
}

func getTaskSize(taskType models.ChainTaskType, taskArgs string) (uint64, error) {
	if taskType == models.TaskTypeSD {
		num, err := models.GetTaskConfigNumImages(taskArgs)
		if err != nil {
			return 0, err
		}
		return uint64(num), nil
	} else {
		return 1, nil
	}
}

func getTaskFee(taskType models.ChainTaskType, baseTaskFee, cap uint64) uint64 {
	if taskType == models.TaskTypeSD {
		return baseTaskFee * cap
	} else {
		return baseTaskFee * cap
	}
}

var clientRateLimiters map[string]*rate.Limiter = make(map[string]*rate.Limiter)

func getClientRateLimiter(clientID string) *rate.Limiter {
	limiter, ok := clientRateLimiters[clientID]
	if !ok {
		var interval time.Duration = time.Minute
		limiter = rate.NewLimiter(rate.Every(interval), 20)
		clientRateLimiters[clientID] = limiter
	}
	return limiter
}

func buildTasks(in *TaskInput, client *models.Client, clientTask *models.ClientTask, appConfig *config.AppConfig) ([]*models.InferenceTask, error) {
	taskType := *in.TaskType

	result, err := models.ValidateTaskArgsJsonStr(in.TaskArgs, taskType)
	if err != nil {
		return nil, response.NewExceptionResponse(err)
	}

	if result != nil {
		return nil, response.NewValidationErrorResponse("task_args", result.Error())
	}

	var minVram uint64

	if in.MinVram == nil {
		// task args has been validated, so there should be no error
		minVram, _ = getDefaultMinVram(taskType, in.TaskArgs)
	} else {
		minVram = *in.MinVram
	}

	var taskVersion = "2.5.0"
	if in.TaskVersion != nil {
		taskVersion = *in.TaskVersion
	}

	// task args has been validated, so there should be no error
	taskSize, _ := getTaskSize(taskType, in.TaskArgs)
	taskFee := getTaskFee(taskType, appConfig.Task.TaskFee, taskSize) // unit: GWei

	repeatNum := appConfig.Task.RepeatNum
	if in.RepeatNum != nil {
		repeatNum = *in.RepeatNum
	}

	modelIDs, err := models.GetTaskConfigModelIDs(in.TaskArgs, taskType)
	if err != nil {
		return nil, response.NewExceptionResponse(err)
	}

	taskIDBytes := make([]byte, 32)
	rand.Read(taskIDBytes)
	taskID := hexutil.Encode(taskIDBytes)

	tasks := make([]*models.InferenceTask, 0)
	for i := 0; i < repeatNum; i++ {
		task := &models.InferenceTask{
			Client:          *client,
			ClientTask:      *clientTask,
			TaskArgs:        in.TaskArgs,
			TaskType:        taskType,
			TaskModelIDs:    modelIDs,
			TaskVersion:     taskVersion,
			TaskFee:         taskFee,
			MinVram:         minVram,
			RequiredGPU:     in.RequiredGPU,
			RequiredGPUVram: in.RequiredGPUVram,
			TaskSize:        taskSize,
			TaskID:          taskID,
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func DoCreateTask(ctx context.Context, in *TaskInput) (*TaskResponse, error) {
	appConfig := config.GetConfig()
	db := config.GetDB()

	// check rate limit
	limiter := getClientRateLimiter(in.ClientID)
	if !limiter.Allow() {
		err := errors.New("CREATE TASK TOO FREQUENTLY")
		return nil, response.NewExceptionResponse(err)
	}

	// get Client, if not exist, create a new one
	client, err := tools.GetClient(ctx, db, in.ClientID)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.NewExceptionResponse(err)
		}
	}

	// create ClientTask for client
	clientTask, err := tools.CreateClientTask(ctx, db, client)
	if err != nil {
		return nil, response.NewExceptionResponse(err)
	}

	// build interface tasks
	tasks, err := buildTasks(in, client, clientTask, appConfig)
	if err != nil {
		return nil, err
	}

	// save tasks to local db
	err = models.SaveTasks(ctx, config.GetDB(), tasks)
	if err != nil {
		return nil, response.NewExceptionResponse(err)
	}

	return &TaskResponse{Data: clientTask}, nil
}

func CreateTask(c *gin.Context, in *TaskInput) (*TaskResponse, error) {
	ctx := c.Request.Context()

	return DoCreateTask(ctx, in)
}
