package main

import (
	"fmt"

	"github.com/lu-css/android-tui/src/files"
	"github.com/lu-css/android-tui/src/translate-xml"
)

func main() {

	manifestPath := files.FindManifestPath()
	data, err := files.ReadManifest(manifestPath)

	if err != nil {
		fmt.Println(err)
	}

	manifest, _ := translate_xml.ParseManifest(data)

	files.PrinParsedManifest(manifest)
}
