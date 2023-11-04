package dtos

import "pattern/Splitwise/entities"

type GetUserGroupsResponseDto struct {
	Groups []entities.IGroup
	err error
}

func NewGetUserGroupsResponseDto() *GetUserGroupsResponseDto {
	return &GetUserGroupsResponseDto{
		Groups: make([]entities.IGroup, 0),
	}
}

func NewGetUserGroupsResponseDtoWithGroups(groups []entities.IGroup) *GetUserGroupsResponseDto {
	return &GetUserGroupsResponseDto{
		Groups: groups,
	}
}

func (gugr *GetUserGroupsResponseDto) SetGroups(groups []entities.IGroup) *GetUserGroupsResponseDto {
	gugr.Groups = groups
	return gugr
}

func (gugr *GetUserGroupsResponseDto) GetGroups() []entities.IGroup {
	return gugr.Groups
}

func (gugr *GetUserGroupsResponseDto) AddGroup(group entities.IGroup) *GetUserGroupsResponseDto {
	gugr.Groups = append(gugr.Groups, group)
	return gugr
}

func (gugr *GetUserGroupsResponseDto) SetErr(err error) *GetUserGroupsResponseDto {
	gugr.err = err
	return gugr
}

func (gugr *GetUserGroupsResponseDto) GetErr() error {
	return gugr.err
}
