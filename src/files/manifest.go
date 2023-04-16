package files

import (
	"fmt"
	_ "io/fs"
	"os"

	// "path/filepath"
	"path/filepath"
	_ "path/filepath"
)

func FindManifest() string {
	return "teste"
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

	println(file)

	return ""
}
