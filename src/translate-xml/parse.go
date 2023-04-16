package translate_xml

import (
	"encoding/xml"
)

func ParseManifest(data []byte) (Manifest, error) {
	// Create an empty Book struct
	var manifest Manifest

	// Unmarshal the XML data into the struct
	err := xml.Unmarshal(data, &manifest)
	if err != nil {
		return manifest, err
	}

	return manifest, nil
}
