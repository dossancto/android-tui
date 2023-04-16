package commands

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/lu-css/android-tui/src/files"
)

var runCommands = []Command{
	{
		"open",
		Open,
		"Open the app in a device",
	},
	{
		"run",
		run,
		"Compiles and open the project",
	},
}

func RunCommand(args []string) error {
	if len(args) <= 1 {
		runHelp()
		return nil
	}

	newArgs := args[1:]

	commandName := newArgs[0]

	if UseCommand(commandName, newArgs, runCommands) {
		return nil
	}

	s := fmt.Sprintf("\"%s\" not found", commandName)
	fmt.Println(s)

	return nil
}

func runHelp() {
	for _, cmd := range runCommands {
		fmt.Println(cmd.Name)
	}

}

func Open(args []string) error {
	manifest := files.GetManifest()
	launchActivityName := files.GetLaunchActitity(manifest.Application.Activities).Name
	launchString := manifest.Package + "/" + launchActivityName

	cmd := exec.Command("adb", "shell", "am", "start", "-n", launchString)

	out, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Printf("%s\n", out)
		log.Fatal(err)
	}

	fmt.Printf("%s", out)

	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func run(args []string) error {
	// Windows version
	cmd := exec.Command("./gradlew", "installDebug")

	out, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Printf("%s\n", out)
		log.Fatal(err)
	}

	fmt.Printf("%s", out)

	if err != nil {
		log.Fatal(err)
	}

  Open(nil)

	return nil
}
