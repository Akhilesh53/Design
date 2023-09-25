package loggersystem

import (
	"fmt"
	"pattern/LoggerSystem/logger"
)

func CallLogger() {
	logger, err := logger.GetLoggerInstance()
	if err != nil {
		fmt.Println(err)
		return
	}

	logger.Info("this is info statement")
	logger.Debug("this is debug statement")
	logger.Error("this is error statement")
}
