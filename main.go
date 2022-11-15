//go:generate goversioninfo -icon=resource/icon.ico -manifest=resource/goversioninfo.exe.manifest
package main

import (
	"MasterClient/HttpServer"
	"MasterClient/Logs"
	"MasterClient/Minio"
	"MasterClient/UnityServer"
)

func main() {
	Logs.Loggers().Print("欢迎使用解析服务器客户端！！！")
	Minio.InitMinio()
	clientUrl := UnityServer.InitClient()
	HttpServer.ListenAndServer(clientUrl)
}
