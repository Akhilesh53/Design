package entities

type TaskType int

const (
	BUG TaskType = iota
	FEATURE
	IMPROVEMENT
	STORY
)
