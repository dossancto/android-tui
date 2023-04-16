package commands

import "fmt"

var BaseComands = []Command{
	{
		Name:        "manifest",
		Action:      ManifestCommand,
		Description: "Manifest commands",
	},
	{
		Name:        "run",
		Action:      RunCommand,
		Description: "compile commands",
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
		s := fmt.Sprintf("%1s %-20s %-60s", "", command.Name, command.Description)

		fmt.Println(s)
	}
}
