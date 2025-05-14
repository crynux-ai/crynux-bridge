package models

import "encoding/json"

type SDModelArgs struct {
	Name    string `json:"name"`
	Variant string `json:"variant,omitempty"`
}

type SDTaskConfig struct {
	ImageWidth    int     `json:"image_width"`
	ImageHeight   int     `json:"image_height"`
	Steps         int     `json:"steps"`
	Seed          int     `json:"seed"`
	NumImages     int     `json:"num_images"`
	SafetyChecker bool    `json:"safety_checker"`
	Cfg           float64 `json:"cfg"`
}

type SDLoraArgs struct {
	Model          string `json:"model"`
	Weight         int    `json:"weight,omitempty"`
	WeightFileName string `json:"weight_file_name,omitempty"`
}

type SDControlnetArgs struct {
	Model        string      `json:"model"`
	Variant      string      `json:"variant,omitempty"`
	ImageDataurl string      `json:"image_dataurl,omitempty"`
	Weight       int         `json:"weight,omitempty"`
	Preprocess   interface{} `json:"preprocess,omitempty"`
}

type Scheduler interface {
	GetMethod() SDSchedulerMethod
}

type SDSchedulerMethod string

const (
	SDSchedulerMethodEulerAncestral SDSchedulerMethod = "EulerAncestralDiscreteScheduler"
	SDSchedulerMethodLCM            SDSchedulerMethod = "LCMScheduler"
	SDSchedulerMethodDPM            SDSchedulerMethod = "DPMSolverMultistepScheduler"
)

type EulerAncestralDiscrete struct {
	NumTrainTimesteps   int     `json:"num_train_timesteps,omitempty"`
	BetaStart           float64 `json:"beta_start,omitempty"`
	BetaEnd             float64 `json:"beta_end,omitempty"`
	BetaSchedule        string  `json:"beta_schedule,omitempty"`
	PredictionType      string  `json:"prediction_type,omitempty"`
	TimestepSpacing     string  `json:"timestep_spacing,omitempty"`
	StepsOffset         int     `json:"steps_offset,omitempty"`
	RescaleBetasZeroSNR float64 `json:"rescale_betas_zero_snr,omitempty"`
}

func (s *EulerAncestralDiscrete) GetMethod() SDSchedulerMethod {
	return SDSchedulerMethodEulerAncestral
}

func (s EulerAncestralDiscrete) MarshalJSON() ([]byte, error) {
	type Alias EulerAncestralDiscrete
	return json.Marshal(&struct {
		Method string `json:"method"`
		Args   *Alias `json:"args,omitempty"`
	}{
		Method: string(SDSchedulerMethodEulerAncestral),
		Args:   (*Alias)(&s),
	})
}

type LCM struct {
	OriginalInferenceSteps   int     `json:"original_inference_steps,omitempty"`
	ClipSamples              int     `json:"clip_samples,omitempty"`
	ClipSamplesRange         int     `json:"clip_samples_range,omitempty"`
	SetAlphaToOne            bool    `json:"set_alpha_to_one,omitempty"`
	Thresholding             bool    `json:"thresholding,omitempty"`
	DynamicThresholdingRatio float64 `json:"dynamic_thresholding_ratio,omitempty"`
	SampleMaxValue           float64 `json:"sample_max_value,omitempty"`
	TimestepScaling          float64 `json:"timestep_scaling,omitempty"`
}

func (s *LCM) GetMethod() SDSchedulerMethod {
	return SDSchedulerMethodLCM
}

func (s LCM) MarshalJSON() ([]byte, error) {
	type Alias LCM
	return json.Marshal(&struct {
		Method string `json:"method"`
		Args   *Alias `json:"args,omitempty"`
	}{
		Method: string(SDSchedulerMethodLCM),
		Args:   (*Alias)(&s),
	})
}

type DPMSolverMultistep struct {
	SolverOrder              int     `json:"solver_order,omitempty"`
	Thresholding             bool    `json:"thresholding,omitempty"`
	DynamicThresholdingRatio float64 `json:"dynamic_thresholding_ratio,omitempty"`
	SampleMaxValue           float64 `json:"sample_max_value,omitempty"`
	AlgorithmType            string  `json:"algorithm_type,omitempty"`
	SolverType               string  `json:"solver_type,omitempty"`
	LowerOrderFinal          bool    `json:"lower_order_final,omitempty"`
	EulerAtFinal             bool    `json:"euler_at_final,omitempty"`
	UseKarrasSigmas          bool    `json:"use_karras_sigmas,omitempty"`
	UseLuLambdas             bool    `json:"use_lu_lambdas,omitempty"`
	FinalSigmasType          string  `json:"final_sigmas_type,omitempty"`
	LambdaMinClipped         float64 `json:"lambda_min_clipped,omitempty"`
	VarianceType             string  `json:"variance_type,omitempty"`
}

func (s *DPMSolverMultistep) GetMethod() SDSchedulerMethod {
	return SDSchedulerMethodDPM
}

func (s DPMSolverMultistep) MarshalJSON() ([]byte, error) {
	type Alias DPMSolverMultistep
	return json.Marshal(&struct {
		Method string `json:"method"`
		Args   *Alias `json:"args,omitempty"`
	}{
		Method: string(SDSchedulerMethodDPM),
		Args:   (*Alias)(&s),
	})
}

type RefinerArgs struct {
	Model           string `json:"model"`
	Variant         string `json:"variant,omitempty"`
	DenoisingCutoff int    `json:"denoising_cutoff,omitempty"`
	Steps           int    `json:"steps,omitempty"`
}

type SDTaskArgs struct {
	BaseModel        SDModelArgs       `json:"base_model"`
	Prompt           string            `json:"prompt"`
	TaskConfig       SDTaskConfig      `json:"task_config"`
	Dtype            DType             `json:"dtype,omitempty"`
	Unet             string            `json:"unet,omitempty"`
	NegativePrompt   string            `json:"negative_prompt,omitempty"`
	Lora             *SDLoraArgs       `json:"lora,omitempty"`
	Controlnet       *SDControlnetArgs `json:"controlnet,omitempty"`
	Scheduler        Scheduler        `json:"scheduler,omitempty"`
	Vae              string            `json:"vae,omitempty"`
	Refiner          *RefinerArgs      `json:"refiner,omitempty"`
	TextualInversion string            `json:"textual_inversion,omitempty"`
}
