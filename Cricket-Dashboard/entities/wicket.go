package entities

import "pattern/Cricket-Dashboard/enums"

type Wicket struct {
	wicketType enums.WicketType
	playerOut  Player
	bowledBy   Player
	caughtBy   Player
	runoutBy   Player
	stumpedBy  Player
}
