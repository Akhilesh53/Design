package models

type IShowSeat interface {
	
}

type ShowSeat struct {
	Show         IShow
	Seat         ISeat
	IsReserved   bool
	ShowSeatType IShowSeatType
}
