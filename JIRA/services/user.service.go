package services

import (
	"pattern/JIRA/dtos"
	"pattern/JIRA/entities"
	"sync"
)

var userServiceOnce sync.Once
var userService *UserService

// UserService struct
type UserService struct {
	userRepository repositories.UserRepository
}

// NewUserService function
func NewUserService() *UserService {
	userServiceOnce.Do(func() {
		userService = &UserService{
			userRepository: repositories.NewUserRepository(),
		}
	})
	return userService
}

// CreateUser function with input as user dto
func (us *UserService) CreateUser(createUserDto dtos.CreateUserDto) (*entities.User, error) {
	return us.userRepository.CreateUser(createUserDto)
}

// GetUser function with input as GetUserDto
func (us *UserService) GetUser(getUserDto dtos.GetUserDto) (*entities.User, error) {
	return us.userRepository.GetUser(getUserDto.GetUserID())
}
