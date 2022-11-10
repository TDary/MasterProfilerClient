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
	UnityPath        []string
	UnityProjectPath string
	MinioServerPath  string
}
