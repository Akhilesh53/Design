package controllers

import (
	"pattern/JIRA/dtos"
	"pattern/JIRA/entities"
	"pattern/JIRA/services"
	"sync"
)

var usercontrollerOnce sync.Once
var usercontroller *UserController

// UserController struct
type UserController struct {
	userService *services.UserService
}

// NewUserController function
func NewUserController() *UserController {
	usercontrollerOnce.Do(func() {
		usercontroller = &UserController{
			userService: services.NewUserService(),
		}
	})
	return usercontroller
}

// CreateUser function with input as user dto
func (uc *UserController) CreateUser(createUserDto *dtos.CreateUserDto) (*entities.User, error) {
	return uc.userService.CreateUser(createUserDto.GetName(), createUserDto.GetEmail())
}

// GetUser function with input as GetUserDto
func (uc *UserController) GetUser(getUserDto *dtos.GetUserDto) (*entities.User, error) {
	return uc.userService.GetUser(getUserDto.GetUserID())
}
