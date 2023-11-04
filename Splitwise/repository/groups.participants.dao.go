package repository

type GroupParticipantsRepository interface {
	AddMemberToGroup(groupId int, userId int)
}

type groupParticipantsRepository struct {
	groupParticipantsRepository map[int][]int
}

func NewGroupParticipantsRepository() GroupParticipantsRepository {
	return &groupParticipantsRepository{
		groupParticipantsRepository: make(map[int][]int),
	}
}

func (g *groupParticipantsRepository) AddMemberToGroup(groupId int, userId int) {
	if _, ok := g.groupParticipantsRepository[groupId]; !ok {
		g.groupParticipantsRepository[groupId] = make([]int, 0)
	}
	g.groupParticipantsRepository[groupId] = append(g.groupParticipantsRepository[groupId], userId)
}
