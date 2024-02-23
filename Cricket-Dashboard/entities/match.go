package entities

import (
	"pattern/Cricket-Dashboard/enums"
	"time"
)

type Match struct {
	macthId      uint16
	tournament   Tournament
	teamsBetween TeamsBetween
	matchType    enums.MatchType
	venue        Venue
	winnerTeam   Team
	startTime    time.Time
	endTime      time.Time
	matchStatus  enums.MatchStatus
	innings      map[int]Innings
	scorers      []Scorer
}

func NewMatch(macthId uint16, matchType enums.MatchType) Match {
	return Match{
		macthId:   macthId,
		matchType: matchType,
	}
}

func (m Match) GetMatchId() uint16 {
	return m.macthId
}

func (m Match) PlayingBetween(teamsBetween TeamsBetween) Match {
	m.teamsBetween = teamsBetween
	return m
}

func (m Match) MatchToBePlayedAt(venue Venue) Match {
	m.venue = venue
	return m
}
