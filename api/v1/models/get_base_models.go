package models

import (
	"crynux_bridge/api/v1/response"
	"crynux_bridge/config"
	"crynux_bridge/models"

	"github.com/gin-gonic/gin"
)

type GetBaseModelsOutput struct {
	response.Response
	Data []models.BaseModel `json:"data"`
}

func GetBaseModels(*gin.Context) (*GetBaseModelsOutput, error) {

	var baseModels []models.BaseModel

	if err := config.GetDB().Model(&models.BaseModel{}).Find(&baseModels).Error; err != nil {
		return nil, response.NewExceptionResponse(err)
	}

	return &GetBaseModelsOutput{
		Data: baseModels,
	}, nil
}
