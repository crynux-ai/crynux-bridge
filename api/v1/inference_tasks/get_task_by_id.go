package inference_tasks

import (
	"crynux_bridge/api/v1/response"
	"crynux_bridge/api/v1/tools"
	"crynux_bridge/config"
	"crynux_bridge/models"
	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GetTaskInput struct {
	ClientID     string `path:"client_id" json:"client_id" description:"Client id" validate:"required"`
	ClientTaskID uint   `path:"client_task_id" json:"client_task_id" description:"Client task id" validate:"required"`
}

type GetTaskResponse struct {
	response.Response
	Data *models.InferenceTask `json:"data"`
}

func GetTaskById(c *gin.Context, in *GetTaskInput) (*GetTaskResponse, error) {
	ctx := c.Request.Context()
	db := config.GetDB()

	// get Client from local db
	client, err := tools.GetClient(ctx, db, in.ClientID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.NewValidationErrorResponse("client_id", "Client not found")
		} else {
			return nil, response.NewExceptionResponse(err)
		}
	}

	// get ClientTask from local db
	clientTask, err := tools.GetClientTask(ctx, db, client.ID, in.ClientTaskID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.NewValidationErrorResponse("client_task_id", "Client task not found")
		} else {
			return nil, response.NewExceptionResponse(err)
		}
	}
	if len(clientTask.InferenceTasks) == 0 {
		return nil, response.NewValidationErrorResponse("client_task_id", "Client task has no associated tasks")
	}

	task := clientTask.InferenceTasks[0]
	for _, t := range clientTask.InferenceTasks[1:] {
		if t.Status == models.InferenceTaskResultDownloaded {
			if task.Status != models.InferenceTaskResultDownloaded {
				task = t
			} else if task.UpdatedAt.Sub(t.UpdatedAt) > 0 {
				task = t
			}
		} else if task.Status == models.InferenceTaskEndAborted && t.Status != models.InferenceTaskEndAborted {
			task = t
		}
	}

	return &GetTaskResponse{
		Data: &task,
	}, nil
}
