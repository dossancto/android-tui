package main

import (
	"fmt"

	"os"

	"github.com/lu-css/android-tui/src/cli/commands"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("HELP")
		return
	}

	args := os.Args[1:]

	if os.Args[1] == "help" {
		commands.Help()
		return
	}

	if !commands.UseCommand(os.Args[1], args, commands.BaseComands) {
		println("Command not found")
	}
}
