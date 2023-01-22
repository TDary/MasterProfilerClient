package UnityServer

import (
	"MasterClient/Logs"
	"os/exec"
)

// 启动解析进程
func StartAnalyze(data AnalyzeData) (int, string) {
	logPath := config.FilePath + "/" + data.UUID + "/" + data.RawFile + ".log"
	rawPath := config.FilePath + "/" + data.UUID + "/" + data.RawFile
	csvPath := config.FilePath + "/" + data.UUID + "/" + data.RawFile + ".csv"
	funjsonPath := config.FilePath + "/" + data.UUID + "/" + data.RawFile + "_fun.json"
	funRowPath := config.FilePath + "/" + data.UUID + "/" + data.RawFile + "_funrow.json"
	renderRowPath := config.FilePath + "/" + data.UUID + "/" + data.RawFile + "_renderrow.json"
	funhashPath := config.FilePath + "/" + data.UUID + "/" + data.RawFile + "_funhash.json"
	shield := "false"
	//判断unity版本然后进行选取
	Unity_Name := GetUnityVerison(data)
	if Unity_Name == "" {
		Logs.Loggers().Print("无可使用Unity版本----")
		return -1, csvPath
	}
	//寻找可用的Unity工程进行解析
	UnityPjPath, num := GetUnityProject()
	if UnityPjPath == "" {
		Logs.Loggers().Print("无可用Unity工程----")
		return -1, csvPath
	}
	argu := " -quit -batchmode -nographics "
	argu = argu + "-projectPath " + UnityPjPath + " "
	argu = argu + "-executeMethod Entrance.EntranceParseBegin "
	argu = argu + "-logFile " + logPath + " "
	argu = argu + "-rawPath " + rawPath + " "
	argu = argu + "-csvPath " + csvPath + " "
	argu = argu + "-funjsonPath " + funjsonPath + " "
	argu = argu + "-funrowjsonPath " + funRowPath + " "
	argu = argu + "-funrenderrowjsonPath " + renderRowPath + " "
	argu = argu + "-funhashPath " + funhashPath + " "
	argu = argu + "-shieldSwitch " + shield + " "
	argu = Unity_Name + argu
	cmd := exec.Command("cmd.exe", "/c", "start "+argu)
	er := cmd.Start()
	if er != nil { // 运行命令
		Logs.Loggers().Print(er)
	}
	SuccessBegin(data, num)
	return cmd.Process.Pid, csvPath
}

// 占用一个Unity工程
func GetUnityProject() (string, int) {
	config.lock.Lock() //加锁防止多个线程同时启动时获取到同一个Unity解析工程
	defer config.lock.Unlock()
	for _, val := range config.UnityProjectPath {
		if val.State {
			val.State = false
			return val.Path, val.Numb
		}
	}
	return "", 0
}

// 释放Unity工程使用状态
func RleaseUnityProject(num int) {
	config.lock.Lock() //加锁
	defer config.lock.Unlock()
	for _, val := range config.UnityProjectPath {
		if val.Numb == num {
			val.State = true
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
