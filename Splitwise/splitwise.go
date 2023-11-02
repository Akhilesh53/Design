package main

import (
	"pattern/Splitwise/cmds"
	"pattern/Splitwise/cmds/registry"
)

func main() {
	commadndsRegistry := registry.NewCommandRegistry()

	commadndsRegistry.AddCommand(cmds.NewUserRegisterCommand())

	var input string = "RegisterUser Akhilesh 9906877909 Akhilesh@123"
	commadndsRegistry.ExecuteCommand(input)
}
