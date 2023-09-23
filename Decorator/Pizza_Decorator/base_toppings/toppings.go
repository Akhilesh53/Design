package base_toppings

type ExtraCheeze struct {
	pizza IPizza
}

func AddExtraCheeze(basePizza IPizza) IPizza {
	return &ExtraCheeze{pizza: basePizza}
}

func (e *ExtraCheeze) GetPrice() float64 {
	return e.pizza.GetPrice() + 50
}

type Mushrooms struct {
	pizza IPizza
}

func AddMushrooms(basePizza IPizza) IPizza {
	return &ExtraCheeze{pizza: basePizza}
}

func (e *Mushrooms) GetPrice() float64 {
	return e.pizza.GetPrice() + 25
}
