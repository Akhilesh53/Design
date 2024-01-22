package dto

type GenerateBillRequestDto struct {
	vehiclenumber string
	gateId        uint16
}

func (r *GenerateBillRequestDto) GetVehicleNumber() string { return r.vehiclenumber }

func (r *GenerateBillRequestDto) GetGateId() uint16 { return r.gateId }
