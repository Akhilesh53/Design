package interfaces

import "pattern/VendingMachine/constants"

type Item interface {
	GetCode() int
	SetCode(int)

	GetCount() int
	SetCount(int)

	SetType(constants.ItemType)
	GetType() constants.ItemType

	GetPrice() int
	SetPrice(price int)
}
