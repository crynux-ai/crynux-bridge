package main

import (
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

	// Send tasks to the Blockchain
	go tasks.StartSendTaskOnChain()

	// Get the task creation transactions status from the blockchain
	go tasks.StartGetTaskCreationResult()

	// Download the result images from the relay server
	go tasks.StartDownloadResults()

	// Upload the task params to the relay server
	go tasks.StartUploadTaskParams()

	// Sync block to update task status
	go tasks.StartSyncBlock()

	// Auto create task every 5 minutes
	go tasks.StartAutoCreateTask()

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
