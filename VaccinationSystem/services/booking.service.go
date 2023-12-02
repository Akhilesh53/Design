package services

import (
	"pattern/VaccinationSystem/dao"
	"pattern/VaccinationSystem/entities"

	"google.golang.org/genproto/googleapis/type/date"
)

type bookingService struct {
	bookingDao dao.BookingDao
}

func NewBookingService() *bookingService {
	return &bookingService{
		bookingDao: *dao.NewBookingDao(),
	}
}

func (bs *bookingService) AddBooking(bookingId int, bookingDate *date.Date, user entities.IUser, vc entities.IVC) {
	booking := entities.NewBooking(bookingId, user, vc, bookingDate)
	bs.bookingDao.AddBooking(booking)
}

func (bs *bookingService) GetBooking(bookingId int) entities.IBooking {
	return bs.bookingDao.GetBooking(bookingId)
}

