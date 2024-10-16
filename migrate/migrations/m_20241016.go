package migrations

import (
	"crynux_bridge/models"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func M20241016(db *gorm.DB) *gormigrate.Gormigrate {
	return gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "M20241016",
			Migrate: func(tx *gorm.DB) error {
				err := tx.Model(&models.BaseModel{}).
					Where(&models.BaseModel{Key: "crynux-ai/stable-diffusion-xl-base-1.0"}).
					Updates(&models.BaseModel{
						Name: "SDXL Turbo", 
						Key: "crynux-ai/sdxl-turbo", 
						Link: "https://huggingface.co/crynux-ai/sdxl-turbo", 
						Type: models.ModelType_SDXL_Turbo,
						Description: "SDXL Turbo. Inference Only.",
					}).
					Error
				if err != nil {
					return err
				}

				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				err := tx.Model(&models.BaseModel{}).
					Where(&models.BaseModel{Key: "crynux-ai/sdxl-turbo"}).
					Updates(&models.BaseModel{
						Name: "Stable Diffusion XL",
						Key: "crynux-ai/stable-diffusion-xl-base-1.0", 
						Link: "https://huggingface.co/crynux-ai/stable-diffusion-xl-base-1.0",
						Type: models.ModelType_SD_XL,
						Description: "A combination of Stable Diffusion XL base and refiner models to generate image directly",
					}).
					Error
				if err != nil {
					return err
				}

				return nil
			},
		},
	})
}
