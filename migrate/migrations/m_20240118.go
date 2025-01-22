package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func M20240115(db *gorm.DB) *gormigrate.Gormigrate {
	type Client struct {
		gorm.Model
		ClientId string `json:"client_id"`
	}

	type ChainTaskType uint8

	type InferenceTask struct {
		gorm.Model
		ClientID    uint `gorm:"index"`
		Client      Client
		TaskArgs    string
		Status      int
		TxHash      string `gorm:"index;size:256"`
		TaskId      uint64 `gorm:"index"`
		ResultNode  string
		AbortReason string
		TaskType    ChainTaskType
		VramLimit   uint64
	}

	return gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "M20240115",
			Migrate: func(tx *gorm.DB) error {

				if err := tx.Migrator().AddColumn(&InferenceTask{}, "TaskType"); err != nil {
					return err
				}

				return tx.Migrator().AddColumn(&InferenceTask{}, "VramLimit")
			},
			Rollback: func(tx *gorm.DB) error {

				if err := tx.Migrator().DropColumn(&InferenceTask{}, "TaskType"); err != nil {
					return err
				}

				return tx.Migrator().DropColumn(&InferenceTask{}, "VramLimit")
			},
		},
	})
}
