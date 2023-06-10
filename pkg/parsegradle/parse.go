package parsegradle

import (
	"bufio"
	// "fmt"
	"log"
	"strings"
)

func Parse(file string) AndroidConfig {
	reader := strings.NewReader(file)
	scan := bufio.NewScanner(reader)

	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}

	// Iterate over each line of the file

	return parseBlocks(scan, 0, "")
}

func parseBlocks(scan *bufio.Scanner, originalDepth int, contextFather string) AndroidConfig {
	var gradle AndroidConfig

	depth := originalDepth
	blockName := contextFather

	for scan.Scan() {
		line := scan.Text()

		if countOpenBreakts(line, &depth) {
			blockName = clearBlockName(line)
			if strings.Contains("android", blockName) {
				return parseAndroidConfig(scan)
			}
			continue
		}

		if countClosingBreakts(line, &depth) {
			continue
		}

		// k, v := spliGradleKeyValue(line)

		// A valid key value struct
		// if v != "" {
		// 	android := parseKeyValue(k, v, blockName)
		// 	fmt.Println(android)
		// }

		if depth < originalDepth {
			break
		}
	}

	return gradle
}

func parseKeyValue(key string, value string, context string) AndroidConfig {
	if strings.Contains(context, "android") {
		return AndroidConfig{}
	}

	return AndroidConfig{}
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
