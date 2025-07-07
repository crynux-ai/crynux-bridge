package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func M20250706(db *gorm.DB) *gormigrate.Gormigrate {
	type ClientTask struct {
		Status      string `json:"status" gorm:"index"`
		FailedCount int    `json:"failed_count" gorm:"default:0"`
	}
	return gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "M20250706",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.Migrator().AddColumn(&ClientTask{}, "Status"); err != nil {
					return err
				}
				if err := tx.Migrator().AddColumn(&ClientTask{}, "FailedCount"); err != nil {
					return err
				}
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				if err := tx.Migrator().DropColumn(&ClientTask{}, "Status"); err != nil {
					return err
				}
				if err := tx.Migrator().DropColumn(&ClientTask{}, "FailedCount"); err != nil {
					return err
				}
				return nil
			},
		},
	})
}
