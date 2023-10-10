package models

type IShowSeatType interface {
}

type ShowSeatType struct {
	Show     IShow
	SeatType SeatType
	Price    float64
}
