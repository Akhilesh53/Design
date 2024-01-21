package entities

import "pattern/ParkingLot/enums"

type ParkingFloor struct {
	floorId        uint16
	parkingSpots   map[enums.SpotType][]ParkingSpot
	availableSpots uint16
	status         enums.SpotStatus
}

func NewParkingFloor(floorId uint16) *ParkingFloor {
	return &ParkingFloor{
		floorId: floorId,
	}
}
