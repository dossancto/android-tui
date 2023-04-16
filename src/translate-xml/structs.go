package translate_xml

type ActivityMetaData struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"valeu,attr"`
}

type Activity struct {
	MetaData ActivityMetaData `xml:"meta-data"`
	Exported bool             `xml:"exported,attr"`
	Name     string           `xml:"name,attr"`
}

type Application struct {
	AllowBackup         string     `xml:"allowBackup,attr"`
	DataExtractionRules string     `xml:"dataExtractionRules,attr"`
	FullBackupContent   string     `xml:"fullBackupContent,attr"`
	Icon                string     `xml:"icon,attr"`
	Label               string     `xml:"label,attr"`
	RoundIcon           string     `xml:"roundIcon,attr"`
	SupportsRtl         string     `xml:"supportsRtl,attr"`
	TargetApi           int        `xml:"targetApi,attr"`
	Theme               string     `xml:"theme,attr"`
	Activities          []Activity `xml:"activity"`
}

type UsesPermission struct {
	Name string `xml:"name,attr"`
}

type UsesFeature struct {
	Name     string `xml:"name,attr"`
	Required bool   `xml:"required,attr"`
}

type Manifest struct {
	Tools       string           `xml:"tools,attr"`
	Package     string           `xml:"package,attr"`
	Permissions []UsesPermission `xml:"uses-permission"`
	Features    []UsesFeature    `xml:"uses-feature"`
	Application Application      `xml:"application"`
}
