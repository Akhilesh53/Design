package repositories

import "pattern/ParkingLot/entities"

type BillRepository struct {
	billRepo map[uint16]*entities.Bill
}

func NewBillRepository() *BillRepository {
	return &BillRepository{
		billRepo: make(map[uint16]*entities.Bill),
	}
}
