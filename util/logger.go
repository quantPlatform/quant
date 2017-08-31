package util

import (
	"os"

	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func InitLogger() {
	logger = logrus.New()
	logger.Out = os.Stderr
}

func Logger() *logrus.Logger {
	return logger
}
