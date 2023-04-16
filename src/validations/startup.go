package validations

import (
	"log"
	"os"
)

func GradleFileExists() bool {
	currentPath, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	files, err := os.ReadDir(currentPath)

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {

		if file.Name() == "gradlew" {
			return true
		}

	}

	return false
}
