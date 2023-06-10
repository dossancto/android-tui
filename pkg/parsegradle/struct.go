package parsegradle

type AndroidBuildTypesReleases struct {
	minifyEnabled bool
}
type AndroidBuildTypes struct {
	Releases AndroidBuildTypesReleases
}

type AndroidCompileOptions struct {
	SourceCompatibility string
	TargetCompatibility string
}

type AndroidDefaultConfig struct {
	ApplicationId string
	MinSdk        int
	TargetSdk     int
	VersionCode   int
	VersionName   string

	testInstrumentationRunner string
}

type AndroidConfig struct {
	Namespace  string
	CompileSdk int

	DefaultConfig  AndroidDefaultConfig
	BuildTypes     AndroidBuildTypes
	CompileOptions AndroidCompileOptions
}

type DependenciesConfig struct {
	Implementations           []string
	TestImplementation        []string
	androidTestImplementation []string
}

type GradleProject struct {
	Android      AndroidConfig
}
