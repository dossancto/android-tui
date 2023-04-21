package validations

import (
	"errors"
)

func NonBlankInput(input string) error {
	if input == "" {
		return errors.New("Blank")
	}
	return nil
}
