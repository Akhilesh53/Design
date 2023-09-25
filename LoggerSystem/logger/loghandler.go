package logger

import (
	"errors"
	"fmt"
)

type ILogHandler interface {
	GetNext() ILogHandler
	SetNext(next ILogHandler)
	GetLevel() string
	SetLevel(level LogType)
	LogMessage(logType LogType, msg string)
}

func GetLogHandler(logType LogType, next ILogHandler) (ILogHandler, error) {
	switch logType {
	case INFO:
		return getInfoLogHandler(next), nil
	default:
		return nil, errors.New("invalid log type")
	}
}


func displayMessage(h ILogHandler, msg string) {
	fmt.Println(h.GetLevel(), " ", msg)
}

func doChaning() ILogHandler {
	infoHandler := getDefaultInfoLogHandler()
	errorHandler := getDefaultErrorLogHandler()
	debugHandler := getDefaultDebugLogHandler()

	infoHandler.SetNext(errorHandler)
	errorHandler.SetNext(debugHandler)
	debugHandler.SetNext(nil)
	return infoHandler
}
