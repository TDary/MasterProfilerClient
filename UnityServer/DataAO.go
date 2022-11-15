package UnityServer

var config Config        //客户端解析配置
var handCase HandingCase //正在进行解析的案例记录

type HandingCase struct {
	Case []OneCase
}

type OneCase struct {
	UUID    string
	RawFile string
	Numb    int //对应解析工程的Numb
}

type AnalyzeData struct {
	UUID          string
	RawFile       string
	UnityVersion  string
	AnalyzeBucket string
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
