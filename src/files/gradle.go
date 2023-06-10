package files

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type AndroidConfig struct {
	Namespace string `gradle:"namespace"`
	CompileSdk int `gradle:"compileSdk"`
}

type GradleProject struct {
	Android AndroidConfig `gradle:"android"`
}

func readGradleFile() string {
	gradlePath := "/home/lucss/program/java/mobile/api-android/app/build.gradle"

	file, err := ioutil.ReadFile(gradlePath)

	if err != nil {
		log.Fatal(err)
	}

	return string(file)
}

func GetGradleBuildConfig() {
	// Open the .gradle file
	file := readGradleFile()

	reader := strings.NewReader(file)
	scan := bufio.NewScanner(reader)

	// Create a variable to store the parsed GradleProject
	var project GradleProject

	// Iterate over each line of the file
	for scan.Scan() {
		line := scan.Text()

		fmt.Println(line)
		// Enters the android{} section
		if strings.HasPrefix(line, "android") {
			parseAndroidSection(&project, scan)
		}
	}

	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}
}

func parseAndroidSection(project *GradleProject, scan *bufio.Scanner) {
	brackets := 1

	for scan.Scan() {
		line := scan.Text()
		line = strings.TrimSpace(line)

		// // Remove "project" and surrounding parentheses
		// line = strings.TrimPrefix(line, "project(")
		// line = strings.TrimSuffix(line, ")")

		// Start of Section
		if countOpenBreakts(line, &brackets) {
			fmt.Println(line)
		}

		// End of Section
		if countClosingBreakts(line, &brackets) {

		}

		// Ends the Section
		if brackets == 0 {
			break
		}
	}
}

func countOpenBreakts(line string, count *int) bool {
	if strings.ContainsAny(line, "{") {
		*count++
		return true
	}

	return false
}

func countClosingBreakts(line string, count *int) bool {
	if strings.ContainsAny(line, "}") && *count > 0 {
		*count--
		return true
	}

	return false
}
