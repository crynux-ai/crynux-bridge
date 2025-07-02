package sdftlora

import (
	"crynux_bridge/api/ratelimit"
	"crynux_bridge/api/v1/inference_tasks"
	"crynux_bridge/api/v1/response"
	"crynux_bridge/api/v1/tools"
	"crynux_bridge/config"
	"crynux_bridge/models"
	"encoding/json"
	"fmt"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type SDFinetuneLoraRequest struct {
	SDFinetuneLoraTaskParams
	Authorization string `header:"Authorization" validate:"required" description:"API key"`
}

func CreateSDFinetuneLoraTask(c *gin.Context, in *SDFinetuneLoraRequest) (*response.Response, error) {
	ctx := c.Request.Context()
	db := config.GetDB()

	// validate request (apiKey)
	apiKey, err := tools.ValidateAuthorization(ctx, db, in.Authorization)
	if err != nil {
		return nil, err
	}

	allowed, waitTime, err := ratelimit.APIRateLimiter.CheckRateLimit(ctx, apiKey.ClientID, apiKey.RateLimit, time.Minute)
	if err != nil {
		return nil, response.NewExceptionResponse(err)
	}
	if !allowed {
		return nil, response.NewValidationErrorResponse("rate_limit", fmt.Sprintf("rate limit exceeded, please wait %.2f seconds", waitTime))
	}

	in.SetDefaultValues()

	taskArgs := &models.FinetuneLoraTaskArgs{
		Model: models.ModelArgs{
			Name:     in.ModelName,
			Variant:  *in.ModelVariant,
			Revision: in.ModelRevision,
		},
		Dataset: models.DatasetArgs{
			Url:           *in.DatasetUrl,
			Name:          *in.DatasetName,
			ConfigName:    *in.DatasetConfigName,
			ImageColumn:   in.DatasetImageColumn,
			CaptionColumn: in.DatasetCaptionColumn,
		},
		Validation: models.ValidationArgs{
			Prompt:    *in.ValidationPrompt,
			NumImages: in.ValidationNumImages,
		},
		TrainArgs: models.TrainArgs{
			LearningRate:              in.LearningRate,
			BatchSize:                 in.BatchSize,
			GradientAccumulationSteps: in.GradientAccumulationSteps,
			PredictionType:            in.PredictionType,
			MaxGradNorm:               in.MaxGradNorm,
			NumTrainEpochs:            in.NumTrainEpochs,
			NumTrainSteps:             in.NumTrainSteps,
			MaxTrainEpochs:            in.MaxTrainEpochs,
			MaxTrainSteps:             in.MaxTrainSteps,
			ScaleLR:                   *in.ScaleLR,
			Resolution:                in.Resolution,
			NoiseOffset:               in.NoiseOffset,
			SNRGamma:                  in.SNRGamma,
			LRScheduler: models.LRSchedulerArgs{
				LRScheduler:   in.LRScheduler,
				LRWarmupSteps: in.LRWarmupSteps,
			},
			AdamArgs: models.AdamOptimizerArgs{
				Beta1:       in.AdamBeta1,
				Beta2:       in.AdamBeta2,
				WeightDecay: in.AdamWeightDecay,
				Epsilon:     in.AdamEpsilon,
			},
		},
		Lora: models.LoraArgs{
			Rank:            in.Rank,
			InitLoraWeights: in.InitLoraWeights,
			TargetModules:   in.TargetModules,
		},
		Transforms: models.TransformArgs{
			CenterCrop: in.CenterCrop,
			RandomFlip: in.RandomFlip,
		},
		MixedPrecision: in.MixedPrecision,
		Seed:           in.Seed,
	}

	form, err := c.MultipartForm()
	if err != nil {
		return nil, response.NewExceptionResponse(err)
	}

	if files, ok := form.File["checkpoint"]; ok {
		if len(files) != 1 {
			return nil, response.NewValidationErrorResponse("checkpoint", "More than one checkpoint file uploaded")
		}

		checkpoint := files[0]
		appConfig := config.GetConfig()

		uuid := uuid.New().String()
		checkpointFilename := filepath.Join(appConfig.DataDir.InferenceTasks, fmt.Sprintf("%s_checkpoint.zip", uuid))
		if err = c.SaveUploadedFile(checkpoint, checkpointFilename); err != nil {
			return nil, response.NewExceptionResponse(err)
		}
		taskArgs.Checkpoint = checkpointFilename
	}

	taskArgsStr, err := json.Marshal(taskArgs)
	if err != nil {
		return nil, response.NewExceptionResponse(err)
	}

	taskType := models.TaskTypeSDFTLora
	minVram := uint64(24)
	taskFee := uint64(15000000000)
	repeatNum := 1
	task := &inference_tasks.TaskInput{
		ClientID: apiKey.ClientID,
		TaskArgs: string(taskArgsStr),
		TaskType: &taskType,
		MinVram:  &minVram,
		TaskFee:  &taskFee,
		RepeatNum: &repeatNum,
	}

	_, _, err = inference_tasks.ProcessSDFTLoraTask(ctx, db, task)
	if err != nil {
		return nil, err
	}

	return &response.Response{}, nil
}
