package UnityServer

import (
	"MasterClient/Logs"
	"MasterClient/Minio"
	"io/ioutil"
)

func DownLoadFile(getdata AnalyzeData) bool {
	filePath := config.FilePath + "/" + getdata.UUID + "/" + getdata.RawFile
	var isExit, _ = ioutil.ReadFile(filePath)
	if isExit != nil {
		Logs.Loggers().Print("已存在源文件:" + getdata.RawFile)
		return true
	}
	fileName := getdata.RawFile
	bucket := getdata.AnalyzeBucket
	isSuccess := Minio.DownLoadFile(fileName, filePath, "", bucket)
	if isSuccess {
		return isSuccess
	}
	return false
}
