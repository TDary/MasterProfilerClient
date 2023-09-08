package UnityServer

import (
	"MasterClient/Logs"
	"os/exec"
	"strings"
	"syscall"
)

// 启动解析进程
func StartAnalyze(data AnalyzeData) int {
	logPath := config.FilePath + "/" + data.UUID + "/" + data.RawFile + ".log"
	rawPath := config.FilePath + "/" + data.UUID + "/" + data.RawFile
	csvPath := config.FilePath + "/" + data.UUID + "/" + data.RawFile + ".csv"
	funPath := config.FilePath + "/" + data.UUID + "/" + data.RawFile + "_fun.bin"
	funRowPath := config.FilePath + "/" + data.UUID + "/" + data.RawFile + "_funrow.bin"
	renderRowPath := config.FilePath + "/" + data.UUID + "/" + data.RawFile + "_renderrow.bin"
	funhashPath := config.FilePath + "/" + data.UUID + "/" + data.RawFile + "_funhash.bin"
	analyzeType := data.AnalyzeType
	shield := "false"
	//判断unity版本然后进行选取
	Unity_Name := GetUnityVerison(data)
	if Unity_Name == "" {
		Logs.Loggers().Print("无可使用Unity版本----")
		return -1
	}
	//寻找可用的Unity工程进行解析
	UnityPjPath, num := GetUnityProject()
	if UnityPjPath == "" {
		Logs.Loggers().Print("无可用Unity工程----")
		return -1
	}
	var Startargs strings.Builder
	Startargs.WriteString(Unity_Name)
	Startargs.WriteString(" -quit -batchmode -nographics ")
	Startargs.WriteString("-projectPath ")
	Startargs.WriteString(UnityPjPath)
	Startargs.WriteString(" -executeMethod Entrance.EntranceParseBegin ")
	Startargs.WriteString("-logFile ")
	Startargs.WriteString(logPath)
	Startargs.WriteString(" -rawPath ")
	Startargs.WriteString(rawPath)
	Startargs.WriteString(" -csvPath ")
	Startargs.WriteString(csvPath)
	Startargs.WriteString(" -funPath ")
	Startargs.WriteString(funPath)
	Startargs.WriteString(" -funrowPath ")
	Startargs.WriteString(funRowPath)
	Startargs.WriteString(" -funrenderrowPath ")
	Startargs.WriteString(renderRowPath)
	Startargs.WriteString(" -funhashPath ")
	Startargs.WriteString(funhashPath)
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
