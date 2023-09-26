package logger

import (
	"errors"
)

type ILogHandler interface {
	GetNext() ILogHandler
	SetNext(next ILogHandler)
	GetLevel() string
	SetLevel(level LogType)
	LogMessage(logType LogType, msg string)

	AddObserver(ISinkObserver)
	RemoveObserver(ISinkObserver)
	Notify(mssg string)
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
	h.Notify(msg)
}

func doChaining() ILogHandler {
	infoHandler := getDefaultInfoLogHandler()
	errorHandler := getDefaultErrorLogHandler()
	debugHandler := getDefaultDebugLogHandler()

	fileObserver := getFileObserver()
	consoleObserver := getConsoleObserver()

	infoHandler.AddObserver(fileObserver)
	infoHandler.AddObserver(consoleObserver)

	errorHandler.AddObserver(fileObserver)

	debugHandler.AddObserver(consoleObserver)
	infoHandler.SetNext(errorHandler)
	errorHandler.SetNext(debugHandler)
	debugHandler.SetNext(nil)
	return infoHandler
}
