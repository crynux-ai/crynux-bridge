package models

import (
	"crynux_bridge/api/v1/response"
	"crynux_bridge/config"
	"crynux_bridge/models"

	"github.com/gin-gonic/gin"
)

type GetLoraModelsInput struct {
	Type models.ModelType `json:"type" query:"type"`
}

type GetLoraModelsOutput struct {
	response.Response
	Data []models.LoraModel `json:"data"`
}

func GetLoraModels(_ *gin.Context, in *GetLoraModelsInput) (*GetLoraModelsOutput, error) {

	var loraModels []models.LoraModel

	if !models.IsModelTypeValid(in.Type) {
		return nil, response.NewValidationErrorResponse("type", "Invalid model type")
	}

	if err := config.GetDB().Where(&models.LoraModel{
		Type: in.Type,
	}).Find(&loraModels).Error; err != nil {
		return nil, response.NewExceptionResponse(err)
	}

	return &GetLoraModelsOutput{
		Data: loraModels,
	}, nil
}
