package repository

import "pattern/Cricket-Dashboard/entities"

var (
	TournamentRepository map[int]entities.Tournament
	MatchRepository      map[int]entities.Match
	TeamRepository       map[int]entities.Team
	PlayerRepository     map[int]entities.Player
	ScorerRepository     map[int]entities.Scorer
	VenueMap             map[int]entities.Venue
	ScoreCardRepository  map[int]entities.ScoreCard
)
