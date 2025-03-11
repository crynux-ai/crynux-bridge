package models

type RelayTask struct {
	TaskArgs         string          `json:"task_args"`
	TaskIDCommitment string          `json:"task_id_commitment" gorm:"index"`
	Creator          string          `json:"creator"`
	Status           ChainTaskStatus `json:"status"`
	TaskType         ChainTaskType   `json:"task_type" gorm:"index"`
	MinVRAM          uint64          `json:"min_vram"`
	RequiredGPU      string          `json:"required_gpu"`
	RequiredGPUVRAM  uint64          `json:"required_gpu_vram"`
	TaskFee          uint64          `json:"task_fee"`
	TaskSize         uint64          `json:"task_size"`
	ModelIDs         StringArray     `json:"model_ids" gorm:"type:text"`
	AbortReason      TaskAbortReason `json:"abort_reason"`
	TaskError        TaskError       `json:"task_error"`
	SelectedNode     string          `json:"selected_node"`

	Sequence     uint64 `json:"sequence"`
	SamplingSeed string `json:"sampling_seed"`
}
