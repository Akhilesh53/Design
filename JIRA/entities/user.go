package entities

var userId uint16

func init() {
	userId = 0
}

func getUserId() uint16 {
	userId++
	return userId
}

type User struct {
	userid uint16
	name   string
	email  string
}

func NewUser(name string, email string) *User {
	return &User{
		userid: getUserId(),
		name:   name,
		email:  email,
	}
}

func (u *User) GetUserID() uint16 {
	return u.userid
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) GetEmail() string {
	return u.email
}

func (u *User) SetUserID(userid uint16) *User {
	u.userid = userid
	return u
}

func (u *User) SetName(name string) *User {
	u.name = name
	return u
}

func (u *User) SetEmail(email string) *User {
	u.email = email
	return u
}
