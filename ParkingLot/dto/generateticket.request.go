package dto

import (
	"pattern/ParkingLot/enums"
)

type GenerateTicketRequestDto struct {
	vehicleNumber string
	vehicleType   enums.VehicleType
	gateId        uint16
}

func NewGenerateTicketRequestDto(vehiclenumber string, vehicletype enums.VehicleType, gateid uint16) *GenerateTicketRequestDto {
	return &GenerateTicketRequestDto{
		vehicleNumber: vehiclenumber,
		vehicleType:   vehicletype,
		gateId:        gateid,
	}
}

func (d *GenerateTicketRequestDto) GetVehicleNumber() string {
	return d.vehicleNumber
}

func (d *GenerateTicketRequestDto) GetVehicleType() enums.VehicleType {
	return d.vehicleType
}

func (d *GenerateTicketRequestDto) GetGateId() uint16 {
	return d.gateId
}
