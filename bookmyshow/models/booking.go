package models

type IBooking interface{

}

type Booking struct {
	User IUser
	Show IShow
	ShowSeats []IShowSeat
	Amount float64
	Payment string //todo: there will be seaparate payment object
	BookingStatus BookingStatus
}