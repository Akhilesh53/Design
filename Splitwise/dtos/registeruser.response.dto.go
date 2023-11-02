package dtos

import "pattern/Splitwise/entities"

type RegisterUserResponseDto struct {
	user entities.IUser
	err  error
}

func NewRegisterUserResponseDto() *RegisterUserResponseDto {
	return &RegisterUserResponseDto{}
}

func NewRegisterUserResponseDtoWithUser(user entities.IUser) *RegisterUserResponseDto {
	return &RegisterUserResponseDto{
		user: user,
	}
}

func (rurd *RegisterUserResponseDto) GetUser() entities.IUser {
	return rurd.user
}

func (rurd *RegisterUserResponseDto) SetUser(user entities.IUser) *RegisterUserResponseDto {
	rurd.user = user
	return rurd
}

func (rurd *RegisterUserResponseDto) GetErr() error {
	return rurd.err
}

func (rurd *RegisterUserResponseDto) SetErr(err error) *RegisterUserResponseDto {
	rurd.err = err
	return rurd
}
