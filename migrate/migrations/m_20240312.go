package migrations

import (
	"crynux_bridge/models"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func M20240312(db *gorm.DB) *gormigrate.Gormigrate {

	return gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "M20240312",
			Migrate: func(tx *gorm.DB) error {

				if err := tx.Migrator().AddColumn(&models.InferenceTask{}, "TaskFee"); err != nil {
					return err
				}

				return tx.Migrator().AddColumn(&models.InferenceTask{}, "Cap")
			},
			Rollback: func(tx *gorm.DB) error {

				if err := tx.Migrator().DropColumn(&models.InferenceTask{}, "TaskFee"); err != nil {
					return err
				}

				return tx.Migrator().DropColumn(&models.InferenceTask{}, "Cap")
			},
		},
	})
}
