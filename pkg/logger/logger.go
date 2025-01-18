package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

func Init(level, file string) {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logLevel, err := logrus.ParseLevel(level)
	if err != nil {
		logLevel = logrus.InfoLevel
	}
	logrus.SetLevel(logLevel)

	logFile, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err == nil {
		logrus.SetOutput(logFile)
	} else {
		logrus.SetOutput(os.Stdout)
	}
}
