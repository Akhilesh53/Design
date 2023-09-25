package logger

import (
	"errors"
	"sync"
)

var LoggerInstance *Logger

type Logger struct {
	// since a singletone object has to be created
	// initialse a global logger variable

	// to maintain chain of responsibility b/w log types
	LogHandler ILogHandler
}

func GetLoggerInstance() (*Logger, error) {
	var mu sync.Mutex

	if LoggerInstance == nil {
		mu.Lock()
		defer mu.Unlock()
		LoggerInstance = &Logger{
			LogHandler: doChaning(),
		}
		return LoggerInstance, nil
	}
	return nil, errors.New("logger instance already created")
}

func (l *Logger) Info(mssg string) {
	l.createLog(INFO, mssg)
}

func (l *Logger) Error(mssg string) {
	l.createLog(ERROR, mssg)
}

func (l *Logger) Debug(mssg string) {
	l.createLog(DEBUG, mssg)
}

func (l *Logger) createLog(logType LogType, mssg string) {
	l.LogHandler.LogMessage(logType, mssg)
}
