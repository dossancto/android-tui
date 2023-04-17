package translate_xml_test

import (
	"testing"

	translate_xml "github.com/lu-css/android-tui/src/translate-xml"
)

var testManifest = []byte(`<?xml version="1.0" encoding="utf-8"?>
<manifest xmlns:android="http://schemas.android.com/apk/res/android"
    xmlns:tools="http://schemas.android.com/tools"
    package="com.example.testapp">

    <uses-feature
        android:name="android.hardware.sensor.proximity"
        android:required="true" />

    <uses-permission android:name="android.permission.ACCESS_FINE_LOCATION" />
    <uses-permission android:name="android.permission.ACCESS_COARSE_LOCATION" />
    <uses-permission android:name="android.permission.ACCESS_NETWORK_STATE" />
    <uses-permission android:name="android.permission.INTERNET" />
    <uses-permission android:name="android.permission.WRITE_EXTERNAL_STORAGE" />
    <uses-permission android:name="android.permission.READ_EXTERNAL_STORAGE" />

    <application
        android:allowBackup="true"
        android:dataExtractionRules="@xml/data_extraction_rules"
        android:fullBackupContent="@xml/backup_rules"
        android:icon="@drawable/iconn"
        android:label="@string/app_name"
        android:roundIcon="@drawable/iconn"
        android:supportsRtl="true"
        android:theme="@style/Theme.AppCompat.Light.NoActionBar"
        tools:targetApi="31">
        <activity
            android:name=".activity1"
            android:exported="false" />
        <activity
            android:name=".activity2"
            android:exported="false" />
        <activity
            android:name=".activity3"
            android:exported="false" />
        <activity
            android:name=".activity4"
            android:exported="false" />
        <activity
            android:name=".activity5"
            android:exported="false" />
        <activity
            android:name=".activity6"
            android:exported="false" />
        <activity
            android:name=".MainActivity"
            android:exported="true">
            <intent-filter>
                <action android:name="android.intent.action.MAIN" />

                <category android:name="android.intent.category.LAUNCHER" />
            </intent-filter>
        </activity>

        <meta-data
            android:name="preloaded_fonts"
            android:resource="@array/preloaded_fonts" />
    </application>

</manifest>`)

func TestParseManifest(t *testing.T) {

	expectManifest := translate_xml.Manifest{
		Tools:       "http://schemas.android.com/tools",
		Package:     "com.example.testapp",
		Application: translate_xml.Application{
      false,
    },
	}
	manifest, err := translate_xml.ParseManifest(testManifest)

	if err != nil {
		t.Errorf("Something went wrong %s", err)
	}

	if manifest != expectManifest {
		t.Errorf("Unespected manifest decode")
	}

}
