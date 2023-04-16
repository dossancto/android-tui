package commands

import (
	"fmt"

	"github.com/lu-css/android-tui/src/files"
	"github.com/lu-css/android-tui/src/translate-xml"
	"strings"
)

var manifestCommands = []Command{
	{
		"permittion",
		Permissions,
		"Show all permittions",
	},
	{
		"activity",
		Activities,
		"Show all permittions",
	},
	{
		"feature",
		Features,
		"Show all features",
	},
}

func ManifestCommand(args []string) error {
	if len(args) <= 1 {
		println("HELP MANIFEST")
		return nil
	}

	newArgs := args[1:]

	commandName := newArgs[0]

	if UseCommand(commandName, newArgs, manifestCommands) {
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

func Features(args []string) error {
	data, err := files.ReadManifest()

	if err != nil {
		fmt.Println(err)
	}

	manifest, _ := translate_xml.ParseManifest(data)

	s := fmt.Sprintf("%-70s %-10s", "Name", "required")

	fmt.Println(s)
	fmt.Println("===================================================================================")

	for _, feature := range manifest.Features {
		s := fmt.Sprintf("%-70s %-10t", feature.Name, feature.Required)

		fmt.Println(s)
	}

	return nil
}

func Activities(args []string) error {
	data, err := files.ReadManifest()

	if err != nil {
		fmt.Println(err)
	}

	manifest, _ := translate_xml.ParseManifest(data)

	s := fmt.Sprintf("%1s %-70s", "", "Name")

	fmt.Println(s)
	fmt.Println("===================================================================================")

	for _, activity := range manifest.Application.Activities {
		activityName := strings.Replace(activity.Name, manifest.Package+".", "", 1)
		s := fmt.Sprintf("%1s %-20s %s", "", activityName, activity.Filter.Category.Name)

		fmt.Println(s)
	}

	return nil
}
