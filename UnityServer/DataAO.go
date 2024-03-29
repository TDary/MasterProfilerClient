package UnityServer

import "sync"

var config Config //客户端解析配置
var taskMutex sync.Mutex
var removeLock sync.Mutex
var m_State MachineState

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
	RawFileName   string
	UnityVersion  string
	AnalyzeBucket string
	AnalyzeNum    int
	AnalyzeType   string
}

type MinioConfig struct {
	MinioServerPath string
	MinioBucket     string
	MinioRawBucket  string
	UserName        string
	PassWord        string
}

type Config struct {
	FilePath         string
	UnityPath        []UnityConfig
	UnityProjectPath []UnityProject
	Minioconfig      MinioConfig
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

type MachineState struct {
	Ip    string
	State string //三个状态：忙碌，空闲和离线  idle:空闲 busy:繁忙 out:离线  离线通常是在合并服务器上显示，如合并服务器启动了，但是解析器还没启动
	Num   int    //可用的解析进程数量
}
