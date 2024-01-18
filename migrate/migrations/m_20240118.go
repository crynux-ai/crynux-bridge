package migrations

import (
	"crynux_bridge/models"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func M20240115(db *gorm.DB) *gormigrate.Gormigrate {

	return gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "M20240115",
			Migrate: func(tx *gorm.DB) error {

				if err := tx.Migrator().AddColumn(&models.InferenceTask{}, "TaskType"); err != nil {
					return err
				}

				return tx.Migrator().AddColumn(&models.InferenceTask{}, "VramLimit")
			},
			Rollback: func(tx *gorm.DB) error {

				if err := tx.Migrator().DropColumn(&models.InferenceTask{}, "TaskType"); err != nil {
					return err
				}

				return tx.Migrator().DropColumn(&models.InferenceTask{}, "VramLimit")
			},
		},
	})
}
