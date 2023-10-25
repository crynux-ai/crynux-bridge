package config

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
)

func InitLog(appConfig *AppConfig) error {

	println("Initializing logger...")

	logrus.SetFormatter(&logrus.TextFormatter{})

	level, err := logrus.ParseLevel(appConfig.Log.Level)
	if err != nil {
		return err
	}
	logrus.SetLevel(level)

	logrus.SetReportCaller(true)

	if appConfig.Log.Output == "" || appConfig.Log.Output == "stderr" {
		logrus.SetOutput(os.Stderr)
	} else if appConfig.Log.Output == "stdout" {
		logrus.SetOutput(os.Stdout)
	} else {

		logWriter := &lumberjack.Logger{
			Filename: appConfig.Log.Output,
			Compress: true,
		}

		if appConfig.Log.MaxFileSize == 0 {
			logWriter.MaxSize = 500
		} else {
			logWriter.MaxSize = appConfig.Log.MaxFileSize
		}

		if appConfig.Log.MaxDays == 0 {
			logWriter.MaxAge = 30
		} else {
			logWriter.MaxAge = appConfig.Log.MaxDays
		}

		if appConfig.Log.MaxFileNum == 0 {
			logWriter.MaxBackups = 10
		} else {
			logWriter.MaxBackups = appConfig.Log.MaxFileNum
		}

		mw := io.MultiWriter(os.Stdout, logWriter)
		logrus.SetOutput(mw)
	}

	return nil
}
