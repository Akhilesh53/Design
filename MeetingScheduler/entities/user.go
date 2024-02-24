package entities

import "fmt"

type User struct {
	userId uint16
	name   string
	email  string
	phone  string
}

type Host struct {
	User
}

type Attendee struct {
	User
}

func NewUser(userId uint16, name string, email string, phone string) *User {
	return &User{userId: userId, name: name, email: email, phone: phone}
}

func (u *User) SendNotificationToUser(desc string) {
	fmt.Println("sending notificaion for meeting :: ", desc)
}
