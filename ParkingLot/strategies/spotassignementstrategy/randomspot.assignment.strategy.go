package spotassignementstrategy

import (
	"pattern/ParkingLot/entities"
	"pattern/ParkingLot/enums"
)

type RandomSpotAssignmentStrategy struct{}

func (s *RandomSpotAssignmentStrategy) AssignSpot(gateid uint16, vehicletype enums.VehicleType) *entities.ParkingSpot {
	// get []spots for the floor
	// traverse the list of spots
	// check if for the given vehicle type if there is any spot available
	// assign the very first spot that is found
	// else no spot available, return nil
	return nil
}
