package dtos

type GetUserGroupsRequestDto struct {
	UserId string
}

func NewGetUserGroupsRequestDto() *GetUserGroupsRequestDto {
	return &GetUserGroupsRequestDto{}
}

func NewGetUserGroupsRequestDtoWithUserId(userId string) *GetUserGroupsRequestDto {
	return &GetUserGroupsRequestDto{
		UserId: userId,
	}
}

func (gugr *GetUserGroupsRequestDto) SetUserId(userId string) *GetUserGroupsRequestDto {
	gugr.UserId = userId
	return gugr
}

func (gugr *GetUserGroupsRequestDto) GetUserId() string {
	return gugr.UserId
}
