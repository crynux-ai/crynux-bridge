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
	ClientID  string                `json:"client_id" description:"Client id" validate:"required"`
	TaskArgs  string                `json:"task_args" description:"Task args" validate:"required"`
	TaskType  *models.ChainTaskType `json:"task_type" description:"Task type. 0 - SD task, 1 - LLM task" validate:"required"`
	VramLimit *uint64               `json:"vram_limit,omitempty" description:"Task minimal vram requirement" validate:"omitempty"`
	RepeatNum int                   `json:"repeat_num,omitempty" description:"Task repeat number" default:"2" validate:"omitempty,gt=0"`
}

type TaskResponse struct {
	response.Response
	Data *models.ClientTask `json:"data"`
}

func getDefaultVramLimit(taskType models.ChainTaskType, taskArgs string) (uint64, error) {
	if taskType == models.TaskTypeSD {
		baseModel, err := models.GetTaskConfigBaseModel(taskArgs)
		if err != nil {
			return 0, err
		}
		if baseModel == "crynux-ai/stable-diffusion-v1-5" {
			return 8, nil
		} else {
			return 10, nil
		}
	} else {
		return 8, nil
	}
}

func getTaskCap(taskType models.ChainTaskType, taskArgs string) (uint64, error) {
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

	result, err := models.ValidateTaskArgsJsonStr(in.TaskArgs, *in.TaskType)
	if err != nil {
		return nil, response.NewExceptionResponse(err)
	}

	if result != nil {
		return nil, response.NewValidationErrorResponse("task_args", result.Error())
	}

	var vramLimit uint64

	if in.VramLimit == nil {
		// task args has been validated, so there should be no error
		vramLimit, _ = getDefaultVramLimit(*in.TaskType, in.TaskArgs)
	} else {
		vramLimit = *in.VramLimit
	}

	clientTask := models.ClientTask{
		Client: client,
	}
	if err := config.GetDB().Create(&clientTask).Error; err != nil {
		return nil, response.NewExceptionResponse(err)
	}

	// task args has been validated, so there should be no error
	cap, _ := getTaskCap(*in.TaskType, in.TaskArgs)
	taskFee := getTaskFee(*in.TaskType, appConfig.Task.TaskFee, cap) // unit: GWei

	for i := 0; i < in.RepeatNum; i++ {
		task := &models.InferenceTask{
			Client:     client,
			ClientTask: clientTask,
			TaskArgs:   in.TaskArgs,
			TaskType:   *in.TaskType,
			VramLimit:  vramLimit,
			TaskFee:    taskFee,
			Cap:        cap,
		}

		if err := config.GetDB().Create(task).Error; err != nil {
			return nil, response.NewExceptionResponse(err)
		}
	}

	return &TaskResponse{Data: &clientTask}, nil
}
