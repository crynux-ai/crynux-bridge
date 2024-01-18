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
	TaskID   uint   `path:"task_id" description:"Task id" validate:"required"`
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

	task := &models.InferenceTask{
		ClientID: client.ID,
	}

	if err := config.GetDB().Where(task).Omit("Pose.DataURL").First(task, in.TaskID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.NewValidationErrorResponse("task_id", "Task not found")
		} else {
			return response.NewExceptionResponse(err)
		}
	}

	appConfig := config.GetConfig()
	imageFile := filepath.Join(
		appConfig.DataDir.InferenceTasks,
		strconv.FormatUint(uint64(task.ID), 10),
		in.ImageNum+".png",
	)

	if _, err := os.Stat(imageFile); err != nil {
		return response.NewValidationErrorResponse("image_num", "Image not found")
	}

	ctx.Header("Content-Description", "File Transfer")
	ctx.Header("Content-Transfer-Encoding", "binary")
	ctx.Header("Content-Disposition", "attachment; filename="+in.ImageNum+".png")
	ctx.Header("Content-Type", "application/octet-stream")
	ctx.File(imageFile)

	return nil
}
