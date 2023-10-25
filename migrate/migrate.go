package migrate

import (
	"github.com/go-gormigrate/gormigrate/v2"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"ig_server/migrate/migrations"
)

var migrationScripts []*gormigrate.Gormigrate

func Migrate() error {
	for _, migrationScript := range migrationScripts {
		if err := migrationScript.Migrate(); err != nil {
			log.Errorf("Migrate: %v", err)
			return err
		}
	}

	return nil
}

func Rollback() error {
	lastMigration := migrationScripts[len(migrationScripts)-1]

	if err := lastMigration.RollbackLast(); err != nil {
		log.Errorf("Migrate: %v", err)
		return err
	}

	return nil
}

func InitMigration(db *gorm.DB) {
	db.Set("gorm:table_options", "CHARSET=utf8mb4")

	// Add new migrations here
	migrationScripts = append(migrationScripts, migrations.M20230810(db))
	migrationScripts = append(migrationScripts, migrations.M20230830(db))
}
