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

	uploadTaskChan := make(chan int)
	go tasks.StartUploadTaskParamsWithTerminateChannel(uploadTaskChan)

	sendTaskChan := make(chan int)
	go tasks.StartSendTaskOnChainWithTerminateChannel(sendTaskChan)

	getTaskCreationResultChan := make(chan int)
	go tasks.StartGetTaskCreationResultWithTerminateChannel(getTaskCreationResultChan)

	syncBlockChan := make(chan int)
	go tasks.StartSyncBlockWithTerminateChannel(syncBlockChan)

	t.Cleanup(func() {
		tests.ClearDB()
	})

	t.Cleanup(func() {
		uploadTaskChan <- 1
		sendTaskChan <- 1
		getTaskCreationResultChan <- 1
		syncBlockChan <- 1
	})

	addresses, privateKeys, err := tests.PrepareAccounts()
	assert.Equal(t, nil, err, "error preparing accounts")

	err = tests.PrepareNetwork(addresses, privateKeys)
	assert.Equal(t, nil, err, "error preparing network")

	err = tests.PrepareTaskCreatorAccount(addresses[0], privateKeys[0])
	assert.Equal(t, nil, err, "error preparing task creator account")

	t.Cleanup(func() {
		err := tests.ClearNetwork(addresses, privateKeys)
		if err != nil {
			t.Error(err)
		}
	})

	for _, taskType := range tests.TaskTypes {
		task, err := tests.NewTask(taskType)
		assert.Equal(t, nil, err, "error creating task")
		log.Debugln("Task created in db with pk: " + strconv.FormatUint(uint64(task.ID), 10))
	
		time.Sleep(20 * time.Second)
	
		task = tests.AssertTaskStatus(t, task.ID, models.InferenceTaskParamsUploaded)
	
		assert.NotZero(t, task.TaskId, "TaskId on chain is zero")
	
		log.Debugln("Task created on chain")
		log.Debugln("Now lets submit the task results from the nodes")
	
		err = tests.SuccessTaskOnChain(big.NewInt(int64(task.TaskId)), addresses, privateKeys)
		assert.Equal(t, nil, err, "error submitting result on chain")
	
		log.Debugln("Task results submitted")
	
		time.Sleep(40 * time.Second)
		tests.AssertTaskStatus(t, task.ID, models.InferenceTaskPendingResult)
	}
}

func TestTaskAbortedResult(t *testing.T) {

	err := tests.SyncToLatestBlock()
	assert.Equal(t, nil, err, "catchup error")

	uploadTaskChan := make(chan int)
	go tasks.StartUploadTaskParamsWithTerminateChannel(uploadTaskChan)

	sendTaskChan := make(chan int)
	go tasks.StartSendTaskOnChainWithTerminateChannel(sendTaskChan)

	getTaskCreationResultChan := make(chan int)
	go tasks.StartGetTaskCreationResultWithTerminateChannel(getTaskCreationResultChan)

	syncBlockChan := make(chan int)
	go tasks.StartSyncBlockWithTerminateChannel(syncBlockChan)

	t.Cleanup(func() {
		tests.ClearDB()
	})

	t.Cleanup(func() {
		uploadTaskChan <- 1
		sendTaskChan <- 1
		getTaskCreationResultChan <- 1
		syncBlockChan <- 1
	})

	addresses, privateKeys, err := tests.PrepareAccounts()
	assert.Equal(t, nil, err, "error preparing accounts")

	err = tests.PrepareNetwork(addresses, privateKeys)
	assert.Equal(t, nil, err, "error preparing network")

	err = tests.PrepareTaskCreatorAccount(addresses[0], privateKeys[0])
	assert.Equal(t, nil, err, "error preparing task creator account")

	t.Cleanup(func() {
		err := tests.ClearNetwork(addresses, privateKeys)
		if err != nil {
			t.Error(err)
		}
	})

	for _, taskType := range tests.TaskTypes {
		task, err := tests.NewTask(taskType)
		assert.Equal(t, nil, err, "error creating task")
	
		time.Sleep(20 * time.Second)
		task = tests.AssertTaskStatus(t, task.ID, models.InferenceTaskParamsUploaded)
	
		err = tests.AbortTaskOnChain(big.NewInt(int64(task.TaskId)), addresses, privateKeys)
		assert.Equal(t, nil, err, "error submitting result on chain")
	
		time.Sleep(40 * time.Second)
		tests.AssertTaskStatus(t, task.ID, models.InferenceTaskAborted)
	}
}
