package services

import (
	"pattern/Splitwise/entities"
	"pattern/Splitwise/repository"
	"sync"
)

type GroupService struct {
	groupRepository             repository.GroupRepository
	groupParticipantsRepository repository.GroupParticipantsRepository
}

var groupServiceInstance *GroupService
var groupServiceOnce sync.Once
var groupRepo repository.GroupRepository
var groupParticipantsRepo repository.GroupParticipantsRepository

func init() {
	groupRepo = repository.NewGroupRepository()
	groupParticipantsRepo = repository.NewGroupParticipantsRepository()
}

func NewGroupService() *GroupService {
	groupServiceOnce.Do(func() {
		groupServiceInstance = &GroupService{
			groupRepository:             groupRepo,
			groupParticipantsRepository: groupParticipantsRepo,
		}
	})
	return groupServiceInstance
}

func (gs *GroupService) CreateGroup(groupId int, groupName, groupDescription string, createdBy entities.IUser) entities.IGroup {
	group := entities.NewDefaultGroup().SetId(groupId).SetName(groupName).SetDescription(groupDescription).SetCreatedBy(createdBy)
	return gs.groupRepository.CreateGroup(group)
}

func (gs *GroupService) GetGroup(groupId int) (entities.IGroup, error) {
	return gs.groupRepository.GetGroup(groupId)
}

func (gs *GroupService) AddMember(group entities.IGroup, user entities.IUser) (entities.IGroup, error) {
	group.AddParticipant(user)
	gs.groupParticipantsRepository.AddMemberToGroup(group.GetId(), user.GetId())
	return group, nil
}
