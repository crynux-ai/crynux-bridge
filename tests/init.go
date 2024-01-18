package tests

import (
	"crynux_bridge/api"
	"crynux_bridge/config"
	"crynux_bridge/migrate"
	"os"

	"github.com/gin-gonic/gin"
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
