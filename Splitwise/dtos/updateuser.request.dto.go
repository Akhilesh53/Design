package dtos

type UpdateUserRequestDto struct {
	userId   string
	password string
}

func NewUpdateUserRequestDto(userId string, pwd string) *UpdateUserRequestDto {
	return &UpdateUserRequestDto{
		userId:   userId,
		password: pwd,
	}
}

func (u *UpdateUserRequestDto) GetUserId() string {
	return u.userId
}

func (u *UpdateUserRequestDto) GetUserPaasword() string {
	return u.password
}

func (u *UpdateUserRequestDto) SetUserId(userId string) *UpdateUserRequestDto {
	u.userId = userId
	return u
}

func (u *UpdateUserRequestDto) SetUserPassword(pwd string) *UpdateUserRequestDto {
	u.password = pwd
	return u
}
