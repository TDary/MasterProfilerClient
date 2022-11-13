//go:generate goversioninfo -icon=resource/icon.ico -manifest=resource/goversioninfo.exe.manifest
package main

import (
	"MasterClient/HttpServer"
	"MasterClient/Minio"
	"MasterClient/UnityServer"
	"fmt"
)

func main() {
	fmt.Print("欢迎使用解析服务器客户端！！！")
	UnityServer.InitClient()
	Minio.InitMinio()
	HttpServer.ListenAndServer("")
}
