package cmds

import (
	"fmt"
	"pattern/Splitwise/controllers"
	"pattern/Splitwise/dtos"
	"strings"
	"sync"
)

type getGroupCommand struct {
	groupController controllers.IGroupController
}

var getGroupCommandInstance ICommmand
var getGroupCommandInstanceOnce sync.Once

func NewGetGroupsCommand() ICommmand {
	getGroupCommandInstanceOnce.Do(func() {
		getGroupCommandInstance = &getGroupCommand{
			groupController: groupContoller,
		}
	})
	return getGroupCommandInstance
}

func (gcd *getGroupCommand) Execute(command string) {
	commandTokens := strings.Split(command, " ")
	requestDto := dtos.NewGetUserGroupsRequestDtoWithUserId(commandTokens[0])
	response := gcd.groupController.GetGroupsForUser(requestDto)
	if response.GetErr() != nil {
		fmt.Println("Error while getting groups", response.GetErr())
		return
	}
	fmt.Println("Groups for user", commandTokens[0], "are: ", len(response.GetGroups()))
}

// <userid> Groups
func (gcd *getGroupCommand) Parse(command string) bool {
	commandsToken := strings.Split(command, " ")
	if len(commandsToken) != 2 ||
		!strings.EqualFold(commandsToken[1], GET_GROUPS_COMMAND) ||
		len(commandsToken[0]) == 0 {
		return false
	}
	return true
}
