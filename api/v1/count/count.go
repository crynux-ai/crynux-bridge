package count

import (
	"crynux_bridge/api/v1/response"
	"crynux_bridge/config"
	"crynux_bridge/models"
	"time"

	"github.com/gin-gonic/gin"
)

type CountInput struct {
	Duration int `query:"duration" description:"recent time duration in hours to count"`
}

type CountOutput struct {
	TotalTaskCount      int `json:"total_task_count" description:"total task count in the duration"`
	SuccessTaskCount    int `json:"success_task_count" description:"success task count in the duration"`
	AbortedTaskCount    int `json:"aborted_task_count" description:"aborted task count in the duration"`
	UnfinishedTaskCount int `json:"unfinished_task_count" description:"unfinished task count in the duration"`
	AvgTaskTime         int `json:"avg_task_time" description:"average successful task execution time, in seconds"`
}

type CountResponse struct {
	response.Response
	Data *CountOutput `json:"data"`
}

func CountTask(_ *gin.Context, input *CountInput) (*CountResponse, error) {
	duration := time.Duration(input.Duration) * time.Hour

	startTime := time.Now().Add(-duration)
	var allTasks []models.InferenceTask

	offset := 0
	limit := 100
	for {
		var tasks []models.InferenceTask
		if err := config.GetDB().Model(&models.InferenceTask{}).Where("created_at >= ?", startTime).Order("id").Offset(offset).Limit(limit).Find(&tasks).Error; err != nil {
			return nil, response.NewExceptionResponse(err)
		}
		allTasks = append(allTasks, tasks...)
		if len(tasks) < limit {
			break
		}
		offset += limit
	}

	result := CountOutput{}
	result.TotalTaskCount = len(allTasks)

	totalTaskTime := time.Duration(0)

	for _, task := range allTasks {
		if task.Status == models.InferenceTaskSuccess {
			result.SuccessTaskCount += 1
			totalTaskTime += task.UpdatedAt.Sub(task.CreatedAt)
		} else if task.Status == models.InferenceTaskAborted {
			result.AbortedTaskCount += 1
		} else {
			result.UnfinishedTaskCount += 1
		}
	}

	result.AvgTaskTime = int(totalTaskTime/time.Second) / result.TotalTaskCount

	return &CountResponse{Data: &result}, nil
}
