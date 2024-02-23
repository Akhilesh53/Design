package enums

type MatchStatus int

const (
	PLAYING MatchStatus = iota
	YET_TO_BE_PLAY
	TIED
	COMPLETED
)
