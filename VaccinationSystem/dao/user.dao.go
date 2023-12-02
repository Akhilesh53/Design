package dao

import "pattern/VaccinationSystem/entities"

type UserDao struct {
	userDao map[int]entities.IUser
}

func NewUserDao() *UserDao {
	return &UserDao{
		userDao: make(map[int]entities.IUser),
	}
}

func (ud *UserDao) AddUser(user entities.IUser) {
	ud.userDao[user.GetId()] = user
}

func (ud *UserDao) GetUser(id int) entities.IUser {
	return ud.userDao[id]
}

func (ud *UserDao) RemoveUser(id int) {
	delete(ud.userDao, id)
}

func (ud *UserDao) GetAllUsers() map[int]entities.IUser {
	return ud.userDao
}
