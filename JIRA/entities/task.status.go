package entities

type TaskStatus int

const (
	OPEN TaskStatus = iota
	IN_PROGRESS
	COMPLETED
	DISABLED
)
