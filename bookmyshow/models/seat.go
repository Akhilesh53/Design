package models

type ISeat interface {
	GetSeatType() SeatType
	SetSeatType(SeatType) ISeat

	GetSeatRow() string
	SetSeatRow(string) ISeat

	GetSeatNumber() string
	SetSeatNumber(string) ISeat
}

type Seat struct {
	seatRow    string
	seatNumber string
	seatType   SeatType
}

func (s *Seat) GetSeatType() SeatType {
	return s.seatType
}

func (s *Seat) SetSeatType(seatType SeatType) ISeat {
	s.seatType = seatType
	return s
}

func (s *Seat) GetSeatRow() string {
	return s.seatRow
}

func (s *Seat) SetSeatRow(row string) ISeat {
	s.seatRow = row
	return s
}

func (s *Seat) GetSeatNumber() string {
	return s.seatNumber
}

func (s *Seat) SetSeatNumber(number string) ISeat {
	s.seatNumber = number
	return s
}
