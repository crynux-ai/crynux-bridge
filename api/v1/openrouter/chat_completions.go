package openrouter

import (
	"crynux_bridge/api/v1/inference_tasks"
	"crynux_bridge/api/v1/openrouter/structs"
	"crynux_bridge/api/v1/openrouter/utils"
	"crynux_bridge/api/v1/response"
	"crynux_bridge/models"
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin"
)

// build TaskInput from ChatCompletionsRequest, create task, wait for task to finish, get task result, then return ChatCompletionsResponse
func ChatCompletions(c *gin.Context, in *structs.ChatCompletionsRequest) (*structs.ChatCompletionsResponse, error) {

	/* 1. Build TaskInput from ChatCompletionsRequest */
	in.SetDefaultValues() // set default values for some fields

	clientID := "openrouter"

	messages := make([]structs.Message, len(in.Messages))
	for i, m := range in.Messages {
		messages[i] = utils.CCReqMessageToMessage(m)
	}

	generationConfig := &structs.GPTGenerationConfig{
		MaxNewTokens:       in.MaxComletionTokens,
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
		Tools:            in.Tools,
		GenerationConfig: generationConfig,
		Seed:             in.Seed,
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
	gptTaskResponse, resultDownloadedTask, err := ProcessGPTTask(c, task)
	if err != nil {
		return nil, err
	}

	/* 3. Wrap GPTTaskResponse into ChatCompletionsResponse and return */
	choices := make([]structs.CCResChoice, len(gptTaskResponse.Choices))
	for i, c := range gptTaskResponse.Choices {
		choices[i] = utils.ResponseChoiceToCCResChoice(c)
	}
	ccResponse := &structs.ChatCompletionsResponse{
		Id:      resultDownloadedTask.TaskID,
		Created: resultDownloadedTask.CreatedAt.Unix(),
		Model:   gptTaskResponse.Model,
		Choices: choices,
		Usage:   utils.UsageToCCResUsage(gptTaskResponse.Usage),
		// Object:  "text",
		// ServiceTier: "",
	}

	return ccResponse, nil

}
