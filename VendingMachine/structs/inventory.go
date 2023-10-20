package structs

import (
	"errors"
	"pattern/VendingMachine/interfaces"
)

type inventory struct {
	items map[int]interfaces.Item
}

func GetNewInventory() interfaces.Inventory {
	return &inventory{
		items: make(map[int]interfaces.Item),
	}
}

func (i *inventory) GetItems() map[int]interfaces.Item {
	return i.items
}

func (i *inventory) GetItem(productCode int) interfaces.Item {
	if _, ok := i.items[productCode]; !ok {
		return getDefaultItem()
	}
	return i.items[productCode]
}

func (i *inventory) SetItem(productCode int, item interfaces.Item) {
	i.items[productCode] = item
}

func (i *inventory) HasProduct(productCode int) bool {
	if _, ok := i.items[productCode]; !ok {
		return ok
	}
	return i.GetItem(productCode).GetCount() > 0
}

func (i *inventory) DispenseProduct(productCode int) error {
	if !i.HasProduct(productCode) {
		return errors.New(" product not found")
	}
	i.GetItem(productCode).SetCount(i.GetItem(productCode).GetCount() - 1)
	return nil
}

func (i *inventory) AddItemInInventory(productCode int, item interfaces.Item) {
	if !i.HasProduct(productCode) {
		i.SetItem(productCode, item)
	}
	i.GetItem(productCode).SetCount(i.GetItem(productCode).GetCount() + item.GetCount())
}

func (i *inventory) GetItemDetailsFromInventory(productCode int) interfaces.Item {
	return i.GetItem(productCode)
}

func (i *inventory) DecrementProductCount(productCode int) {
	i.GetItem(productCode).SetCount(i.items[productCode].GetCount() - 1)
}
