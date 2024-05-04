package services

import (
	"pattern/JIRA/dtos"
	"pattern/JIRA/entities"
	"pattern/JIRA/repositories"
	"sync"
)

var userServiceOnce sync.Once
var userService *UserService

// UserService struct
type UserService struct {
	userRepository *repositories.UserRepo
}

// NewUserService function
func NewUserService() *UserService {
	userServiceOnce.Do(func() {
		userService = &UserService{
			userRepository: repositories.NewUserRepo(),
		}
	})
	return userService
}

// CreateUser function with input as user dto
func (us *UserService) CreateUser(name, email string) (*entities.User, error) {
	user := entities.NewUser(name, email)
	return us.userRepository.Save(user)
}

// GetUser function with input as GetUserDto
func (us *UserService) GetUser(getUserDto dtos.GetUserDto) (*entities.User, error) {
	return us.userRepository.FindByID(getUserDto.GetUserID())
}
