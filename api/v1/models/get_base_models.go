package models

import (
	"github.com/gin-gonic/gin"
	"ig_server/api/v1/response"
	"ig_server/config"
	"ig_server/models"
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
