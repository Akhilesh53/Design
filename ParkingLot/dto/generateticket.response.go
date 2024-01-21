package dto

import (
	"pattern/ParkingLot/entities"
	"pattern/ParkingLot/exceptions"
)

type GenerateTicketResponseDto struct {
	responseStatus exceptions.Error
	ticket         *entities.Ticket
}

func NewGenerateTicketResponseDto(responseStatus exceptions.Error, ticket *entities.Ticket) *GenerateTicketResponseDto {
	return &GenerateTicketResponseDto{
		responseStatus: responseStatus,
		ticket:         ticket,
	}
}

func NewDefaultGenerateTicketResponseDto() *GenerateTicketResponseDto {
	return &GenerateTicketResponseDto{}
}

func (d *GenerateTicketResponseDto) SetResponseStatus(status exceptions.Error) *GenerateTicketResponseDto {
	d.responseStatus = status
	return d
}

func (d *GenerateTicketResponseDto) SetTicket(ticket *entities.Ticket) *GenerateTicketResponseDto {
	d.ticket = ticket
	return d
}
