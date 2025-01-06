package tasks_test

import (
	"crynux_bridge/models"
	"crynux_bridge/tasks"
	"crynux_bridge/tests"
	"math/big"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestUploadRightTask(t *testing.T) {

	err := tests.SyncToLatestBlock()
	assert.Equal(t, nil, err, "catchup error")

	uploadTaskChan := make(chan int)
	go tasks.StartUploadTaskParamsWithTerminateChannel(uploadTaskChan)

	sendTaskChan := make(chan int)
	go tasks.StartSendTaskOnChainWithTerminateChannel(sendTaskChan)

	getTaskCreationResultChan := make(chan int)
	go tasks.StartGetTaskCreationResultWithTerminateChannel(getTaskCreationResultChan)

	t.Cleanup(func() {
		tests.ClearDB()
	})

	t.Cleanup(func() {
		uploadTaskChan <- 1
		sendTaskChan <- 1
		getTaskCreationResultChan <- 1
	})

	addresses, privateKeys, err := tests.PrepareAccounts()
	assert.Nil(t, err, "error preparing accounts")

	err = tests.PrepareNetwork(addresses, privateKeys)
	assert.Nil(t, err, "error preparing the network")

	err = tests.PrepareTaskCreatorAccount(addresses[0], privateKeys[0])
	assert.Nil(t, err, "error preparing the task creator account")

	t.Cleanup(func() {
		err := tests.ClearNetwork(addresses, privateKeys)
		if err != nil {
			t.Error(err)
		}
	})

	for _, taskType := range tests.TaskTypes {
		task, err := tests.NewTask(taskType)
		assert.Nil(t, err, "error creating task")

		time.Sleep(40 * time.Second)
		task = tests.AssertTaskStatus(t, task.ID, models.InferenceTaskParamsUploaded)

		// Task must be finished before clearing the network
		err = tests.SuccessTaskOnChain(big.NewInt(int64(task.TaskID)), addresses, privateKeys)
		assert.Equal(t, nil, err, "error submitting result on chain")
	}
}
