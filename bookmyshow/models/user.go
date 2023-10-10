package models

type IUser interface {
}

type User struct {
	Name string
	// address, list<bookings>, previous bookings
}
