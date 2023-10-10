package models

type BookingStatus int

const (
	RESERVED BookingStatus = iota
	BOOKED
	CANCELLED
)

func (b BookingStatus) String() string {
	switch b {
	case RESERVED:
		return "reserved"
	case CANCELLED:
		return "cancelled"
	case BOOKED:
		return "booked"
	default:
		return "unknown"
	}
}
