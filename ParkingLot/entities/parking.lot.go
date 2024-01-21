package entities

import "pattern/ParkingLot/enums"

type ParkingLot struct {
	parkinglotId  uint16
	parkingFloors []ParkingFloor
	gates         []Gate
	status        enums.SpotStatus
	name          string
	address       string
}
