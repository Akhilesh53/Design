package entities

import "time"

type Tournament struct {
	tounamentId       uint32
	tournamentName    string
	particpatingTeams []Team
	scheduledMatches  []Match
	startDate         time.Time
	endDate           time.Time
	winner            Team
	runner            Team
}

func NewTournament(tounamentId uint32, tournamentName string, startDate, endDate time.Time) Tournament {
	return Tournament{
		tounamentId:       tounamentId,
		tournamentName:    tournamentName,
		startDate:         startDate,
		endDate:           endDate,
		particpatingTeams: make([]Team, 0),
		scheduledMatches:  make([]Match, 0),
	}
}

func (t *Tournament) AddTeam(team Team) {
	t.particpatingTeams = append(t.particpatingTeams, team)
}

func (t *Tournament) AddMatch(match Match) {
	match.tournament = *t
	t.scheduledMatches = append(t.scheduledMatches, match)
}
