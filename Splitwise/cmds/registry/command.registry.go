package registry

import "pattern/Splitwise/cmds"

type ICommandRegistry interface {
	AddCommand(command cmds.ICommmand)
	GetCommand(command string) cmds.ICommmand
	RemoveCommand(inputCommand string)
	GetAllCommands() []cmds.ICommmand
	ExecuteCommand(inputCommand string)
}

type commandRegistry struct {
	commands []cmds.ICommmand
}

func NewCommandRegistry() ICommandRegistry {
	return &commandRegistry{
		commands: make([]cmds.ICommmand, 0),
	}
}

func (cr *commandRegistry) AddCommand(command cmds.ICommmand) {
	cr.commands = append(cr.commands, command)
}

func (cr *commandRegistry) GetCommand(command string) cmds.ICommmand {
	for _, cmd := range cr.commands {
		if cmd.Parse(command) {
			return cmd
		}
	}
	return nil
}

func (cr *commandRegistry) RemoveCommand(inputCommand string) {
	for idx, command := range cr.commands {
		if command.Parse(inputCommand) {
			cr.commands = append(cr.commands[:idx], cr.commands[idx+1:]...)
			break
		}
	}
}

func (cr *commandRegistry) GetAllCommands() []cmds.ICommmand {
	return cr.commands
}

func (cr *commandRegistry) ExecuteCommand(inputCommand string) {
	for _, command := range cr.commands {
		if command.Parse(inputCommand) {
			command.Execute()
			break
		}
	}
}
