package tasks

import (
	"archive/zip"
	"context"
	"crynux_bridge/config"
	"crynux_bridge/models"
	"crypto/rand"
	"errors"
	"os"
	"path/filepath"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	log "github.com/sirupsen/logrus"
)

func ProcessSDFTTasks(ctx context.Context) {
	type result struct {
		ClientTaskID uint `json:"client_task_id"`
		ID           uint `json:"id"`
	}

	lastID := uint(0)
	limit := 100

	for {
		tasks, err := func(ctx context.Context) ([]*models.InferenceTask, error) {
			var results []result

			dbCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
			defer cancel()
			err := config.GetDB().WithContext(dbCtx).Model(&models.InferenceTask{}).
				Select("client_task_id,max(id) as id").
				Where("task_type = ?", models.TaskTypeSDFTLora).
				Where("id > ?", lastID).
				Group("client_task_id").
				Order("id ASC").
				Limit(limit).
				Find(&results).
				Error
			if err != nil {
				return nil, err
			}

			var tasks []*models.InferenceTask
			var ids []uint
			for _, result := range results {
				ids = append(ids, result.ID)
			}

			if len(ids) > 0 {
				dbCtx1, cancel1 := context.WithTimeout(ctx, 10*time.Second)
				defer cancel1()
				err := config.GetDB().WithContext(dbCtx1).Model(&models.InferenceTask{}).
					Where("id IN (?)", ids).
					Order("id ASC").
					Find(&tasks).
					Error
				if err != nil {
					return nil, err
				}
			}

			return tasks, nil
		}(ctx)
		if err != nil {
			log.Errorf("GetSDFTTasks: cannot get tasks: %v", err)
			time.Sleep(5 * time.Second)
			continue
		}

		if len(tasks) == 0 {
			time.Sleep(5 * time.Second)
			continue
		}

		for _, task := range tasks {
			go func(ctx context.Context, task *models.InferenceTask) {
				duration := time.Duration(task.Timeout) * time.Second
				ctx, cancel := context.WithTimeout(ctx, duration)
				defer cancel()
				err := processSDFTTaskWithRetry(ctx, task)
				if err != nil {
					log.Errorf("ProcessSDFTTasks: cannot process task %d: %v", task.ID, err)
				}
			}(ctx, task)
		}

		lastID = tasks[len(tasks)-1].ID
	}
}

func processSDFTTaskWithRetry(ctx context.Context, task *models.InferenceTask) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			err := processSDFTTask(ctx, task)
			if err != nil {
				time.Sleep(3 * time.Second)
			} else {
				return nil
			}
		}
	}	
}

func processSDFTTask(ctx context.Context, task *models.InferenceTask) error {
	taskGroup, err := models.WaitTaskGroup(ctx, config.GetDB(), task)
	if err != nil {
		return err
	}

	resultDownloadedTask, err := models.WaitResultTask(ctx, config.GetDB(), taskGroup)

	if err == models.ErrTaskEndWithoutResult {
		taskFailedCount, err := models.GetSDFTTaskFailedCount(ctx, config.GetDB(), task.ClientTaskID)
		if err != nil {
			return err
		}
		if taskFailedCount <= 3 {
			return processFailedSDFTTask(ctx, task)
		}
		return nil
	}
	if err != nil {
		return err
	}

	if resultDownloadedTask != nil {
		return processResultDownloadedSDFTTask(ctx, resultDownloadedTask)
	}

	return nil
}

func processResultDownloadedSDFTTask(ctx context.Context, task *models.InferenceTask) error {
	appConfig := config.GetConfig()

	resultFilePath := filepath.Join(appConfig.DataDir.InferenceTasks, task.TaskIDCommitment, "result.zip")
	if _, err := os.Stat(resultFilePath); !os.IsNotExist(err) {
		return nil
	}

	checkpointFilePath := filepath.Join(appConfig.DataDir.InferenceTasks, task.TaskIDCommitment, "checkpoint.zip")
	if _, err := os.Stat(checkpointFilePath); os.IsNotExist(err) {
		log.Errorf("processSDFTTasks: checkpoint file of task %s not found", task.TaskIDCommitment)
		return errors.New("checkpoint file not found")
	}

	// unzip checkpoint file and check if FINISH file exists
	finished, err := func() (bool, error) {
		zipFile, err := zip.OpenReader(checkpointFilePath)
		if err != nil {
			log.Errorf("processSDFTTasks: cannot open checkpoint file of task %s: %v", task.TaskIDCommitment, err)
			return false, err
		}
		defer zipFile.Close()

		for _, file := range zipFile.File {
			if file.Name == "FINISH" {
				return true, nil
			}
		}
		return false, nil
	}()

	if err != nil {
		log.Errorf("processSDFTTasks: cannot check checkpoint file of task %s: %v", task.TaskIDCommitment, err)
		return err
	}

	if finished {
		log.Infof("processSDFTTasks: task %s is finished", task.TaskIDCommitment)
		// rename the checkpoint file to result.zip
		return os.Rename(checkpointFilePath, filepath.Join(appConfig.DataDir.InferenceTasks, task.TaskIDCommitment, "result.zip"))
	} else {
		// sd ft task is not finished, create a new task with the same client task id and task args, except the checkpoint file
		newTaskArgs, err := models.ChangeSDFTTaskArgsCheckpoint(task.TaskArgs, checkpointFilePath)
		if err != nil {
			log.Errorf("processSDFTTasks: cannot change task args of task %s: %v", task.TaskIDCommitment, err)
			return err
		}
		taskIDBytes := make([]byte, 32)
		rand.Read(taskIDBytes)
		newTaskID := hexutil.Encode(taskIDBytes)

		newTask := &models.InferenceTask{
			ClientID:        task.ClientID,
			ClientTaskID:    task.ClientTaskID,
			TaskArgs:        newTaskArgs,
			TaskType:        task.TaskType,
			TaskModelIDs:    task.TaskModelIDs,
			TaskVersion:     task.TaskVersion,
			TaskFee:         task.TaskFee,
			MinVram:         task.MinVram,
			RequiredGPU:     task.RequiredGPU,
			RequiredGPUVram: task.RequiredGPUVram,
			TaskSize:        task.TaskSize,
			TaskID:          newTaskID,
			Timeout:         task.Timeout,
		}

		err = newTask.Save(ctx, config.GetDB())
		if err != nil {
			log.Errorf("processSDFTTasks: cannot save new task %s: %v", newTaskID, err)
			return err
		}

		return nil
	}
}

func processFailedSDFTTask(ctx context.Context, task *models.InferenceTask) error {
	taskIDBytes := make([]byte, 32)
	rand.Read(taskIDBytes)
	newTaskID := hexutil.Encode(taskIDBytes)

	newTask := &models.InferenceTask{
		ClientID:        task.ClientID,
		ClientTaskID:    task.ClientTaskID,
		TaskArgs:        task.TaskArgs,
		TaskType:        task.TaskType,
		TaskModelIDs:    task.TaskModelIDs,
		TaskVersion:     task.TaskVersion,
		TaskFee:         task.TaskFee,
		MinVram:         task.MinVram,
		RequiredGPU:     task.RequiredGPU,
		RequiredGPUVram: task.RequiredGPUVram,
		TaskSize:        task.TaskSize,
		TaskID:          newTaskID,
		Timeout:         task.Timeout,
	}

	err := newTask.Save(ctx, config.GetDB())
	if err != nil {
		return err
	}

	return nil
}
