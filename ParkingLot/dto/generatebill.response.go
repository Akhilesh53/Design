package dto

import (
	"pattern/ParkingLot/entities"
	"pattern/ParkingLot/exceptions"
)

type GenerateBillResponseDto struct {
	responseStatus exceptions.Error
	bill           *entities.Bill
}

func GenerateBill() *GenerateBillResponseDto {
	return &GenerateBillResponseDto{}
}

func (d *GenerateBillResponseDto) SetResponseStatus(status exceptions.Error) *GenerateBillResponseDto {
	d.responseStatus = status
	return d
}

func (d *GenerateBillResponseDto) SetBill(bill *entities.Bill) *GenerateBillResponseDto {
	d.bill = bill
	return d
}
