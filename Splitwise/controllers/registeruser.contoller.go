package controllers

import (
	"pattern/Splitwise/dtos"
	"pattern/Splitwise/services"
)

type IUserContoller interface {
}

type userContoller struct {
	userservice services.UserService // for controllers to work, it need an instance of service
}

func NewUserContoller() IUserContoller {
	return &userContoller{}
}

func (uc *userContoller) RegisterUser(req dtos.RegisterUserRequestDto) *dtos.RegisterUserResponseDto {
	user := uc.userservice.RegisterUser(101, req.GetName(), req.GetPassword(), req.GetPhone())
	return dtos.NewRegisterUserResponseDtoWithUser(user)
}
