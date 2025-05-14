package pkg

import (
	"fmt"
	"runtime"

	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func InitLogger() {
	Logger = logrus.New()
	Logger.SetFormatter(&logrus.JSONFormatter{})
	Logger.SetLevel(logrus.InfoLevel)
}

func LogError(msg string, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		funcName := extractFunctionName()
		formattedMsg := fmt.Sprintf("Error in %s: %s, occurred in %s:%d", funcName, msg, file, line)
		Logger.WithFields(logrus.Fields{
			"error": err,
		}).Error(formattedMsg)
	}
}

func LogInfo(msg string) {
	_, file, line, _ := runtime.Caller(1)
	funcName := extractFunctionName()
	formattedMsg := fmt.Sprintf("Info in %s: %s, occurred in %s:%d", funcName, msg, file, line)
	Logger.Info(formattedMsg)
}

func extractFunctionName() string {
	pc, _, _, ok := runtime.Caller(2)
	if !ok {
		return "unknown"
	}

	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return "unknown"
	}

	return fn.Name()
}
