package scaffold

import (
	"fmt"
	"os"
	"strings"

	"github.com/paulrademacher/climenu"
)

func baseTemplateLayout(activityName string) string {
	cleanActivityName := strings.Trim(activityName, "\n")

	template := fmt.Sprintf(`<?xml version="1.0" encoding="utf-8"?>
    <androidx.constraintlayout.widget.ConstraintLayout xmlns:android="http://schemas.android.com/apk/res/android"
    xmlns:app="http://schemas.android.com/apk/res-auto"
    xmlns:tools="http://schemas.android.com/tools"
    android:layout_width="match_parent"
    android:layout_height="match_parent"
    tools:context=".%v">

    </androidx.constraintlayout.widget.ConstraintLayout>`, cleanActivityName)

	return template
}

func ChooseGenerate(args []string) error {
	menu := climenu.NewButtonMenu("Scaffold", "Choose item to generate ")

	menu.AddMenuItem("Activity", "activity")
	menu.AddMenuItem("Fragment", "fragment")
	menu.AddMenuItem("Folder", "folder")

	action, escaped := menu.Run()

	if escaped {
		os.Exit(0)
	}

	switch action {
	case "activity":
		GenLayout(args)
		break

	}

	return nil
}

func GenLayout(args []string) error {
	activityName := climenu.GetText("ActivityName", "nothing")

	layout := baseTemplateLayout(activityName)

	fmt.Println(layout)

	return nil
}
