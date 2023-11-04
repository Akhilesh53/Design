package main

import (
	"pattern/Splitwise/cmds"
	"pattern/Splitwise/cmds/registry"
)

func main() {
	commadndsRegistry := registry.NewCommandRegistry()

	commadndsRegistry.AddCommand(cmds.NewUserRegisterCommand())
	commadndsRegistry.AddCommand(cmds.NewUpdateUserCommand())

	var input string = "RegisterUser Akhilesh 9906877909 Akhilesh@123"
	commadndsRegistry.ExecuteCommand(input)

	input = "101 UpdateUser Akhilesh"
	commadndsRegistry.ExecuteCommand(input)
}
