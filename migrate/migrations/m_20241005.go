package migrations

import (
	"crynux_bridge/models"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)


func M20241005(db *gorm.DB) *gormigrate.Gormigrate {

	type ClientTask struct {
		gorm.Model
		ClientID       uint `gorm:"index"`
		Client         models.Client
		InferenceTasks []models.InferenceTask
	}
	
	type InferenceTask struct {
		gorm.Model
		ClientID     uint
		Client       models.Client
		ClientTaskID uint `gorm:"index"`
		ClientTask   ClientTask
		TaskArgs     string
		Status       models.TaskStatus
		TxHash       string
		TaskId       uint64
		ResultNode   string
		AbortReason  string
		TaskType     models.ChainTaskType
		VramLimit    uint64
		TaskFee      uint64
		Cap          uint64
	}


	return gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "M20241005",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.Migrator().CreateTable(&ClientTask{}); err != nil {
					return err
				}

				if err := tx.Migrator().AddColumn(&InferenceTask{}, "ClientTaskID"); err != nil {
					return err
				}

				offset := 0
				limit := 100
				var allTasks []InferenceTask
				for {
					var tasks []InferenceTask
					if err := tx.Model(&InferenceTask{}).Order("id").Offset(offset).Limit(limit).Find(&tasks).Error; err != nil {
						return err
					}
					allTasks = append(allTasks, tasks...)
					if len(tasks) < limit {
						break
					}
					offset += limit
				}

				var clientTasks []ClientTask
				for _, task := range allTasks {
					clientTask := ClientTask{
						ClientID: task.ClientID,
					}
					clientTasks = append(clientTasks, clientTask)
				}

				if err := tx.CreateInBatches(&clientTasks, 100).Error; err != nil {
					return err
				}

				for i, task := range allTasks {
					if err := tx.Model(&task).Update("client_task_id", clientTasks[i].ID).Error; err != nil {
						return err
					}
				}
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				if err := tx.Migrator().DropTable(&ClientTask{}); err != nil {
					return err
				}
				if err := tx.Migrator().DropColumn(&InferenceTask{}, "ClientTaskID"); err != nil {
					return err
				}
				return nil
			},
		},
	})
}
