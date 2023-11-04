package controllers

import (
	"pattern/Splitwise/dtos"
	"pattern/Splitwise/services"
	"strconv"
	"sync"
)

type groupController struct {
	groupService *services.GroupService
}

var groupService *services.GroupService
var groupControllerInstance *groupController
var groupControllerOnce sync.Once

func init() {
	groupService = services.NewGroupService()
}

func NewGroupController() IGroupController {
	groupControllerOnce.Do(func() {
		groupControllerInstance = &groupController{
			groupService: groupService,
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
	group := gc.groupService.CreateGroup(201, req.GetName(), req.GetDescription(), user)
	return dtos.NewAddGroupResponseDtoWithGroup(group)

}
