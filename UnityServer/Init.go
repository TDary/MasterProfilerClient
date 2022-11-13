package UnityServer

import (
	"MasterClient/Logs"
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
}
