package controllers

import (
	"pattern/Splitwise/dtos"
	"pattern/Splitwise/services"
	"sync"
)

var once sync.Once
var userService *services.UserService
var userContollerInstance IUserContoller

func init() {
	userService = services.NewUserService()
}

type userContoller struct {
	userservice *services.UserService // for controllers to work, it need an instance of service
}

func NewUserContoller() IUserContoller {
	once.Do(func() {
		userContollerInstance = &userContoller{
			userservice: userService,
		}
	})
	return userContollerInstance
}

func (uc *userContoller) RegisterUser(req dtos.RegisterUserRequestDto) *dtos.RegisterUserResponseDto {
	user := uc.userservice.RegisterUser(101, req.GetName(), req.GetPassword(), req.GetPhone())
	return dtos.NewRegisterUserResponseDtoWithUser(user)
}
