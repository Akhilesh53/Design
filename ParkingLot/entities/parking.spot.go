package entities

import "pattern/ParkingLot/enums"

type ParkingSpot struct {
	spotId           uint16
	spotStatus       enums.SpotStatus
	parkingSpotFloor uint16
	parkingSpotType  enums.SpotType
}

func NewParkingSpot() *ParkingSpot {
	return &ParkingSpot{}
}
