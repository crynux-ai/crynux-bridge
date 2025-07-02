package image

import "math/rand"

type SDFinetuneLoraTaskParams struct {
	// model
	ModelName     string  `json:"model_name" form:"model_name" validate:"required" description:"Model name"`                        // Model name
	ModelVariant  *string `json:"model_variant" form:"model_variant" description:"Model variant, optional"`                            // Model variant, optional
	ModelRevision string  `json:"model_revision" form:"model_revision" description:"Model revision, defaults to 'main'" default:"main"` // Model revision, defaults to "main"

	// dataset
	DatasetUrl           *string `json:"dataset_url" form:"dataset_url" description:"Dataset url, optional"`
	DatasetName          *string `json:"dataset_name" form:"dataset_name" description:"Dataset name, optional"`
	DatasetConfigName    *string `json:"dataset_config_name" form:"dataset_config_name" description:"Dataset config name, optional"`
	DatasetImageColumn   string  `json:"dataset_image_column" form:"dataset_image_column" description:"Dataset image column name, defaults to 'image'" default:"image"`
	DatasetCaptionColumn string  `json:"dataset_caption_column" form:"dataset_caption_column" description:"Dataset caption column name, defaults to 'text'" default:"text"`

	// validation
	ValidationPrompt    *string `json:"validation_prompt" form:"validation_prompt" description:"Validation prompt, optional"`                                // Validation prompt, optional
	ValidationNumImages int     `json:"validation_num_images" form:"validation_num_images" description:"Number of validation images, defaults to 4" default:"4"` // Number of validation images, defaults to 4

	// transforms
	CenterCrop bool `json:"center_crop" form:"center_crop" description:"Whether to center crop, defaults to false" default:"false"` // Whether to center crop, defaults to false
	RandomFlip bool `json:"random_flip" form:"random_flip" description:"Whether to random flip, defaults to false" default:"false"` // Whether to random flip, defaults to false

	// lora
	Rank            int         `json:"rank" form:"rank" description:"LoRA rank, defaults to 8" default:"8"`                                                                     // LoRA rank, defaults to 8
	InitLoraWeights interface{} `json:"init_lora_weights" form:"init_lora_weights" description:"Initialize LoRA weights, can be bool or 'gaussian'/'loftq', defaults to true" default:"true"` // Initialize LoRA weights, can be bool or "gaussian"/"loftq", defaults to true
	TargetModules   interface{} `json:"target_modules" form:"target_modules" description:"Target modules, can be []string or string or nil, defaults to nil" default:"nil"`                // Target modules, can be []string or string or nil, defaults to nil

	// train
	LearningRate              float64  `json:"learning_rate" form:"learning_rate" description:"Learning rate, defaults to 1e-4" default:"1e-4"`                            // Learning rate, defaults to 1e-4
	BatchSize                 int      `json:"batch_size" form:"batch_size" description:"Batch size, defaults to 16" default:"16"`                                      // Batch size, defaults to 16
	GradientAccumulationSteps int      `json:"gradient_accumulation_steps" form:"gradient_accumulation_steps" description:"Gradient accumulation steps, defaults to 1" default:"1"`      // Gradient accumulation steps, defaults to 1
	PredictionType            *string  `json:"prediction_type" form:"prediction_type" description:"Prediction type, optional 'epsilon' or 'v_prediction'" default:"epsilon"` // Prediction type, optional "epsilon" or "v_prediction"
	MaxGradNorm               float64  `json:"max_grad_norm" form:"max_grad_norm" description:"Maximum gradient norm, defaults to 1.0" default:"1.0"`                      // Maximum gradient norm, defaults to 1.0
	NumTrainEpochs            int      `json:"num_train_epochs" form:"num_train_epochs" description:"Number of training epochs, defaults to 1" default:"1"`                   // Number of training epochs, defaults to 1
	NumTrainSteps             *int     `json:"num_train_steps" form:"num_train_steps" description:"Number of training steps, optional" default:"nil"`                        // Number of training steps, optional
	MaxTrainEpochs            int      `json:"max_train_epochs" form:"max_train_epochs" description:"Maximum training epochs, defaults to 1" default:"1"`                     // Maximum training epochs, defaults to 1
	MaxTrainSteps             *int     `json:"max_train_steps" form:"max_train_steps" description:"Maximum training steps, optional" default:"nil"`                          // Maximum training steps, optional
	ScaleLR                   *bool    `json:"scale_lr" form:"scale_lr" description:"Whether to scale learning rate, defaults to true" default:"true"`                // Whether to scale learning rate, defaults to true
	Resolution                int      `json:"resolution" form:"resolution" description:"Resolution, defaults to 512" default:"512"`                                    // Resolution, defaults to 512
	NoiseOffset               float64  `json:"noise_offset" form:"noise_offset" description:"Noise offset, defaults to 0" default:"0"`                                    // Noise offset, defaults to 0
	SNRGamma                  *float64 `json:"snr_gamma" form:"snr_gamma" description:"SNR gamma, optional" default:"nil"`                                             // SNR gamma, optional

	// lr_scheduler
	LRScheduler   string `json:"lr_scheduler" form:"lr_scheduler" description:"Learning rate scheduler type, defaults to 'constant'" default:"constant"` // Learning rate scheduler type, defaults to "constant"
	LRWarmupSteps int    `json:"lr_warmup_steps" form:"lr_warmup_steps" description:"Learning rate warmup steps, defaults to 500" default:"500"`            // Learning rate warmup steps, defaults to 500

	// adam
	AdamBeta1       float64 `json:"adam_beta1" form:"adam_beta1" description:"Adam beta1, defaults to 0.9" default:"0.9"`                 // Adam beta1, defaults to 0.9
	AdamBeta2       float64 `json:"adam_beta2" form:"adam_beta2" description:"Adam beta2, defaults to 0.999" default:"0.999"`             // Adam beta2, defaults to 0.999
	AdamWeightDecay float64 `json:"adam_weight_decay" form:"adam_weight_decay" description:"Adam weight decay, defaults to 1e-2" default:"1e-2"` // Adam weight decay, defaults to 1e-2
	AdamEpsilon     float64 `json:"adam_epsilon" form:"adam_epsilon" description:"Adam epsilon, defaults to 1e-8" default:"1e-8"`           // Adam epsilon, defaults to 1e-8

	// other
	MixedPrecision string `json:"mixed_precision" form:"mixed_precision" description:"Mixed precision, optional 'no'/'fp16'/'bf16', defaults to 'no'" default:"no"` // Mixed precision, optional "no"/"fp16"/"bf16", defaults to "no"
	Seed           int    `json:"seed" form:"seed" description:"Random seed, defaults to 0" default:"0"`                                                 // Random seed, defaults to 0
}

func (p *SDFinetuneLoraTaskParams) SetDefaultValues() {
	if p.ModelRevision == "" {
		p.ModelRevision = "main"
	}
	if p.DatasetImageColumn == "" {
		p.DatasetImageColumn = "image"
	}
	if p.DatasetCaptionColumn == "" {
		p.DatasetCaptionColumn = "text"
	}
	if p.ValidationNumImages <= 0 {
		p.ValidationNumImages = 4
	}
	if p.Rank <= 0 {
		p.Rank = 8
	}
	if p.InitLoraWeights == nil {
		p.InitLoraWeights = true
	}
	if p.TargetModules == nil {
		p.TargetModules = nil
	}
	if p.LearningRate <= 0 {
		p.LearningRate = 1e-4
	}
	if p.BatchSize <= 0 {
		p.BatchSize = 16
	}
	if p.GradientAccumulationSteps <= 0 {
		p.GradientAccumulationSteps = 1
	}
	if p.MaxGradNorm <= 0 {
		p.MaxGradNorm = 1.0
	}
	if p.NumTrainEpochs <= 0 {
		p.NumTrainEpochs = 1
	}
	if p.MaxTrainEpochs <= 0 {
		p.MaxTrainEpochs = 1
	}
	if p.ScaleLR == nil {
		scaleLR := true
		p.ScaleLR = &scaleLR
	}
	if p.Resolution <= 0 {
		p.Resolution = 512
	}
	if p.LRScheduler == "" {
		p.LRScheduler = "constant"
	}
	if p.LRWarmupSteps <= 0 {
		p.LRWarmupSteps = 500
	}
	if p.AdamBeta1 <= 0 {
		p.AdamBeta1 = 0.9
	}
	if p.AdamBeta2 <= 0 {
		p.AdamBeta2 = 0.999
	}
	if p.AdamWeightDecay <= 0 {
		p.AdamWeightDecay = 1e-2
	}
	if p.AdamEpsilon <= 0 {
		p.AdamEpsilon = 1e-8
	}
	if p.MixedPrecision == "" {
		p.MixedPrecision = "no"
	}
	if p.Seed <= 0 {
		p.Seed = rand.Intn(100000000)
	}
}
