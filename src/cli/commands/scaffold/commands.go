package scaffold

import (
	"fmt"
	"github.com/lu-css/android-tui/src/cli/commands/more"
)

var ScaffoldCommands = []more.Command{
	{
		Name:        "layout",
		Action:      GenLayout,
		Description: "Manifest commands",
	},
}

func scaffoldHelp() {
	for _, cmd := range ScaffoldCommands {
		fmt.Println(cmd.Name)
	}

}

func ScaffoldCommand(args []string) error {

	if len(args) <= 1 {
		scaffoldHelp()
		return nil
	}

	newArgs := args[1:]

	commandName := newArgs[0]

	if more.UseCommand(commandName, newArgs, ScaffoldCommands) {
		return nil
	}

	s := fmt.Sprintf("\"%s\" not found", commandName)
	fmt.Println(s)

	return nil
}
