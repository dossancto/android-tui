package more

func UseCommand(commandName string, args []string, commands []Command) bool {
	for _, command := range commands {
		if command.Name == commandName {
			command.Action(args)
			return true
		}
	}

	return false
}
