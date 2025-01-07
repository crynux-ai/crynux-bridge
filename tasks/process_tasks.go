package tasks

import (
	"context"
	"crynux_bridge/blockchain"
	"crynux_bridge/blockchain/bindings"
	"crynux_bridge/config"
	"crynux_bridge/models"
	"crynux_bridge/utils"
	"database/sql"
	"errors"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	log "github.com/sirupsen/logrus"
)

func getChainTask(ctx context.Context, taskIDCommitmentBytes [32]byte) (*bindings.VSSTaskTaskInfo, error) {
	callCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	return blockchain.GetTaskByCommitment(callCtx, taskIDCommitmentBytes)
}

func reportTaskParamsUploaded(ctx context.Context, task *models.InferenceTask) error {
	taskIDCommitmentBytes, err := utils.HexStrToCommitment(task.TaskIDCommitment)
	if err != nil {
		return nil
	}

	txHash, err := func() (string, error) {
		callCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()
		return blockchain.ReportTaskParamsUploaded(callCtx, *taskIDCommitmentBytes)
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
		log.Errorf("ProcessTasks: %s reportTaskParamsUploaded failed: %s", task.TaskIDCommitment, errMsg)
		return errors.New(errMsg)
	}
	return nil
}

func reportTaskResultUploaded(ctx context.Context, task *models.InferenceTask) error {
	taskIDCommitmentBytes, err := utils.HexStrToCommitment(task.TaskIDCommitment)
	if err != nil {
		return nil
	}

	txHash, err := func() (string, error) {
		callCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()
		return blockchain.ReportTaskResultUploaded(callCtx, *taskIDCommitmentBytes)
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
		log.Errorf("ProcessTasks: %s reportTaskResultUploaded failed: %s", task.TaskIDCommitment, errMsg)
		return errors.New(errMsg)
	}
	return nil
}

func doWaitTaskStatus(ctx context.Context, taskIDCommitment string, status models.TaskStatus) error {
	for {
		task, err := models.GetTaskByIDCommitment(ctx, taskIDCommitment)
		if err != nil {
			return err
		}
		if task.Status == status {
			return nil
		}
		time.Sleep(time.Second)
	}
}

func waitTaskStatus(ctx context.Context, taskIDCommitment string, status models.TaskStatus) error {
	c := make(chan error, 1)
	go func() {
		c <- doWaitTaskStatus(ctx, taskIDCommitment, status)
	}()
	select {
	case err := <-c:
		return err
	case <-ctx.Done():
		log.Errorf("ProcessTasks: check task %s status %d timed out", taskIDCommitment, status)
		return ctx.Err()
	}
}

func syncTask(ctx context.Context, task *models.InferenceTask) (*bindings.VSSTaskTaskInfo, error) {
	taskIDCommitmentBytes, err := utils.HexStrToCommitment(task.TaskIDCommitment)
	if err != nil {
		return nil, err
	}

	chainTask, err := getChainTask(ctx, *taskIDCommitmentBytes)
	if err != nil {
		return nil, err
	}

	changed := false
	newTask := &models.InferenceTask{}
	chainTaskStatus := models.ChainTaskStatus(chainTask.Status)
	abortReason := models.TaskAbortReason(chainTask.AbortReason)
	taskError := models.TaskError(chainTask.Error)
	scoreReadyTimestamp := chainTask.ScoreReadyTimestamp.Int64()

	if scoreReadyTimestamp > 0 && !task.ScoreReadyTime.Valid {
		newTask.ScoreReadyTime = sql.NullTime{
			Time:  time.Unix(scoreReadyTimestamp, 0).UTC(),
			Valid: true,
		}
		changed = true
	}
	if abortReason != task.AbortReason {
		newTask.AbortReason = abortReason
		changed = true
	}
	if taskError != task.TaskError {
		newTask.TaskError = taskError
		changed = true
	}

	if chainTaskStatus == models.ChainTaskParametersUploaded {
		if task.Status != models.InferenceTaskParamsUploaded {
			newTask.Status = models.InferenceTaskParamsUploaded
			changed = true
		}
	} else if chainTaskStatus == models.ChainTaskValidated || chainTaskStatus == models.ChainTaskGroupValidated {
		if !task.ValidatedTime.Valid {
			newTask.ValidatedTime = sql.NullTime{
				Time:  time.Now().UTC(),
				Valid: true,
			}
			changed = true
		}
	} else if chainTaskStatus == models.ChainTaskEndAborted {
		if task.Status != models.InferenceTaskEndAborted {
			newTask.Status = models.InferenceTaskEndAborted
			changed = true
		}
		if !task.ValidatedTime.Valid {
			newTask.ValidatedTime = sql.NullTime{
				Time:  time.Now().UTC(),
				Valid: true,
			}
			changed = true
		}
	} else if chainTaskStatus == models.ChainTaskEndInvalidated {
		if task.Status != models.InferenceTaskEndInvalidated {
			newTask.Status = models.InferenceTaskEndInvalidated
			changed = true
		}
		if !task.ValidatedTime.Valid {
			newTask.ValidatedTime = sql.NullTime{
				Time:  time.Now().UTC(),
				Valid: true,
			}
			changed = true
		}
	} else if chainTaskStatus == models.ChainTaskEndGroupRefund {
		if task.Status != models.InferenceTaskEndGroupRefund {
			newTask.Status = models.InferenceTaskEndGroupRefund
			changed = true
		}
		if !task.ValidatedTime.Valid {
			newTask.ValidatedTime = sql.NullTime{
				Time:  time.Now().UTC(),
				Valid: true,
			}
			changed = true
		}
	}

	if changed {
		if err := task.Update(ctx, newTask); err != nil {
			return nil, err
		}
	}
	return chainTask, nil
}

func processOneTask(ctx context.Context, task *models.InferenceTask) error {
	// sync task from blockchain first
	_, err := syncTask(ctx, task)
	if err != nil {
		return err
	}

	// report task params is uploaded to blochchain
	if task.Status == models.InferenceTaskCreated {
		if err := reportTaskParamsUploaded(ctx, task); err != nil {
			return err
		}

		newTask := &models.InferenceTask{
			Status: models.InferenceTaskParamsUploaded,
			StartTime: sql.NullTime{
				Time:  time.Now().UTC(),
				Valid: true,
			},
		}

		if err := task.Update(ctx, newTask); err != nil {
			return err
		}
		log.Infof("ProcessTasks: report task %s params uploaded", task.TaskIDCommitment)
	}

	// wait task has been validated or end
	needResult := false
	for {
		chainTask, err := syncTask(ctx, task)
		if err != nil {
			return err
		}
		chainTaskStatus := models.ChainTaskStatus(chainTask.Status)
		needResult = (chainTaskStatus == models.ChainTaskValidated || chainTaskStatus == models.ChainTaskGroupValidated)
		if task.Status != models.InferenceTaskParamsUploaded || task.ValidatedTime.Valid {
			break
		}
		time.Sleep(time.Second)
	}

	// report task result is uploaded to blockchain
	if needResult {
		// wait task result is ready
		log.Infof("ProcessTasks: task %s is validated", task.TaskIDCommitment)
		err := func() error {
			timeCtx, cancel := context.WithTimeout(ctx, 5*time.Minute)
			defer cancel()
			return waitTaskStatus(timeCtx, task.TaskIDCommitment, models.InferenceTaskResultsReady)
		}()
		if err != nil {
			return err
		}
		log.Infof("ProcessTasks: task %s result is uploaded", task.TaskIDCommitment)
		// task result is uploaded
		if err := reportTaskResultUploaded(ctx, task); err != nil {
			return err
		}

		newTask := &models.InferenceTask{
			Status: models.InferenceTaskEndSuccess,
			ResultUploadedTime: sql.NullTime{
				Time:  time.Now().UTC(),
				Valid: true,
			},
		}

		if err := task.Update(ctx, newTask); err != nil {
			return err
		}
		log.Infof("ProcessTasks: report task %s result is uploaded", task.TaskIDCommitment)
	} else {
		log.Infof("ProcessTasks: task %s finished with status: %d", task.TaskIDCommitment, task.Status)
	}
	return nil
}

func ProcessTasks(ctx context.Context) {
	limit := 100
	lastID := uint(0)

	interval := 1

	for {
		tasks, err := func(ctx context.Context) ([]models.InferenceTask, error) {
			var tasks []models.InferenceTask

			dbCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
			defer cancel()
			err := config.GetDB().WithContext(dbCtx).Model(&models.InferenceTask{}).
				Where("status != ?", models.InferenceTaskEndAborted).
				Where("status != ?", models.InferenceTaskEndInvalidated).
				Where("status != ?", models.InferenceTaskEndSuccess).
				Where("id > ?", lastID).
				Find(&tasks).
				Order("id ASC").
				Limit(limit).
				Error
			if err != nil {
				return nil, err
			}
			return tasks, nil
		}(ctx)
		if err != nil {
			log.Errorf("ProcessTasks: cannot get unprocessed tasks: %v", err)
			time.Sleep(time.Duration(interval) * time.Second)
			continue
		}

		if len(tasks) > 0 {
			lastID = tasks[len(tasks)-1].ID

			for _, task := range tasks {
				go func(ctx context.Context, task models.InferenceTask) {
					log.Infof("ProcessTasks: start processing task %s", task.TaskIDCommitment)
					var ctx1 context.Context
					var cancel context.CancelFunc
					if !task.StartTime.Valid {
						ctx1, cancel = context.WithTimeout(ctx, 10*time.Minute)
					} else {
						deadline := task.StartTime.Time.Add(10 * time.Minute)
						ctx1, cancel = context.WithDeadline(ctx, deadline)
					}
					defer cancel()

					for {
						c := make(chan error, 1)
						go func() {
							c <- processOneTask(ctx1, &task)
						}()

						select {
						case err := <-c:
							if err != nil {
								log.Errorf("ProcessTasks: process task %s error %v, retry", task.TaskIDCommitment, err)
							} else {
								log.Infof("ProcessTasks: process task %s successfully", task.TaskIDCommitment)
								return
							}
						case <-ctx1.Done():
							err := ctx1.Err()
							log.Errorf("ProcessTasks: process task %s timeout %v, finish", task.TaskIDCommitment, err)
							// set task status to aborted to avoid processing it again
							if err == context.DeadlineExceeded {
								newTask := &models.InferenceTask{Status: models.InferenceTaskEndAborted}
								if err := task.Update(ctx, newTask); err != nil {
									log.Errorf("ProcessTasks: save task %s error %v", task.TaskIDCommitment, err)
								}
							}
							return
						}
					}
				}(ctx, task)
			}
		}

		time.Sleep(time.Second)
	}
}
