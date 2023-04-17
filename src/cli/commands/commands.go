package commands

import (
	"fmt"
	"github.com/lu-css/android-tui/src/cli/commands/more"
	"github.com/lu-css/android-tui/src/cli/commands/scaffold"
)

var BaseComands = []more.Command{
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
	{
		Name:        "scaffold",
		Action:      scaffold.ScaffoldCommand,
		Description: "compile commands",
	},
}

func Help() {
	for _, command := range BaseComands {
		s := fmt.Sprintf("%1s %-20s %-60s", "", command.Name, command.Description)

		fmt.Println(s)
	}
}
