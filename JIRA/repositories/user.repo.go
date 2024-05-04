package repositories

import (
	"errors"
	"pattern/JIRA/entities"
	"sync"
)

var userRepoOnce sync.Once
var userRepoInstance *UserRepo

type UserRepo struct {
	userMap map[uint16]*entities.User
}

func NewUserRepo() *UserRepo {
	userRepoOnce.Do(func() {
		userRepoInstance = &UserRepo{
			userMap: make(map[uint16]*entities.User),
		}
	})
	return userRepoInstance
}

func (r *UserRepo) Save(user *entities.User) (*entities.User, error) {
	r.userMap[user.GetUserID()] = user
	return user, nil
}

func (r *UserRepo) FindByID(userID uint16) (*entities.User, error) {
	if user, ok := r.userMap[userID]; ok {
		return user, nil
	}
	return nil, errors.New("user not found")
}

func (r *UserRepo) FindAll() []*entities.User {
	users := make([]*entities.User, 0)
	for _, user := range r.userMap {
		users = append(users, user)
	}
	return users
}

func (r *UserRepo) Delete(userID uint16) {
	delete(r.userMap, userID)
}

func (r *UserRepo) Update(user *entities.User) {
	r.userMap[user.GetUserID()] = user
}
