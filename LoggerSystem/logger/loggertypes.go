package logger

type LogType int

const (
	INFO LogType = iota
	ERROR
	DEBUG
)

func (l LogType) String() string {
	switch l {
	case INFO:
		return "INFO"
	case ERROR:
		return "ERROR"
	case DEBUG:
		return "DEBUG"
	default:
		return "Invalid LogType"
	}
}
