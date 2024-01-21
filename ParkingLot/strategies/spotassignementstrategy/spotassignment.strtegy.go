package spotassignementstrategy

import (
	"pattern/ParkingLot/entities"
	"pattern/ParkingLot/enums"
)

type SpotAssignmentStrategy interface {
	AssignSpot(gateid uint16, vehicletype enums.VehicleType) *entities.ParkingSpot
}
