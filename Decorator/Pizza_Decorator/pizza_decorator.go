package pizzadecorator

import (
	"fmt"
	pd "pattern/Decorator/Pizza_Decorator/base_toppings"
)

/*
  Design Question:
  ===================

  Consider a coffee machine/Pizza Shop. People wants to order  coffee/Pizza.
  It can be a normal one or one with the toppings/ additional items.
  Add necessary items what ever is selected and return the cost for the same.

  Design system for the same.
*/

func CreatePizzaDecorators() {

	basePizza, err := pd.GetPizza("Veggie Delight")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(" Base Pizza Price :: ", basePizza.GetPrice())

	basePizza = pd.AddExtraCheeze(basePizza)

	fmt.Println(" Price of piza with extra cheese :: ", basePizza.GetPrice())

}
