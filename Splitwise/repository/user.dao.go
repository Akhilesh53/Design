package repository

import (
	"errors"
	"pattern/Splitwise/entities"
)

// table <user id, name , phone number, password>

type UserRepository interface {
	CreateUser(user entities.IUser) (entities.IUser)
	GetUser(id int) (entities.IUser, error)
	DeleteUser(id int)
	UpdateUser(id int, name string, phoneNumber string, password string) (entities.IUser, error)
	GetAllUsers() []entities.IUser
	UpdateUserByParam(paramName, val string, id int) (entities.IUser, error)
}

type userRepository struct {
	userRepository map[int]entities.IUser
}

func NewUserRepository() UserRepository {
	return &userRepository{
		userRepository: make(map[int]entities.IUser),
	}
}

func (u *userRepository) CreateUser(user entities.IUser) (entities.IUser) {
	u.addUser(user)
	return user
}

func (u *userRepository) addUser(user entities.IUser)  {
	u.userRepository[user.GetId()] = user
}

func (u *userRepository) GetUser(id int) (entities.IUser, error) {
	if user, ok := u.userRepository[id]; ok {
		return user, nil
	}
	return nil, errors.New("user not found")
}

func (u *userRepository) DeleteUser(id int) {
	if _, ok := u.userRepository[id]; ok {
		delete(u.userRepository, id)
		return
	}
}

func (u *userRepository) UpdateUser(id int, name string, phoneNumber string, password string) (entities.IUser, error) {
	if user, ok := u.userRepository[id]; ok {
		user.SetName(name).SetPhoneNumber(phoneNumber).SetPassword(password)
		return user, nil
	}
	return nil, errors.New("user not found")
}

func (u *userRepository) GetAllUsers() []entities.IUser {
	var users []entities.IUser
	for _, user := range u.userRepository {
		users = append(users, user)
	}
	return users
}

func (u *userRepository) UpdateUserByParam(paramName, val string, id int) (entities.IUser, error) {
	if user, ok := u.userRepository[id]; ok {
		switch paramName {
		case "name":
			user.SetName(val)
		case "phoneNumber":
			user.SetPhoneNumber(val)
		case "password":
			user.SetPassword(val)
		}
		return user, nil
	}
	return nil, errors.New("user not found")
}
