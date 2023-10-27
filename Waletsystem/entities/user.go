package entities

type IUser interface {
	GetName() string
	SetName(name string) IUser
	GetUserId() int
	SetUserId(userId int) IUser
}

func NewUser(name string, userId int) IUser {
	return &user{
		name: name,
		id:   userId,
	}
}

type user struct {
	name string
	id   int
}

func (u *user) GetName() string {
	return u.name
}

func (u *user) SetName(name string) IUser {
	u.name = name
	return u
}

func (u *user) GetUserId() int {
	return u.id
}

func (u *user) SetUserId(userId int) IUser {
	u.id = userId
	return u
}
