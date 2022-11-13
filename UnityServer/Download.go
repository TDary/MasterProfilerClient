package UnityServer

import "MasterClient/Minio"

func DownLoadFile(getdata AnalyzeData) bool {
	filePath := config.FilePath + "/" + getdata.UUID + "/" + getdata.RawFile
	fileName := getdata.RawFile
	bucket := getdata.AnalyzeBucket
	isSuccess := Minio.DownLoadFile(fileName, filePath, "", bucket)
	if isSuccess {
		return true
	}
	return false
}
