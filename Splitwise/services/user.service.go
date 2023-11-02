package services

import (
	"pattern/Splitwise/entities"
	"pattern/Splitwise/repository"
)

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService() *UserService {
	return &UserService{
		userRepository: repository.NewUserRepository(),
	}
}

func (us *UserService) RegisterUser(id int, name, pwd, phone string) entities.IUser {
	user := entities.NewUser(id, name, pwd, phone)
	return us.userRepository.CreateUser(user)
}
