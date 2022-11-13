package UnityServer

import (
	"MasterClient/Logs"
	"os/exec"
)

//启动解析进程
func StartAnalyze(data AnalyzeData) (int, string) {
	logPath := config.FilePath + "/" + data.UUID + "/" + data.RawFile + ".log"
	rawPath := config.FilePath + "/" + data.UUID + "/" + data.RawFile + ".raw" //TODO:前提要下载文件成功
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
	argu := "-quit -batchmode -nographics "
	argu = argu + "-projectPath " + config.UnityProjectPath + " " // todo:多个工程进行判断是否已使用
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
	er := cmd.Run()
	if er != nil { // 运行命令
		Logs.Loggers().Print(er)
	}
	return cmd.Process.Pid, csvPath
}

func GetUnityVerison(data AnalyzeData) string {
	subVer := data.UnityVersion
	subVer = subVer[:6]
	for _, val := range config.UnityPath {
		Version := val.Version
		Version = Version[:6]
		if subVer == Version { //TODO：判断前面的字符版本即可
			return val.Path
		}
	}
	return ""
}
