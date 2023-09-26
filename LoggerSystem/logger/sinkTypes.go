package logger

type SinkType int

const (
	FILE SinkType = iota
	CONSOLE
)

func (s SinkType) String() string {
	switch s {
	case FILE:
		return "FILE"
	case CONSOLE:
		return "CONSOLE"
	default:
		return "Invalid Sink Type"
	}
}
