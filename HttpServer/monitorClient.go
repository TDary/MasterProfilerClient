package HttpServer

import (
	"MasterClient/Logs"
	"MasterClient/ParseServer"
	"encoding/json"
	"net/http"
	"strings"
)

func ListenAndServer(address string) {
	http.HandleFunc("/analyze", RequestProfiler)
	http.HandleFunc("/stop", SuccessAnalyze)
	//Http监听函数
	http.ListenAndServe(address, nil)
}

//Http请求处理模块
func DealReceivedMessage(msg string) int {
	if strings.Contains(msg, "analyze") {
		BeginMsg := strings.Split(msg, "?")[1]
		go ParseServer.GetAnalyzeMes(BeginMsg)
		Logs.Loggers().Print("接收到解析任务的消息----")
		return 200
	} else if strings.Contains(msg, "stop") {
		SuccessMsg := strings.Split(msg, "?")[1]
		go ParseServer.AnalyzeSuccess(SuccessMsg)
		Logs.Loggers().Print("接收到解析完毕的消息----")
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
		resData = "{'code':200,'msg':'ok'}"
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
		resData = "{'code':200,'msg':'ok receive'}"
	} else {
		resData = "Request Fail"
	}
	w.Header().Set("Content-Type", "application/json") //设置响应内容
	jsonByte, _ := json.Marshal(resData)               //转string
	w.Write(jsonByte)
}
