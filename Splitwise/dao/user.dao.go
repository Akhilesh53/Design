package dao

import (
	"errors"
	"pattern/Splitwise/entities"
)

// table <user id, name , phone number, password>

type UserDao interface {
	CreateUser(id int, name string, phoneNumber string, password string) (entities.IUser, error)
	GetUser(id int) (entities.IUser, error)
	DeleteUser(id int)
	UpdateUser(id int, name string, phoneNumber string, password string) (entities.IUser, error)
	GetAllUsers() []entities.IUser
	UpdateUserByParam(paramName, val string, id int) (entities.IUser, error)
}

type userDao struct {
	userDao map[int]entities.IUser
}

func NewUserDao() UserDao {
	return &userDao{
		userDao: make(map[int]entities.IUser),
	}
}

func (u *userDao) CreateUser(id int, name string, phoneNumber string, password string) (entities.IUser, error) {
	user := entities.NewUser(id, name, phoneNumber, password)
	err := u.addUser(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userDao) addUser(user entities.IUser) error {
	if _, ok := u.userDao[user.GetId()]; ok {
		return errors.New("user already exists")
	}
	u.userDao[user.GetId()] = user
	return nil
}

func (u *userDao) GetUser(id int) (entities.IUser, error) {
	if user, ok := u.userDao[id]; ok {
		return user, nil
	}
	return nil, errors.New("user not found")
}

func (u *userDao) DeleteUser(id int) {
	if _, ok := u.userDao[id]; ok {
		delete(u.userDao, id)
		return
	}
}

func (u *userDao) UpdateUser(id int, name string, phoneNumber string, password string) (entities.IUser, error) {
	if user, ok := u.userDao[id]; ok {
		user.SetName(name)
		user.SetPhoneNumber(phoneNumber)
		user.SetPassword(password)
		return user, nil
	}
	return nil, errors.New("user not found")
}

func (u *userDao) GetAllUsers() []entities.IUser {
	var users []entities.IUser
	for _, user := range u.userDao {
		users = append(users, user)
	}
	return users
}

func (u *userDao) UpdateUserByParam(paramName, val string, id int) (entities.IUser, error) {
	if user, ok := u.userDao[id]; ok {
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
