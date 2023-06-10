package translate_xml

import (
	"errors"
	"fmt"
)

func (application Application) GetLaunchActitity() (Activity, error) {

	for _, activity := range application.Activities {
		if activity.Filter.Category.Name == "android.intent.category.LAUNCHER" {
			return activity, nil
		}
	}

	return Activity{}, errors.New("Launch Activity does not exists.")
}

func (manifest Manifest) PrinParsedManifest() {
	fmt.Println(manifest.Tools)
	fmt.Println(manifest.Package)

	println("\nFeatures\n")

	for _, features := range manifest.Features {
		s := fmt.Sprintf("%s - %t", features.Name, features.Required)
		fmt.Println(s)
	}

	println("\nPermittions\n")

	for _, features := range manifest.Permissions {
		s := fmt.Sprintf("%s", features.Name)
		fmt.Println(s)
	}

	println("\nApplication infos\n")
	b := manifest.Application

	fmt.Println(b.AllowBackup)
	fmt.Println(b.DataExtractionRules)
	fmt.Println(b.FullBackupContent)
	fmt.Println(b.Icon)
	fmt.Println(b.Label)
	fmt.Println(b.RoundIcon)
	fmt.Println(b.SupportsRtl)
	fmt.Println(b.Theme)
	fmt.Println(b.TargetApi)

	println("\nActivities infos\n")
	for _, ac := range b.Activities {
		fmt.Println("Name:", ac.Name)
		fmt.Println("Expoted:", ac.Exported)

		fmt.Println("Meta-data Name:", ac.MetaData.Name)
		fmt.Println("Meta-data Value:", ac.MetaData.Value)
		fmt.Println()

	}
}

func (m *ManifestFile) copyPermittionsToFile(permittions []UsesPermission) {
	for _, p := range permittions {
    permittion := UsesPermissionFile{
    	Name: p.Name,
    }
		m.Permissions = append(m.Permissions, permittion)
	}
}

func (m *ManifestFile) copyFeaturesToFile(features []UsesFeature) {
	for _, f := range features {
		feature := UsesFeatureFile{
			Name:     f.Name,
			Required: f.Required,
		}
		m.Features = append(m.Features, feature)
	}
}

func (m *ManifestFile) copyApplication(application Application) {
	m.Application.AllowBackup = application.AllowBackup
	m.Application.DataExtractionRules = application.DataExtractionRules
	m.Application.FullBackupContent = application.FullBackupContent
	m.Application.Icon = application.Icon
	m.Application.Label = application.Label
	m.Application.RoundIcon = application.RoundIcon
	m.Application.SupportsRtl = application.SupportsRtl
	m.Application.TargetApi = application.TargetApi
	m.Application.Theme = application.Theme
}

func (m *ManifestFile) copyActivities(activities []Activity) {
	for _, a := range activities {
		action := IntentActionFile{
			Name: a.Filter.Action.Name,
		}

		filter := IntentFilterFile{
			Action:   action,
			Category: IntentCategoryFile{},
		}

		metaData := ActivityMetaDataFile{
			Name:  a.MetaData.Name,
			Value: a.MetaData.Value,
		}

		activity := ActivityFile{
			Name:     a.Name,
			Exported: a.Exported,
			MetaData: metaData,
			Filter:   filter,
		}

		m.Application.Activities = append(m.Application.Activities, activity)
	}
}
