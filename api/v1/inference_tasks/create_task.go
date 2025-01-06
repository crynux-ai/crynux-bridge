package inference_tasks

import (
	"crynux_bridge/api/v1/response"
	"crynux_bridge/config"
	"crynux_bridge/models"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"gorm.io/gorm"
)

type TaskInput struct {
	ClientID        string               `json:"client_id" description:"Client id" validate:"required"`
	TaskArgs        string               `json:"task_args" description:"Task args" validate:"required"`
	TaskType        models.ChainTaskType `json:"task_type" description:"Task type. 0 - SD task, 1 - LLM task, 2 - SD Finetune task" validate:"required"`
	TaskVersion     *string              `json:"task_version,omitempty" description:"Task version. Default is 3.0.0" validate:"omitempty"`
	MinVram         *uint64              `json:"min_vram,omitempty" description:"Task minimal vram requirement" validate:"omitempty"`
	RequiredGPU     string               `json:"required_gpu,omitempty" description:"Task required GPU name" validate:"omitempty"`
	RequiredGPUVram uint64               `json:"required_gpu_vram,omitempty" description:"Task required GPU Vram" validate:"omitempty"`
	RepeatNum       *int                 `json:"repeat_num,omitempty" description:"Task repeat number" validate:"omitempty"`
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

func CreateTask(_ *gin.Context, in *TaskInput) (*TaskResponse, error) {
	appConfig := config.GetConfig()
	client := models.Client{ClientId: in.ClientID}

	limiter := getClientRateLimiter(in.ClientID)
	if !limiter.Allow() {
		err := errors.New("CREATE TASK TOO FREQUENTLY")
		return nil, response.NewExceptionResponse(err)
	}

	if err := config.GetDB().Where(&client).First(&client).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.NewExceptionResponse(err)
		}
	}

	result, err := models.ValidateTaskArgsJsonStr(in.TaskArgs, in.TaskType)
	if err != nil {
		return nil, response.NewExceptionResponse(err)
	}

	if result != nil {
		return nil, response.NewValidationErrorResponse("task_args", result.Error())
	}

	var minVram uint64

	if in.MinVram == nil {
		// task args has been validated, so there should be no error
		minVram, _ = getDefaultMinVram(in.TaskType, in.TaskArgs)
	} else {
		minVram = *in.MinVram
	}

	var taskVersion = "3.0.0"
	if in.TaskVersion != nil {
		taskVersion = *in.TaskVersion
	}

	clientTask := models.ClientTask{
		Client: client,
	}
	if err := config.GetDB().Create(&clientTask).Error; err != nil {
		return nil, response.NewExceptionResponse(err)
	}

	// task args has been validated, so there should be no error
	taskSize, _ := getTaskSize(in.TaskType, in.TaskArgs)
	taskFee := getTaskFee(in.TaskType, appConfig.Task.TaskFee, taskSize) // unit: GWei

	repeatNum := appConfig.Task.RepeatNum
	if in.RepeatNum != nil {
		repeatNum = *in.RepeatNum
	}

	modelIDs, err := models.GetTaskConfigModelIDs(in.TaskArgs, in.TaskType)
	if err != nil {
		return nil, response.NewExceptionResponse(err)
	}

	for i := 0; i < repeatNum; i++ {
		task := &models.InferenceTask{
			Client:     client,
			ClientTask: clientTask,
			TaskArgs:   in.TaskArgs,
			TaskType:   in.TaskType,
			TaskModelIDs: modelIDs,
			TaskVersion: taskVersion,
			TaskFee: taskFee,
			MinVram: minVram,
			RequiredGPU: in.RequiredGPU,
			RequiredGPUVram: in.RequiredGPUVram,
			TaskSize: taskSize,
		}

		if err := config.GetDB().Create(task).Error; err != nil {
			return nil, response.NewExceptionResponse(err)
		}
	}

	return &TaskResponse{Data: &clientTask}, nil
}
