package main

import (
	"fmt"
	"pattern/FacadePattern/entities"
)

func main() {
	waiter := entities.NewWaiterImpl()

	vegMenu := waiter.GetVegMenu()
	fmt.Println(vegMenu)

	nonVegMenu := waiter.GetNonVegMenu()
	fmt.Println(nonVegMenu)

	bothMenu := waiter.GetBothVegNonVegMenu()
	fmt.Println(bothMenu)
}
