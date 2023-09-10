package UnityServer

import (
	"MasterClient/Logs"
	"MasterClient/RabbitMqServer"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"
)

//发送解析失败的消息
func SendFailMessage(rawfile string, uuid string) {
	request_Url := "http://" + config.MasterServerUrl.Ip + ":" + config.MasterServerUrl.Port +
		"/FailledProfiler" + "?" + "uuid=" + uuid + "&rawfile=" + rawfile + "&ip=" + config.ClientUrl.Ip
	//超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(request_Url)
	if err != nil {
		Logs.Loggers().Print(err)
		return
	}
	defer resp.Body.Close()
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			Logs.Loggers().Print(err)
		}
	}
	if strings.Contains(result.String(), "ok") {
		Logs.Loggers().Print("中枢服务器接收到解析成功消息----")
	} else {
		Logs.Loggers().Print("中枢服务器未成功接收到消息----")
		return
	}
}

//发送成功解析的消息
func GetSucessData(rawfile string, uuid string) {
	request_Url := "http://" + config.MasterServerUrl.Ip + ":" + config.MasterServerUrl.Port +
		"/SuccessProfiler" + "?" + "uuid=" + uuid + "&rawfile=" + rawfile + "&ip=" + config.ClientUrl.Ip
	//超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(request_Url)
	if err != nil {
		Logs.Loggers().Print(err)
		return
	}
	defer resp.Body.Close()
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			Logs.Loggers().Print(err)
		}
	}
	if strings.Contains(result.String(), "ok") {
		Logs.Loggers().Print("中枢服务器接收到解析成功消息----")
	} else {
		Logs.Loggers().Print("中枢服务器未成功接收到消息----")
		return
	}
}

//发送重新解析消息
func SendReProfiler(rawfile string, uuid string) {
	request_Url := "http://" + config.MasterServerUrl.Ip + ":" + config.MasterServerUrl.Port +
		"/ReAnalyze" + "?" + "uuid=" + uuid + "&rawfile=" + rawfile
	//超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(request_Url)
	if err != nil {
		Logs.Loggers().Print(err)
		return
	}
	defer resp.Body.Close()
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			Logs.Loggers().Print(err)
		}
	}
	if strings.Contains(result.String(), "success") {
		Logs.Loggers().Print("中枢服务器接收到重新解析消息----")
	} else {
		Logs.Loggers().Print("中枢服务器未成功接收到消息----")
		return
	}
}

//发送启动解析器消息，通知master服务器我启动了
func SendStartMess() {
	request_Url := "http://" + config.MasterServerUrl.Ip + ":" + config.MasterServerUrl.Port +
		"/RquestClient" + "?" + "ip=" + config.ClientUrl.Ip
	//超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(request_Url)
	if err != nil {
		Logs.Loggers().Print("该地址不存在----", err)
		return
	}
	defer resp.Body.Close()
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			Logs.Loggers().Print(err)
		}
	}
	if strings.Contains(result.String(), "success") {
		Logs.Loggers().Print("中枢服务器接收到开启消息----", result.String())
	} else {
		Logs.Loggers().Print("中枢服务器未成功接收到消息----", result.String())
		return
	}
}

//发送请求 去获取空闲的解析器列表
func SendToGetAnalyzer() []MachineState {
	request_Url := "http://" + config.MasterServerUrl.Ip + ":" + config.MasterServerUrl.Port +
		"/requestidles" + "?" + "ip=" + config.ClientUrl.Ip
	//超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(request_Url)
	if err != nil {
		Logs.Loggers().Print("该地址不存在----", err)
		return nil
	}
	defer resp.Body.Close()
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			Logs.Loggers().Print(err)
		}
	}
	if result != nil {
		var receive []MachineState
		err = json.Unmarshal(result.Bytes(), &receive)
		if err != nil {
			Logs.Loggers().Print("json反序列化失败了", err)
		}
		return receive
	}
	return nil
}

//发送请求轮转给其他解析器进行解析
func SendAnalyzeToOther(data string, ip string, taskpath string) {
	request_Url := "http://" + ip + ":" + config.ClientUrl.Port +
		"/analyze" + "?" + data
	//超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(request_Url)
	if err != nil {
		Logs.Loggers().Print(err)
		return
	}
	defer resp.Body.Close()
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			Logs.Loggers().Print(err)
		}
	}
	if strings.Contains(result.String(), "ok") {
		Logs.Loggers().Print(ip + ",该解析器收到解析消息----")
	} else {
		Logs.Loggers().Print(ip + ",该解析器未成功接收到消息----")
		//重新回到队列中
		RabbitMqServer.PutData(taskpath, data)
		return
	}
}
