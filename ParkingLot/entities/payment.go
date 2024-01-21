package entities

import "pattern/ParkingLot/enums"

type Payment struct {
	paymentId uint16
	paymentMode enums.PaymentMode 
	amount float64
	paymentStatus enums.BillStatus
}