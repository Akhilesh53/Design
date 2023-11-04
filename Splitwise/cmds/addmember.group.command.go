package cmds

import (
	"fmt"
	"pattern/Splitwise/controllers"
	"strings"
	"sync"
)

type AddMemberGroupCommand struct {
	groupController controllers.IGroupController
}

var addMemberGroupInstance ICommmand
var addMemberGroupOnce sync.Once

func NewAddMemberGroupCommand() ICommmand {
	addMemberGroupOnce.Do(func() {
		addMemberGroupInstance = &AddMemberGroupCommand{
			groupController: groupContoller,
		}
	})
	return addMemberGroupInstance
}

func (cmd *AddMemberGroupCommand) Execute(command string) {
	fmt.Println("AddMemberGroupCommand called")
}

// <userid_whoisadding> AddUser <userid_tobeadded> <groupid>
func (cmd *AddMemberGroupCommand) Parse(command string) bool {
	commandTokens := strings.Split(command, " ")
	if len(commandTokens) != 4 ||
		commandTokens[1] != ADD_MEMBER_COMMAND ||
		len(commandTokens[0]) == 0 ||
		len(commandTokens[2]) == 0 ||
		len(commandTokens[3]) == 0 {
		return false

	}
	return true
}
