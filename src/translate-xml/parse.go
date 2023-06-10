package translate_xml

import (
	"encoding/xml"
  "github.com/go-xmlfmt/xmlfmt"
	"log"
	"strings"
)

func ParseManifest(data []byte) (Manifest, error) {
	var manifest Manifest

	err := xml.Unmarshal(data, &manifest)
	if err != nil {
		return manifest, err
	}

	return manifest, nil
}

func ToManifestFile(manifesto Manifest) ManifestFile {
	var manifestFile ManifestFile

	manifestFile.Android = manifesto.Android
	manifestFile.Tools = manifesto.Tools
	manifestFile.copyPermittionsToFile(manifesto.Permissions)
	manifestFile.copyFeaturesToFile(manifesto.Features)
	manifestFile.copyApplication(manifesto.Application)
	manifestFile.copyActivities(manifesto.Application.Activities)

	return manifestFile
}

func (manifest ManifestFile) GetXmlFile() string {
	a, err := xml.Marshal(manifest)

	if err != nil {
		log.Fatal(err)
	}

	xml := string(a)
  ChangeActivityTagName(&xml)

	return xmlfmt.FormatXML(xml, "\t", " ")
}

func ChangeActivityTagName(xml *string) {
	*xml = strings.ReplaceAll(*xml, "<ManifestFile", "<manifest")
	*xml = strings.ReplaceAll(*xml, "</ManifestFile>", "</manifest>")
}
