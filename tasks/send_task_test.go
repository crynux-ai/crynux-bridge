package tasks_test

import (
	"github.com/stretchr/testify/assert"
	"ig_server/models"
	"ig_server/tasks"
	"ig_server/tests"
	"math/big"
	"testing"
	"time"
)

func TestNotEnoughNodes(t *testing.T) {

	sendTaskChan := make(chan int)
	go tasks.StartSendTaskOnChainWithTerminateChannel(sendTaskChan)

	getTaskCreationResultChan := make(chan int)
	go tasks.StartGetTaskCreationResultWithTerminateChannel(getTaskCreationResultChan)

	addresses, privateKeys, err := tests.PrepareAccounts()
	assert.Equal(t, nil, err, "error preparing accounts")

	err = tests.PrepareTaskCreatorAccount(addresses[0], privateKeys[0])
	assert.Equal(t, nil, err, "error preparing task creator account")

	task, err := tests.NewTask()
	assert.Equal(t, nil, err, "error creating task")

	time.Sleep(20 * time.Second)
	tests.AssertTaskStatus(t, task.ID, models.InferenceTaskAborted)

	t.Cleanup(func() {
		sendTaskChan <- 1
		getTaskCreationResultChan <- 1
		tests.ClearDB()
	})
}

func TestNotEnoughToken(t *testing.T) {

	sendTaskChan := make(chan int)
	go tasks.StartSendTaskOnChainWithTerminateChannel(sendTaskChan)

	getTaskCreationResultChan := make(chan int)
	go tasks.StartGetTaskCreationResultWithTerminateChannel(getTaskCreationResultChan)

	addresses, privateKeys, err := tests.PrepareAccounts()
	assert.Equal(t, nil, err, "error preparing accounts")

	err = tests.PrepareNetwork(addresses, privateKeys)
	assert.Equal(t, nil, err, "error preparing network nodes")

	task, err := tests.NewTask()
	assert.Equal(t, nil, err, "error creating task")

	time.Sleep(20 * time.Second)
	tests.AssertTaskStatus(t, task.ID, models.InferenceTaskAborted)

	t.Cleanup(func() {
		sendTaskChan <- 1
		getTaskCreationResultChan <- 1
		err := tests.ClearNetwork(addresses, privateKeys)
		assert.Equal(t, nil, err, "error clearing blockchain network")
		tests.ClearDB()
	})

}

func TestSuccessCreation(t *testing.T) {

	sendTaskChan := make(chan int)
	go tasks.StartSendTaskOnChainWithTerminateChannel(sendTaskChan)

	getTaskCreationResultChan := make(chan int)
	go tasks.StartGetTaskCreationResultWithTerminateChannel(getTaskCreationResultChan)

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

	// The results must be submitted in order to free the 3 nodes from the network
	err = tests.SuccessTaskOnChain(big.NewInt(int64(task.TaskId)), addresses, privateKeys)
	assert.Equal(t, nil, err, "error submitting result on chain")

	t.Cleanup(func() {
		sendTaskChan <- 1
		getTaskCreationResultChan <- 1
		err := tests.ClearNetwork(addresses, privateKeys)
		assert.Equal(t, nil, err, "error clearing blockchain network")
		tests.ClearDB()
	})
}
