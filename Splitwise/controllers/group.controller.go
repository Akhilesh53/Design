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
