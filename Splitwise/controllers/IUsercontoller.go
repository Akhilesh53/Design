package controllers

import "pattern/Splitwise/dtos"

type IUserContoller interface {
	RegisterUser(req dtos.RegisterUserRequestDto) *dtos.RegisterUserResponseDto
	UpdateUser(req *dtos.UpdateUserRequestDto) *dtos.RegisterUserResponseDto
	GetUser(req *dtos.GetUserRequestDto) *dtos.RegisterUserResponseDto
}
