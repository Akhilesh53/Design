package entities

import "pattern/Cricket-Dashboard/enums"

type Team struct {
	teamId           uint16
	teamName         string
	teamMembers      []Player
	representingType enums.TeamRepresentingType
	representingName string
	supportStaff     []SupportStaff
}

func NewTeam(teamId uint16, teamName string, representingType enums.TeamRepresentingType, represntingName string) Team {
	return Team{teamId: teamId, teamName: teamName}
}

func (t *Team) AddPlayer(player Player) {
	t.teamMembers = append(t.teamMembers, player)
}
