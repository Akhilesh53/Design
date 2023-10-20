package constants

type ItemType int

const (
	PEPSI ItemType = iota
	COKE
	JUICE
	ITEM_TYPE_NONE ItemType = -1
)
