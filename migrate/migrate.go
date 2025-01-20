package migrate

import (
	"crynux_bridge/migrate/migrations"

	"github.com/go-gormigrate/gormigrate/v2"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
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
	migrationScripts = append(migrationScripts, migrations.M20240115(db))
	migrationScripts = append(migrationScripts, migrations.M20240312(db))
	migrationScripts = append(migrationScripts, migrations.M20240910(db))
	migrationScripts = append(migrationScripts, migrations.M20240919(db))
	migrationScripts = append(migrationScripts, migrations.M20241005(db))
	migrationScripts = append(migrationScripts, migrations.M20241016(db))
	migrationScripts = append(migrationScripts, migrations.M20250120(db))
}
