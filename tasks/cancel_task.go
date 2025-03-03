package tasks

import (
	"context"
	"crynux_bridge/blockchain"
	"crynux_bridge/config"
	"crynux_bridge/models"
	"crynux_bridge/utils"
	"errors"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	log "github.com/sirupsen/logrus"
)

func cancelTaskOnChain(ctx context.Context, task *models.InferenceTask) error {
	if len(task.TaskIDCommitment) == 0 {
		return nil
	}

	txHash, err := func() (string, error) {
		callCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
		defer cancel()
		return blockchain.CancelTask(callCtx, task)
	}()
	if err != nil {
		return err
	}

	receipt, err := func() (*types.Receipt, error) {
		callCtx, cancel := context.WithTimeout(ctx, 120*time.Second)
		defer cancel()
		return blockchain.WaitTxReceipt(callCtx, common.HexToHash(txHash))
	}()
	if err != nil {
		return err
	}

	if receipt.Status == 0 {
		errMsg, err := blockchain.GetErrorMessageFromReceipt(ctx, receipt)
		if err != nil {
			return err
		}
		log.Errorf("ProcessTasks: %d cancelTask failed: %s", task.ID, errMsg)
		return errors.New(errMsg)
	}
	return nil
}

func cancelTask(ctx context.Context, task *models.InferenceTask) error {
	log.Infof("CancelTasks: start to cancel task %d", task.ID)
	newTask := &models.InferenceTask{}
	if len(task.TaskIDCommitment) == 0 {
		newTask.Status = models.InferenceTaskEndAborted
		newTask.AbortReason = models.TaskAbortTimeout
		if err := task.Update(ctx, config.GetDB(), newTask); err != nil {
			log.Errorf("CancelTasks: cannot save task %d status: %v", task.ID, err)
			return err
		}
		log.Infof("CancelTasks: task %d canceled successfully", task.ID)
		return nil
	}
	taskIDCommitment, _ := utils.HexStrToBytes32(task.TaskIDCommitment)
	chainTask, err := blockchain.GetTaskByCommitment(ctx, *taskIDCommitment)
	if err != nil {
		log.Errorf("CancelTasks: cannot get task %d from chain: %v", task.ID, err)
		return err
	}
	if hexutil.Encode(chainTask.TaskIDCommitment[:]) != task.TaskIDCommitment {
		newTask.Status = models.InferenceTaskEndAborted
		newTask.AbortReason = models.TaskAbortTimeout
		if err := task.Update(ctx, config.GetDB(), newTask); err != nil {
			log.Errorf("CancelTasks: cannot save task %d status: %v", task.ID, err)
			return err
		}
		log.Infof("CancelTasks: task %d canceled successfully", task.ID)
		return nil
	}

	chainTaskStatus := models.ChainTaskStatus(chainTask.Status)
	if chainTaskStatus == models.ChainTaskEndSuccess || chainTaskStatus == models.ChainTaskEndGroupSuccess {
		newTask.Status = models.InferenceTaskEndSuccess
	} else if chainTaskStatus == models.ChainTaskEndGroupRefund {
		newTask.Status = models.InferenceTaskEndGroupRefund
	} else if chainTaskStatus == models.ChainTaskEndInvalidated {
		newTask.Status = models.InferenceTaskEndInvalidated
	} else if chainTaskStatus == models.ChainTaskEndAborted {
		newTask.Status = models.InferenceTaskEndAborted
		newTask.AbortReason = models.TaskAbortReason(chainTask.AbortReason)
	} else {
		if err := cancelTaskOnChain(ctx, task); err != nil {
			log.Errorf("CancelTasks: cannot cancel task %d on chain: %v", task.ID, err)
			return err
		}
		newTask.Status = models.InferenceTaskEndAborted
		newTask.AbortReason = models.TaskAbortTimeout
	}
	if err := task.Update(ctx, config.GetDB(), newTask); err != nil {
		log.Errorf("CancelTasks: cannot save task %d status: %v", task.ID, err)
		return err
	}
	log.Infof("CancelTasks: task %d canceled successfully", task.ID)
	return nil
}

func getTasksNeedCancel(ctx context.Context) ([]models.InferenceTask, error) {
	limit := 100
	offset := 0

	allTasks := []models.InferenceTask{}
	for {
		tasks, err := func(ctx context.Context, offset, limit int) ([]models.InferenceTask, error) {
			var tasks []models.InferenceTask

			dbCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
			defer cancel()
			err := config.GetDB().WithContext(dbCtx).Model(&models.InferenceTask{}).
				Where("status = ?", models.InferenceTaskNeedCancel).
				Order("id ASC").
				Limit(limit).
				Offset(offset).
				Find(&tasks).
				Error
			if err != nil {
				return nil, err
			}
			return tasks, nil
		}(ctx, offset, limit)
		if err != nil {
			return nil, err
		}
		allTasks = append(allTasks, tasks...)
		if len(tasks) < limit {
			break
		}
		offset += limit
	}

	return allTasks, nil
}

func CancelTasks(ctx context.Context) {
	for {
		tasks, err := getTasksNeedCancel(ctx)
		if err != nil {
			log.Errorf("CancelTasks: cannot get tasks need to cancel: %v", err)
			time.Sleep(2 * time.Second)
			continue
		}
		log.Infof("CancelTasks: %d tasks need to cancel", len(tasks))

		for _, task := range tasks {
			if err := cancelTask(ctx, &task); err != nil {
				log.Errorf("CancelTasks: cannot cancel task %d due to %v", task.ID, err)
				continue
			}
		}

		time.Sleep(time.Minute)
	}
}