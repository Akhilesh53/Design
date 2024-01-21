package services

import (
	"errors"
	"pattern/ParkingLot/entities"
	"pattern/ParkingLot/enums"
	repo "pattern/ParkingLot/repositories"
	"pattern/ParkingLot/strategies/spotassignementstrategy"
	"time"
)

type TicketService struct {
	vehicleService         *VehicleService
	gateService            *GateService
	spotAssignmentStrategy spotassignementstrategy.SpotAssignmentStrategy
	ticketRepository       *repo.TicketRepository
}

func NewTicketService(spotAssignmentStrategy spotassignementstrategy.SpotAssignmentStrategy) *TicketService {
	return &TicketService{
		ticketRepository:       repo.NewTicketRepository(),
		vehicleService:         NewVehicleService(),
		gateService:            NewGateService(),
		spotAssignmentStrategy: spotAssignmentStrategy,
	}
}

func (ts *TicketService) GenerateTicket(vehiclenumber string, vehicleType enums.VehicleType, gateid uint16) (*entities.Ticket, error) {
	// get details for the vehicle
	vehicle := ts.vehicleService.FindVehicleById(vehiclenumber)
	var err error
	if vehicle == nil {
		vehicle, err = ts.vehicleService.RegisterVehicle(vehiclenumber, vehicleType)
	}

	if err != nil {
		return nil, errors.New("error registering vehicle")
	}

	// get details of the gate
	gate := ts.gateService.GetGate(gateid)
	if gate == nil {
		return nil, errors.New("gate not found")
	}

	// assign spot to the vehicle using spot assignemnt strategy
	spot := ts.spotAssignmentStrategy.AssignSpot(gateid, vehicleType)

	// now generate ticket and store it in repository
	ticket := entities.NewTicket().ForVehicle(vehicle).FromGate(gate).HavingEntryTime(time.Now()).WithParkingSpot(spot).Build()
	return ticket, nil
}
