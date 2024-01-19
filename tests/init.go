package tests

import (
	"crynux_bridge/api"
	"crynux_bridge/config"
	"crynux_bridge/migrate"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

var Application *gin.Engine = nil

func init() {
	wd, _ := os.Getwd()
	wd = strings.SplitAfter(wd, "crynux-bridge")[0]

	if err := os.Chdir(wd); err != nil {
		print(err.Error())
		os.Exit(1)
	}

	if err := config.InitConfig("tests"); err != nil {
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
