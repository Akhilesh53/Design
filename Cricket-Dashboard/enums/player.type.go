package enums

type PlayerType int

const (
	BATSMAN PlayerType = iota
	BOWLER
	ALL_ROUNDER
	WICKET_KEEPER
)
