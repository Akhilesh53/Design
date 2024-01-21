package services

import "pattern/ParkingLot/entities"

type GateService struct {
	//gate repository
}

func NewGateService() *GateService {
	return &GateService{}
}

func (gs *GateService) GetGate(gateId uint16) *entities.Gate {
	// get gate details by gate id
	// if found return gate
	// else return nil
	return nil
}
