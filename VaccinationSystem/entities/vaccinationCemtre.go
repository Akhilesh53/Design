package entities

import (
	"errors"

	"google.golang.org/genproto/googleapis/type/date"
)

type IVC interface {
	GetId() int
	GetDistrict() string
	GetAddress() string
	GetBookingsForADay(bookingDate *date.Date) []IUser
	GetBookings() map[*date.Date][]IUser
	GetCapacityForADay(bookingDate *date.Date) int
	GetCapacity() map[*date.Date]int
	AddBookingForADay(bookingDate *date.Date, booking IUser) error
	RemoveBookingForADay(bookingDate *date.Date, booking IUser) error
	SetId(id int) IVC
	SetDistrict(district string) IVC
	SetAddress(address string) IVC
	SetBookings(bookings map[*date.Date][]IUser) IVC
	SetCapacity(capacity map[*date.Date]int) IVC
	SetCapacityForADay(bookingDate *date.Date, capacity int) IVC
}

type vaccinationCentre struct {
	id       int
	district string
	address  string
	bookings map[*date.Date][]IUser
	capacity map[*date.Date]int
}

func NewVaccinationCentre(id int, district string, address string) IVC {
	return &vaccinationCentre{
		id:       id,
		district: district,
		address:  address,
		bookings: make(map[*date.Date][]IUser),
		capacity: make(map[*date.Date]int),
	}
}

func (vc *vaccinationCentre) GetId() int {
	return vc.id
}

func (vc *vaccinationCentre) GetDistrict() string {
	return vc.district
}

func (vc *vaccinationCentre) GetAddress() string {
	return vc.address
}

func (vc *vaccinationCentre) GetBookingsForADay(bookingDate *date.Date) []IUser {
	return vc.bookings[bookingDate]
}

func (vc *vaccinationCentre) GetBookings() map[*date.Date][]IUser {
	return vc.bookings
}

func (vc *vaccinationCentre) GetCapacityForADay(bookingDate *date.Date) int {
	return vc.capacity[bookingDate]
}

func (vc *vaccinationCentre) GetCapacity() map[*date.Date]int {
	return vc.capacity
}

func (vc *vaccinationCentre) AddBookingForADay(bookingDate *date.Date, booking IUser) error {
	if vc.capacity[bookingDate] == 0 {
		return errors.New("No capacity for the given date")
	}
	vc.bookings[bookingDate] = append(vc.bookings[bookingDate], booking)
	vc.capacity[bookingDate]--
	return nil
}

func (vc *vaccinationCentre) RemoveBookingForADay(bookingDate *date.Date, booking IUser) error {
	if _, ok := vc.bookings[bookingDate]; ok {
		delete(vc.bookings, bookingDate)
	}
	return errors.New("Booking date not found")
}

func (vc *vaccinationCentre) SetId(id int) IVC {
	vc.id = id
	return vc
}

func (vc *vaccinationCentre) SetDistrict(district string) IVC {
	vc.district = district
	return vc
}

func (vc *vaccinationCentre) SetAddress(address string) IVC {
	vc.address = address
	return vc
}

func (vc *vaccinationCentre) SetBookings(bookings map[*date.Date][]IUser) IVC {
	vc.bookings = bookings
	return vc
}

func (vc *vaccinationCentre) SetCapacity(capacity map[*date.Date]int) IVC {
	vc.capacity = capacity
	return vc
}

func (vc *vaccinationCentre) SetCapacityForADay(bookingDate *date.Date, capacity int) IVC {
	vc.capacity[bookingDate] = capacity
	return vc
}