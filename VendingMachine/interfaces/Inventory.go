package interfaces

type Inventory interface {
	HasProduct(productCode int) bool
	DispenseProduct(productCode int) error
	AddItemInInventory(productCode int, item Item)
	GetItemDetailsFromInventory(productCode int) Item
	GetItems() map[int]Item
	GetItem(productCode int) Item
	SetItem(productCode int, item Item)

	//todo:
	// removeitem( productcode int)
	// updatecompletedetailsofitem(productcode int)
}
