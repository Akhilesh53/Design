package vendingmachine

import (
	"fmt"
	"pattern/VendingMachine/constants"
	"pattern/VendingMachine/interfaces"
	"pattern/VendingMachine/structs"
)

func CallVendingMachine() {

	// fmt.Println("\nTest for less amount \n==================================")
	// fmt.Println("\nInventory details before")
	// VendingMachine := GetVendingMachine()
	// fmt.Println(VendingMachine.GetInventory().GetItemDetailsFromInventory(102), "\n", VendingMachine.GetMachineCoins())
	// TestForLessAmount(VendingMachine)
	// fmt.Println("\nInventory details after \n==================================")
	// fmt.Println(VendingMachine.GetInventory().GetItemDetailsFromInventory(102), "\n", VendingMachine.GetMachineCoins())

	fmt.Println("\nTest for exact amount \n==================================")
	fmt.Println("\nInventory details before")
	VendingMachine := GetVendingMachine()
	fmt.Println(VendingMachine.GetInventory().GetItemDetailsFromInventory(102),"\n",VendingMachine.GetMachineCoins())
	TestForExactAmount(VendingMachine)
	fmt.Println("\nInventory details after \n==================================")
	fmt.Println(VendingMachine.GetInventory().GetItemDetailsFromInventory(102), "\n", VendingMachine.GetMachineCoins())

	// fmt.Println("\nTest for more amount \n==================================")
	// fmt.Println("\nInventory details before ")
	// VendingMachine = GetVendingMachine()
	// fmt.Println(VendingMachine.GetInventory().GetItemDetailsFromInventory(102), "\n", VendingMachine.GetMachineCoins())
	// TestForMoreAmount(VendingMachine)
	// fmt.Println("\nInventory details after \n==================================")
	// fmt.Println(VendingMachine.GetInventory().GetItemDetailsFromInventory(102), "\n", VendingMachine.GetMachineCoins())

	// fmt.Println("\nTest for invalid product \n==================================")
	// fmt.Println("\nInventory details before ")
	// fmt.Println(VendingMachine.GetInventory().GetItemDetailsFromInventory(102), "\n", VendingMachine.GetMachineCoins())
	// VendingMachine = GetVendingMachine()
	// TestForInvalidProduct(VendingMachine)
	// fmt.Println("\nInventory details after \n==================================")
	// fmt.Println(VendingMachine.GetInventory().GetItemDetailsFromInventory(102), "\n", VendingMachine.GetMachineCoins())

}

func GetVendingMachine() interfaces.IVendingMachine {
	//create VendingMachine
	var VendingMachine = structs.NewVendingMachine()

	// create items
	var coke = structs.GetItem(constants.COKE, 101, 10, 25)
	var pepsi = structs.GetItem(constants.PEPSI, 102, 5, 30)
	var juice = structs.GetItem(constants.JUICE, 103, 1, 50)

	// add inventory to vending machine
	VendingMachine.GetInventory().AddItemInInventory(coke.GetCode(), coke)
	VendingMachine.GetInventory().AddItemInInventory(pepsi.GetCode(), pepsi)
	VendingMachine.GetInventory().AddItemInInventory(juice.GetCode(), juice)

	// add coins to vending machine
	VendingMachine.GetMachineCoins()[constants.Dime] = 10
	VendingMachine.GetMachineCoins()[constants.Nickel] = 5
	VendingMachine.GetMachineCoins()[constants.Penny] = 100
	VendingMachine.GetMachineCoins()[constants.Quarter] = 25

	return VendingMachine
}

// product price is 30 but inserted amount is 10
func TestForLessAmount(VendingMachine interfaces.IVendingMachine) {

	// click on insert coin button
	if err := VendingMachine.GetMachineState().ClickOnInsertCoinButton(VendingMachine); err != nil {
		fmt.Println(err)
		return
	}

	// insert coins
	if err := VendingMachine.GetMachineState().InsertCoin(VendingMachine, constants.Dime); err != nil {
		fmt.Println(err)
		return
	}

	// click on select product button
	if err := VendingMachine.GetMachineState().ClickOnInsertProdctCodeButton(VendingMachine); err != nil {
		fmt.Println(err)
		return
	}

	// select product
	if err := VendingMachine.GetMachineState().ChooseProduct(VendingMachine, 102); err != nil {
		fmt.Println(err)
		return
	}

	// dispense product
	if err := VendingMachine.GetMachineState().DispenseProduct(VendingMachine, 102); err != nil {
		fmt.Println(err)
		return
	}
}

// product price is 30 and inserted amount is 30
func TestForExactAmount(VendingMachine interfaces.IVendingMachine) {

	// click on insert coin button
	if err := VendingMachine.GetMachineState().ClickOnInsertCoinButton(VendingMachine); err != nil {
		fmt.Println(err)
		return
	}

	// insert coins
	if err := VendingMachine.GetMachineState().InsertCoin(VendingMachine, constants.Dime); err != nil {
		fmt.Println(err)
		return
	}
	if err := VendingMachine.GetMachineState().InsertCoin(VendingMachine, constants.Dime); err != nil {
		fmt.Println(err)
		return
	}
	if err := VendingMachine.GetMachineState().InsertCoin(VendingMachine, constants.Dime); err != nil {
		fmt.Println(err)
		return
	}

	// click on select product button
	if err := VendingMachine.GetMachineState().ClickOnInsertProdctCodeButton(VendingMachine); err != nil {
		fmt.Println(err)
		return
	}

	// select product
	if err := VendingMachine.GetMachineState().ChooseProduct(VendingMachine, 102); err != nil {
		fmt.Println(err)
		return
	}

	// dispense product
	if err := VendingMachine.GetMachineState().DispenseProduct(VendingMachine, 102); err != nil {
		fmt.Println(err)
		return
	}
}

// product price is 30 and inserted amount is 30
func TestForMoreAmount(VendingMachine interfaces.IVendingMachine) {

	// click on insert coin button
	if err := VendingMachine.GetMachineState().ClickOnInsertCoinButton(VendingMachine); err != nil {
		fmt.Println(err)
		return
	}

	// insert coins
	if err := VendingMachine.GetMachineState().InsertCoin(VendingMachine, constants.Dime); err != nil {
		fmt.Println(err)
		return
	}
	if err := VendingMachine.GetMachineState().InsertCoin(VendingMachine, constants.Dime); err != nil {
		fmt.Println(err)
		return
	}
	if err := VendingMachine.GetMachineState().InsertCoin(VendingMachine, constants.Quarter); err != nil {
		fmt.Println(err)
		return
	}

	// click on select product button
	if err := VendingMachine.GetMachineState().ClickOnInsertProdctCodeButton(VendingMachine); err != nil {
		fmt.Println(err)
		return
	}

	// select product
	if err := VendingMachine.GetMachineState().ChooseProduct(VendingMachine, 102); err != nil {
		fmt.Println(err)
		return
	}

	// dispense product
	if err := VendingMachine.GetMachineState().DispenseProduct(VendingMachine, 102); err != nil {
		fmt.Println(err)
		return
	}
}

// product code 105 is not present in inventory
func TestForInvalidProduct(VendingMachine interfaces.IVendingMachine) {

	// click on insert coin button
	if err := VendingMachine.GetMachineState().ClickOnInsertCoinButton(VendingMachine); err != nil {
		fmt.Println(err)
		return
	}

	// insert coins
	if err := VendingMachine.GetMachineState().InsertCoin(VendingMachine, constants.Dime); err != nil {
		fmt.Println(err)
		return
	}
	if err := VendingMachine.GetMachineState().InsertCoin(VendingMachine, constants.Dime); err != nil {
		fmt.Println(err)
		return
	}
	if err := VendingMachine.GetMachineState().InsertCoin(VendingMachine, constants.Dime); err != nil {
		fmt.Println(err)
		return
	}

	// click on select product button
	if err := VendingMachine.GetMachineState().ClickOnInsertProdctCodeButton(VendingMachine); err != nil {
		fmt.Println(err)
		return
	}

	// select product
	if err := VendingMachine.GetMachineState().ChooseProduct(VendingMachine, 105); err != nil {
		fmt.Println(err)
		return
	}

	// dispense product
	if err := VendingMachine.GetMachineState().DispenseProduct(VendingMachine, 105); err != nil {
		fmt.Println(err)
		return
	}
}
