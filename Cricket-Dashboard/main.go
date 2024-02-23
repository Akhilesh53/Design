package main

import (
	"pattern/Cricket-Dashboard/entities"
	"pattern/Cricket-Dashboard/enums"
	matchtype "pattern/Cricket-Dashboard/enums"
	playertype "pattern/Cricket-Dashboard/enums"
	teamrepresnting "pattern/Cricket-Dashboard/enums"
	"time"
)

func main() {

	//create tournament
	tournament := entities.NewTournament(123, "Indian Premier League", time.Now().AddDate(2, 2, 5), time.Now().AddDate(2, 2, 25))

	// add teams in the tournament
	team1 := entities.NewTeam(456, "Chennai Super Kings", teamrepresnting.STATE, "CSK")
	team2 := entities.NewTeam(454, "Mumbai Indians", teamrepresnting.STATE, "MI")

	// add players to the team
	player1 := entities.NewPlayer("MS Dhoni", playertype.WICKET_KEEPER, false)
	team1.AddPlayer(player1)

	player2 := entities.NewPlayer("Rohit Sharma", playertype.BATSMAN, false)
	team2.AddPlayer(player2)

	tournament.AddTeam(team1)
	tournament.AddTeam(team2)

	//schedule a match
	match := entities.NewMatch(789, matchtype.IPL)
	match.PlayingBetween(entities.MatchTeamsBetween(team1, team2))
	match.MatchToBePlayedAt(entities.NewVenue(entities.NewAddressInfo("somehwre in chennai", "chennai", "245678", "Tamil Nadu", "India"), "chinnaswamy stadium", 1234))
	// add further  details related to match

	tournament.AddMatch(match)

	innings1 := entities.NewInnings(1, match.GetMatchId())

	over1 := entities.NewOver(1)
	ball1 := entities.NewBall(1).OfBallType(enums.NORMAL).GivingRun(enums.ONE)
	// add other details to vall like bowled by and faced by

	over1.AddBall(ball1)

	innings1.AddOver(over1)
	// similarly add details  for the scores of the players who are playing
	// on the field
}
