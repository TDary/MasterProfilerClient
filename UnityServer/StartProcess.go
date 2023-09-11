package UnityServer

import (
	"MasterClient/Logs"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

//启动csv解析进程
func StartAnalyzeForCsvProfiler(data AnalyzeData, analyzeProject string, num int) int {
	var rawPath strings.Builder
	var analyzePath string
	var logPath strings.Builder
	var csvPath strings.Builder
	//————————————————————————————————rwaPath
	rawPath.WriteString(config.FilePath)
	rawPath.WriteString("/")
	rawPath.WriteString(data.UUID)
	rawPath.WriteString("/")
	rawPath.WriteString(data.RawFile)
	analyzePath = strings.Split(rawPath.String(), ".")[0]
	_, err := os.Stat(analyzePath)
	if err != nil {
		//文件夹不存在 需要创建一个
		os.Mkdir(analyzePath, 0755)
	}
	//——————————————————————————————————logPath
	logPath.WriteString(config.FilePath)
	logPath.WriteString("/")
	logPath.WriteString(data.UUID)
	logPath.WriteString("/")
	logPath.WriteString(data.RawFile)
	logPath.WriteString(".log")
	//————————————————————————————————————csvPath
	csvPath.WriteString(analyzePath)
	csvPath.WriteString("/")
	csvPath.WriteString(data.RawFile)
	csvPath.WriteString(".csv")
	analyzeType := data.AnalyzeType
	//判断unity版本然后进行选取
	Unity_Name := GetUnityVerison(data)
	if Unity_Name == "" {
		Logs.Loggers().Print("无可使用Unity版本----")
		return -1
	}
	var Startargs strings.Builder
	Startargs.WriteString(Unity_Name)
	Startargs.WriteString(" -quit -batchmode -nographics ")
	Startargs.WriteString("-projectPath ")
	Startargs.WriteString(analyzeProject)
	Startargs.WriteString(" -executeMethod Entrance.EntranceParseBegin ")
	Startargs.WriteString("-logFile ")
	Startargs.WriteString(logPath.String())
	Startargs.WriteString(" -rawPath ")
	Startargs.WriteString(rawPath.String())
	Startargs.WriteString(" -csvPath ")
	Startargs.WriteString(csvPath.String())
	Startargs.WriteString(" -analyzeType ")
	Startargs.WriteString(analyzeType)
	Startargs.WriteString(" -UUID ")
	Startargs.WriteString(data.UUID)
	Startargs.WriteString(" -ServerUrl ")
	Startargs.WriteString(config.ClientUrl.Ip)
	Startargs.WriteString(":")
	Startargs.WriteString(config.ClientUrl.Port)
	cmd := exec.Command("cmd.exe", "/c", "start "+Startargs.String())
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	er := cmd.Start()
	if er != nil { // 运行命令
		Logs.Loggers().Print(er.Error())
	}
	SuccessBegin(data, num)
	return num
}

// 启动函数堆栈解析进程
func StartAnalyzeForFunProfiler(data AnalyzeData, analyzeProject string, num int) int {
	var rawPath strings.Builder
	var analyzePath string
	var logPath strings.Builder
	var funPath strings.Builder
	var funRowPath strings.Builder
	var renderRowPath strings.Builder
	//——————————————————————————————————————rawPath
	rawPath.WriteString(config.FilePath)
	rawPath.WriteString("/")
	rawPath.WriteString(data.UUID)
	rawPath.WriteString("/")
	rawPath.WriteString(data.RawFile)
	analyzePath = strings.Split(rawPath.String(), ".")[0]
	_, err := os.Stat(analyzePath)
	if err != nil {
		//文件夹不存在 需要创建一个
		os.Mkdir(analyzePath, 0755)
	}
	//——————————————————————————————————logPath
	logPath.WriteString(config.FilePath)
	logPath.WriteString("/")
	logPath.WriteString(data.UUID)
	logPath.WriteString("/")
	logPath.WriteString(data.RawFile)
	logPath.WriteString(".log")
	//————————————————————————————————————funPath
	funPath.WriteString(analyzePath)
	funPath.WriteString("/")
	funPath.WriteString(data.RawFile)
	funPath.WriteString("_fun.bin")
	//————————————————————————————————————funRowPath
	funRowPath.WriteString(analyzePath)
	funRowPath.WriteString("/")
	funRowPath.WriteString(data.RawFile)
	funRowPath.WriteString("_funrow.bin")
	//————————————————————————————————————funrenrowPath
	renderRowPath.WriteString(analyzePath)
	renderRowPath.WriteString("/")
	renderRowPath.WriteString(data.RawFile)
	renderRowPath.WriteString("_renderrow.bin")
	analyzeType := data.AnalyzeType
	shield := "false"
	//判断unity版本然后进行选取
	Unity_Name := GetUnityVerison(data)
	if Unity_Name == "" {
		Logs.Loggers().Print("无可使用Unity版本----")
		return -1
	}
	var Startargs strings.Builder
	Startargs.WriteString(Unity_Name)
	Startargs.WriteString(" -quit -batchmode -nographics ")
	Startargs.WriteString("-projectPath ")
	Startargs.WriteString(analyzeProject)
	Startargs.WriteString(" -executeMethod Entrance.EntranceParseBegin ")
	Startargs.WriteString("-logFile ")
	Startargs.WriteString(logPath.String())
	Startargs.WriteString(" -rawPath ")
	Startargs.WriteString(rawPath.String())
	Startargs.WriteString(" -funPath ")
	Startargs.WriteString(funPath.String())
	Startargs.WriteString(" -funrowPath ")
	Startargs.WriteString(funRowPath.String())
	Startargs.WriteString(" -funrenderrowPath ")
	Startargs.WriteString(renderRowPath.String())
	Startargs.WriteString(" -funhashPath ")
	Startargs.WriteString(" -analyzeType ")
	Startargs.WriteString(analyzeType)
	Startargs.WriteString(" -shieldSwitch ")
	Startargs.WriteString(shield)
	Startargs.WriteString(" -UUID ")
	Startargs.WriteString(data.UUID)
	Startargs.WriteString(" -ServerUrl ")
	Startargs.WriteString(config.ClientUrl.Ip)
	Startargs.WriteString(":")
	Startargs.WriteString(config.ClientUrl.Port)
	cmd := exec.Command("cmd.exe", "/c", "start "+Startargs.String())
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	er := cmd.Start()
	if er != nil { // 运行命令
		Logs.Loggers().Print(er.Error())
	}
	SuccessBegin(data, num)
	return num
}

// 占用一个Unity工程
func GetUnityProject() (string, int) {
	taskMutex.TryLock() //加锁防止多个线程同时启动时获取到同一个Unity解析工程
	for key, val := range config.UnityProjectPath {
		if val.State {
			config.UnityProjectPath[key].State = false
			taskMutex.Unlock()
			return val.Path, val.Numb
		}
	}
	taskMutex.Unlock()
	return "", -1
}

// 释放Unity工程使用状态
func RleaseUnityProject(num int) {
	for key, val := range config.UnityProjectPath {
		if val.Numb == num {
			config.UnityProjectPath[key].State = true
			break
		}
	}
}

// 检查Unity版本
func GetUnityVerison(data AnalyzeData) string {
	if data.UnityVersion != "" {
		subVer := data.UnityVersion
		subVer = subVer[:6]
		for _, val := range config.UnityPath {
			Version := val.Version
			Version = Version[:6]
			if subVer == Version { //判断前面的字符版本大版本2021.3即可
				return val.Path
			}
		}
	}
	return ""
}

//检查是否有空闲可解析工程
func CheckFreeAnalyze() bool {
	for _, val := range config.UnityProjectPath {
		if val.State {
			return true
		}
	}
	return false
}
