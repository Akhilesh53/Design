package entities

type IUser interface {
	GetId() int
	GetName() string
	GetPhoneNumber() string
	GetPassword() string
	SetId(int) IUser
	SetName(string) IUser
	SetPhoneNumber(string) IUser
	SetPassword(string) IUser
}

func NewUser(id int, name string, phoneNumber string, password string) IUser {
	return &user{
		id:          id,
		name:        name,
		phoneNumber: phoneNumber,
		password:    password,
	}
}

type user struct {
	id          int
	name        string
	phoneNumber string
	password    string
}

func (u *user) GetId() int {
	return u.id
}

func (u *user) GetName() string {
	return u.name
}

func (u *user) GetPhoneNumber() string {
	return u.phoneNumber
}

func (u *user) GetPassword() string {
	return u.password
}

func (u *user) SetId(id int) IUser {
	u.id = id
	return u
}

func (u *user) SetName(name string) IUser {
	u.name = name
	return u
}

func (u *user) SetPhoneNumber(phoneNumber string) IUser {
	u.phoneNumber = phoneNumber
	return u
}

func (u *user) SetPassword(password string) IUser {
	u.password = password
	return u
}
