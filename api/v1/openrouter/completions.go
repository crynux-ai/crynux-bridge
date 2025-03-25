package openrouter

import (
	"crynux_bridge/api/v1/inference_tasks"
	"crynux_bridge/api/v1/openrouter/structs"
	"crynux_bridge/api/v1/openrouter/utils"
	"crynux_bridge/api/v1/response"
	"crynux_bridge/api/v1/tools"
	"crynux_bridge/config"
	"crynux_bridge/models"
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin"
)

// build TaskInput from CompletionsRequest, create task, wait for task to finish, get task result, then return CompletionsResponse
func Completions(c *gin.Context, in *structs.CompletionsRequest) (*structs.CompletionsResponse, error) {
	ctx := c.Request.Context()
	db := config.GetDB()

	/* 1. Build TaskInput from CompletionsRequest */
	in.SetDefaultValues() // set default values for some fields

	clientID := "openrouter"
	// if client does not exist, create a new client
	if _, err := tools.CreateClientIfNotExist(ctx, db, clientID); err != nil {
		return nil, response.NewExceptionResponse(err)
	}

	messages := make([]structs.Message, 1)
	messages[0] = structs.Message{
		Role:    structs.RoleUser,
		Content: utils.RawMessageToString(in.Prompt),
		// ToolCallID: "",
		// ToolCalls:  nil,
	}

	generationConfig := &structs.GPTGenerationConfig{
		MaxNewTokens:       in.MaxTokens,
		DoSample:           true,
		Temperature:        in.Temperature,
		TopP:               in.TopP,
		RepetitionPenalty:  in.FrequencyPenalty,
		NumReturnSequences: in.N,
		// NumBeams:           1,
		// TypicalP:           0.95,
		// TopK:               50,
	}

	taskArgs := structs.GPTTaskArgs{
		Model:            in.Model,
		Messages:         messages,
		GenerationConfig: generationConfig,
		Seed:             in.Seed,
		// Tools:            in.Tools,
		// DType:            structs.DTypeAuto,
		// QuantizeBits:     structs.QuantizeBits8,
	}
	taskArgsStr, err := json.Marshal(taskArgs)
	if err != nil {
		err := errors.New("failed to marshal taskArgs")
		return nil, response.NewExceptionResponse(err)
	}

	taskType := models.TaskTypeLLM

	task := &inference_tasks.TaskInput{
		ClientID:        clientID,
		TaskArgs:        string(taskArgsStr),
		TaskType:        &taskType,
		TaskVersion:     nil,
		MinVram:         nil,
		RequiredGPU:     "",
		RequiredGPUVram: 0,
		RepeatNum:       nil,
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

	return ccResponse, nil
}
