package parsegradle

import (
	"strconv"
)

func androidConfig(gradle *GradleProject, key string, value string) error {
	switch key {
	case ("namespace"):
		gradle.Android.Namespace = value
		break

	case ("compileSdk"):
		i, _ := strconv.Atoi(value)
		gradle.Android.CompileSdk = i
		break

	case ("m"):
		i, _ := strconv.Atoi(value)
		gradle.Android.CompileSdk = i
		break
	}

	return nil
}
