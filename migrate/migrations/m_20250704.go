package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func M20250704(db *gorm.DB) *gormigrate.Gormigrate {
	type InferenceTask struct {
		Timeout         uint64        `json:"timeout"`

	}
	return gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "M20250704",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.Migrator().AddColumn(&InferenceTask{}, "Timeout"); err != nil {
					return err
				}
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				if err := tx.Migrator().DropColumn(&InferenceTask{}, "Timeout"); err != nil {
					return err
				}
				return nil
			},
		},
	})
}
