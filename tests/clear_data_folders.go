package tests

import (
	"ig_server/config"
	"os"
)

func ClearDataFolders() error {
	appConfig := config.GetConfig()
	return removeAllContent(appConfig.DataDir.InferenceTasks)
}

func removeAllContent(dir string) error {
	if err := os.RemoveAll(dir); err != nil {
		return err
	}

	if err := os.MkdirAll(dir, os.ModeDir); err != nil {
		return err
	}

	return nil
}
