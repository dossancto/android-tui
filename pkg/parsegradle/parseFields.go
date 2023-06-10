package parsegradle

import (
	"bufio"
	"strconv"
	"strings"
)

func parseAndroidConfig(scan *bufio.Scanner) AndroidConfig {
	var android AndroidConfig

	for scan.Scan() {
		line := scan.Text()

		k, v := spliGradleKeyValue(line)

		if v == "" {
			continue
		}

		switch k {
		case ("targetSdk"):
			android.DefaultConfig.TargetSdk = toint(v)

		case ("minSdk"):
			android.DefaultConfig.MinSdk = toint(v)

		case ("applicationId"):
			android.DefaultConfig.ApplicationId = strings.Trim(v, "\"")

		case ("namespace"):
			android.Namespace = strings.Trim(v, "\"")

		case ("compileSdk"):
			android.CompileSdk = toint(v)
		}
	}

	return android
}

func toint(v string) int {
	i, err := strconv.Atoi(v)

	if err != nil {
		return 0
	}

	return i
}
