package services

import (
	"pattern/VaccinationSystem/dao"
	"pattern/VaccinationSystem/entities"

	"google.golang.org/genproto/googleapis/type/date"
)

type UserService struct {
	userDao dao.UserDao
}

func NewUserService() *UserService {
	return &UserService{
		userDao: *dao.NewUserDao(),
	}
}

func (us *UserService) AddUser(userId int, name string, dob *date.Date, isVaccinated bool) {
	user := entities.NewUser(userId, name, dob, isVaccinated, make(map[string]entities.IBooking))
	us.userDao.AddUser(user)
}

func (us *UserService) GetUser(userId int) entities.IUser {
	return us.userDao.GetUser(userId)
}

func (us *UserService) RemoveUser(userId int) {
	us.userDao.RemoveUser(userId)
}
