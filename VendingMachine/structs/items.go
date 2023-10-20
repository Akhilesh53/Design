package structs

import (
	"pattern/VendingMachine/constants"
	interfaces "pattern/VendingMachine/interfaces"
)

type item struct {
	code     int
	itemType constants.ItemType
	count    int
	price    int
}

func GetItem(itemType constants.ItemType, code, count, price int) interfaces.Item {
	return &item{
		code:     code,
		count:    count,
		itemType: itemType,
		price:    price,
	}
}

func getDefaultItem() interfaces.Item {
	return &item{
		code:     -1,
		count:    0,
		itemType: constants.ITEM_TYPE_NONE,
		price:    0,
	}
}

func (i *item) GetCode() int {
	return i.code
}

func (i *item) GetCount() int {
	return i.count
}

func (i *item) SetCode(code int) {
	i.code = code
}

func (i *item) SetCount(count int) {
	i.count = count
}

func (i *item) SetType(itemType constants.ItemType) {
	i.itemType = itemType
}

func (i *item) GetType() constants.ItemType {
	return i.itemType
}

func (i *item) GetPrice() int {
	return i.price
}

func (i *item) SetPrice(price int) {
	i.price = price
}
