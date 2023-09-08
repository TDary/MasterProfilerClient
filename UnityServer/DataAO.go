package UnityServer

import "sync"

var config Config //客户端解析配置
var taskMutex sync.Mutex
var removeLock sync.Mutex

type OneCase struct {
	UUID    string
	RawFile string
	Numb    int //对应解析工程的Numb
}

type AnalyzeState struct {
	UUID        string
	RawFile     string
	AnalyzeType string
	State       string
}

type AnalyzeData struct {
	UUID          string
	RawFile       string
	UnityVersion  string
	AnalyzeBucket string
	AnalyzeNum    int
	AnalyzeType   string
}

type Config struct {
	FilePath         string
	UnityPath        []UnityConfig
	UnityProjectPath []UnityProject
	MinioServerPath  string
	MasterServerUrl  ServerConfig
	ClientUrl        ServerConfig
}

type UnityProject struct {
	Path  string
	Numb  int
	State bool
}

type UnityConfig struct {
	Version string
	Path    string
}

type ServerConfig struct {
	Ip   string
	Port string
}
