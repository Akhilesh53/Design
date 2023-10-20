package structs

import (
	"errors"
	"pattern/VendingMachine/constants"
	interfaces "pattern/VendingMachine/interfaces"
)

type hasCoinState struct {
}

func NewHasCoinState() interfaces.IState {
	return &hasCoinState{}
}

func (h *hasCoinState) InsertCoin(vendingMachine interfaces.IVendingMachine, coin constants.Coins) error {
	vendingMachine.AddCoin(coin)
	return nil
}

func (h *hasCoinState) ChooseProduct(vendingMachine interfaces.IVendingMachine, productCode int) error {
	return errors.New(" please press the insert product code button first")
}

func (h *hasCoinState) ReturnChange(vendingMachine interfaces.IVendingMachine, changeAmount int) error {
	return errors.New(" please press the insert product code button first")
}

func (h *hasCoinState) DispenseProduct(vendingMachine interfaces.IVendingMachine, productCode int) error {
	return errors.New(" please press the insert product code button first")
}

func (h *hasCoinState) RefundFullMoney(vendingMachine interfaces.IVendingMachine, refundAmount int) error {
	return errors.New(" please press the insert product code button first")
}

func (h *hasCoinState) ClickOnInsertCoinButton(vendingMachine interfaces.IVendingMachine) error {
	return nil
}

func (h *hasCoinState) ClickOnInsertProdctCodeButton(vendingMachine interfaces.IVendingMachine) error {
	vendingMachine.SetMachineState(NewHasProductState())
	return nil
}
