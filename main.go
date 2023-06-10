package main

import (
	"fmt"
	"os"

	"github.com/lu-css/android-tui/src/cli/commands"
	"github.com/lu-css/android-tui/src/cli/commands/more"
	// "github.com/lu-css/android-tui/src/files"
	"github.com/lu-css/android-tui/src/validations"
)

func main() {

	if !validations.GradleFileExists() {
		s := fmt.Errorf("Gradle Project not found, place enter the app directory")
		fmt.Println(s)
		return

	}

	if len(os.Args) < 2 {
		fmt.Println("HELP")
		return
	}

	args := os.Args[1:]

	if os.Args[1] == "help" {
		commands.Help()
		return
	}

	if !more.UseCommand(os.Args[1], args, commands.BaseComands) {
		println("Command not found")
	}
}
