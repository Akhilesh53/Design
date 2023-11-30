package entities

type Waiter interface {
	GetVegMenu() []string
	GetNonVegMenu() []string
	GetBothVegNonVegMenu() []string
}
