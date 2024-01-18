package tasks_test

import (
	"crynux_bridge/models"
	"crynux_bridge/tasks"
	"crynux_bridge/tests"
	"math/big"
	"strconv"
	"testing"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestTaskSuccessResult(t *testing.T) {

	err := tests.SyncToLatestBlock()
	assert.Equal(t, nil, err, "catchup error")

	sendTaskChan := make(chan int)
	go tasks.StartSendTaskOnChainWithTerminateChannel(sendTaskChan)

	getTaskCreationResultChan := make(chan int)
	go tasks.StartGetTaskCreationResultWithTerminateChannel(getTaskCreationResultChan)

	syncBlockChan := make(chan int)
	go tasks.StartSyncBlockWithTerminateChannel(syncBlockChan)

	addresses, privateKeys, err := tests.PrepareAccounts()
	assert.Equal(t, nil, err, "error preparing accounts")

	err = tests.PrepareNetwork(addresses, privateKeys)
	assert.Equal(t, nil, err, "error preparing network")

	err = tests.PrepareTaskCreatorAccount(addresses[0], privateKeys[0])
	assert.Equal(t, nil, err, "error preparing task creator account")

	task, err := tests.NewTask()
	assert.Equal(t, nil, err, "error creating task")
	log.Debugln("Task created in db with pk: " + strconv.FormatUint(uint64(task.ID), 10))

	time.Sleep(20 * time.Second)

	task = tests.AssertTaskStatus(t, task.ID, models.InferenceTaskBlockchainConfirmed)

	assert.NotZero(t, task.TaskId, "TaskId on chain is zero")

	log.Debugln("Task created on chain")
	log.Debugln("Now lets submit the task results from the nodes")

	err = tests.SuccessTaskOnChain(big.NewInt(int64(task.TaskId)), addresses, privateKeys)
	assert.Equal(t, nil, err, "error submitting result on chain")

	log.Debugln("Task results submitted")

	time.Sleep(40 * time.Second)
	tests.AssertTaskStatus(t, task.ID, models.InferenceTaskPendingResult)

	t.Cleanup(func() {
		sendTaskChan <- 1
		getTaskCreationResultChan <- 1
		syncBlockChan <- 1
		err := tests.ClearNetwork(addresses, privateKeys)
		assert.Equal(t, nil, err, "error clearing blockchain network")
		tests.ClearDB()
	})
}

func TestTaskAbortedResult(t *testing.T) {

	err := tests.SyncToLatestBlock()
	assert.Equal(t, nil, err, "catchup error")

	sendTaskChan := make(chan int)
	go tasks.StartSendTaskOnChainWithTerminateChannel(sendTaskChan)

	getTaskCreationResultChan := make(chan int)
	go tasks.StartGetTaskCreationResultWithTerminateChannel(getTaskCreationResultChan)

	syncBlockChan := make(chan int)
	go tasks.StartSyncBlockWithTerminateChannel(syncBlockChan)

	addresses, privateKeys, err := tests.PrepareAccounts()
	assert.Equal(t, nil, err, "error preparing accounts")

	err = tests.PrepareNetwork(addresses, privateKeys)
	assert.Equal(t, nil, err, "error preparing network")

	err = tests.PrepareTaskCreatorAccount(addresses[0], privateKeys[0])
	assert.Equal(t, nil, err, "error preparing task creator account")

	task, err := tests.NewTask()
	assert.Equal(t, nil, err, "error creating task")

	time.Sleep(20 * time.Second)
	task = tests.AssertTaskStatus(t, task.ID, models.InferenceTaskBlockchainConfirmed)

	err = tests.AbortTaskOnChain(big.NewInt(int64(task.TaskId)), addresses, privateKeys)
	assert.Equal(t, nil, err, "error submitting result on chain")

	time.Sleep(40 * time.Second)
	tests.AssertTaskStatus(t, task.ID, models.InferenceTaskAborted)

	t.Cleanup(func() {
		sendTaskChan <- 1
		getTaskCreationResultChan <- 1
		syncBlockChan <- 1
		err := tests.ClearNetwork(addresses, privateKeys)
		assert.Equal(t, nil, err, "error clearing blockchain network")
		tests.ClearDB()
	})
}
