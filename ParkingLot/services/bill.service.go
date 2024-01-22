package services

import (
	"pattern/ParkingLot/entities"
	"pattern/ParkingLot/enums"
	"pattern/ParkingLot/repositories"
)

type BillService struct {
	ticketService  *TicketService
	paymentService *PaymentService
	gateService    *GateService
	billRepos      *repositories.BillRepository
	vehicleService *VehicleService
}

func NewBillService() *BillService {
	return &BillService{}
}

func (bs *BillService) GenerateBill(vehiclenumber string, gateId uint16) (*entities.Bill, error) {
	// get ticket details from vehicle ticket mapping
	ticket, err := bs.vehicleService.FindTicketAssociatedWithVehicle(vehiclenumber)
	if err != nil {
		return nil, err
	}

	gate := bs.gateService.GetGate(gateId)

	// DO PAYMENT
	payment, err := bs.paymentService.MakePayment(enums.CREDIT_CARD, 100)
	if err != nil {
		return nil, err
	}
	bill := entities.NewBill().HavingTicket(ticket).AtGate(gate).WithPayment(payment).Generate(enums.SUCCESSFUL)
	return bill, nil
}
