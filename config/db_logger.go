package config

import (
	"context"
	"errors"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
)

type Logger struct {
	SlowThreshold         time.Duration
	SourceField           string
	SkipErrRecordNotFound bool
	Debug                 bool
}

var logrusInstance *log.Logger

func NewDBLogger() *Logger {

	logrusInstance = log.New()

	err := configLogrusInstance()
	if err != nil {
		println("init db logger failed!")
		println(err)
		os.Exit(1)
	}

	return &Logger{
		SkipErrRecordNotFound: true,
		Debug:                 true,
	}
}

func configLogrusInstance() error {
	logrusInstance.SetFormatter(&log.TextFormatter{})

	level, err := log.ParseLevel(appConfig.Db.Log.Level)
	if err != nil {
		return err
	}
	logrusInstance.SetLevel(level)

	logrusInstance.SetReportCaller(true)

	if appConfig.Db.Log.Output == "" || appConfig.Db.Log.Output == "stderr" {
		logrusInstance.SetOutput(os.Stderr)
	} else if appConfig.Db.Log.Output == "stdout" {
		logrusInstance.SetOutput(os.Stdout)
	} else {

		logWriter := &lumberjack.Logger{
			Filename: appConfig.Db.Log.Output,
			Compress: true,
		}

		if appConfig.Db.Log.MaxFileSize == 0 {
			logWriter.MaxSize = 500
		} else {
			logWriter.MaxSize = appConfig.Db.Log.MaxFileSize
		}

		if appConfig.Db.Log.MaxDays == 0 {
			logWriter.MaxAge = 30
		} else {
			logWriter.MaxAge = appConfig.Db.Log.MaxDays
		}

		if appConfig.Db.Log.MaxFileNum == 0 {
			logWriter.MaxBackups = 10
		} else {
			logWriter.MaxBackups = appConfig.Db.Log.MaxFileNum
		}

		logrusInstance.SetOutput(logWriter)
	}
	return nil
}

func (l *Logger) LogMode(gormlogger.LogLevel) gormlogger.Interface {
	return l
}

func (l *Logger) Info(ctx context.Context, s string, args ...interface{}) {
	logrusInstance.WithContext(ctx).Infof(s, args...)
}

func (l *Logger) Warn(ctx context.Context, s string, args ...interface{}) {
	logrusInstance.WithContext(ctx).Warnf(s, args...)
}

func (l *Logger) Error(ctx context.Context, s string, args ...interface{}) {
	logrusInstance.WithContext(ctx).Errorf(s, args...)
}

func (l *Logger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	sql, _ := fc()
	fields := log.Fields{}
	if l.SourceField != "" {
		fields[l.SourceField] = utils.FileWithLineNum()
	}
	if err != nil && !(errors.Is(err, gorm.ErrRecordNotFound) && l.SkipErrRecordNotFound) {
		fields[log.ErrorKey] = err
		logrusInstance.WithContext(ctx).WithFields(fields).Errorf("%s [%s]", sql, elapsed)
		return
	}

	if l.SlowThreshold != 0 && elapsed > l.SlowThreshold {
		logrusInstance.WithContext(ctx).WithFields(fields).Warnf("%s [%s]", sql, elapsed)
		return
	}

	if l.Debug {
		logrusInstance.WithContext(ctx).WithFields(fields).Debugf("%s [%s]", sql, elapsed)
	}
}
