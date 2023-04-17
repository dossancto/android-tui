package scaffold

import (
	"strings"
)

func ValidActivityName(activityName string) bool {
	cleanActivityName := strings.Trim(activityName, "\n")

	return cleanActivityName != ""
}
