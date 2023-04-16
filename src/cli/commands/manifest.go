package commands

import (
	"fmt"

	"github.com/lu-css/android-tui/src/files"
	"github.com/lu-css/android-tui/src/translate-xml"
)

var ManifestCommands = []Command{
	{
		"permittion",
		Permissions,
		"Show all permittions",
	},
}

func ManifestCommand(args []string) error {
	fmt.Println(args)

	if len(args) <= 1 {
		println("HELP MANIFEST")
		return nil
	}

	newArgs := args[1:]

	commandName := newArgs[0]

	if UseCommand(commandName, newArgs, ManifestCommands) {
		return nil
	}

	s := fmt.Sprintf("\"%s\" not found", commandName)
	fmt.Println(s)

	return nil
}

func Permissions(args []string) error {
	data, err := files.ReadManifest()

	if err != nil {
		fmt.Println(err)
	}

	manifest, _ := translate_xml.ParseManifest(data)

	println("Listions permittions in project\n")
	for _, permittion := range manifest.Permissions {
		println(permittion.Name)
	}

	return nil
}
