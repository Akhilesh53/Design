package entities

import "pattern/Cricket-Dashboard/enums"

type PlayerScore struct {
	player     Player
	runsScored []uint16
	wicket     enums.WicketType
	bowler     Player
	isOut      bool
	fours      uint16
	sixes      uint16
}
