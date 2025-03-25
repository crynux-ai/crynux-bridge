package openrouter

import (
	"crynux_bridge/api/v1/inference_tasks"
	"crynux_bridge/api/v1/openrouter/structs"
	"crynux_bridge/api/v1/response"
	"crynux_bridge/config"
	"crynux_bridge/models"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

func ProcessGPTTask(c *gin.Context, in *inference_tasks.TaskInput) (*structs.GPTTaskResponse, *models.InferenceTask, error) {
	/* 1. Create GPT task by function CreateTask */
	taskResponse, err := inference_tasks.CreateTask(c, in)
	if err != nil {
		return nil, nil, err
	}

	/* 2. Get tasks, wait until they are finished and the taks result is downloaded  */
	tasks := taskResponse.Data.InferenceTasks
	if len(tasks) == 0 {
		err := errors.New("no tasks created")
		return nil, nil, response.NewExceptionResponse(err)
	}
	// one goroutine for each task to monitor its status
	var waitGroup sync.WaitGroup
	resultChan := make(chan struct {
		task   *models.InferenceTask
		status models.TaskStatus
		err    error
	}, len(tasks))

	for _, t := range tasks {
		waitGroup.Add(1)

		go func(t *models.InferenceTask) {
			defer waitGroup.Done()

			status, err := waitForTaskFinish(c, t.Client.ClientId, t.ClientTaskID)

			// Store the result in the resultChan
			resultChan <- struct {
				task   *models.InferenceTask
				status models.TaskStatus
				err    error
			}{t, status, err}
		}(&t)
	}

	// Wait for all goroutines to finish
	go func() {
		waitGroup.Wait()
		close(resultChan)
	}()

	// Get the result from the resultChan
	var resultDownloadedTask *models.InferenceTask = nil
	for result := range resultChan {
		if result.err != nil {
			return nil, nil, result.err
		}
		if result.status == models.InferenceTaskResultDownloaded {
			resultDownloadedTask = result.task
			break
		}
	}
	if resultDownloadedTask == nil {
		err := errors.New("all tasks end without result downloaded")
		return nil, nil, response.NewExceptionResponse(err)
	}

	/* 3. Read task result and return */
	results, err := readGPTTaskResults(resultDownloadedTask)
	if err != nil {
		return nil, nil, response.NewExceptionResponse(err)
	}

	// without stream, the response is a single object, i.e. results[0]
	gptTaskResponse := results[0]

	return &gptTaskResponse, resultDownloadedTask, nil
}

func waitForTaskFinish(c *gin.Context, clientID string, clientTaskID uint) (models.TaskStatus, error) {
	for {
		// 1. get task by id
		getTaskInput := &inference_tasks.GetTaskInput{
			ClientID:     clientID,
			ClientTaskID: clientTaskID,
		}
		getTaskResponse, err := inference_tasks.GetTaskById(c, getTaskInput)
		if err != nil {
			return models.InferenceTaskEndAborted, err
		}

		// 2. check task status
		taskStatus := getTaskResponse.Data.Status
		// task end without result downloaded
		if taskStatus == models.InferenceTaskEndInvalidated || taskStatus == models.InferenceTaskEndGroupRefund || taskStatus == models.InferenceTaskEndAborted {
			return taskStatus, nil
		}
		// task end with result downloaded
		if taskStatus == models.InferenceTaskResultDownloaded {
			return taskStatus, nil
		}
		// task not end, then sleep and continue looping
		time.Sleep(time.Second)
	}
}

func readGPTTaskResults(task *models.InferenceTask) ([]structs.GPTTaskResponse, error) {
	if task.TaskType != models.TaskTypeLLM {
		err := errors.New("unsupported task type")
		return nil, err
	}

	appConfig := config.GetConfig()
	taskFolder := path.Join(appConfig.DataDir.InferenceTasks, task.TaskIDCommitment)
	ext := "json"

	// 1. create a results array
	results := make([]structs.GPTTaskResponse, task.TaskSize)
	var wg sync.WaitGroup
	errCh := make(chan error, int(task.TaskSize))

	// 2. read all result files in parallel
	for i := uint64(0); i < task.TaskSize; i++ {
		wg.Add(1)
		go func(index uint64) {
			defer wg.Done()

			filename := path.Join(taskFolder, fmt.Sprintf("%d.%s", index, ext))
			data, err := os.ReadFile(filename)
			if err != nil {
				errCh <- fmt.Errorf("readGPTTaskResults: failed to read file %s, error: %w", filename, err)
				return
			}

			var result structs.GPTTaskResponse

			if err := json.Unmarshal(data, &result); err != nil {
				errCh <- fmt.Errorf("failed to unmarshal json file %s, error: %w", filename, err)
				return
			}

			// 3. save result to results array, according to the index
			results[index] = result
		}(i)
	}

	// 4. wait for all goroutines to finish
	go func() { wg.Wait(); close(errCh) }()
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	return results, nil
}
