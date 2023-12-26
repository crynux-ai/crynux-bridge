package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
	"ig_server/models"
)

func M20230830(db *gorm.DB) *gormigrate.Gormigrate {
	type BaseModel struct {
		gorm.Model
		Name        string
		Key         string `gorm:"type:varchar(128);uniqueIndex"`
		Description string
		Link        string `json:"link"`
		Type        models.ModelType
	}

	type LoraModel struct {
		gorm.Model
		Name         string
		Description  string
		Type         models.ModelType
		DisplayLink  string `json:"display_link"`
		DownloadLink string `json:"download_link"`
	}

	var baseModels = []BaseModel{
		{
			Name:        "Stable Diffusion 1.5",
			Key:         "runwayml/stable-diffusion-v1-5",
			Description: "Stable Diffusion 1.5. Inference Only.",
			Link:        "https://huggingface.co/runwayml/stable-diffusion-v1-5",
			Type:        models.ModelType_SD_1_5,
		},
		{
			Name:        "Stable Diffusion XL",
			Key:         "stabilityai/stable-diffusion-xl-base-1.0",
			Description: "A combination of Stable Diffusion XL base and refiner models to generate image directly",
			Link:        "https://huggingface.co/stabilityai/stable-diffusion-xl-base-1.0",
			Type:        models.ModelType_SD_XL,
		},
		{
			Name:        "ChilloutMix",
			Key:         "emilianJR/chilloutmix_NiPrunedFp32Fix",
			Description: "Base model that generates Asian beauties based on Stable Diffusion 1.5",
			Link:        "https://huggingface.co/emilianJR/chilloutmix_NiPrunedFp32Fix",
			Type:        models.ModelType_SD_1_5,
		},
	}

	var loraModels = []LoraModel{
		{
			Name:         "Korean Doll Likeness v20",
			Description:  "Generate girl images with Korean faces",
			DisplayLink:  "https://civitai.com/models/26124/koreandolllikeness-v20",
			DownloadLink: "https://civitai.com/api/download/models/31284",
			Type:         models.ModelType_SD_1_5,
		},
		{
			Name:         "Chinese Doll Likeness v10",
			Description:  "Generate girl images with Chinese faces",
			DisplayLink:  "https://civitai.com/models/9434/lora-chinese-doll-likeness",
			DownloadLink: "https://civitai.com/api/download/models/11195",
			Type:         models.ModelType_SD_1_5,
		},
		{
			Name:         "Japanese Doll Likeness v15",
			Description:  "Generate girl images with Japanese faces",
			DisplayLink:  "https://civitai.com/models/28811/japanesedolllikeness-v15",
			DownloadLink: "https://civitai.com/api/download/models/34562",
			Type:         models.ModelType_SD_1_5,
		},
	}

	return gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "M20230830",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.Migrator().CreateTable(&BaseModel{}); err != nil {
					return err
				}
				if err := tx.Migrator().CreateTable(&LoraModel{}); err != nil {
					return err
				}

				if err := db.Create(baseModels).Error; err != nil {
					return err
				}

				if err := db.Create(loraModels).Error; err != nil {
					return err
				}

				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				if err := tx.Migrator().DropTable("base_models"); err != nil {
					return err
				}
				return tx.Migrator().DropTable("lora_models")
			},
		},
	})
}
