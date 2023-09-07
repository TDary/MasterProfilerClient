package ParseServer

import (
	"MasterClient/Logs"
	"MasterClient/UnityServer"
	"strings"
	"time"
)

//获取解析成功的数据
func AnalyzeSuccess(data string) {
	ParseSuccessData(data)
}

// 开始启动解析前的准备工作
func Analyze(data string) {
	var getdata UnityServer.AnalyzeData
	getdata = ParseData(data, getdata) //解析传入的数据
	successDownLoad := UnityServer.DownLoadFile(getdata)
	if successDownLoad {
		projectID := UnityServer.StartAnalyze(getdata)
		CheckProcessState(getdata, projectID) //监控解析进程
		//完成解析
		Logs.Loggers().Print("解析流程完毕----")
		UnityServer.SuccessAnalyze(getdata)
		//完成解析消息回传发送准备
		UnityServer.GetSucessData(getdata.RawFile, getdata.UUID)
	} else {
		Logs.Loggers().Print("下载源文件" + getdata.RawFile + "失败----")
	}
}

//循环检测进程
func CheckProcessState(getdata UnityServer.AnalyzeData, id int) {
	var count int
	for {
		time.Sleep(10 * time.Second)
		if CheckAnalyzeState(getdata) == "success" {
			Logs.Loggers().Print("UUID:" + getdata.UUID + ",rawFile:" + getdata.RawFile + "解析成功----")
			UnityServer.RleaseUnityProject(id)
			break
		} else if CheckAnalyzeState(getdata) == "failed" {
			Logs.Loggers().Print("UUID:" + getdata.UUID + ",rawFile:" + getdata.RawFile + "解析失败----")
			UnityServer.RleaseUnityProject(id)
			break
		} else {
			//超过一定的等待时间即代表着已经解析出问题了
			if count >= 12 {
				//释放unity解析池组
				UnityServer.RleaseUnityProject(id)
				break
			}
		}
		count++
	}
}

//检查解析完毕的数组是否有对应的
func CheckAnalyzeState(getdata UnityServer.AnalyzeData) string {
	logicMutex.Lock()
	defer logicMutex.Unlock()
	for i := 0; i < len(analyzeData); i++ {
		if analyzeData[i].RawFile == getdata.RawFile && analyzeData[i].UUID == getdata.UUID &&
			analyzeData[i].AnalyzeType == getdata.AnalyzeType && analyzeData[i].State == "success" {
			analyzeData = append(analyzeData[:i], analyzeData[i+1:]...)
			return "success"
		} else if analyzeData[i].RawFile == getdata.RawFile && analyzeData[i].UUID == getdata.UUID &&
			analyzeData[i].AnalyzeType == getdata.AnalyzeType && analyzeData[i].State == "failed" {
			analyzeData = append(analyzeData[:i], analyzeData[i+1:]...)
			return "failed"
		}
	}
	return "wait"
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
		} else if strings.Contains(current[i], "analyzeType") {
			cdata := strings.Split(current[i], "=")
			gdata.AnalyzeType = cdata[1]
		}
	}
	return gdata
}

//将回传的成功http消息进行处理
func ParseSuccessData(data string) {
	var gdata UnityServer.AnalyzeState
	current := strings.Split(data, "&")
	for i := 0; i < len(current); i++ {
		if strings.Contains(current[i], "uuid") {
			cdata := strings.Split(current[i], "=")
			gdata.UUID = cdata[1]
		} else if strings.Contains(current[i], "rawfile") {
			cdata := strings.Split(current[i], "=")
			gdata.RawFile = cdata[1]
		} else if strings.Contains(current[i], "anaType") {
			cdata := strings.Split(current[i], "=")
			gdata.AnalyzeType = cdata[1]
		} else if strings.Contains(current[i], "state") {
			cdata := strings.Split(current[i], "=")
			gdata.State = cdata[1]
		}
	}
	analyzeData = append(analyzeData, gdata)
}
