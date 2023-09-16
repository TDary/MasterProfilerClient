package ParseServer

import (
	"MasterClient/Logs"
	"MasterClient/UnityServer"
	"net"
	"strings"
)

var conn net.Conn
var err error

//初始化连接对象
func InitSocketClient() {
	address := UnityServer.GetConfig().MasterServerUrl.Ip + ":" + UnityServer.GetConfig().MasterServerUrl.Port
	// 连接到服务器
	conn, err = net.Dial("tcp", address)
	if err != nil {
		Logs.Loggers().Fatalf("Failed to connect to server: %s", err.Error())
	}
	Logs.Loggers().Print("Connect Master successful!")
	// 发送消息到服务器
	message := "markeid?anaclient"
	_, err = conn.Write([]byte(message))
	if err != nil {
		Logs.Loggers().Printf("Error sending message to server: %s", err.Error())
		return
	}
	for {
		// 接收从服务器返回的消息
		buffer := make([]byte, 2048)
		n, err := conn.Read(buffer)
		if err != nil {
			Logs.Loggers().Printf("Error receiving message from server: %s", err.Error())
			return
		}
		if len(buffer) != 0 {
			res := string(buffer[:n])
			if strings.Contains(res, "analyze") {
				Logs.Loggers().Print("接收到解析任务的消息----", res)
				BeginMsg := strings.Split(res, "?")[1]
				go GetAnalyzeMes(BeginMsg)
				message := "Hello, server!"
				conn.Write([]byte(message))
			} else {
				Logs.Loggers().Printf("Receive Data From Server: %s", res)
			}
		}
	}
}

//关闭socket连接
func CloseConnection() {
	conn.Close()
}

//获取连接对象
func GetConn() net.Conn {
	return conn
}
