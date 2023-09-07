package UnityServer

import (
	"MasterClient/Logs"
	"encoding/gob"
	"os"
	"strconv"
)

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
	var getdata AnalyzeData
	currentNum := strconv.Itoa(data.AnalyzeNum)
	getFile := "./analyzing" + "/" + currentNum + ".gob"
	err := readGob(getFile, getdata)
	if err != nil {
		Logs.Loggers().Print(err)
	}
	if data.UUID == getdata.UUID && data.RawFile == getdata.RawFile && data.AnalyzeType == getdata.AnalyzeType {
		os.Remove(getFile)
	}
}

//序列化写入文件
func writeGob(filePath string, object interface{}) error {
	file, err := os.Create(filePath)
	if err == nil {
		encoder := gob.NewEncoder(file)
		encoder.Encode(object)
	}
	defer file.Close()
	return err
}

//反序列化
func readGob(filePath string, object interface{}) error {
	file, err := os.Open(filePath)
	if err == nil {
		decoder := gob.NewDecoder(file)
		err = decoder.Decode(object)
	}
	defer file.Close()
	return err
}
