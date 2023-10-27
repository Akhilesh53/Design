package dao

import (
	"errors"
	"pattern/Waletsystem/entities"
)

type IUserDAO interface {
	FindUserById(id int) (entities.IUser, error)
	SetUserById(id int, user entities.IUser) error
	UpdateUserById(id int, user entities.IUser) IUserDAO
	GetAllUsers() []entities.IUser
}

func NewUserDAO() IUserDAO {
	return &userDAO{
		userIdDetailsMap: make(map[int]entities.IUser),
	}
}

type userDAO struct {
	userIdDetailsMap map[int]entities.IUser
}

func (u *userDAO) FindUserById(id int) (entities.IUser, error) {
	if user, ok := u.userIdDetailsMap[id]; ok {
		return user, nil
	}
	return nil, errors.New("user not found")
}

func (u *userDAO) SetUserById(id int, user entities.IUser) error {
	if _, ok := u.userIdDetailsMap[id]; ok {
		return errors.New("user Id already exists")
	}
	u.userIdDetailsMap[id] = user
	return nil
}

func (u *userDAO) UpdateUserById(id int, user entities.IUser) IUserDAO {
	u.userIdDetailsMap[id] = user
	return u
}

func (u *userDAO) GetAllUsers() []entities.IUser {
	var users []entities.IUser
	for _, user := range u.userIdDetailsMap {
		users = append(users, user)
	}
	return users
}
