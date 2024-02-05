package enums

type Status int

const (
	MOVING Status = iota
	IDLE
	STOPPED
	NOT_AVAILABLE

	OPEN
	CLOSED
)
