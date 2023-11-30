package entities

type bothVegNonVegHotel struct {
	hotelName string
	menus     []string
}

func NewBothVegNonVegHotel() *bothVegNonVegHotel {
	return &bothVegNonVegHotel{
		menus: make([]string, 0),
	}
}

func (h *bothVegNonVegHotel) SetHotelName(name string) *bothVegNonVegHotel {
	h.hotelName = name
	return h
}

func (h *bothVegNonVegHotel) GetHotelName() string {
	return h.hotelName
}

func (h *bothVegNonVegHotel) AddItemToMenus(item string) *bothVegNonVegHotel {
	h.menus = append(h.menus, item)
	return h
}

func (h *bothVegNonVegHotel) AddMenu(menu []string) *bothVegNonVegHotel {
	h.menus = menu
	return h
}

func (h *bothVegNonVegHotel) GetMenu() []string {
	return h.menus
}
