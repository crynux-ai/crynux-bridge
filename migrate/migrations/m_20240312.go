package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func M20240312(db *gorm.DB) *gormigrate.Gormigrate {
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
		TaskFee     uint64
		Cap         uint64
	}

	return gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "M20240312",
			Migrate: func(tx *gorm.DB) error {

				if err := tx.Migrator().AddColumn(&InferenceTask{}, "TaskFee"); err != nil {
					return err
				}

				return tx.Migrator().AddColumn(&InferenceTask{}, "Cap")
			},
			Rollback: func(tx *gorm.DB) error {

				if err := tx.Migrator().DropColumn(&InferenceTask{}, "TaskFee"); err != nil {
					return err
				}

				return tx.Migrator().DropColumn(&InferenceTask{}, "Cap")
			},
		},
	})
}
