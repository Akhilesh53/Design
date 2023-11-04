package cmds

import (
	"fmt"
	"pattern/Splitwise/controllers"
	"pattern/Splitwise/dtos"
	"strconv"
	"strings"
	"sync"
)

type addGroupCommand struct {
	groupContoller controllers.IGroupController
}

var addGroupCommandInstance ICommmand
var addGroupCommandOnce sync.Once
var groupContoller controllers.IGroupController

func init() {
	groupContoller = controllers.NewGroupController()
}

func NewAddGroupCommand() ICommmand {
	addGroupCommandOnce.Do(func() {
		addGroupCommandInstance = &addGroupCommand{
			groupContoller: groupContoller,
		}
	})
	return addGroupCommandInstance
}

func (agc *addGroupCommand) Execute(inputCommand string) {
	commandTokens := strings.Split(inputCommand, " ")
	request := dtos.NewAddGroupRequestDto(commandTokens[0], commandTokens[2], commandTokens[3])
	response := agc.groupContoller.AddGroup(request)

	if response.GetErr() != nil {
		fmt.Println(response.GetErr())
		return
	}
	addMemberGroupRequestDto := dtos.NewAddMemberGroupRequestDto().SetAdminUserId(commandTokens[0]).SetUserIdToBeAdded(commandTokens[0]).SetGroupId(strconv.Itoa(response.GetGroup().GetId()))
	response = agc.groupContoller.AddMember(addMemberGroupRequestDto)

	fmt.Println("Added Group Name: ", response.GetGroup().GetName(), response.GetGroup().GetCreatedBy(), response.GetGroup().GetId())
}

// Example: userid AddGroup <groupname> <description>
func (agc *addGroupCommand) Parse(command string) bool {
	commandTokens := strings.Split(command, " ")
	if len(commandTokens) != 4 ||
		!strings.EqualFold(commandTokens[1], ADD_GROUP_COMMAND) ||
		len(commandTokens[0]) == 0 ||
		len(commandTokens[1]) == 0 ||
		len(commandTokens[2]) == 0 ||
		len(commandTokens[3]) == 0 {
		return false
	}
	return true
}
