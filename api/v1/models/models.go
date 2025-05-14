package models

import (
	"crynux_bridge/config"
	"encoding/json"
	"os"

	"github.com/gin-gonic/gin"
)

type Model struct {
	ID                  string `json:"id"`
	Name                string `json:"name"`
	Created             uint64 `json:"created"`
	Description         string `json:"description,omitempty"`
	ContextLength       uint64 `json:"context_length"`
	MaxCompletionTokens uint64 `json:"max_completion_tokens"`
	Quantization        string `json:"quantization"`
	Pricing             struct {
		Prompt     string `json:"prompt"`
		Completion string `json:"completion"`
		Image      string `json:"image"`
		Request    string `json:"request"`
	}
}

var modelsList []Model

func readModels(modelsFile string) ([]Model, error) {
	var models []Model
	file, err := os.Open(modelsFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&models)
	if err != nil {
		return nil, err
	}
	return models, nil
}

func getModels(modelsFile string) ([]Model, error) {
	if len(modelsList) == 0 {
		models, err := readModels(modelsFile)
		if err != nil {
			return nil, err
		}
		modelsList = models
	}
	return modelsList, nil
}

type ModelListResponse struct {
	Data []Model `json:"data"`
}

func GetOpenrouterModels(c *gin.Context) (*ModelListResponse, error) {
	appConfig := config.GetConfig()
	models, err := getModels(appConfig.OpenRouter.ModelsFile)
	if err != nil {
		return nil, err
	}
	return &ModelListResponse{Data: models}, nil
}
