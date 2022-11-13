package UnityServer

var config Config

type AnalyzeData struct {
	UUID          string
	RawFile       string
	UnityVersion  string
	AnalyzeBucket string
}

type Config struct {
	FilePath         string
	UnityPath        []UnityConfig
	UnityProjectPath string
	MinioServerPath  string
	MasterServerUrl  string
	ClientUrl        string
}

type UnityConfig struct {
	Version string
	Path    string
}
