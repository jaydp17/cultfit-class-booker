package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

func New() *logrus.Logger {
	log := logrus.New()
	if isAWSLambda() {
		log.SetFormatter(&logrus.JSONFormatter{})
	} else {
		log.SetFormatter(&logrus.TextFormatter{})
	}
	return log
}

func isAWSLambda() bool {
	fnName := os.Getenv("AWS_LAMBDA_FUNCTION_NAME")
	return len(fnName) > 0
}
