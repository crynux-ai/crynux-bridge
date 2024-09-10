package migrations

import (
	"crynux_bridge/models"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func M20240910(db *gorm.DB) *gormigrate.Gormigrate {
	return gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "M20230830",
			Migrate: func(tx *gorm.DB) error {
				var err error
				err = tx.Model(&models.BaseModel{}).
					Where(&models.BaseModel{Type: models.ModelType_SD_1_5}).
					Updates(&models.BaseModel{Key: "crynux-ai/stable-diffusion-v1-5", Link: "https://huggingface.co/crynux-ai/stable-diffusion-v1-5"}).
					Error
				if err != nil {
					return err
				}
				err = tx.Model(&models.BaseModel{}).
					Where(&models.BaseModel{Type: models.ModelType_SD_XL}).
					Updates(&models.BaseModel{Key: "crynux-ai/stable-diffusion-xl-base-1.0", Link: "https://huggingface.co/crynux-ai/stable-diffusion-xl-base-1.0"}).
					Error
				if err != nil {
					return err
				}

				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				var err error
				err = tx.Model(&models.BaseModel{}).
					Where(&models.BaseModel{Type: models.ModelType_SD_1_5}).
					Updates(&models.BaseModel{Key: "runway/stable-diffusion-v1-5", Link: "https://huggingface.co/runwayml/stable-diffusion-v1-5"}).
					Error
				if err != nil {
					return err
				}
				err = tx.Model(&models.BaseModel{}).
					Where(&models.BaseModel{Type: models.ModelType_SD_XL}).
					Updates(&models.BaseModel{Key: "stabilityai/stable-diffusion-xl-base-1.0", Link: "https://huggingface.co/stabilityai/stable-diffusion-xl-base-1.0"}).
					Error
				if err != nil {
					return err
				}

				return nil
			},
		},
	})
}
