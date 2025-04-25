package openrouter

import (
	"context"
	"crynux_bridge/api/v1/inference_tasks"
	"crynux_bridge/api/v1/openrouter/structs"
	"crynux_bridge/api/v1/response"
	"crynux_bridge/config"
	"crynux_bridge/models"
	"crynux_bridge/utils"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"gorm.io/gorm"
)

func ProcessGPTTask(ctx context.Context, db *gorm.DB, in *inference_tasks.TaskInput) (*structs.GPTTaskResponse, *models.InferenceTask, error) {
	/* 1. Create GPT task by function CreateTask */
	taskResponse, err := inference_tasks.DoCreateTask(ctx, in)
	if err != nil {
		return nil, nil, err
	}

	/* 2. Get tasks, wait until they are finished and the taks result is downloaded  */
	tasks := taskResponse.Data.InferenceTasks
	if len(tasks) == 0 {
		err := errors.New("no task created")
		return nil, nil, response.NewExceptionResponse(err)
	}
	taskGroups, err := waitAllTaskGroup(ctx, db, tasks)
	if err != nil {
		return nil, nil, response.NewExceptionResponse(err)
	}
	resultDownloadedTask, err := waitResultTask(ctx, db, taskGroups)
	if err != nil {
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

func waitTaskGroup(ctx context.Context, db *gorm.DB, task *models.InferenceTask) ([]models.InferenceTask, error) {
	for {
		err := task.Sync(ctx, db)
		if err != nil {
			return nil, err
		}
		if len(task.VRFNumber) > 0 {
			break
		}
		time.Sleep(time.Second)
	}
	vrfNumber, _ := hexutil.Decode(task.VRFNumber)
	if utils.VrfNeedValidation(vrfNumber) {
		taskGroup, err := models.GetTaskGroup(ctx, db, task.TaskID)
		if err != nil {
			return nil, err
		}
		return taskGroup, nil
	} else {
		return []models.InferenceTask{*task}, nil
	}
}

func waitAllTaskGroup(ctx context.Context, db *gorm.DB, tasks []models.InferenceTask) ([]models.InferenceTask, error) {
	var waitGroup sync.WaitGroup
	taskGroupChan := make(chan []models.InferenceTask, len(tasks))
	for _, task := range tasks {
		waitGroup.Add(1)
		go func(task models.InferenceTask) {
			defer waitGroup.Done()
			taskGroup, err := waitTaskGroup(ctx, db, &task)
			if err != nil {
				return
			}
			taskGroupChan <- taskGroup
		}(task)
	}
	go func() {
		waitGroup.Wait()
		close(taskGroupChan)
	}()

	taskGroups := make([]models.InferenceTask, 0)
	for taskGroup := range taskGroupChan {
		taskGroups = append(taskGroups, taskGroup...)
	}
	return taskGroups, nil
}

func waitForTaskFinish(ctx context.Context, db *gorm.DB, task *models.InferenceTask) (models.TaskStatus, error) {
	for {
		// 1. get task by id
		err := task.Sync(ctx, db)
		if err != nil {
			return task.Status, err
		}

		// 2. check task status
		taskStatus := task.Status
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

func waitResultTask(ctx context.Context, db *gorm.DB, tasks []models.InferenceTask) (*models.InferenceTask, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Minute)
	defer cancel()

	resultChan := make(chan *models.InferenceTask, len(tasks))
	errChan := make(chan error, len(tasks))
	doneChan := make(chan struct{})

	var wg sync.WaitGroup
	wg.Add(len(tasks))

	for _, t := range tasks {
		task := t
		go func(ctx context.Context, t *models.InferenceTask) {
			defer wg.Done()

			select {
			case <-doneChan:
				return
			default:
				status, err := waitForTaskFinish(ctx, db, t)
				if err != nil {
					select {
					case errChan <- err:
					case <-doneChan:
					}
					return
				}
				if status == models.InferenceTaskResultDownloaded {
					select {
					case resultChan <- t:
						close(doneChan)
					case <-doneChan:
					}
					return
				}
				select {
				case errChan <- nil:
				case <-doneChan:
				}
			}
		}(ctx, &task)
	}

	var errs []error
	for i := 0; i < len(tasks); i++ {
		select {
		case err := <-errChan:
			if err != nil {
				errs = append(errs, err)
			}
		case result := <-resultChan:
			go func() {
				wg.Wait()
				close(errChan)
				close(resultChan)
			}()
			return result, nil
		case <-ctx.Done():
			close(doneChan)
			wg.Wait()
			return nil, fmt.Errorf("timeout after 3 minutes: %w", ctx.Err())
		}
	}

	if len(errs) > 0 {
		return nil, fmt.Errorf("all tasks failed: %v", errs)
	}
	return nil, errors.New("all tasks end without result downloaded")
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
