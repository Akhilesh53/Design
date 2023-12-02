package entities

import (
	"errors"
)

type IVC interface {
	GetId() int
	GetDistrict() string
	GetAddress() string
	GetBookingsForADay(bookingDate string) []IUser
	GetBookings() map[string][]IUser
	GetCapacityForADay(bookingDate string) int
	GetCapacity() map[string]int
	AddBookingForADay(bookingDate string, booking IUser) error
	RemoveBookingForADay(bookingDate string, booking IUser) error
	SetId(id int) IVC
	SetDistrict(district string) IVC
	SetAddress(address string) IVC
	SetBookings(bookings map[string][]IUser) IVC
	SetCapacity(capacity map[string]int) IVC
	SetCapacityForADay(bookingDate string, capacity int) IVC
}

type vaccinationCentre struct {
	id       int
	district string
	address  string
	bookings map[string][]IUser
	capacity map[string]int
}

func NewVaccinationCentre(id int, district string, address string) IVC {
	return &vaccinationCentre{
		id:       id,
		district: district,
		address:  address,
		bookings: make(map[string][]IUser),
		capacity: make(map[string]int),
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

func (vc *vaccinationCentre) GetBookingsForADay(bookingDate string) []IUser {
	return vc.bookings[bookingDate]
}

func (vc *vaccinationCentre) GetBookings() map[string][]IUser {
	return vc.bookings
}

func (vc *vaccinationCentre) GetCapacityForADay(bookingDate string) int {
	return vc.capacity[bookingDate]
}

func (vc *vaccinationCentre) GetCapacity() map[string]int {
	return vc.capacity
}

func (vc *vaccinationCentre) AddBookingForADay(bookingDate string, booking IUser) error {
	if vc.capacity[bookingDate] == 0 {
		return errors.New("No capacity for the given date")
	}
	vc.bookings[bookingDate] = append(vc.bookings[bookingDate], booking)
	vc.capacity[bookingDate]--
	return nil
}

func (vc *vaccinationCentre) RemoveBookingForADay(bookingDate string, booking IUser) error {
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

func (vc *vaccinationCentre) SetBookings(bookings map[string][]IUser) IVC {
	vc.bookings = bookings
	return vc
}

func (vc *vaccinationCentre) SetCapacity(capacity map[string]int) IVC {
	vc.capacity = capacity
	return vc
}

func (vc *vaccinationCentre) SetCapacityForADay(bookingDate string, capacity int) IVC {
	vc.capacity[bookingDate] = capacity
	return vc
}
