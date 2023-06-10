package files

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/lu-css/android-tui/pkg/parsegradle"
)

func getGradleBuildPath() string {
	currentPath, err := os.Getwd()

	if err != nil {
		fmt.Println(err)
	}

	file := find(currentPath, "build.gradle")

	return file
}

func readGradleBuildFile() ([]byte, error) {
	gradlePath := getGradleBuildPath()

	file, err := ioutil.ReadFile(gradlePath)

	if err != nil {
		return nil, err
	}

	return file, nil

}

func GetGradleBuildConfig() parsegradle.AndroidConfig {
	file, err := readGradleBuildFile()

	if err != nil {
		log.Fatal(err)
	}

	androidConfig := parsegradle.Parse(string(file))

	return androidConfig
}
