package inference_tasks

import (
	"crynux_bridge/api/v1/response"
	"crynux_bridge/config"
	"crynux_bridge/models"
	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GetTaskInput struct {
	ClientID string `path:"client_id" json:"client_id" description:"Client id" validate:"required"`
	TaskID   uint   `path:"task_id" json:"task_id" description:"Task id" validate:"required"`
}

type GetTaskResponse struct {
	response.Response
	Data *models.InferenceTask `json:"data"`
}

func GetTaskById(_ *gin.Context, in *GetTaskInput) (*GetTaskResponse, error) {

	client := &models.Client{ClientId: in.ClientID}

	if err := config.GetDB().Where(client).First(client).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.NewValidationErrorResponse("client_id", "Client not found")
		} else {
			return nil, response.NewExceptionResponse(err)
		}
	}

	task := &models.InferenceTask{
		ClientID: client.ID,
	}

	if err := config.GetDB().Where(task).Omit("Pose.DataURL").First(task, in.TaskID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.NewValidationErrorResponse("task_id", "Task not found")
		} else {
			return nil, response.NewExceptionResponse(err)
		}
	}

	return &GetTaskResponse{
		Data: task,
	}, nil
}
