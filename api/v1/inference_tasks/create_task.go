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
}

type TaskResponse struct {
	response.Response
	Data models.InferenceTask `json:"data"`
}

func getDefaultVramLimit(taskType models.ChainTaskType, taskArgs string) (uint64, error) {
	if taskType == models.TaskTypeSD {
		baseModel, err := models.GetTaskConfigBaseModel(taskArgs)
		if err != nil {
			return 0, err
		}
		if baseModel == "runwayml/stable-diffusion-v1-5" {
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

func getTaskFee(taskType models.ChainTaskType) uint64 {
	if taskType == models.TaskTypeSD {
		return 5100000000
	} else {
		return 5200000000
	}
}

var clientRateLimiters map[string]*rate.Limiter = make(map[string]*rate.Limiter)

func getClientRateLimiter(clientID string) *rate.Limiter {
	limiter, ok := clientRateLimiters[clientID]
	if !ok {
		var interval time.Duration = time.Minute
		limiter = rate.NewLimiter(rate.Every(interval), 1)
		clientRateLimiters[clientID] = limiter
	}
	return limiter
}

func CreateTask(_ *gin.Context, in *TaskInput) (*TaskResponse, error) {

	client := &models.Client{ClientId: in.ClientID}

	limiter := getClientRateLimiter(in.ClientID)
	if !limiter.Allow() {
		err := errors.New("CREATE TASK TOO FREQUENTLY")
		return nil, response.NewExceptionResponse(err)
	}

	if err := config.GetDB().Where(client).First(&client).Error; err != nil {
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

	// task args has been validated, so there should be no error
	cap, _ := getTaskCap(*in.TaskType, in.TaskArgs)
	taskFee := getTaskFee(*in.TaskType) // unit: GWei

	task := &models.InferenceTask{
		Client:    *client,
		TaskArgs:  in.TaskArgs,
		TaskType:  *in.TaskType,
		VramLimit: vramLimit,
		TaskFee:   taskFee,
		Cap:       cap,
	}

	if err := config.GetDB().Create(task).Error; err != nil {
		return nil, response.NewExceptionResponse(err)
	}

	return &TaskResponse{Data: *task}, nil
}
