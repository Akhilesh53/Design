package structs

import (
	"errors"
	"pattern/VendingMachine/constants"
	interfaces "pattern/VendingMachine/interfaces"
)

type idleState struct {
}

func NewIdleState() interfaces.IState {
	return &idleState{}
}

func (i *idleState) InsertCoin(vendingMachine interfaces.IVendingMachine, coin constants.Coins) error {
	return errors.New(" please press the insert coin button first")
}

func (i *idleState) ChooseProduct(vendingMachine interfaces.IVendingMachine, productCode int) error {
	return errors.New(" please press the insert coin button first")
}

func (i *idleState) ReturnChange(vendingMachine interfaces.IVendingMachine, changeAmount int) error {
	return errors.New(" please press the insert coin button first")
}

func (i *idleState) DispenseProduct(vendingMachine interfaces.IVendingMachine, productCode int) error {
	return errors.New(" please press the insert coin button first")
}

func (i *idleState) RefundFullMoney(vendingMachine interfaces.IVendingMachine, refundAmount int) error {
	return errors.New(" please press the insert coin button first")
}

func (i *idleState) ClickOnInsertCoinButton(vendingMachine interfaces.IVendingMachine) error {
	vendingMachine.SetMachineState(NewHasCoinState())
	return nil
}

func (i *idleState) ClickOnInsertProdctCodeButton(vendingMachine interfaces.IVendingMachine) error {
	return errors.New(" please press the insert coin button first")
}
