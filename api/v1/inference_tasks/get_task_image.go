package inference_tasks

import (
	"crynux_bridge/api/v1/response"
	"crynux_bridge/config"
	"crynux_bridge/models"
	"errors"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GetTaskImageInput struct {
	ClientID string `path:"client_id" description:"Client id" validate:"required"`
	ClientTaskID   uint   `path:"client_task_id" description:"Client task id" validate:"required"`
	ImageNum string `path:"image_num" description:"Image number" validate:"required"`
}

func GetTaskImage(ctx *gin.Context, in *GetTaskImageInput) error {

	client := &models.Client{ClientId: in.ClientID}

	if err := config.GetDB().Where(client).First(client).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.NewValidationErrorResponse("client_id", "Client not found")
		} else {
			return response.NewExceptionResponse(err)
		}
	}

	clientTask := &models.ClientTask{
		RootModel: models.RootModel{ID: in.ClientTaskID},
		ClientID: client.ID,
	}

	if err := config.GetDB().Model(&clientTask).Where(&clientTask).Preload("InferenceTasks").First(&clientTask).Error; err != nil {
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
		if t.Status == models.InferenceTaskSuccess  {
			if task.Status != models.InferenceTaskSuccess {
				task = t
			} else if task.UpdatedAt.Sub(t.UpdatedAt) > 0 {
				task = t
			}
		}
	}

	var fileExt string
	if task.TaskType == models.TaskTypeSD {
		fileExt = ".png"
	} else {
		fileExt = ".json"
	}

	appConfig := config.GetConfig()
	imageFile := filepath.Join(
		appConfig.DataDir.InferenceTasks,
		strconv.FormatUint(uint64(task.ID), 10),
		in.ImageNum+fileExt,
	)

	if _, err := os.Stat(imageFile); err != nil {
		return response.NewValidationErrorResponse("image_num", "Image not found")
	}

	ctx.Header("Content-Description", "File Transfer")
	ctx.Header("Content-Transfer-Encoding", "binary")
	ctx.Header("Content-Disposition", "attachment; filename="+in.ImageNum+fileExt)
	ctx.Header("Content-Type", "application/octet-stream")
	ctx.File(imageFile)

	return nil
}
