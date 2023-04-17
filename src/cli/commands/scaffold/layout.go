package scaffold

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func BaseTemplateLayout(activityName string) string {
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

func readActivityName() string {
	var err error
	var activityName string

	reader := bufio.NewReader(os.Stdin)

	for true {
		fmt.Print("Activity name: ")
		activityName, err = reader.ReadString('\n')

		if ValidActivityName(activityName) {
			return activityName
		}

		fmt.Println(err)
	}

	return activityName
}

func GenLayout(args []string) error {

	activityName := readActivityName()

	layout := BaseTemplateLayout(activityName)

	fmt.Println(layout)

	return nil
}
