package generate

import (
	"fmt"
	"os"
	"strings"

	"github.com/lu-css/android-tui/src/files"

	"github.com/paulrademacher/climenu"
)

func baseTemplateLayout(activityName string) (string, string) {
	cleanActivityName := "activity_" + strings.Trim(activityName, "\n")

	template := fmt.Sprintf(`<?xml version="1.0" encoding="utf-8"?>
    <androidx.constraintlayout.widget.ConstraintLayout xmlns:android="http://schemas.android.com/apk/res/android"
    xmlns:app="http://schemas.android.com/apk/res-auto"
    xmlns:tools="http://schemas.android.com/tools"
    android:layout_width="match_parent"
    android:layout_height="match_parent"
    tools:context=".%v">

    </androidx.constraintlayout.widget.ConstraintLayout>`, cleanActivityName)

	return template, cleanActivityName
}

func baseJavaLayout(ActivityName string, layoutName string) string {
	manifest := files.GetManifest()
	className := ActivityName

	template := fmt.Sprintf(`
    package %s;

    import androidx.appcompat.app.AppCompatActivity;

    import android.os.Bundle;

    public class %s extends AppCompatActivity {

        @Override
        protected void onCreate(Bundle savedInstanceState) {
            super.onCreate(savedInstanceState);
            setContentView(R.layout.%s);
        }
    } `, manifest.Package, className, layoutName)

	return template

}
func ChooseActivity() error {
	menu := climenu.NewButtonMenu("Activity", "Choose a activity template to generate ")

	menu.AddMenuItem("Empty", "empty")

	action, escaped := menu.Run()

	if escaped {
		os.Exit(0)
	}

	switch action {
	case "empty":
		genEmptyActivity()
		break
	}

	return nil

}

func genEmptyActivity() {
	activityName := climenu.GetText("ActivityName", "nothing")

	layout, layoutName := baseTemplateLayout(activityName)
	javaCode := baseJavaLayout(activityName, layoutName)

	fmt.Println(layout)
	fmt.Println(javaCode)
}
