package cmds

type ICommmand interface {
	// execute will execute the respective function
	Execute()

	//parse will return whether command entered is valid or not
	Parse(command string) bool
}
