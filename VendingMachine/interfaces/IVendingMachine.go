package interfaces

import "pattern/VendingMachine/constants"

type IVendingMachine interface {
	SetMachineState(IState) IVendingMachine
	GetMachineState() IState
	GetInventory() Inventory
	SetInventory(Inventory) IVendingMachine
	GetMachineCoins() map[constants.Coins]int
	GetMachineCoinCount(constants.Coins) int
	IncrementCoinCount(coin constants.Coins) IVendingMachine
	DecrementCoinCount(coin constants.Coins) IVendingMachine
	AddCoin(coin constants.Coins) IVendingMachine
	GetCoins() []constants.Coins
}
