package UnityServer

import (
	"MasterClient/Logs"
	"MasterClient/RabbitMqServer"
	"encoding/gob"
	"os"
	"strconv"
	"time"
)

//检查Unity版本
func CheckUnityVersion(data AnalyzeData) string {
	for i := 0; i < len(config.UnityPath); i++ {
		if data.UnityVersion == config.UnityPath[i].Version {
			return config.UnityPath[i].Path
		}
	}
	return ""
}

// 成功启动的话将进行写入进行中队列，独特的队列
func SuccessBegin(data AnalyzeData, num int) {
	data.AnalyzeNum = num
	filepath := "./analyzing"
	_, err := os.Stat(filepath)
	if err != nil {
		os.Mkdir(filepath, 0755)
	}
	currentNum := strconv.Itoa(num)
	writeFile := filepath + "/" + currentNum + ".gob"
	err = writeGob(writeFile, data)
	if err != nil {
		Logs.Loggers().Println(err)
	}
}

// 解析完毕的话去除文件中进行的
func SuccessAnalyze(data AnalyzeData) {
	removeLock.TryLock()
	var getdata AnalyzeData
	currentNum := strconv.Itoa(data.AnalyzeNum)
	getFile := "./analyzing" + "/" + currentNum + ".gob"
	err := readGob(getFile, &getdata)
	if err != nil {
		Logs.Loggers().Print(err)
	}
	if data.UUID == getdata.UUID && data.RawFile == getdata.RawFile && data.AnalyzeType == getdata.AnalyzeType {
		os.Remove(getFile)
	}
	removeLock.Unlock()
}

//序列化写入文件
func writeGob(filePath string, object interface{}) error {
	file, err := os.Create(filePath)
	if err == nil {
		encoder := gob.NewEncoder(file)
		encoder.Encode(object)
	}
	file.Close()
	return err
}

//反序列化
func readGob(filePath string, object interface{}) error {
	file, err := os.Open(filePath)
	if err == nil {
		decoder := gob.NewDecoder(file)
		err = decoder.Decode(object)
	}
	file.Close()
	return err
}

//获取当前解析状态
func GetAnalyzeProjState() MachineState {
	for i := 0; i < len(config.UnityProjectPath); i++ {
		if config.UnityProjectPath[i].State {
			m_State.State = "idle"
			return m_State
		}
	}
	m_State.State = "busy"
	return m_State
}

//获取空闲状态的解析工程进程
func GetIdleAnalyzer() int {
	num := 0
	for i := 0; i < len(config.UnityProjectPath); i++ {
		if config.UnityProjectPath[i].State {
			num++
		}
	}
	return num
}

//任务轮转通信 检测
func TaskTransfer() {
	m_State.Ip = config.ClientUrl.Ip
	m_State.State = "idle" //刚开始赋值为空闲状态
	for {
		if GetAnalyzeProjState().State == "busy" {
			//将任务进行轮转至其他解析器上
			idleTable := SendToGetAnalyzer()
			if len(idleTable) > 0 {
				for i := 0; i < len(idleTable); i++ {
					if idleTable[i].State == "idle" && idleTable[i].Num > 0 {
						for j := 0; j < idleTable[i].Num; j++ {
							taskPath := "./AnalyTask/ParseQue"
							data := RabbitMqServer.GetData(taskPath)
							SendAnalyzeToOther(data, idleTable[i].Ip, taskPath) //轮转任务
						}
					}
				}
			}
		}
		time.Sleep(10 * time.Minute) //每隔十分钟检测一次
	}
}
