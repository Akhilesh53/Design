package entities

import "pattern/Cricket-Dashboard/enums"

type Toss struct {
	tossedBy   Player
	askedTo    Player
	wonBy      Team
	tossAction enums.TossAction
}
