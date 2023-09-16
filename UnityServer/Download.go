package UnityServer

import (
	"MasterClient/Logs"
	"MasterClient/Minio"
	"io/ioutil"
	"os"
	"strings"
)

func DownLoadRawFile(getdata AnalyzeData) bool { //todo:字符串拼接优化
	var filePath strings.Builder
	//——————————————————————————————————filePath
	filePath.WriteString(config.FilePath)
	filePath.WriteString("/")
	filePath.WriteString(getdata.UUID)
	filePath.WriteString("/")
	filePath.WriteString(getdata.RawFile)
	//——————————————————————————————————srcPath
	createPath := config.FilePath + "/" + getdata.UUID
	var isExit, _ = ioutil.ReadFile(filePath.String())
	if isExit != nil {
		Logs.Loggers().Print("已存在源文件:" + getdata.RawFile)
		err := ExtractZip(filePath.String(), createPath)
		if err != nil {
			Logs.Loggers().Print("解压源文件夹失败----")
			return false
		}
		return true
	} else {
		_, err := os.Stat(createPath)
		if err != nil {
			Logs.Loggers().Print("当前文件夹不存在：" + createPath)
			Logs.Loggers().Printf("重新创建文件夹%s----\n", createPath)
			os.Mkdir(createPath, 0755) //创建文件夹
		}
		isSuccess := Minio.DownLoadFile(getdata.RawFileName, filePath.String(), "application/zip")
		if isSuccess {
			err := ExtractZip(filePath.String(), createPath)
			if err != nil {
				Logs.Loggers().Print("解压源文件夹失败----")
				return false
			}
			return true
		}
	}
	return false
}

// func DownLoadCsvFile(getdata AnalyzeData) bool {
// 	filePath := config.FilePath + "/" + getdata.UUID + "/" + getdata.RawFile
// 	var isExit, _ = ioutil.ReadFile(filePath)
// 	if isExit != nil {
// 		Logs.Loggers().Print("已存在源文件:" + getdata.RawFile)
// 		return true
// 	}
// 	fileName := getdata.RawFile
// 	isSuccess := Minio.DownLoadFile(fileName, filePath, "application/csv")
// 	if isSuccess {
// 		return isSuccess
// 	}
// 	return false
// }

// func DownLoadBinFile(getdata AnalyzeData) bool {
// 	filePath := config.FilePath + "/" + getdata.UUID + "/" + getdata.RawFile
// 	var isExit, _ = ioutil.ReadFile(filePath)
// 	if isExit != nil {
// 		Logs.Loggers().Print("已存在源文件:" + getdata.RawFile)
// 		return true
// 	}
// 	fileName := getdata.RawFile
// 	isSuccess := Minio.DownLoadFile(fileName, filePath, "application/Bin")
// 	if isSuccess {
// 		return isSuccess
// 	}
// 	return false
// }
