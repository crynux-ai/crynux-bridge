package migrations

import (
	"crynux_bridge/models"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func M20240919(db *gorm.DB) *gormigrate.Gormigrate {
	return gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "M20240919",
			Migrate: func(tx *gorm.DB) error {
				err := tx.Model(&models.BaseModel{}).
					Where(&models.BaseModel{Key: "runwayml/stable-diffusion-v1-5"}).
					Updates(&models.BaseModel{Key: "crynux-ai/stable-diffusion-v1-5", Link: "https://huggingface.co/crynux-ai/stable-diffusion-v1-5"}).
					Error
				if err != nil {
					return err
				}
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				err := tx.Model(&models.BaseModel{}).
					Where(&models.BaseModel{Key: "crynux-ai/stable-diffusion-v1-5"}).
					Updates(&models.BaseModel{Key: "runwayml/stable-diffusion-v1-5", Link: "https://huggingface.co/runwayml/stable-diffusion-v1-5"}).
					Error
				if err != nil {
					return err
				}
				return nil
			},
		},
	})
}
