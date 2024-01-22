package controllers

import (
	"pattern/ParkingLot/dto"
	"pattern/ParkingLot/exceptions"
	"pattern/ParkingLot/services"
)

type BillController struct {
	billService *services.BillService
}

func NewBillController() *BillController {
	return &BillController{}
}

func (bc *BillController) GenerateBill(request *dto.GenerateBillRequestDto) *dto.GenerateBillResponseDto {
	//get ticket from vehicle ticket mapping
	response := &dto.GenerateBillResponseDto{}
	bill, err := bc.billService.GenerateBill(request.GetVehicleNumber(), request.GetGateId())
	if err != nil {
		return response.SetBill(nil).SetResponseStatus(exceptions.ErrGeneratingBillFailure)
	}
	return response.SetBill(bill).SetResponseStatus(exceptions.ErrGeneratingBillSuccess)
}
