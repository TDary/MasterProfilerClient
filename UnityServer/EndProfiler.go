package UnityServer

import (
	"MasterClient/Logs"
	"bytes"
	"io"
	"net/http"
	"time"
)

func GetSucessData(rawfile string, uuid string, csvpath string) {
	request_Url := "http://" + config.MasterServerUrl + ":" +
		"/SuccessProfiler" + "?" + "uuid=" + uuid + "&rawfile=" + rawfile + "&csvpath=" + csvpath
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
	if result.String() == "ok" {
		Logs.Loggers().Print("中枢服务器接收到解析成功消息----")
	} else {
		Logs.Loggers().Print("中枢服务器未成功接收到消息----")
		return
	}
}

func SendReProfiler(rawfile string, uuid string) {
	request_Url := "http://" + config.MasterServerUrl + ":" +
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
	if result.String() == "ok" {
		Logs.Loggers().Print("中枢服务器接收到重新解析消息----")
	} else {
		Logs.Loggers().Print("中枢服务器未成功接收到消息----")
		return
	}
}
