package entities

import (
	"pattern/ParkingLot/enums"
	"time"
)

type Bill struct {
	billingId  uint16
	ticket     *Ticket
	exitTime   time.Time
	payment    *Payment
	gate       *Gate
	billStatus enums.BillStatus
}

func NewBill() *Bill {
	return &Bill{}
}

func GenerateBill() {
	return
}

func (b *Bill) HavingTicket(ticket *Ticket) *Bill {
	b.ticket = ticket
	return b
}

func (b *Bill) WithPayment(payment *Payment) *Bill {
	b.payment = payment
	return b
}

func (b *Bill) AtGate(gate *Gate) *Bill {
	b.gate = gate
	return b
}

func (b *Bill) Generate(status enums.BillStatus) *Bill {
	b.billStatus = status
	b.exitTime = time.Now()
	return b
}
