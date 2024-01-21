package entities

import "pattern/ParkingLot/enums"

type Gate struct {
	gateNumber uint16
	gateType   enums.GateType
}

func NewGate(gateNumber uint16, gateType enums.GateType) *Gate {
	return &Gate{
		gateNumber: gateNumber,
		gateType:   gateType,
	}
}
