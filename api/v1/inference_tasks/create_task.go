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
	ClientID string `json:"client_id" description:"Client id" validate:"required"`
	TaskArgs string `json:"task_args" description:"Task args" validate:"required"`
}

type TaskResponse struct {
	response.Response
	Data models.InferenceTask `json:"data"`
}

func CreateTask(_ *gin.Context, in *TaskInput) (*TaskResponse, error) {

	client := &models.Client{ClientId: in.ClientID}

	if err := config.GetDB().Where(client).First(&client).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.NewExceptionResponse(err)
		}
	}

	result, err := models.ValidateTaskArgsJsonStr(in.TaskArgs)
	if err != nil {
		return nil, response.NewExceptionResponse(err)
	}

	if result != nil {
		return nil, response.NewValidationErrorResponse("task_args", result.Error())
	}

	task := &models.InferenceTask{
		Client:   *client,
		TaskArgs: in.TaskArgs,
	}

	if err := config.GetDB().Create(task).Error; err != nil {
		return nil, response.NewExceptionResponse(err)
	}

	return &TaskResponse{Data: *task}, nil
}
