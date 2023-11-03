package cmds

import (
	"fmt"
	"pattern/Splitwise/controllers"
	"pattern/Splitwise/dtos"
	"strings"
)

type userRegisterCommand struct {
	userContoller controllers.IUserContoller
}

func NewUserRegisterCommand() ICommmand {
	return &userRegisterCommand{
		userContoller: controllers.NewUserContoller(),
	}
}

func (urc *userRegisterCommand) Execute(inputCommand string) {
	commandTokens := strings.Split(inputCommand, " ")
	requestDto := dtos.NewRegisterUserRequestDto(commandTokens[1], commandTokens[2], commandTokens[3])
	response := urc.userContoller.RegisterUser(*requestDto)

	fmt.Println("User Registered Successfully", response.GetUser().GetName(), response.GetUser().GetId())
}

// Example: RegisterUser <name> <phone number> <password>
func (urc *userRegisterCommand) Parse(command string) bool {
	commandTokens := strings.Split(command, " ")
	if len(commandTokens) != 4 ||
		!strings.EqualFold(commandTokens[0], USER_REGISTER_COMMAND) ||
		len(commandTokens[1]) == 0 ||
		len(commandTokens[2]) == 0 ||
		len(commandTokens[3]) == 0 {

		fmt.Println("Invalid Command")
		return false
	}

	//todo: add more validations for other params
	return true
}