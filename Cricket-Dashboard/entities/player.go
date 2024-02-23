package entities

import "pattern/Cricket-Dashboard/enums"

type Player struct {
	Person
	playerResponsibilty enums.PlayerResponsibilty
	playerType          enums.PlayerType
	isSubstitue         bool
}

func NewPlayer(name string, playerType enums.PlayerType, isSubstitue bool) Player {
	return Player{
		Person:      NewPerson(name),
		playerType:  playerType,
		isSubstitue: isSubstitue,
	}
}
