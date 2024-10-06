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
	var allClientTasks []models.ClientTask

	offset := 0
	limit := 100

	for {
		var clientTasks []models.ClientTask
		if err := config.GetDB().Model(&models.ClientTask{}).Preload("InferenceTasks").Where("created_at >= ?", startTime).Order("id").Offset(offset).Limit(limit).Find(&clientTasks).Error; err != nil {
			return nil, response.NewExceptionResponse(err)
		}
		allClientTasks = append(allClientTasks, clientTasks...)
		if len(clientTasks) < limit {
			break
		}
		offset += limit
	}

	result := CountOutput{}
	result.TotalTaskCount = len(allClientTasks)

	if result.TotalTaskCount > 0 {
		totalTaskTime := time.Duration(0)
	
		for _, clientTask := range allClientTasks {
			var successCount, abortedCount int
			var successTaskTime time.Duration
			for _, task := range clientTask.InferenceTasks {
				if task.Status == models.InferenceTaskSuccess {
					successCount += 1
					t := task.UpdatedAt.Sub(task.CreatedAt)
					if successTaskTime == 0 || t < successTaskTime {
						successTaskTime = t
					}
				} else if task.Status == models.InferenceTaskAborted {
					abortedCount += 1
				}
			}

			if successCount > 0 {
				result.SuccessTaskCount += 1
				totalTaskTime += successTaskTime
			} else if abortedCount == len(clientTask.InferenceTasks) {
				result.AbortedTaskCount += 1
			} else {
				result.UnfinishedTaskCount += 1
			}
		}
	
		result.AvgTaskTime = int(totalTaskTime/time.Second) / result.TotalTaskCount
	}

	return &CountResponse{Data: &result}, nil
}
