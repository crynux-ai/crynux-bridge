package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func M20250424(db *gorm.DB) *gormigrate.Gormigrate {
	type ClientAPIKey struct {
		RateLimit  int64     `json:"rate_limit" gorm:"default:1"`
	}
	
	return gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "M20250424",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.Migrator().AddColumn(&ClientAPIKey{}, "RateLimit"); err != nil {
					return err
				}

				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				if err := tx.Migrator().DropColumn(&ClientAPIKey{}, "RateLimit"); err != nil {
					return err
				}
				return nil
			},
		},
	})
}
