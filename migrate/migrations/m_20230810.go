package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func M20230810(db *gorm.DB) *gormigrate.Gormigrate {
	type Client struct {
		gorm.Model
		ClientId string `json:"client_id"`
	}

	type InferenceTask struct {
		gorm.Model
		ClientID    uint `gorm:"index"`
		Client      Client
		TaskArgs    string
		Status      int
		TxHash      string `gorm:"uniqueIndex;size:256"`
		TaskId      uint64 `gorm:"index"`
		ResultNode  string
		AbortReason string
	}

	type SyncedBlock struct {
		gorm.Model
		BlockNumber uint64
	}

	return gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "M20230810",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.Migrator().CreateTable(&Client{}); err != nil {
					return err
				}

				if err := tx.Migrator().CreateTable(&InferenceTask{}); err != nil {
					return err
				}

				return tx.Migrator().CreateTable(&SyncedBlock{})
			},
			Rollback: func(tx *gorm.DB) error {
				if err := tx.Migrator().DropTable("inference_tasks"); err != nil {
					return err
				}
				if err := tx.Migrator().DropTable("clients"); err != nil {
					return err
				}
				return tx.Migrator().DropTable("synced_blocks")
			},
		},
	})
}
