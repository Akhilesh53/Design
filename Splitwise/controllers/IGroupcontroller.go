package controllers

import "pattern/Splitwise/dtos"

type IGroupController interface {
	AddGroup(req *dtos.AddGroupRequestDto) *dtos.AddGroupResponseDto
}
