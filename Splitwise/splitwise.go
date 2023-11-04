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
	commadndsRegistry.AddCommand(cmds.NewAddMemberGroupCommand())

	// RegisterUser <name> <phone> <password>
	var input string = "RegisterUser Akhilesh1 9906877909 Akhilesh@123"
	commadndsRegistry.ExecuteCommand(input)

	// <userid> UpdateUser <password>
	input = "1 UpdateUser Akhilesh"
	commadndsRegistry.ExecuteCommand(input)

	input = "RegisterUser Akhilesh2 9906876909 Akhilesh@123"
	commadndsRegistry.ExecuteCommand(input)

	// <userid> AddGroup <groupname> <description>
	input = "1 AddGroup GoaTrip GoaTripExpenses"
	commadndsRegistry.ExecuteCommand(input)

	// <userid_whoisadding> AddUser <userid_tobeadded> <groupid>
	input = "1 AddMember 2 1"
	commadndsRegistry.ExecuteCommand(input)
}
