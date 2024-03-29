package inference_tasks

import (
	"crynux_bridge/api/v1/response"
	"crynux_bridge/config"
	"crynux_bridge/models"
	"errors"

	"github.com/gin-gonic/gin"
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

func CreateTask(_ *gin.Context, in *TaskInput) (*TaskResponse, error) {

	client := &models.Client{ClientId: in.ClientID}

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

	task := &models.InferenceTask{
		Client:    *client,
		TaskArgs:  in.TaskArgs,
		TaskType:  *in.TaskType,
		VramLimit: vramLimit,
		TaskFee:   30,
		Cap:       cap,
	}

	if err := config.GetDB().Create(task).Error; err != nil {
		return nil, response.NewExceptionResponse(err)
	}

	return &TaskResponse{Data: *task}, nil
}
