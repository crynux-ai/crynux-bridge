package migrations

import (
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func M20250414(db *gorm.DB) *gormigrate.Gormigrate {

	type ClientAPIKey struct {
		ID         uint           `gorm:"primarykey"`
		CreatedAt  time.Time      `gorm:"index"`
		UpdatedAt  time.Time      `gorm:"index"`
		DeletedAt  gorm.DeletedAt `gorm:"index"`
		ClientID   string         `json:"client_id"`
		KeyPrefix  string         `json:"key_prefix" gorm:"index;type:string;size:255"`
		KeyHash    string         `json:"key_hash"`
		ExpiresAt  time.Time      `json:"expires_at"`
		LastUsedAt time.Time      `json:"last_used_at"`
	}

	return gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "M20250414",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.Migrator().CreateTable(&ClientAPIKey{}); err != nil {
					return err
				}

				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				if err := tx.Migrator().DropTable(&ClientAPIKey{}); err != nil {
					return err
				}
				return nil
			},
		},
	})
}
