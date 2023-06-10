package parsegradle

import (
	"io/ioutil"
	"log"
)


func readGradleFile() string {
	gradlePath := "/home/lucss/program/java/mobile/api-android/app/build.gradle"

	file, err := ioutil.ReadFile(gradlePath)

	if err != nil {
		log.Fatal(err)
	}

	return string(file)
}

func GetGradleBuildConfig() {
  var gradleConfig GradleProject
	// Open the .gradle file
	file := readGradleFile()

	parse(file, &gradleConfig)
}
