package interfaces

import (
	"pattern/VendingMachine/constants"
)

type IState interface {
	InsertCoin(vendingMachine IVendingMachine, coin constants.Coins) error
	ChooseProduct(vendingMachine IVendingMachine, productCode int) error
	ReturnChange(IVendingMachine, int) error
	DispenseProduct(IVendingMachine, int) error
	RefundFullMoney(vendingMachine IVendingMachine, amount int) error
	ClickOnInsertCoinButton(IVendingMachine) error
	ClickOnInsertProdctCodeButton(IVendingMachine) error
}
