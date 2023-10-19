package utils

import (
	"jobs-api/config"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func InitialiseLogger() {
	logger = logrus.New()
	logger.Formatter = &logrus.TextFormatter{}
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.InfoLevel)

	// Optionally, you can configure logrus to log to a file
	file, err := os.OpenFile(GetLogFilePath(config.GetConfig().Log.Path, config.GetConfig().App.Name), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		logger.SetOutput(file)
	} else {
		logger.Info("Failed to log to file, using default stdout")
	}
}

func GetLogFilePath(logPath, appName string) string {
	var logDir string

	if logPath == "root" {
		logDir = "" // Root directory
	} else if logPath == "var" {
		logDir = filepath.Join("/var", "logs", appName)
	} else if logPath == "temp" {
		logDir = os.TempDir()
	} else {
		// Default to the root directory
		logDir = ""
	}

	if logDir != "" {
		err := os.MkdirAll(logDir, 0755)
		if err != nil {
			panic(err)
		}
	}

	// Build the full log file path
	logFilePath := filepath.Join(logDir, "app.log")

	return logFilePath
}

func GetLogger() *logrus.Logger {
	return logger
}
