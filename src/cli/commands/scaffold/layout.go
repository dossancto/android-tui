package scaffold

import (
	"fmt"
	"strings"

	"github.com/lu-css/android-tui/src/cli/commands/scaffold/generate"
	"github.com/manifoldco/promptui"
)

func ChooseGenerate(args []string) error {
	generateTypes := []string{
		"Activity",
		"Fragment",
		"Folder",
	}

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "\U000027A1 {{ . | cyan }}",
		Inactive: "  {{ . | cyan }} ",
		Selected: "Generete: {{ . | red | cyan }}",
	}

	prompt := promptui.Select{
		Label:     "Select a Model",
		Items:     generateTypes,
		Templates: templates,
		Size:      4,
	}

	i, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return nil
	}

	switch strings.ToLower(generateTypes[i]) {
	case "activity":
		generate.ChooseActivity()
		break
	}

	return nil
}
