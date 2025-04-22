package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func M20250422_1(db *gorm.DB) *gormigrate.Gormigrate {
	type ClientAPIKey struct {
		UsedCount  int64     `json:"used_count" gorm:"default:0"`
		UseLimit   int64     `json:"use_limit" gorm:"default:20"`
	}
	
	return gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "M20250422_1",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.Migrator().AddColumn(&ClientAPIKey{}, "UsedCount"); err != nil {
					return err
				}
				if err := tx.Migrator().AddColumn(&ClientAPIKey{}, "UseLimit"); err != nil {
					return err
				}

				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				if err := tx.Migrator().DropColumn(&ClientAPIKey{}, "UsedCount"); err != nil {
					return err
				}
				if err := tx.Migrator().DropColumn(&ClientAPIKey{}, "UseLimit"); err != nil {
					return err
				}
				return nil
			},
		},
	})
}
