package classes

type GameStatus int

const (
	IN_PROGRESS GameStatus = iota
	DRAW
	ENDED
)
