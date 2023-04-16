package translate_xml

import (
	"encoding/xml"
)

func ParseManifest(data []byte) (Manifest, error) {
	var manifest Manifest

	err := xml.Unmarshal(data, &manifest)
	if err != nil {
		return manifest, err
	}

	return manifest, nil
}
