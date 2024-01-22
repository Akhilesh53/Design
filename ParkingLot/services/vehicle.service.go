package services

import (
	"pattern/ParkingLot/entities"
	"pattern/ParkingLot/enums"
)

type VehicleService struct {
	// vehicle repository
	// vehicle ticket repository
}

func NewVehicleService() *VehicleService {
	return &VehicleService{}
}

func (vs *VehicleService) FindVehicleById(vechilenumber string) (vehicle *entities.Vehicle) {
	//get vechile from vehicle repository
	// if found, return vehicle
	// else return nil
	return nil
}

func (vs *VehicleService) RegisterVehicle(vehiclenumber string, vechiletype enums.VehicleType) (*entities.Vehicle, error) {
	// store vehicle in vehicle repository
	return nil, nil
}

func (vs *VehicleService) FindTicketAssociatedWithVehicle(vehiclenumber string) (*entities.Ticket, error) {
	// get detail of tickeet associate with vehicle from vehicle ticket repo
	return nil, nil
}
