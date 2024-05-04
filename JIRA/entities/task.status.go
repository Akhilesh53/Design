package entities

type TaskStatus int

const (
	OPEN TaskStatus = iota
	IN_PROGRESS
	COMPLETED
	DISABLED
)

func (ts TaskStatus) String() string {
	return [...]string{"OPEN", "IN_PROGRESS", "COMPLETED", "DISABLED"}[ts]
}
