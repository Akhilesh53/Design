package entities

type WaiterImpl struct {
}

func NewWaiterImpl() Waiter {
	return &WaiterImpl{}
}

func (impl *WaiterImpl) GetVegMenu() []string {
	vegRest := NewVegHotel()
	return vegRest.GetMenu()
}

func (impl *WaiterImpl) GetNonVegMenu() []string {
	vegRest := NewNonVegHotel()
	return vegRest.GetMenu()
}

func (impl *WaiterImpl) GetBothVegNonVegMenu() []string {
	vegRest := NewBothVegNonVegHotel()
	return vegRest.GetMenu()
}
