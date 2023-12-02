package entities

import (
	"errors"

	"google.golang.org/genproto/googleapis/type/date"
)

type IUser interface {
	GetId() int
	GetName() string
	GetDob() *date.Date
	GetIsVaccinated() bool
	GetBookingsForADay(bookingDate *date.Date) IBooking
	GetBookings() map[*date.Date]IBooking
	AddBoookingForADay(bookingDate *date.Date, booking IBooking) error
	RemoveBookingForADay(bookingDate *date.Date, booking IBooking) error
	SetId(id int) IUser
	SetName(name string) IUser
	SetDob(dob *date.Date) IUser
	SetIsVaccinated(isVaccinated bool) IUser
	SetBookings(bookings map[*date.Date]IBooking) IUser
}

type user struct {
	id           int
	name         string
	dob          *date.Date
	isVaccinated bool
	bookings     map[*date.Date]IBooking
}

func NewUser(id int, name string, dob *date.Date, isVaccinated bool, bookings map[*date.Date]IBooking) IUser {
	return &user{
		id:           id,
		name:         name,
		dob:          dob,
		isVaccinated: isVaccinated,
		bookings:     bookings,
	}
}

func (u *user) GetId() int {
	return u.id
}

func (u *user) GetName() string {
	return u.name
}

func (u *user) GetDob() *date.Date {
	return u.dob
}

func (u *user) GetIsVaccinated() bool {
	return u.isVaccinated
}

func (u *user) GetBookingsForADay(bookingDate *date.Date) IBooking {
	return u.bookings[bookingDate]
}

func (u *user) GetBookings() map[*date.Date]IBooking {
	return u.bookings
}

func (u *user) AddBoookingForADay(bookingDate *date.Date, booking IBooking) error {
	if _, ok := u.bookings[bookingDate]; ok {
		u.bookings[bookingDate] = booking
		return nil
	}
	return errors.New("Booking date not found")
}

func (u *user) RemoveBookingForADay(bookingDate *date.Date, booking IBooking) error {
	if _, ok := u.bookings[bookingDate]; ok {
		delete(u.bookings, bookingDate)
	}
	return errors.New("Booking date not found")
}

func (u *user) SetId(id int) IUser {
	u.id = id
	return u
}

func (u *user) SetName(name string) IUser {
	u.name = name
	return u
}

func (u *user) SetDob(dob *date.Date) IUser {
	u.dob = dob
	return u
}

func (u *user) SetIsVaccinated(isVaccinated bool) IUser {
	u.isVaccinated = isVaccinated
	return u
}

func (u *user) SetBookings(bookings map[*date.Date]IBooking) IUser {
	u.bookings = bookings
	return u
}
