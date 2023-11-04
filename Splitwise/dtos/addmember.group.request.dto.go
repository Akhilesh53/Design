package dtos

type AddMemberGroupRequestDto struct {
	AdminUserId     string `json:"admin_user_id"`
	GroupId         string `json:"group_id"`
	UserIdToBeAdded string `json:"user_id_to_be_added"`
}

func NewAddMemberGroupRequestDto() *AddMemberGroupRequestDto {
	return &AddMemberGroupRequestDto{}
}

func (dto *AddMemberGroupRequestDto) SetAdminUserId(adminUserId string) *AddMemberGroupRequestDto {
	dto.AdminUserId = adminUserId
	return dto
}

func (dto *AddMemberGroupRequestDto) SetGroupId(groupId string) *AddMemberGroupRequestDto {
	dto.GroupId = groupId
	return dto
}

func (dto *AddMemberGroupRequestDto) SetUserIdToBeAdded(userIdToBeAdded string) *AddMemberGroupRequestDto {
	dto.UserIdToBeAdded = userIdToBeAdded
	return dto
}

func (dto *AddMemberGroupRequestDto) GetAdminUserId() string {
	return dto.AdminUserId
}

func (dto *AddMemberGroupRequestDto) GetGroupId() string {
	return dto.GroupId
}

func (dto *AddMemberGroupRequestDto) GetUserIdToBeAdded() string {
	return dto.UserIdToBeAdded
}