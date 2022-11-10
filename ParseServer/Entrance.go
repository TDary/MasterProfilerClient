package ParseServer

import (
	"MasterClient/UnityServer"
	"strings"
)

func Analyze(data string) {
	var getdata UnityServer.AnalyzeData
	ParseData(data, getdata)
	UnityServer.StartAnalyze(getdata)
}

func ParseData(data string, gdata UnityServer.AnalyzeData) {
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
}
