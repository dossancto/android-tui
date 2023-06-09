package translate_xml

import (
	"encoding/xml"
	"log"
)

func ParseManifest(data []byte) (Manifest, error) {
	var manifest Manifest

	err := xml.Unmarshal(data, &manifest)
	if err != nil {
		return manifest, err
	}

	return manifest, nil
}

func ToManifestFile(manifesto Manifest) manifest {
	var manifestFile manifest

	manifestFile.Android = manifesto.Android
	manifestFile.Tools = manifesto.Tools
	manifestFile.copyPermittionsToFile(manifesto.Permissions)
	manifestFile.copyFeaturesToFile(manifesto.Features)
	manifestFile.copyApplication(manifesto.Application)
	manifestFile.copyActivities(manifesto.Application.Activities)

	return manifestFile
}

func (manifest manifest) GetXmlFile() string {
	a, err := xml.Marshal(manifest)

	if err != nil {
		log.Fatal(err)
	}

	return string(a)
}
