package entities

type vegHotel struct {
	hotelName string
	menus     []string
}

func NewVegHotel() *vegHotel {
	return &vegHotel{
		menus: make([]string, 0),
	}
}

func (h *vegHotel) SetHotelName(name string) *vegHotel {
	h.hotelName = name
	return h
}

func (h *vegHotel) GetHotelName() string {
	return h.hotelName
}

func (h *vegHotel) AddItemToMenus(item string) *vegHotel {
	h.menus = append(h.menus, item)
	return h
}

func (h *vegHotel) AddMenu(menu []string) *vegHotel {
	h.menus = menu
	return h
}

func (h *vegHotel) GetMenu() []string {
	return h.menus
}

