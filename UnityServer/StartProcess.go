package UnityServer

import (
	"MasterClient/Logs"
	"MasterClient/Tools"
	"encoding/json"
	"os/exec"
)

//启动解析进程
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
	UnityPjPath, num := CheckUnityProject()
	if UnityPjPath == "" {
		Logs.Loggers().Print("无可用Unity工程----")
		return -1, csvPath
	}
	argu := "-quit -batchmode -nographics "
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
	er := cmd.Run()
	if er != nil { // 运行命令
		Logs.Loggers().Print(er)
	}
	SuccessBegin(data, num)
	return cmd.Process.Pid, csvPath
}

//成功启动的话将进行写入进行中
func SuccessBegin(data AnalyzeData, num int) {
	for _, val := range handCase.Case {
		if val.UUID != "" {
			val.RawFile = data.RawFile
			val.UUID = data.UUID
			val.Numb = num
		}
	}
	//写入到文件中
	str, err := json.Marshal(handCase)
	if err != nil {
		Logs.Loggers().Print("转换json失败----", err)
	}
	Tools.WriteHandFile(string(str))
}

//成功解析完毕的话去除文件中进行的
func SuccessAnalyze(data AnalyzeData) {
	for _, val := range handCase.Case {
		if val.UUID == data.UUID {
			//首先将进程释放开,然后再进行清空
			OpenUnityProject(val.Numb)
			val.RawFile = ""
			val.UUID = ""
			val.Numb = 0
		}
	}
	//写入到文件中
	str, err := json.Marshal(handCase)
	if err != nil {
		Logs.Loggers().Print("转换json失败----", err)
	}
	Tools.WriteHandFile(string(str))
}

//检查Unity的解析工程文件夹
func CheckUnityProject() (string, int) {
	for _, val := range config.UnityProjectPath {
		if !val.State {
			val.State = true
			return val.Path, val.Numb
		}
	}
	return "", 0
}

//重新打开解析工程状态
func OpenUnityProject(num int) {
	for _, val := range config.UnityProjectPath {
		if val.Numb == num {
			val.State = false
		}
	}
}

//检查Unity版本
func GetUnityVerison(data AnalyzeData) string {
	subVer := data.UnityVersion
	subVer = subVer[:6]
	for _, val := range config.UnityPath {
		Version := val.Version
		Version = Version[:6]
		if subVer == Version { //判断前面的字符版本大版本2021.3即可
			return val.Path
		}
	}
	return ""
}
