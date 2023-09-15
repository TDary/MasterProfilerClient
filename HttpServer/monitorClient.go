package HttpServer

import (
	"MasterClient/Logs"
	"MasterClient/ParseServer"
	"MasterClient/UnityServer"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func ListenAndServer(address string) {
	http.HandleFunc("/stop", SuccessAnalyze)
	http.HandleFunc("/manastate", RequestAnalyzeState)
	//Http监听函数
	http.ListenAndServe(address, nil)
}

//Http请求处理模块
func DealReceivedMessage(msg string) int {
	if strings.Contains(msg, "stop") {
		SuccessMsg := strings.Split(msg, "?")[1]
		go ParseServer.AnalyzeSuccess(SuccessMsg)
		Logs.Loggers().Print("接收到解析完毕的消息----", msg)
		return 200
	} else if strings.Contains(msg, "manastate") {
		StMsg := strings.Split(msg, "?")[1]
		go ParseServer.RequestMachineState(StMsg)
		Logs.Loggers().Print("接收到请求获取解析器状态的消息----")
		return 200
	} else {
		return 400
		//TODO:扩展处理模块
	}
}

//请求解析响应模块
func RequestProfiler(w http.ResponseWriter, r *http.Request) {
	var resData string
	RequestUrlData := r.URL.String()
	resMes := DealReceivedMessage(RequestUrlData)
	if resMes == 200 {
		resData = `{"code":200,"msg":"ok"}`
	} else {
		resData = "Request Fail"
	}
	w.Header().Set("Content-Type", "application/json") //设置响应内容
	jsonByte, _ := json.Marshal(resData)               //转string
	w.Write(jsonByte)
}

//接受解析完成消息
func SuccessAnalyze(w http.ResponseWriter, r *http.Request) {
	var resData string
	RequestUrlData := r.URL.String()
	resMes := DealReceivedMessage(RequestUrlData)
	if resMes == 200 {
		resData = `{"code":200,"msg":"ok receive"}`
	} else {
		resData = "Request Fail"
	}
	w.Header().Set("Content-Type", "application/json") //设置响应内容
	jsonByte, _ := json.Marshal(resData)               //转string
	w.Write(jsonByte)
}

//获取解析器的当前状态
func RequestAnalyzeState(w http.ResponseWriter, r *http.Request) {
	var resData string
	RequestUrlData := r.URL.String()
	resMes := DealReceivedMessage(RequestUrlData)
	if resMes == 200 {
		var res strings.Builder
		res.WriteString(`{"code":200,"state":"`)
		res.WriteString(UnityServer.GetAnalyzeProjState().State)
		res.WriteString(`","num":`)
		res.WriteString(strconv.Itoa(UnityServer.GetIdleAnalyzer()))
		res.WriteString("}")
		resData = res.String()
	} else {
		resData = "Request Fail"
	}
	w.Header().Set("Content-Type", "application/json") //设置响应内容
	jsonByte, _ := json.Marshal(resData)               //转string
	w.Write(jsonByte)
}
