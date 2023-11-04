package repository

import "errors"

type GroupParticipantsRepository interface {
	AddMemberToGroup(groupId int, userId int)
	GetGroupIdsForUserId(userId int) ([]int, error)
}

type groupParticipantsRepository struct {
	groupParticipantsRepository map[int][]int //map<groupId, []userId>
	participantGroupRepository  map[int][]int //map<userId, []groupId>
}

func NewGroupParticipantsRepository() GroupParticipantsRepository {
	return &groupParticipantsRepository{
		groupParticipantsRepository: make(map[int][]int),
		participantGroupRepository:  make(map[int][]int),
	}
}

func (g *groupParticipantsRepository) AddMemberToGroup(groupId int, userId int) {
	if _, ok := g.groupParticipantsRepository[groupId]; !ok {
		g.groupParticipantsRepository[groupId] = make([]int, 0)
	}
	g.groupParticipantsRepository[groupId] = append(g.groupParticipantsRepository[groupId], userId)

	if _, ok := g.participantGroupRepository[userId]; !ok {
		g.participantGroupRepository[userId] = make([]int, 0)
	}
	g.participantGroupRepository[userId] = append(g.participantGroupRepository[userId], groupId)
}

func (g *groupParticipantsRepository) GetGroupIdsForUserId(userId int) ([]int, error) {
	if _, ok := g.participantGroupRepository[userId]; !ok {
		return nil, errors.New("user not found")
	}
	return g.participantGroupRepository[userId], nil
}
