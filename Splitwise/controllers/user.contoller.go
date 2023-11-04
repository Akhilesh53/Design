package controllers

import (
	"pattern/Splitwise/dtos"
	"pattern/Splitwise/services"
	"strconv"
	"sync"
)

var once sync.Once
var userService *services.UserService
var userContollerInstance IUserContoller

func init() {
	userService = services.NewUserService()
}

func generateUserId() func() int {
	id := 0

	return func() int {
		id++
		return id
	}
}

type userContoller struct {
	userservice     *services.UserService // for controllers to work, it need an instance of service
	userIdGenerator func() int
}

func NewUserContoller() IUserContoller {
	once.Do(func() {
		userContollerInstance = &userContoller{
			userservice:     userService,
			userIdGenerator: generateUserId(),
		}
	})
	return userContollerInstance
}

func (uc *userContoller) RegisterUser(req dtos.RegisterUserRequestDto) *dtos.RegisterUserResponseDto {
	user := uc.userservice.RegisterUser(uc.userIdGenerator(), req.GetName(), req.GetPassword(), req.GetPhone())
	return dtos.NewRegisterUserResponseDtoWithUser(user)
}

func (uc *userContoller) UpdateUser(req *dtos.UpdateUserRequestDto) *dtos.RegisterUserResponseDto {
	userId, _ := strconv.Atoi(req.GetUserId())
	user, err := uc.userservice.UpdateUser(userId, req.GetUserPaasword())
	if err != nil {
		return dtos.NewRegisterUserResponseDto().SetErr(err)
	}
	return dtos.NewRegisterUserResponseDtoWithUser(user)
}

func (uc *userContoller) GetUser(req *dtos.GetUserRequestDto) *dtos.RegisterUserResponseDto {
	userId, _ := strconv.Atoi(req.GetUserId())
	user, err := uc.userservice.GetUser(userId)
	if err != nil {
		return dtos.NewRegisterUserResponseDto().SetErr(err)
	}
	return dtos.NewRegisterUserResponseDtoWithUser(user)
}
