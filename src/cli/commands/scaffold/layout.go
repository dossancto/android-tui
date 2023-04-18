package scaffold

import (
	"github.com/lu-css/android-tui/src/cli/commands/scaffold/generate"
	"os"

	"github.com/paulrademacher/climenu"
)

func ChooseGenerate(args []string) error {
	menu := climenu.NewButtonMenu("Scaffold", "Choose item to generate ")

	menu.AddMenuItem("Activity", "activity")
	menu.AddMenuItem("Fragment", "fragment")
	menu.AddMenuItem("Folder", "folder")

	action, escaped := menu.Run()

	if escaped {
		os.Exit(0)
	}

	switch action {
	case "activity":
		generate.ChooseActivity()
		break
	}

	return nil
}
