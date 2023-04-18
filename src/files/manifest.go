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

		//TODO: Ignore git ignored files
		if !file.IsDir() || file.Name() == ".git" {
			continue
		}

		if f := find(filePath, target); f != "" {
			return f
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
