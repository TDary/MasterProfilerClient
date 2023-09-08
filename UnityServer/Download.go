package UnityServer

import (
	"MasterClient/Logs"
	"MasterClient/Minio"
	"io/ioutil"
)

func DownLoadRawFile(getdata AnalyzeData) bool {
	filePath := config.FilePath + "/" + getdata.UUID + "/" + getdata.RawFile
	var isExit, _ = ioutil.ReadFile(filePath)
	if isExit != nil {
		Logs.Loggers().Print("已存在源文件:" + getdata.RawFile)
		return true
	}
	fileName := getdata.RawFile
	isSuccess := Minio.DownLoadFile(fileName, filePath, "application/raw")
	if isSuccess {
		return isSuccess
	}
	return false
}

func DownLoadCsvFile(getdata AnalyzeData) bool {
	filePath := config.FilePath + "/" + getdata.UUID + "/" + getdata.RawFile
	var isExit, _ = ioutil.ReadFile(filePath)
	if isExit != nil {
		Logs.Loggers().Print("已存在源文件:" + getdata.RawFile)
		return true
	}
	fileName := getdata.RawFile
	isSuccess := Minio.DownLoadFile(fileName, filePath, "application/csv")
	if isSuccess {
		return isSuccess
	}
	return false
}

func DownLoadBinFile(getdata AnalyzeData) bool {
	filePath := config.FilePath + "/" + getdata.UUID + "/" + getdata.RawFile
	var isExit, _ = ioutil.ReadFile(filePath)
	if isExit != nil {
		Logs.Loggers().Print("已存在源文件:" + getdata.RawFile)
		return true
	}
	fileName := getdata.RawFile
	isSuccess := Minio.DownLoadFile(fileName, filePath, "application/Bin")
	if isSuccess {
		return isSuccess
	}
	return false
}
