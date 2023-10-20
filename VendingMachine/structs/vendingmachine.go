package structs

import (
	"pattern/VendingMachine/constants"
	interfaces "pattern/VendingMachine/interfaces"
)

type VendingMachine struct {
	machineState interfaces.IState
	machineCoins map[constants.Coins]int
	inventory    interfaces.Inventory
	coins        []constants.Coins
}

func NewVendingMachine() interfaces.IVendingMachine {
	return &VendingMachine{
		machineState: NewIdleState(),
		machineCoins: make(map[constants.Coins]int),
		inventory:    GetNewInventory(),
		coins:        make([]constants.Coins, 0),
	}
}

func (v *VendingMachine) SetMachineState(state interfaces.IState) interfaces.IVendingMachine {
	v.machineState = state
	return v
}

func (v *VendingMachine) GetMachineState() interfaces.IState {
	return v.machineState
}

func (v *VendingMachine) GetInventory() interfaces.Inventory {
	return v.inventory
}

func (v *VendingMachine) SetInventory(inventory interfaces.Inventory) interfaces.IVendingMachine {
	v.inventory = inventory
	return v
}

func (v *VendingMachine) GetMachineCoins() map[constants.Coins]int {
	return v.machineCoins
}

func (v *VendingMachine) AddCoin(coin constants.Coins) interfaces.IVendingMachine {
	v.coins = append(v.coins, coin)
	return v
}

func (v *VendingMachine) GetMachineCoinCount(coin constants.Coins) int {
	return v.machineCoins[coin]
}

func (v *VendingMachine) GetCoins() []constants.Coins {
	return v.coins
}

func (v *VendingMachine) IncrementCoinCount(coin constants.Coins) interfaces.IVendingMachine {
	v.machineCoins[coin]++
	return v
}

func (v *VendingMachine) DecrementCoinCount(coin constants.Coins) interfaces.IVendingMachine {
	if v.machineCoins[coin] > 0 {
		v.machineCoins[coin]--
	}
	return v
}
