package models

import (
	"errors"
	"github.com/ethereum/go-ethereum/crypto"
	"gorm.io/gorm"
	"math/big"
)

type TaskStatus int

const (
	InferenceTaskPending TaskStatus = iota
	InferenceTaskTransactionSent
	InferenceTaskBlockchainConfirmed
	InferenceTaskParamsUploaded
	InferenceTaskPendingResult
	InferenceTaskAborted
	InferenceTaskSuccess
)

type InferenceTask struct {
	RootModel
	ClientID    uint       `json:"client_id"`
	Client      Client     `json:"-"`
	TaskArgs    string     `json:"task_args"`
	Status      TaskStatus `json:"status"`
	TxHash      string     `json:"tx_hash"`
	TaskId      uint64     `json:"task_id"`
	ResultNode  string     `json:"result_node"`
	AbortReason string     `json:"abort_reason"`
}

func (t *InferenceTask) BeforeCreate(*gorm.DB) error {
	t.Status = InferenceTaskPending
	return nil
}

func (t *InferenceTask) GetTaskHash() (*[32]byte, error) {

	hash := crypto.Keccak256Hash([]byte(t.TaskArgs))
	byte32Hash := byteArrayToByte32Array(hash.Bytes())
	return byte32Hash, nil
}

func (t *InferenceTask) GetDataHash() (*[32]byte, error) {
	return nil, nil
}

func (t *InferenceTask) AbortWithReason(reason string, db *gorm.DB) error {

	if t.ID == 0 {
		return errors.New("task not saved in the DB")
	}

	t.AbortReason = reason
	t.Status = InferenceTaskAborted

	return db.Model(t).Select("Status", "AbortReason").Updates(t).Error
}

func UpdateStatusForTask(taskId *big.Int, status TaskStatus, db *gorm.DB) (*InferenceTask, error) {
	task := &InferenceTask{TaskId: taskId.Uint64()}
	if err := db.Where(task).Select("ID").First(task).Error; err != nil {
		return nil, err
	}
	return task, db.Model(task).Update("Status", status).Error
}

func byteArrayToByte32Array(input []byte) *[32]byte {
	var output [32]byte
	copy(output[:], input)
	return &output
}
