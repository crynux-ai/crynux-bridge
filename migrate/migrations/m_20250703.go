package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func M20250703(db *gorm.DB) *gormigrate.Gormigrate {
	type InferenceTask struct {
		ClientID        uint          `json:"client_id" gorm:"index"`
		ClientTaskID    uint          `json:"client_task_id" gorm:"index"`

	}
	return gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "M20250703",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.Migrator().CreateIndex(&InferenceTask{}, "ClientID"); err != nil {
					return err
				}
				if err := tx.Migrator().CreateIndex(&InferenceTask{}, "ClientTaskID"); err != nil {
					return err
				}

				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				if err := tx.Migrator().DropIndex(&InferenceTask{}, "ClientID"); err != nil {
					return err
				}
				if err := tx.Migrator().DropIndex(&InferenceTask{}, "ClientTaskID"); err != nil {
					return err
				}
				return nil
			},
		},
	})
}
