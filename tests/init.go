package tests

import (
	"github.com/gin-gonic/gin"
	"ig_server/api"
	"ig_server/config"
	"ig_server/migrate"
	"os"
)

var Application *gin.Engine = nil

func init() {

	if err := config.InitConfig(""); err != nil {
		print(err.Error())
		os.Exit(1)
	}

	testAppConfig := config.GetConfig()

	if err := config.InitLog(testAppConfig); err != nil {
		print(err.Error())
		os.Exit(1)
	}

	err := config.InitDB(testAppConfig)
	if err != nil {
		print(err.Error())
		os.Exit(1)
	}

	migrate.InitMigration(config.GetDB())

	if err := migrate.Migrate(); err != nil {
		print(err.Error())
		os.Exit(1)
	}

	Application = api.GetHttpApplication(testAppConfig)
}
