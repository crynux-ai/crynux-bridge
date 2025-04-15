package migrations

import (

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func M20250415(db *gorm.DB) *gormigrate.Gormigrate {

	type ClientAPIKey struct {
		Roles string `json:"roles" gorm:"type:string;size:255"`
	}

	return gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "M20250415",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.Migrator().AddColumn(&ClientAPIKey{}, "Roles"); err != nil {
					return err
				}

				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				if err := tx.Migrator().DropColumn(&ClientAPIKey{}, "Roles"); err != nil {
					return err
				}
				return nil
			},
		},
	})
}
