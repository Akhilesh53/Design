package repositories

import "pattern/ParkingLot/entities"

type TicketRepository struct {
	ticketRepo map[int]*entities.Ticket
}

func NewTicketRepository() *TicketRepository {
	return &TicketRepository{
		ticketRepo: make(map[int]*entities.Ticket),
	}
}
