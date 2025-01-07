package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
	"gorm.io/gorm"
)

type ChainTaskStatus uint8

const (
	ChainTaskQueued ChainTaskStatus = iota
	ChainTaskStarted
	ChainTaskParametersUploaded
	ChainTaskErrorReported
	ChainTaskScoreReady
	ChainTaskValidated
	ChainTaskGroupValidated
	ChainTaskEndInvalidated
	ChainTaskEndSuccess
	ChainTaskEndAborted
	ChainTaskEndGroupRefund
	ChainTaskEndGroupSuccess
)

type TaskStatus int

const (
	InferenceTaskPending TaskStatus = iota
	InferenceTaskCreated
	InferenceTaskParamsUploaded
	InferenceTaskScoreReady
	InferenceTaskValidated
	InferenceTaskEndAborted
	InferenceTaskEndGroupRefund
	InferenceTaskEndInvalidated
	InferenceTaskEndSuccess
	InferenceTaskResultDownloaded
)

type ChainTaskType uint8

const (
	TaskTypeSD ChainTaskType = iota
	TaskTypeLLM
	TaskTypeSDFTLora
)

type TaskAbortReason uint8

const (
	TaskAbortReasonNone TaskAbortReason = iota
	TaskAbortTimeout
	TaskAbortModelDownloadFailed
	TaskAbortIncorrectResult
	TaskAbortTaskFeeTooLow
)

type TaskError uint8

const (
	TaskErrorNone TaskError = iota
	TaskErrorParametersValidationFailed
)

type StringArray []string

func (arr *StringArray) Scan(val interface{}) error {
	var arrString string
	switch v := val.(type) {
	case string:
		arrString = v
	case []byte:
		arrString = string(v)
	case nil:
		return nil
	default:
		return errors.New(fmt.Sprint("Unable to parse value to StringArray: ", val))
	}
	*arr = strings.Split(arrString, ";")
	return nil
}

func (arr StringArray) Value() (driver.Value, error) {
	res := strings.Join(arr, ";")
	return res, nil
}

func (arr StringArray) MarshalJSON() ([]byte, error) {
	return json.Marshal([]string(arr))
}

func (arr *StringArray) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, (*[]string)(arr))
}

type InferenceTask struct {
	RootModel
	ClientID        uint          `json:"client_id"`
	Client          Client        `json:"-"`
	ClientTaskID    uint          `json:"client_task_id"`
	ClientTask      ClientTask    `json:"-"`
	TaskArgs        string        `json:"task_args"`
	TaskType        ChainTaskType `json:"task_type"`
	TaskModelIDs    StringArray   `json:"task_model_ids"`
	TaskVersion     string        `json:"task_version"`
	TaskFee         uint64        `json:"task_fee"`
	MinVram         uint64        `json:"min_vram"`
	RequiredGPU     string        `json:"required_gpu"`
	RequiredGPUVram uint64        `json:"required_gpu_vram"`
	TaskSize        uint64        `json:"task_size"`

	Status           TaskStatus `json:"status"`
	TaskID           string     `json:"task_id"`
	TaskIDCommitment string     `json:"task_id_commitment"`
	Nonce            string     `json:"nonce"`
	Sequence         uint64     `json:"sequence"`
	NeedResult       bool       `json:"need_result"`

	AbortReason TaskAbortReason `json:"abort_reason"`
	TaskError   TaskError       `json:"task_error"`
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

func byteArrayToByte32Array(input []byte) *[32]byte {
	var output [32]byte
	copy(output[:], input)
	return &output
}
