package entities

import "google.golang.org/genproto/googleapis/type/date"

type IBooking interface {
	GetId() int
	GetIUser() IUser
	GetIVC() IVC
	GetDate() *date.Date
	SetId(int) IBooking
	SetIUser(IUser) IBooking
	SetIVC(IVC) IBooking
	SetDate(*date.Date) IBooking
}

type booking struct {
	id int
	IUser
	IVC
	date *date.Date
}


func NewBooking(id int, user IUser, vc IVC, date *date.Date) IBooking {
	return &booking{
		id: id,
		IUser: user,
		IVC: vc,
		date: date,
	}
}

func(b *booking) GetId() int {
	return b.id
}

func(b *booking) GetIUser() IUser {
	return b.IUser
}

func(b *booking) GetIVC() IVC {
	return b.IVC
}

func(b *booking) GetDate() *date.Date {
	return b.date
}

func(b *booking) SetId(id int) IBooking {
	b.id = id
	return b
}

func(b *booking) SetIUser(user IUser) IBooking {
	b.IUser = user
	return b
}

func(b *booking) SetIVC(vc IVC) IBooking {
	b.IVC = vc
	return b
}

func(b *booking) SetDate(date *date.Date) IBooking {
	b.date = date
	return b
}
