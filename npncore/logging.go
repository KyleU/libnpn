package npncore

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Initializes the logging subsystem
func NewLogger(level logrus.Level, json bool) *logrus.Logger {
	logger := logrus.New()

	logger.SetOutput(os.Stdout)
	logger.SetLevel(level)

	if json {
		logger.SetFormatter(&logrus.JSONFormatter{})
	} else {
		logger.SetFormatter(&logrus.TextFormatter{})
	}

	return logger
}
