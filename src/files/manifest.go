package files

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/lu-css/android-tui/src/translate-xml"
	"path/filepath"
)

func GetManifest() translate_xml.Manifest {
	data, err := ReadManifest()

	if err != nil {
		fmt.Println(err)
		return translate_xml.Manifest{}
	}

	manifest, err := translate_xml.ParseManifest(data)
	if err != nil {
		fmt.Println(err)
		return translate_xml.Manifest{}
	}

	return manifest
}

func find(dirName string, target string) string {
	files, err := os.ReadDir(dirName)

	if err != nil {
		println("error")
		fmt.Println(err)
	}

	for _, file := range files {
		fileName := file.Name()
		filePath := filepath.Join(dirName, fileName)

		if file.Name() == target {
			return filePath
		}

		//TODO: Don't check gitignored files.
		if file.IsDir() && file.Name() != ".git" {
			f := find(filePath, target)

			if f != "" {
				return f
			}

		}

	}

	return ""

}

func FindManifestPath() string {
	currentPath, err := os.Getwd()

	if err != nil {
		fmt.Println(err)
	}

	file := find(currentPath, "AndroidManifest.xml")

	return file
}

func ReadManifest() ([]byte, error) {
	manifestPath := FindManifestPath()

	file, err := ioutil.ReadFile(manifestPath)

	if err != nil {

		return nil, err
	}

	return file, nil

}

func GetLaunchActitity(activities []translate_xml.Activity) translate_xml.Activity {

	for _, activity := range activities {
		if activity.Filter.Category.Name == "android.intent.category.LAUNCHER" {
			return activity
		}
	}

	return translate_xml.Activity{}
}

func PrinParsedManifest(manifest translate_xml.Manifest) {
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
