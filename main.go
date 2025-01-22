package main

import (
	"context"
	"crynux_bridge/api"
	"crynux_bridge/blockchain"
	"crynux_bridge/config"
	"crynux_bridge/migrate"
	"crynux_bridge/tasks"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	if err := config.InitConfig(""); err != nil {
		print("Error reading config file")
		print(err.Error())
		os.Exit(1)
	}

	conf := config.GetConfig()

	if err := config.InitLog(conf); err != nil {
		print("Error initializing log")
		print(err.Error())
		os.Exit(1)
	}

	if err := config.InitDB(conf); err != nil {
		log.Fatalln(err.Error())
	}

	startDBMigration()

	// Check the account balance
	// Approve all the balance to the task contract
	if err := blockchain.ApproveAllBalanceForTaskCreator(); err != nil {
		log.Fatalln(err)
	}

	go tasks.ProcessTasks(context.Background())
	go tasks.AutoCreateTasks(context.Background())

	startServer()
}

func startServer() {
	conf := config.GetConfig()

	app := api.GetHttpApplication(conf)
	address := fmt.Sprintf("%s:%s", conf.Http.Host, conf.Http.Port)

	log.Infoln("Starting application server...")

	if err := app.Run(address); err != nil {
		log.Fatalln(err)
	}
}

func startDBMigration() {

	migrate.InitMigration(config.GetDB())

	if err := migrate.Migrate(); err != nil {
		log.Fatalln(err)
	}

	log.Infoln("DB migrations are done!")
}
