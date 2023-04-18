package more

func UseCommand(commandName string, args []string, commands []Command) bool {
	for _, command := range commands {
		if command.Name == commandName {
			err := command.Action(args)

			if err != nil {
				println(err)
			}

			return true
		}
	}

	return false
}
