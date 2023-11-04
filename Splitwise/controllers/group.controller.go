package controllers

import (
	"pattern/Splitwise/dtos"
	"pattern/Splitwise/services"
	"strconv"
	"sync"
)

type groupController struct {
	groupService     *services.GroupService
	groupIdGenerator func() int
}

var groupService *services.GroupService
var groupControllerInstance *groupController
var groupControllerOnce sync.Once

func init() {
	groupService = services.NewGroupService()
}

func generateGroupId() func() int {
	id := 0

	return func() int {
		id++
		return id
	}
}

func NewGroupController() IGroupController {
	groupControllerOnce.Do(func() {
		groupControllerInstance = &groupController{
			groupService:     groupService,
			groupIdGenerator: generateGroupId(),
		}
	})
	return groupControllerInstance
}

func (gc *groupController) AddGroup(req *dtos.AddGroupRequestDto) *dtos.AddGroupResponseDto {
	userId, _ := strconv.Atoi(req.GetUserId())
	user, err := userService.GetUser(userId)
	if err != nil {
		return dtos.NewAddGroupResponseDto().SetErr(err)
	}
	group := gc.groupService.CreateGroup(gc.groupIdGenerator(), req.GetName(), req.GetDescription(), user)
	return dtos.NewAddGroupResponseDtoWithGroup(group)
}

func (gc *groupController) AddMember(req *dtos.AddMemberGroupRequestDto) *dtos.AddGroupResponseDto {
	//adminUserId, _ := strconv.Atoi(req.GetAdminUserId())
	groupId, _ := strconv.Atoi(req.GetGroupId())
	userIdToBeAdded, _ := strconv.Atoi(req.GetUserIdToBeAdded())

	//todo:
	//check admin user is actually admin of group

	group, err := groupService.GetGroup(groupId)
	if err != nil {
		return dtos.NewAddGroupResponseDto().SetErr(err)
	}

	userToBeAdded, err := userService.GetUser(userIdToBeAdded)
	if err != nil {
		return dtos.NewAddGroupResponseDto().SetErr(err)
	}

	resGroup, err := gc.groupService.AddMember(group, userToBeAdded)
	if err != nil {
		return dtos.NewAddGroupResponseDto().SetErr(err)
	}
	return dtos.NewAddGroupResponseDtoWithGroup(resGroup)
}

func (gc *groupController) GetGroupsForUser(req *dtos.GetUserGroupsRequestDto) *dtos.GetUserGroupsResponseDto {
	userId, _ := strconv.Atoi(req.GetUserId())
	groups, err := gc.groupService.GetGroupsForUser(userId)
	if err != nil {
		return dtos.NewGetUserGroupsResponseDto().SetErr(err)
	}
	return dtos.NewGetUserGroupsResponseDtoWithGroups(groups)
}
