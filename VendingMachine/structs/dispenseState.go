package structs

import (
	"errors"
	"fmt"
	"pattern/VendingMachine/constants"
	interfaces "pattern/VendingMachine/interfaces"
)

type dispenseState struct {
}

func NewDispenseState() interfaces.IState {
	return &dispenseState{}
}

func (d *dispenseState) InsertCoin(vendingMachine interfaces.IVendingMachine, coin constants.Coins) error {
	return errors.New("cannot insert coins in dispense state")
}

func (d *dispenseState) ChooseProduct(vendingMachine interfaces.IVendingMachine, productCode int) error {
	return errors.New("cannot choose product in dispense state")
}

func (d *dispenseState) ReturnChange(vendingMachine interfaces.IVendingMachine, changeAmount int) error {
	return errors.New("cannot get change in dispense state")
}

func (d *dispenseState) DispenseProduct(vendingMachine interfaces.IVendingMachine, productCode int) error {
	fmt.Println("product with product code ", productCode, " is dispensed in collection tray")
	vendingMachine.SetMachineState(NewIdleState())
	if err := vendingMachine.GetInventory().DispenseProduct(productCode); err != nil {
		return err
	}
	for _, coin := range vendingMachine.GetCoins() {
		vendingMachine.IncrementCoinCount(coin)
	}
	return nil
}

func (d *dispenseState) RefundFullMoney(vendingMachine interfaces.IVendingMachine, refundAmount int) error {
	return errors.New("money cannot be refunded in dispense state")
}

func (d *dispenseState) ClickOnInsertCoinButton(vendingMachine interfaces.IVendingMachine) error {
	return errors.New("coins cannot be inserted in dispense state")
}

func (d *dispenseState) ClickOnInsertProdctCodeButton(vendingMachine interfaces.IVendingMachine) error {
	return errors.New("product code cannot be inserted in dispense state")
}
