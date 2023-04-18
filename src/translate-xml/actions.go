package translate_xml

import (
	"errors"
	"fmt"
)

func (application Application) GetLaunchActitity() (Activity, error) {

	for _, activity := range application.Activities {
		if activity.Filter.Category.Name == "android.intent.category.LAUNCHER" {
			return activity, nil
		}
	}

	return Activity{}, errors.New("Launch Activity does not exists.")
}

func (manifest Manifest) PrinParsedManifest() {
	fmt.Println(manifest.Tools)
	fmt.Println(manifest.Package)

	println("\nFeatures\n")

	for _, features := range manifest.Features {
		s := fmt.Sprintf("%s - %t", features.Name, features.Required)
		fmt.Println(s)
	}

	println("\nPermittions\n")

	for _, features := range manifest.Permissions {
		s := fmt.Sprintf("%s", features.Name)
		fmt.Println(s)
	}

	println("\nApplication infos\n")
	b := manifest.Application

	fmt.Println(b.AllowBackup)
	fmt.Println(b.DataExtractionRules)
	fmt.Println(b.FullBackupContent)
	fmt.Println(b.Icon)
	fmt.Println(b.Label)
	fmt.Println(b.RoundIcon)
	fmt.Println(b.SupportsRtl)
	fmt.Println(b.Theme)
	fmt.Println(b.TargetApi)

	println("\nActivities infos\n")
	for _, ac := range b.Activities {
		fmt.Println("Name:", ac.Name)
		fmt.Println("Expoted:", ac.Exported)

		fmt.Println("Meta-data Name:", ac.MetaData.Name)
		fmt.Println("Meta-data Value:", ac.MetaData.Value)
		fmt.Println()

	}
}
