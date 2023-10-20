package structs

import (
	"errors"
	"fmt"
	"pattern/VendingMachine/constants"
	interfaces "pattern/VendingMachine/interfaces"
)

type hasProductState struct {
}

func NewHasProductState() interfaces.IState {
	return &hasProductState{}
}

func (h *hasProductState) InsertCoin(vendingMachine interfaces.IVendingMachine, coin constants.Coins) error {
	return nil
}

func (h *hasProductState) ChooseProduct(vendingMachine interfaces.IVendingMachine, productCode int) error {
	//get total money inserted
	var amountInserted = 0
	for _, coin := range vendingMachine.GetCoins() {
		amountInserted += coin.Value()
	}

	// check if the product is in inventory, then get the price of the product
	if !vendingMachine.GetInventory().HasProduct(productCode) {
		vendingMachine.GetMachineState().RefundFullMoney(vendingMachine, amountInserted)
		return errors.New(" product not found/ not in stock")
	}
	// get the price of product
	productPrice := vendingMachine.GetInventory().GetItemDetailsFromInventory(productCode).GetPrice()

	if amountInserted < productPrice {
		vendingMachine.GetMachineState().RefundFullMoney(vendingMachine, amountInserted)
		return errors.New(" insufficient money added")
	}

	// check if there is any amount to be returned
	if amountInserted > productPrice {
		vendingMachine.GetMachineState().ReturnChange(vendingMachine, amountInserted-productPrice)
	}
	vendingMachine.SetMachineState(NewDispenseState())
	return nil
}

func (h *hasProductState) ReturnChange(vendingMachine interfaces.IVendingMachine, changeAmount int) error {
	fmt.Println(" change amount is dispensed in change tray Rs ", changeAmount)
	//todo:
	//get the denominations which make up the change and decrease its count from machineCoins
	return nil
}

func (h *hasProductState) DispenseProduct(vendingMachine interfaces.IVendingMachine, productCode int) error {
	return errors.New(" please press the insert product code button first")
}

func (h *hasProductState) RefundFullMoney(vendingMachine interfaces.IVendingMachine, refundAmount int) error {
	fmt.Println(" your complate money is refunded Rs ", refundAmount)
	vendingMachine.SetMachineState(NewIdleState())
	return nil
}

func (h *hasProductState) ClickOnInsertCoinButton(vendingMachine interfaces.IVendingMachine) error {
	return errors.New(" coins cannot be inserted in product selection state")
}

func (h *hasProductState) ClickOnInsertProdctCodeButton(vendingMachine interfaces.IVendingMachine) error {
	return nil
}
