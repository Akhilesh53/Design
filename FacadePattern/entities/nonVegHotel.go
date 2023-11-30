package entities

type nonVegHotel struct {
	hotelName string
	menus     []string
}

func NewNonVegHotel() *nonVegHotel {
	return &nonVegHotel{
		menus: make([]string, 0),
	}
}

func (h *nonVegHotel) SetHotelName(name string) *nonVegHotel {
	h.hotelName = name
	return h
}

func (h *nonVegHotel) GetHotelName() string {
	return h.hotelName
}

func (h *nonVegHotel) AddItemToMenus(item string) *nonVegHotel {
	h.menus = append(h.menus, item)
	return h
}

func (h *nonVegHotel) AddMenu(menu []string) *nonVegHotel {
	h.menus = menu
	return h
}

func (h *nonVegHotel) GetMenu() []string {
	return h.menus
}
