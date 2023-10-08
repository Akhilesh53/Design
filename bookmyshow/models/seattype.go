package models

type SeatType int

const (
	NORMAL SeatType = iota
	RECLINER
)

func (s SeatType) String() string {
	switch s {
	case NORMAL:
		return "Normal"
	case RECLINER:
		return "Recliner"
	default:
		return "Invalid Seat Type"
	}
}
