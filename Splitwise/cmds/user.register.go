package cmds

import (
	"fmt"
	"strings"
)

type userRegisterCoomand struct {
}

func NewUserRegisterCommand() ICommmand {
	return &userRegisterCoomand{}
}

func (urc *userRegisterCoomand) Execute() {
	fmt.Println("User Register Command Executed")
}

// Example: RegisterUser <name> <phone number> <password>
func (urc *userRegisterCoomand) Parse(command string) bool {
	commandTokens := strings.Split(command, " ")
	if len(commandTokens) != 4 || !strings.EqualFold(commandTokens[0], USER_REGISTER_COMMAND) || len(commandTokens[1]) == 0 || len(commandTokens[2]) == 0 || len(commandTokens[3]) == 0 {
		fmt.Println("Invalid Command")
		return false
	}

	//todo: add more validations for other params
	return true
}
