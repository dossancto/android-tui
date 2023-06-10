package parsegradle

import (
	"bufio"
	"fmt"
	"log"
	"strings"
)

func parse(file string, v any) any {
	var gradleProject GradleProject

	fmt.Println("Start")

	reader := strings.NewReader(file)
	scan := bufio.NewScanner(reader)

	// Iterate over each line of the file
	parseBlocks(scan, 0, v, &gradleProject, "")

	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}
	return "a"
}

func parseBlocks(scan *bufio.Scanner, originalDepth int, v any, gradleProject *GradleProject, contextFather string) {
	depth := originalDepth
  blockName := contextFather

	for scan.Scan() {
		line := scan.Text()

		if countOpenBreakts(line, &depth) {
			blockName = clearBlockName(line)
			parseBlocks(scan, depth, v, gradleProject, blockName)
			continue
		}

		if countClosingBreakts(line, &depth) {
			continue
		}

		k, v := spliGradleKeyValue(line)

		// A valid key value struct
		if v != "" {
			err := parseKeyValue(k, v, blockName)

			if err != nil {
				log.Fatal(err)
			}
		}

		if depth < originalDepth {
			break
		}
	}
}

func parseKeyValue(key string, value string, context string) error {
	fmt.Println("CONTEXT => " + context)
	var gradle GradleProject

	if strings.Contains(context, "android") {
		androidConfig(&gradle, key, value)
    fmt.Println("CUUUUUUUUUUu")
	}

	return nil
}

func spliGradleKeyValue(line string) (string, string) {
	line = strings.TrimLeft(line, " ")
	keys := strings.Split(line, " ")

	if len(keys) < 2 {
		return "", ""
	}

	return keys[0], keys[1]
}

func clearBlockName(line string) string {
	return strings.ReplaceAll(line, " {", "")
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
