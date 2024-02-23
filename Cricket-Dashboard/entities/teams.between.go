package entities

type TeamsBetween struct {
	team1 Team
	team2 Team
}

func MatchTeamsBetween(team1 Team, team2 Team) TeamsBetween {
	return TeamsBetween{team1, team2}
}
