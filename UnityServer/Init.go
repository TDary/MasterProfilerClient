package UnityServer

import (
	"MasterClient/Logs"
	"MasterClient/Tools"
	"encoding/json"
	"io/ioutil"
)

func InitClient() {
	var data, _ = ioutil.ReadFile("./ClientConfig.json")
	var err = json.Unmarshal(data, &config)
	if err != nil {
		Logs.Loggers().Fatal(err)
	}
	Logs.Loggers().Print("初始化解析客户端配置成功----")
	//测试是否反序列化成功
	// fmt.Print(config)
	//为了避免死机重启后有任务还在运行卡流程，加入一个启动服务器检测的功能
	CheckCaseState()
}

func CheckCaseState() {
	var data, _ = ioutil.ReadFile("./HandingCase.json")
	var err = json.Unmarshal(data, &handCase)
	if err != nil {
		Logs.Loggers().Fatal(err)
	}
	Logs.Loggers().Print("获取文件数据成功,开始检测服务器状态----")
	// fmt.Print(handCase)   //测试是否能够获取
	//检查是否有还在进行解析的任务
	for _, val := range handCase.Case {
		if val.UUID != "" {
			resfilePath := config.FilePath + "/" + val.UUID + "/" + val.RawFile
			var data, _ = ioutil.ReadFile(resfilePath)
			if data != nil {
				Logs.Loggers().Print("UUID：" + val.UUID + "RawFile:" + val.RawFile + "已经解析完毕，正在准备返回解析工程")
				val.UUID = ""
				val.RawFile = ""
				OpenUnityProject(val.Numb)
				val.Numb = 0
				str, err := json.Marshal(handCase)
				if err != nil {
					Logs.Loggers().Print("转换json失败----", err)
				}
				Tools.WriteHandFile(string(str))
			} else { //解析失败，通过回传给中枢服务器请求重新解析
				//TODO:请求回传消息
			}
		} else {
			continue
		}
	}
}
