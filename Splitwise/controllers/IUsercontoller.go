package controllers

import "pattern/Splitwise/dtos"

type IUserContoller interface {
	RegisterUser(req dtos.RegisterUserRequestDto) *dtos.RegisterUserResponseDto
}
