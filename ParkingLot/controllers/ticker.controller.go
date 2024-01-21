package controllers

import (
	"pattern/ParkingLot/dto"
	"pattern/ParkingLot/exceptions"
	"pattern/ParkingLot/services"
	"pattern/ParkingLot/strategies/spotassignementstrategy"
)

type TicketController struct {
	ticketService *services.TicketService
}

func NewTicketController() *TicketController {
	return &TicketController{
		ticketService: services.NewTicketService(new(spotassignementstrategy.RandomSpotAssignmentStrategy)),
	}
}

func (tc *TicketController) GenerateTicket(request *dto.GenerateTicketRequestDto) *dto.GenerateTicketResponseDto {
	response := &dto.GenerateTicketResponseDto{}
	ticket, err := tc.ticketService.GenerateTicket(request.GetVehicleNumber(), request.GetVehicleType(), request.GetGateId())
	if err != nil {
		return response.SetTicket(nil).SetResponseStatus(exceptions.ErrGeneratingTicketFailure)
	}
	return response.SetTicket(ticket).SetResponseStatus(exceptions.ErrGeneratingTicketSuccess)
}
