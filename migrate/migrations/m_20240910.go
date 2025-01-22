package migrations

import (
	"crynux_bridge/models"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func M20240910(db *gorm.DB) *gormigrate.Gormigrate {
	type BaseModel struct {
		gorm.Model
		Name        string
		Key         string `gorm:"type:varchar(128);uniqueIndex"`
		Description string
		Link        string `json:"link"`
		Type        models.ModelType
	}
	
	return gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "M20240910",
			Migrate: func(tx *gorm.DB) error {
				var err error
				err = tx.Model(&BaseModel{}).
					Where(&BaseModel{Key: "runway/stable-diffusion-v1-5"}).
					Updates(&BaseModel{Key: "crynux-ai/stable-diffusion-v1-5", Link: "https://huggingface.co/crynux-ai/stable-diffusion-v1-5"}).
					Error
				if err != nil {
					return err
				}
				err = tx.Model(&BaseModel{}).
					Where(&BaseModel{Key: "stabilityai/stable-diffusion-xl-base-1.0"}).
					Updates(&BaseModel{Key: "crynux-ai/stable-diffusion-xl-base-1.0", Link: "https://huggingface.co/crynux-ai/stable-diffusion-xl-base-1.0"}).
					Error
				if err != nil {
					return err
				}

				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				var err error
				err = tx.Model(&BaseModel{}).
					Where(&BaseModel{Key: "crynux-ai/stable-diffusion-v1-5"}).
					Updates(&BaseModel{Key: "runway/stable-diffusion-v1-5", Link: "https://huggingface.co/runwayml/stable-diffusion-v1-5"}).
					Error
				if err != nil {
					return err
				}
				err = tx.Model(&BaseModel{}).
					Where(&BaseModel{Key: "crynux-ai/stable-diffusion-xl-base-1.0"}).
					Updates(&BaseModel{Key: "stabilityai/stable-diffusion-xl-base-1.0", Link: "https://huggingface.co/stabilityai/stable-diffusion-xl-base-1.0"}).
					Error
				if err != nil {
					return err
				}

				return nil
			},
		},
	})
}
