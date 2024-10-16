package models

type ModelType string

const (
	ModelType_SD_1_5 ModelType = "sd_1_5"
	ModelType_SD_2_1 ModelType = "sd_2_1"
	ModelType_SD_XL  ModelType = "sd_xl"
	ModelType_SDXL_Turbo ModelType = "sdxl_turbo"
)

var BaseModelTypes = []ModelType{
	ModelType_SD_1_5,
	ModelType_SD_2_1,
	ModelType_SD_XL,
	ModelType_SDXL_Turbo,
}

func IsModelTypeValid(modelType ModelType) bool {

	for _, mt := range BaseModelTypes {
		if mt == modelType {
			return true
		}
	}

	return false
}

type BaseModel struct {
	RootModel
	Name        string    `json:"name"`
	Key         string    `json:"key"`
	Description string    `json:"description"`
	Link        string    `json:"link"`
	Type        ModelType `json:"type"`
}
