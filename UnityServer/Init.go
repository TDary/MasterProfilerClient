package UnityServer

import (
	"MasterClient/Logs"
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func InitClient() string {
	var data, _ = ioutil.ReadFile("./ClientConfig.json")
	var err = json.Unmarshal(data, &config)
	if err != nil {
		Logs.Loggers().Fatal(err)
	}
	Logs.Loggers().Print("初始化解析客户端配置成功----")
	//为了避免死机重启后有任务还在运行卡流程，加入一个启动服务器检测的功能
	CheckCaseState()
	//启动客户端解析需要请求一次master服务器
	SendStartMess()
	//检查Unity工程是否存在等
	CheckUnityProject()
	address := config.ClientUrl.Ip + ":" + config.ClientUrl.Port
	return address
}

// 检查有没有宕机重启的情况
func CheckCaseState() {
	//检查有没有解析完成的队列
	Logs.Loggers().Print("正在检查是否有未完成解析的任务----")
	allfilePath := "./analyzing"
	filepath.Walk(allfilePath, func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, ".gob") {
			var getdata AnalyzeData
			err := readGob(path, &getdata)
			if err != nil {
				Logs.Loggers().Fatal(err.Error())
			}
			SendReProfiler(getdata.RawFile, getdata.UUID) //发送失败解析的任务
		}
		return nil
	})
}

func CheckUnityProject() {
	Logs.Loggers().Print("正在检查Unity解析模板以及Unity程序是否存在----")
	for i := 0; i < len(config.UnityProjectPath); i++ {
		_, err := os.Stat(config.UnityProjectPath[i].Path)
		if err != nil {
			Logs.Loggers().Fatal("当前解析模板不存在：", config.UnityProjectPath[i].Path)
		}
	}
	for i := 0; i < len(config.UnityPath); i++ {
		_, err := os.Stat(config.UnityPath[i].Path)
		if err != nil {
			Logs.Loggers().Fatal("当前解析程序不存在：", config.UnityProjectPath[i].Path)
		}
	}
	Logs.Loggers().Print("检查完毕，状态完好!!!")
}
