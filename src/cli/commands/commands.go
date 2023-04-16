package commands

import "fmt"

var BaseComands = []Command{
	{
		Name:        "manifest",
		Action:      ManifestCommand,
		Description: "Manifest commands",
	},
}

func UseCommand(commandName string, args []string, commands []Command) bool {

	for _, command := range commands {
		if command.Name == commandName {
			command.Action(args)
			return true
		}
	}

	return false
}

func Help() {
	for _, command := range BaseComands {
		s := fmt.Sprintf("%s\t%s", command.Name, command.Description)

		fmt.Println(s)
	}
}
