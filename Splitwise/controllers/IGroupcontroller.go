package controllers

import "pattern/Splitwise/dtos"

type IGroupController interface {
	AddGroup(req *dtos.AddGroupRequestDto) *dtos.AddGroupResponseDto
	AddMember(req *dtos.AddMemberGroupRequestDto) *dtos.AddGroupResponseDto
}
