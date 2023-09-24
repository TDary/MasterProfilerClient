//go:generate goversioninfo -icon=resource/icon.ico -manifest=resource/goversioninfo.exe.manifest
package main

import (
	"MasterClient/HttpServer"
	"MasterClient/Logs"
	"MasterClient/ParseServer"
	"MasterClient/UnityServer"
)

func main() {
	//日志初始化
	Logs.Init()
	Logs.Loggers().Print("欢迎使用解析服务器客户端！！！")
	clientUrl := UnityServer.InitClient()
	go ParseServer.AnalyzeRangeCheck() //检测解析任务进行启动解析
	HttpServer.ListenAndServer(clientUrl)
}
