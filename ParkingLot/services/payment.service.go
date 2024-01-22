package services

import (
	"pattern/ParkingLot/entities"
	"pattern/ParkingLot/enums"
)

type PaymentService struct {
}

func (s *PaymentService) MakePayment(paymentmode enums.PaymentMode, amount float32) (*entities.Payment, error) {
	// do all payment related stuff
	return nil, nil
}
