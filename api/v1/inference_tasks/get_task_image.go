package inference_tasks

import (
	"context"
	"crynux_bridge/api/v1/response"
	"crynux_bridge/config"
	"crynux_bridge/models"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GetTaskImageInput struct {
	ClientID     string `path:"client_id" description:"Client id" validate:"required"`
	ClientTaskID uint   `path:"client_task_id" description:"Client task id" validate:"required"`
	Index        *uint64 `path:"index" description:"Result index" validate:"required"`
}

func GetTaskImage(c *gin.Context, in *GetTaskImageInput) error {

	client := &models.Client{ClientId: in.ClientID}

	if err := func() error {
		dbCtx, cancel := context.WithTimeout(c.Request.Context(), time.Second)
		defer cancel()
		return config.GetDB().WithContext(dbCtx).Where(client).First(client).Error
	}(); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.NewValidationErrorResponse("client_id", "Client not found")
		} else {
			return response.NewExceptionResponse(err)
		}
	}

	clientTask := &models.ClientTask{
		RootModel: models.RootModel{ID: in.ClientTaskID},
		ClientID:  client.ID,
	}

	if err := func() error {
		dbCtx, cancel := context.WithTimeout(c.Request.Context(), time.Second)
		defer cancel()
		return config.GetDB().WithContext(dbCtx).Model(&clientTask).Where(&clientTask).Preload("InferenceTasks").First(&clientTask).Error
	}(); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.NewValidationErrorResponse("client_task_id", "Client task not found")
		} else {
			return response.NewExceptionResponse(err)
		}
	}
	if len(clientTask.InferenceTasks) == 0 {
		return response.NewValidationErrorResponse("client_task_id", "Client task has no associated tasks")
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

	if task.Status != models.InferenceTaskResultDownloaded {
		return response.NewValidationErrorResponse("client_task_id", "Client task was not successful")
	}

	ext := "png"
	if task.TaskType == models.TaskTypeLLM {
		ext = "json"
	}
	filename := fmt.Sprintf("%d.%s", *in.Index, ext)

	appConfig := config.GetConfig()
	imageFile := filepath.Join(
		appConfig.DataDir.InferenceTasks,
		task.TaskIDCommitment,
		filename,
	)

	if _, err := os.Stat(imageFile); err != nil {
		return response.NewValidationErrorResponse("image_num", "Image not found")
	}

	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Disposition", "attachment; filename="+filename)
	c.Header("Content-Type", "application/octet-stream")
	c.File(imageFile)

	return nil
}
