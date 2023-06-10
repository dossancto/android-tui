package translate_xml

type ManifestFile struct {
	Android     string               `xml:"xmlns:android,attr"`
	Tools       string               `xml:"xmlns:tools,attr"`
	Permissions []UsesPermissionFile `xml:"uses-permission"`
	Features    []UsesFeatureFile    `xml:"uses-feature"`
	Application ApplicationFile      `xml:"application"`
}

type ActivityMetaDataFile struct {
	Name  string `xml:"android:name,attr"`
	Value string `xml:"android:value,attr"`
}

type IntentActionFile struct {
	Name string `xml:"android:name,attr"`
}

type IntentCategoryFile struct {
	Name string `xml:"android:name,attr"`
}

type IntentFilterFile struct {
	Action   IntentActionFile   `xml:"action"`
	Category IntentCategoryFile `xml:"category"`
}

type ActivityFile struct {
	MetaData ActivityMetaDataFile `xml:"meta-data"`
	Exported bool                 `xml:"android:exported,attr"`
	Name     string               `xml:"android:name,attr"`
	Filter   IntentFilterFile     `xml:"intent-filter"`
}

type ApplicationFile struct {
	AllowBackup         bool           `xml:"android:allowBackup,attr"`
	DataExtractionRules string         `xml:"android:dataExtractionRules,attr"`
	FullBackupContent   string         `xml:"android:fullBackupContent,attr"`
	Icon                string         `xml:"android:icon,attr"`
	Label               string         `xml:"android:label,attr"`
	RoundIcon           string         `xml:"android:roundIcon,attr"`
	SupportsRtl         string         `xml:"android:supportsRtl,attr"`
	TargetApi           int            `xml:"tools:targetApi,attr"`
	Theme               string         `xml:"android:theme,attr"`
	Activities          []ActivityFile `xml:"activity"`
}

type UsesPermissionFile struct {
	Name string `xml:"name,attr"`
}

type UsesFeatureFile struct {
	Name     string `xml:"android:name,attr"`
	Required bool   `xml:"required,attr"`
}
