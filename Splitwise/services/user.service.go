package services

import (
	"pattern/Splitwise/entities"
	"pattern/Splitwise/repository"
	"sync"
)

var once sync.Once
var userRepo repository.UserRepository
var userServiceInstance *UserService

func init() {
	userRepo = repository.NewUserRepository()
}

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService() *UserService {
	once.Do(func() {
		userServiceInstance = &UserService{
			userRepository: userRepo,
		}
	})
	return userServiceInstance
}

func (us *UserService) RegisterUser(id int, name, pwd, phone string) entities.IUser {
	user := entities.NewUser(id, name, phone, pwd)
	return us.userRepository.CreateUser(user)
}

func (us *UserService) UpdateUser(id int, pwd string) (entities.IUser, error) {
	return us.userRepository.UpdateUserByParam("password", pwd, id)
}
