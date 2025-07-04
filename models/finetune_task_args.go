package models

// TransformArgs represents the transformation arguments for data preprocessing
type TransformArgs struct {
	CenterCrop bool `json:"center_crop"`
	RandomFlip bool `json:"random_flip"`
}

// LRSchedulerArgs represents the learning rate scheduler configuration
type LRSchedulerArgs struct {
	LRScheduler   string `json:"lr_scheduler"` // "linear", "cosine", "cosine_with_restarts", "polynomial", "constant", "constant_with_warmup"
	LRWarmupSteps int    `json:"lr_warmup_steps"`
}

// AdamOptimizerArgs represents the Adam optimizer configuration
type AdamOptimizerArgs struct {
	Beta1       float64 `json:"beta1"`
	Beta2       float64 `json:"beta2"`
	WeightDecay float64 `json:"weight_decay"`
	Epsilon     float64 `json:"epsilon"`
}

// LoraArgs represents the LoRA configuration
type LoraArgs struct {
	Rank            int         `json:"rank"`
	InitLoraWeights interface{} `json:"init_lora_weights"` // "true", "false", "gaussian", "loftq" or true, false
	TargetModules   interface{} `json:"target_modules"`    // []string or string or nil
}

// ModelArgs represents the model configuration
type ModelArgs struct {
	Name     string  `json:"name"`
	Variant  *string `json:"variant,omitempty"`
	Revision string  `json:"revision"`
}

// DatasetArgs represents the dataset configuration
type DatasetArgs struct {
	Url           *string `json:"url,omitempty"`
	Name          *string `json:"name,omitempty"`
	ConfigName    *string `json:"config_name,omitempty"`
	ImageColumn   string  `json:"image_column"`
	CaptionColumn string  `json:"caption_column"`
}

// ValidationArgs represents the validation configuration
type ValidationArgs struct {
	Prompt    *string `json:"prompt,omitempty"`
	NumImages int    `json:"num_images"`
}

// TrainArgs represents the training configuration
type TrainArgs struct {
	LearningRate              float64           `json:"learning_rate"`
	BatchSize                 int               `json:"batch_size"`
	GradientAccumulationSteps int               `json:"gradient_accumulation_steps"`
	PredictionType            *string           `json:"prediction_type,omitempty"` // "epsilon", "v_prediction"
	MaxGradNorm               float64           `json:"max_grad_norm"`
	NumTrainEpochs            int               `json:"num_train_epochs"`
	NumTrainSteps             *int              `json:"num_train_steps,omitempty"`
	MaxTrainEpochs            int               `json:"max_train_epochs"`
	MaxTrainSteps             *int              `json:"max_train_steps,omitempty"`
	ScaleLR                   bool              `json:"scale_lr"`
	Resolution                int               `json:"resolution"`
	NoiseOffset               float64           `json:"noise_offset"`
	SNRGamma                  *float64          `json:"snr_gamma,omitempty"`
	LRScheduler               LRSchedulerArgs   `json:"lr_scheduler"`
	AdamArgs                  AdamOptimizerArgs `json:"adam_args"`
}

// FinetuneLoraTaskArgs represents the complete configuration for LoRA fine-tuning
type FinetuneLoraTaskArgs struct {
	Model                ModelArgs      `json:"model"`
	Dataset              DatasetArgs    `json:"dataset"`
	Validation           ValidationArgs `json:"validation"`
	TrainArgs            TrainArgs      `json:"train_args"`
	Lora                 LoraArgs       `json:"lora"`
	Transforms           TransformArgs  `json:"transforms"`
	DataloaderNumWorkers int            `json:"dataloader_num_workers"`
	MixedPrecision       string         `json:"mixed_precision"` // "no", "fp16", "bf16"
	Seed                 int            `json:"seed"`
	Checkpoint           string         `json:"checkpoint,omitempty"`
	Version              string         `json:"version,omitempty"`
}
