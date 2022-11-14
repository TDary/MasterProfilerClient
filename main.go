//go:generate goversioninfo -icon=resource/icon.ico -manifest=resource/goversioninfo.exe.manifest
package main

import (
	"MasterClient/Logs"
	"MasterClient/UnityServer"
)

func main() {
	Logs.Loggers().Print("欢迎使用解析服务器客户端！！！")
	UnityServer.InitClient()
	// Minio.InitMinio()
	// HttpServer.ListenAndServer("")
}
