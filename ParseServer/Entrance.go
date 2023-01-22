package ParseServer

import (
	"MasterClient/Logs"
	"MasterClient/UnityServer"
	"io/ioutil"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

// 开始启动解析前的准备工作
func Analyze(data string) {
	var getdata UnityServer.AnalyzeData
	getdata = ParseData(data, getdata)
	successDownLoad := UnityServer.DownLoadFile(getdata)
	if successDownLoad {
		processID, csvPath := UnityServer.StartAnalyze(getdata)
		CheckProcessState(processID, getdata, csvPath) //监控解析进程
		//完成解析
		UnityServer.SuccessAnalyze(getdata)
		//完成解析消息回传发送准备
		UnityServer.GetSucessData(getdata.RawFile, getdata.UUID, csvPath)
	} else {
		Logs.Loggers().Print("下载源文件" + getdata.RawFile + "失败----")
	}
}

func CheckProcessState(pidID int, getdata UnityServer.AnalyzeData, csvPath string) {
	for true {
		if CheckPid(pidID, getdata, csvPath) {
			time.Sleep(10 * time.Second)
		} else {
			break
		}
	}
}

func CheckPid(pidID int, getdata UnityServer.AnalyzeData, csvpath string) bool {
	pid := strconv.Itoa(pidID)
	cmd := exec.Command("cmd", "/C", "tasklist|findstr "+pid)
	output, _ := cmd.Output()
	if output != nil {
		Logs.Loggers().Print("进程：" + pid + "解析仍在进行中----")
		Logs.Loggers().Print("任务：检测----")
		//检查文件是否已经全部完成
		var data, _ = ioutil.ReadFile(csvpath)
		if data != nil {
			Logs.Loggers().Print("csv文件已解析成功----")
			return false
		}
	}
	return true
}

// 将回传的http消息进行处理
func ParseData(data string, gdata UnityServer.AnalyzeData) UnityServer.AnalyzeData {
	current := strings.Split(data, "&")
	for i := 0; i < len(current); i++ {
		if strings.Contains(current[i], "uuid") {
			cdata := strings.Split(current[i], "=")
			gdata.UUID = cdata[1]
		} else if strings.Contains(current[i], "rawfile") {
			cdata := strings.Split(current[i], "=")
			gdata.RawFile = cdata[1]
		} else if strings.Contains(current[i], "unityversion") {
			cdata := strings.Split(current[i], "=")
			gdata.UnityVersion = cdata[1]
		} else if strings.Contains(current[i], "analyzebucket") {
			cdata := strings.Split(current[i], "=")
			gdata.AnalyzeBucket = cdata[1]
		}
	}
	return gdata
}
