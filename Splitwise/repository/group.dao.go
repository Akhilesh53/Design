package repository

import (
	"errors"
	"pattern/Splitwise/entities"
)

type GroupRepository interface {
	CreateGroup(group entities.IGroup) entities.IGroup
}

type groupRepository struct {
	groupRepository map[int]entities.IGroup
}

func NewGroupRepository() GroupRepository {
	return &groupRepository{
		groupRepository: make(map[int]entities.IGroup),
	}
}

func (g *groupRepository) CreateGroup(group entities.IGroup) entities.IGroup {
	g.addGroup(group)
	return group
}

func (g *groupRepository) addGroup(group entities.IGroup) {
	g.groupRepository[group.GetId()] = group
}

func (g *groupRepository) GetGroup(id int) (entities.IGroup, error) {
	if group, ok := g.groupRepository[id]; ok {
		return group, nil
	}
	return nil, errors.New("group not found")
}
