package cmds

import (
	"fmt"
	"pattern/Splitwise/controllers"
	"pattern/Splitwise/dtos"
	"strings"
	"sync"
)

type updateUserCommand struct {
	userContoller controllers.IUserContoller
}

var userUpdateCommandInstance ICommmand
var userUpdateCommandOnce sync.Once

func NewUpdateUserCommand() ICommmand {
	userUpdateCommandOnce.Do(func() {
		userUpdateCommandInstance = &updateUserCommand{
			userContoller: userContoller,
		}
	})
	return userUpdateCommandInstance
}

func (urc *updateUserCommand) Execute(inputCommand string) {
	commandTokens := strings.Split(inputCommand, " ")
	requestDto := dtos.NewUpdateUserRequestDto(commandTokens[0], commandTokens[2])
	response := urc.userContoller.UpdateUser(requestDto)

	if response.GetErr() != nil {
		fmt.Println("Error while updating user", response.GetErr())
		return
	}
	fmt.Println("User Updated Successfully", response.GetUser().GetName(), response.GetUser().GetId(), response.GetUser().GetPassword())
}

// Example: userid UpdateUser <password>
func (urc *updateUserCommand) Parse(command string) bool {
	commandTokens := strings.Split(command, " ")
	if len(commandTokens) != 3 ||
		!strings.EqualFold(commandTokens[1], USER_UPDATE_COMMAND) ||
		len(commandTokens[0]) == 0 ||
		len(commandTokens[1]) == 0 ||
		len(commandTokens[2]) == 0 {
		return false
	}
	return true
}
