package ParseServer

import (
	"MasterClient/Logs"
	"MasterClient/Minio"
	"MasterClient/RabbitMqServer"
	"MasterClient/UnityServer"
	"os"
	"strings"
	"time"
)

//获取解析成功的数据
func AnalyzeSuccess(data string) {
	ParseSuccessData(data)
}

//从http消息中获取任务并存入队列中
func GetAnalyzeMes(data string) {
	filepath := "./AnalyTask"
	_, err := os.Stat(filepath)
	if err != nil {
		os.Mkdir(filepath, 0755)
	}
	RabbitMqServer.PutData(filepath+"/ParseQue", data)
}

//检测队列中的解析任务并拿出来进行解析
func AnalyzeRangeCheck() {
	taskPath := "./AnalyTask/ParseQue"
	for {
		//有空闲的才去分配解析
		_, err := os.Stat(taskPath)
		if err == nil {
			if UnityServer.CheckFreeAnalyze() {
				taskdata := RabbitMqServer.GetData(taskPath)
				if taskdata != "" {
					go Analyze(taskdata)
				}
			}
		} else {
			//Logs.Loggers().Print("无待解析的队列文件----")
		}
		time.Sleep(3 * time.Second) //每隔3秒开启下一个解析任务
	}
}

// 开始启动解析进程
func Analyze(data string) {
	Logs.Loggers().Print("开始解析任务----")
	var getdata UnityServer.AnalyzeData
	project, num := UnityServer.GetUnityProject()
	getdata = ParseData(data, getdata) //解析传入的数据
	successDownLoad := UnityServer.DownLoadRawFile(getdata)
	if successDownLoad {
		projectID := UnityServer.StartAnalyze(getdata, project, num)
		CheckProcessState(getdata, projectID) //监控解析进程
		Logs.Loggers().Print("解析流程完毕----")
	} else {
		Logs.Loggers().Print("下载源文件" + getdata.RawFile + "失败----")
	}
}

//循环检测进程
func CheckProcessState(getdata UnityServer.AnalyzeData, id int) {
	getdata.AnalyzeNum = id //把工程id赋值
	var count int
	for {
		time.Sleep(5 * time.Second)
		state := CheckAnalyzeState(getdata)
		if state == "success" {
			Logs.Loggers().Print("UUID:" + getdata.UUID + ",rawFile:" + getdata.RawFile + "解析成功----")
			UnityServer.RleaseUnityProject(id)
			UnityServer.SuccessAnalyze(getdata)
			UploadSuccessedData(getdata)
			//完成解析消息回传发送
			UnityServer.GetSucessData(getdata.RawFile, getdata.UUID)
			break
		} else if state == "failed" {
			Logs.Loggers().Print("UUID:" + getdata.UUID + ",rawFile:" + getdata.RawFile + "解析失败----")
			UnityServer.RleaseUnityProject(id)
			UnityServer.SuccessAnalyze(getdata)
			UploadSuccessedData(getdata)
			//解析失败消息上报
			UnityServer.SendFailMessage(getdata.RawFile, getdata.UUID)
			break
		} else {
			//超过一定的等待时间即代表着已经解析出问题了
			if count >= 24 {
				//释放unity解析池组
				UnityServer.RleaseUnityProject(id)
				UnityServer.SuccessAnalyze(getdata)
				//解析失败消息上报
				UnityServer.SendFailMessage(getdata.RawFile, getdata.UUID)
				break
			}
		}
		count++
	}
}

//检查解析完毕的数组是否有对应的
func CheckAnalyzeState(getdata UnityServer.AnalyzeData) string {
	logicMutex.TryLock()
	for i := 0; i < len(analyzeData); i++ {
		if analyzeData[i].RawFile == getdata.RawFile && analyzeData[i].UUID == getdata.UUID &&
			analyzeData[i].AnalyzeType == getdata.AnalyzeType && analyzeData[i].State == "success" {
			analyzeData = append(analyzeData[:i], analyzeData[i+1:]...)
			logicMutex.Unlock()
			return "success"
		} else if analyzeData[i].RawFile == getdata.RawFile && analyzeData[i].UUID == getdata.UUID &&
			analyzeData[i].AnalyzeType == getdata.AnalyzeType && analyzeData[i].State == "failed" {
			analyzeData = append(analyzeData[:i], analyzeData[i+1:]...)
			logicMutex.Unlock()
			return "failed"
		}
	}
	logicMutex.Unlock()
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

//上传解析完成的文件
func UploadSuccessedData(getdata UnityServer.AnalyzeData) {
	uploadMutex.TryLock()
	if getdata.AnalyzeType == "simple" {
		objName := getdata.UUID + "/" + getdata.RawFile + ".csv"
		contentType := "application/csv"
		fPath := UnityServer.GetConfig().FilePath + "/" + getdata.UUID + "/" + getdata.RawFile + ".csv"
		Minio.UploadFile(objName, fPath, contentType)
	} else if getdata.AnalyzeType == "funprofiler" {
		funObjName := getdata.UUID + "/" + getdata.RawFile + "_fun.bin"
		funrowObjName := getdata.UUID + "/" + getdata.RawFile + "_funrow.bin"
		renrowObjName := getdata.UUID + "/" + getdata.RawFile + "_renderrow.bin"
		funhashObjName := getdata.UUID + "/" + getdata.RawFile + "_funhash.bin"
		funPath := UnityServer.GetConfig().FilePath + "/" + getdata.UUID + "/" + getdata.RawFile + "_fun.bin"
		funrowPath := UnityServer.GetConfig().FilePath + "/" + getdata.UUID + "/" + getdata.RawFile + "_funrow.bin"
		renrowPath := UnityServer.GetConfig().FilePath + "/" + getdata.UUID + "/" + getdata.RawFile + "_renderrow.bin"
		funhashPath := UnityServer.GetConfig().FilePath + "/" + getdata.UUID + "/" + getdata.RawFile + "_funhash.bin"
		contentType := "application/bin"
		Minio.UploadFile(funObjName, funPath, contentType)
		Minio.UploadFile(funrowObjName, funrowPath, contentType)
		Minio.UploadFile(renrowObjName, renrowPath, contentType)
		Minio.UploadFile(funhashObjName, funhashPath, contentType)
	} else {
		//todo:deep
	}
	uploadMutex.Unlock()
}
