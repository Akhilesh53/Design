package cmds

import (
	"fmt"
	"pattern/Splitwise/controllers"
	"pattern/Splitwise/dtos"
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
	commandTokens := strings.Split(command, " ")

	addMemberGroupRequestDto := dtos.NewAddMemberGroupRequestDto().SetAdminUserId(commandTokens[0]).SetUserIdToBeAdded(commandTokens[2]).SetGroupId(commandTokens[3])

	response := cmd.groupController.AddMember(addMemberGroupRequestDto)
	if response.GetErr() != nil {
		fmt.Println(response.GetErr())
		return
	}
	fmt.Println("User added to group successfully", response.GetGroup().GetName(), len(response.GetGroup().GetParticipants()))
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
