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

func TestNotEnoughNodes(t *testing.T) {

	sendTaskChan := make(chan int)
	go tasks.StartSendTaskOnChainWithTerminateChannel(sendTaskChan)

	getTaskCreationResultChan := make(chan int)
	go tasks.StartGetTaskCreationResultWithTerminateChannel(getTaskCreationResultChan)

	t.Cleanup(func() {
		sendTaskChan <- 1
		getTaskCreationResultChan <- 1
		tests.ClearDB()
	})

	addresses, privateKeys, err := tests.PrepareAccounts()
	assert.Equal(t, nil, err, "error preparing accounts")

	err = tests.PrepareTaskCreatorAccount(addresses[0], privateKeys[0])
	assert.Equal(t, nil, err, "error preparing task creator account")

	for _, taskType := range tests.TaskTypes {
		task, err := tests.NewTask(taskType)
		assert.Equal(t, nil, err, "error creating task")

		time.Sleep(20 * time.Second)
		tests.AssertTaskStatus(t, task.ID, models.InferenceTaskAborted)
	}
}

func TestNotEnoughToken(t *testing.T) {

	sendTaskChan := make(chan int)
	go tasks.StartSendTaskOnChainWithTerminateChannel(sendTaskChan)

	getTaskCreationResultChan := make(chan int)
	go tasks.StartGetTaskCreationResultWithTerminateChannel(getTaskCreationResultChan)

	t.Cleanup(func() {
		sendTaskChan <- 1
		getTaskCreationResultChan <- 1
		tests.ClearDB()
	})

	addresses, privateKeys, err := tests.PrepareAccounts()
	assert.Equal(t, nil, err, "error preparing accounts")

	err = tests.PrepareNetwork(addresses, privateKeys)
	assert.Equal(t, nil, err, "error preparing network nodes")

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
		tests.AssertTaskStatus(t, task.ID, models.InferenceTaskAborted)
	}
}

func TestSuccessCreation(t *testing.T) {

	sendTaskChan := make(chan int)
	go tasks.StartSendTaskOnChainWithTerminateChannel(sendTaskChan)

	getTaskCreationResultChan := make(chan int)
	go tasks.StartGetTaskCreationResultWithTerminateChannel(getTaskCreationResultChan)

	t.Cleanup(func() {
		sendTaskChan <- 1
		getTaskCreationResultChan <- 1
		tests.ClearDB()
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
		task = tests.AssertTaskStatus(t, task.ID, models.InferenceTaskBlockchainConfirmed)

		// The results must be submitted in order to free the 3 nodes from the network
		err = tests.SuccessTaskOnChain(big.NewInt(int64(task.TaskId)), addresses, privateKeys)
		assert.Equal(t, nil, err, "error submitting result on chain")
	}
}
