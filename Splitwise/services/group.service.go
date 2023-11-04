package services

import (
	"pattern/Splitwise/entities"
	"pattern/Splitwise/repository"
	"sync"
)

type GroupService struct {
	groupRepository repository.GroupRepository
}

var groupServiceInstance *GroupService
var groupServiceOnce sync.Once
var groupRepo repository.GroupRepository

func init() {
	groupRepo = repository.NewGroupRepository()
}

func NewGroupService() *GroupService {
	groupServiceOnce.Do(func() {
		groupServiceInstance = &GroupService{
			groupRepository: groupRepo,
		}
	})
	return groupServiceInstance
}

func (gs *GroupService) CreateGroup(groupId int, groupName, groupDescription string, createdBy entities.IUser) entities.IGroup {
	group := entities.NewDefaultGroup().SetId(groupId).SetName(groupName).SetDescription(groupDescription).SetCreatedBy(createdBy)
	return gs.groupRepository.CreateGroup(group)
}
