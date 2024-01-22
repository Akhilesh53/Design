package repositories

import "pattern/ParkingLot/entities"

type VehicleTicketRepository struct {
	vehicleTicketRepo map[string]*entities.Ticket
}

func NewVehicleTicketRepository() *VehicleTicketRepository {
	return &VehicleTicketRepository{
		vehicleTicketRepo: make(map[string]*entities.Ticket),
	}
}
