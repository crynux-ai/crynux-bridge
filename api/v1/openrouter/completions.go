package openrouter

import (
	"crynux_bridge/api/ratelimit"
	"crynux_bridge/api/v1/inference_tasks"
	"crynux_bridge/api/v1/openrouter/structs"
	"crynux_bridge/api/v1/openrouter/utils"
	"crynux_bridge/api/v1/response"
	"crynux_bridge/config"
	"crynux_bridge/models"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type CompletionsRequest struct {
	structs.CompletionsRequest
	Authorization string `header:"Authorization" validate:"required"`
}

// build TaskInput from CompletionsRequest, create task, wait for task to finish, get task result, then return CompletionsResponse
func Completions(c *gin.Context, in *CompletionsRequest) (*structs.CompletionsResponse, error) {
	ctx := c.Request.Context()
	db := config.GetDB()

	/* 1. Build TaskInput from CompletionsRequest */
	in.SetDefaultValues() // set default values for some fields

	// validate request (apiKey)
	apiKey, err := ValidateRequestApiKey(ctx, db, in.Authorization)
	if err != nil {
		return nil, err
	}

	// 检查速率限制
	allowed, waitTime, err := ratelimit.APIRateLimiter.CheckRateLimit(ctx, apiKey.ClientID, apiKey.RateLimit, time.Second)
	if err != nil {
		return nil, response.NewExceptionResponse(err)
	}
	if !allowed {
		return nil, response.NewValidationErrorResponse("rate_limit", fmt.Sprintf("rate limit exceeded, please wait %.2f seconds", waitTime))
	}

	messages := make([]structs.Message, 1)
	messages[0] = structs.Message{
		Role:    structs.RoleUser,
		Content: in.Prompt,
	}

	generationConfig := &structs.GPTGenerationConfig{
		DoSample:           true,
		Temperature:        in.Temperature,
		NumReturnSequences: in.N,
	}
	if in.MaxTokens != nil {
		generationConfig.MaxNewTokens = *in.MaxTokens
	}
	if in.TopP != nil {
		generationConfig.TopP = *in.TopP
	}
	if in.TopK != nil {
		generationConfig.TopK = *in.TopK
	}
	if in.MinP != nil {
		generationConfig.MinP = *in.MinP
	}
	if in.RepetitionPenalty != nil {
		generationConfig.RepetitionPenalty = *in.RepetitionPenalty
	}
	if len(in.Stop) > 0 {
		generationConfig.StopStrings = in.Stop
	}

	var dtype structs.DType = structs.DTypeAuto
	if strings.HasPrefix(in.Model, "Qwen/Qwen2.5") {
		dtype = structs.DTypeBFloat16
	}

	taskArgs := structs.GPTTaskArgs{
		Model:            in.Model,
		Messages:         messages,
		GenerationConfig: generationConfig,
		Seed:             in.Seed,
		DType:            dtype,
		// Tools:            in.Tools,
		// QuantizeBits:     structs.QuantizeBits8,
	}
	taskArgsStr, err := json.Marshal(taskArgs)
	if err != nil {
		err := errors.New("failed to marshal taskArgs")
		return nil, response.NewExceptionResponse(err)
	}

	taskType := models.TaskTypeLLM
	minVram := uint64(24)
	taskFee := uint64(6000000000)

	task := &inference_tasks.TaskInput{
		ClientID:        apiKey.ClientID,
		TaskArgs:        string(taskArgsStr),
		TaskType:        &taskType,
		TaskVersion:     nil,
		MinVram:         &minVram,
		RequiredGPU:     "",
		RequiredGPUVram: 0,
		RepeatNum:       nil,
		TaskFee:         &taskFee,
	}

	/* 2. Create task, wait until task finish and get task result. Implemented by function ProcessGPTTask */
	gptTaskResponse, resultDownloadedTask, err := ProcessGPTTask(ctx, db, task)
	if err != nil {
		return nil, err
	}

	/* 3. Wrap GPTTaskResponse into CompletionsResponse and return */
	choices := make([]structs.CResChoice, len(gptTaskResponse.Choices))
	for i, c := range gptTaskResponse.Choices {
		choice, err := utils.ResponseChoiceToCResChoice(c)
		if err != nil {
			return nil, response.NewExceptionResponse(err)
		}
		choices[i] = choice
	}
	ccResponse := &structs.CompletionsResponse{
		Id:      resultDownloadedTask.TaskID,
		Created: resultDownloadedTask.CreatedAt.Unix(),
		Model:   gptTaskResponse.Model,
		Choices: choices,
		Usage:   utils.UsageToCResUsage(gptTaskResponse.Usage),
		// Object:  "text",
		// SystemFingerprint: resultDownloadedTask.SystemFingerprint,
	}

	if err := apiKey.Use(ctx, db); err != nil {
		return nil, response.NewExceptionResponse(err)
	}

	return ccResponse, nil
}
