package dtos

import "pattern/Splitwise/entities"

type AddGroupResponseDto struct {
	group entities.IGroup
	err   error
}

func NewAddGroupResponseDtoWithGroup(group entities.IGroup) *AddGroupResponseDto {
	return &AddGroupResponseDto{
		group: group,
	}
}

func NewAddGroupResponseDto() *AddGroupResponseDto {
	return &AddGroupResponseDto{}
}

func (a *AddGroupResponseDto) GetGroup() entities.IGroup {
	return a.group
}

func (a *AddGroupResponseDto) SetGroup(group entities.IGroup) *AddGroupResponseDto {
	a.group = group
	return a
}

func (a *AddGroupResponseDto) GetErr() error {
	return a.err
}

func (a *AddGroupResponseDto) SetErr(err error) *AddGroupResponseDto {
	a.err = err
	return a
}
