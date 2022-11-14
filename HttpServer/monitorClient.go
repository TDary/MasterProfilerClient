package HttpServer

import (
	"MasterClient/ParseServer"
	"encoding/json"
	"net/http"
	"strings"
)

func ListenAndServer(address string) {
	http.HandleFunc("/analyze", RequestProfiler)
	http.HandleFunc("/mergesuccess", SuccessMerge)
	//Http监听函数
	http.ListenAndServe(address, nil)
}

//Http请求处理模块
func DealReceivedMessage(msg string) int {
	if strings.Contains(msg, "analyze") {
		beginMsg := strings.Split(msg, "?")[1]
		go ParseServer.Analyze(beginMsg)
		return 200
	} else if strings.Contains(msg, "mergesuccess") {
		sucMsg := strings.Split(msg, "?")[1]
		go ParseServer.AcceptData(sucMsg)
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
		resData = "ok"
	} else {
		resData = "Request Fail"
	}
	w.Header().Set("Content-Type", "application/json") //设置响应内容
	jsonByte, _ := json.Marshal(resData)               //转string
	w.Write(jsonByte)
}

//接受合并完成消息，删除对应的文件
func SuccessMerge(w http.ResponseWriter, r *http.Request) {
	var resData string
	RequestUrlData := r.URL.String()
	resMes := DealReceivedMessage(RequestUrlData)
	if resMes == 200 {
		resData = "ok"
	} else {
		resData = "Request Fail"
	}
	w.Header().Set("Content-Type", "application/json") //设置响应内容
	jsonByte, _ := json.Marshal(resData)               //转string
	w.Write(jsonByte)
}
