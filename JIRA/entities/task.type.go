package entities

type TaskType int

const (
	BUG TaskType = iota
	FEATURE
	IMPROVEMENT
	STORY
)

func (tt TaskType) String() string {
	return [...]string{"BUG", "FEATURE", "IMPROVEMENT", "STORY"}[tt]
}