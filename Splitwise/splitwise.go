package main

import (
	"pattern/Splitwise/cmds"
	"pattern/Splitwise/cmds/registry"
)

//https://github.com/Naman-Bhalla/splitwise_mar22/tree/a8626d049cb8e4f470288527da74cba66e58efe7

func main() {
	commadndsRegistry := registry.NewCommandRegistry()

	commadndsRegistry.AddCommand(cmds.NewUserRegisterCommand())
	commadndsRegistry.AddCommand(cmds.NewUpdateUserCommand())
	commadndsRegistry.AddCommand(cmds.NewAddGroupCommand())

	// RegisterUser <name> <phone> <password>
	var input string = "RegisterUser Akhilesh 9906877909 Akhilesh@123"
	commadndsRegistry.ExecuteCommand(input)

	// <userid> UpdateUser <password>
	input = "101 UpdateUser Akhilesh"
	commadndsRegistry.ExecuteCommand(input)

	// <userid> AddGroup <groupname> <description>
	input = "101 AddGroup GoaTrip GoaTripExpenses"
	commadndsRegistry.ExecuteCommand(input)
}
