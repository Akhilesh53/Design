package base_toppings

import "errors"

type IPizza interface {
	GetPrice() float64
}

func GetPizza(pizza string) (IPizza, error) {
	switch pizza {
	case "Marghereta":
		return GetMarghereta(), nil
	case "Veggie Delight":
		return GetVeggieDelhight(), nil
	default:
		return nil, errors.New("Pizza type not found: ")
	}
}

type Marghereta struct {
	price float64
}

func GetMarghereta() *Marghereta {
	return &Marghereta{
		price: 100,
	}
}

func (m *Marghereta) GetPrice() float64 {
	return m.price
}

type VeggieDelight struct {
	price float64
}

func GetVeggieDelhight() *VeggieDelight {
	return &VeggieDelight{
		price: 200,
	}
}

func (m *VeggieDelight) GetPrice() float64 {
	return m.price
}
