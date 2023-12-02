package dao

import "pattern/VaccinationSystem/entities"

type BookingDao struct {
	bookingDao map[int]entities.IBooking
}

func NewBookingDao() *BookingDao {
	return &BookingDao{
		bookingDao: make(map[int]entities.IBooking),
	}
}

func (bd *BookingDao) AddBooking(booking entities.IBooking) {
	bd.bookingDao[booking.GetId()] = booking
}

func (bd *BookingDao) GetBooking(id int) entities.IBooking {
	return bd.bookingDao[id]
}

func (bd *BookingDao) RemoveBooking(id int) {
	delete(bd.bookingDao, id)
}

func (bd *BookingDao) GetAllBookings() map[int]entities.IBooking {
	return bd.bookingDao
}

