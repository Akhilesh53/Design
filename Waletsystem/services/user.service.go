package services

import (
	"pattern/Waletsystem/dao"
	"pattern/Waletsystem/entities"
)

type IUserService interface {
	CreateUser(name string) entities.IUser
	GetUserByUserId(userId int) (entities.IUser, error)
	GetAllUsers() []entities.IUser
}

type userService struct {
	userDAO         dao.IUserDAO
	userIdGenerator entities.IUserIdGenerator
}

func NewUserService() IUserService {
	return &userService{
		userDAO:         dao.NewUserDAO(),
		userIdGenerator: entities.NewUserIdGenerator(),
	}
}

func (u *userService) CreateUser(name string) entities.IUser {
	user := entities.NewUser(name, u.userIdGenerator.Generate())
	u.userDAO.SetUserById(user.GetUserId(), user)
	return user
}

func (u *userService) GetUserByUserId(userId int) (entities.IUser, error) {
	return u.userDAO.FindUserById(userId)
}

func (u *userService) GetAllUsers() []entities.IUser {
	users := make([]entities.IUser, 0)
	users = append(users, u.userDAO.GetAllUsers()...)
	return users
}
