package logger

import (
	"errors"
	"fmt"
)

type ISinkObserver interface {
	Update(msg string) // for writing log
}

func GetSinkObserver(sinkType SinkType) (ISinkObserver, error) {
	switch sinkType {
	case FILE:
		return getFileObserver(), nil
	default:
		return nil, errors.New(" unknown sink type ")
	}
}

type FileObserver struct {
	sinkType SinkType
}

func getFileObserver() *FileObserver {
	return &FileObserver{sinkType: FILE}
}

func (fo *FileObserver) Update(mssg string) {
	fmt.Println("Writing in ", fo.GetSinkType(), " ", mssg)
}

func (fo *FileObserver) GetSinkType() string {
	return fo.sinkType.String()
}

type ConsoleObserver struct {
	sinkType SinkType
}

func (co *ConsoleObserver) Update(mssg string) {
	fmt.Println("Writing in ", co.GetSinkType(), " ", mssg)
}

func getConsoleObserver() *ConsoleObserver {
	return &ConsoleObserver{sinkType: CONSOLE}
}

func (co *ConsoleObserver) GetSinkType() string {
	return co.sinkType.String()
}
