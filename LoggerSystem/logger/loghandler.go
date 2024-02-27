package logger

import (
	"errors"
)

type ILogHandler interface {
	GetNext() ILogHandler
	LogMessage(logType LogType, msg string)
	NotifySinkObservers(mssg string)
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
	//fmt.Println(h.GetLevel(), " ", msg)
	h.NotifySinkObservers(msg)
}

func doChaining() ILogHandler {
	infoHandler := getDefaultInfoLogHandler()
	errorHandler := getDefaultErrorLogHandler()
	debugHandler := getDefaultDebugLogHandler()

	fileObserver := getFileObserver()
	consoleObserver := getConsoleObserver()

	infoHandler.AddSinkObserver(fileObserver)
	infoHandler.AddSinkObserver(consoleObserver)

	errorHandler.AddSinkObserver(fileObserver)

	debugHandler.AddSinkObserver(consoleObserver)
	infoHandler.SetNext(errorHandler)
	errorHandler.SetNext(debugHandler)
	debugHandler.SetNext(nil)
	return infoHandler
}
