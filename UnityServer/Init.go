package UnityServer

import (
	"MasterClient/Logs"
	"bufio"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
)

func InitClient() string {
	var data, _ = ioutil.ReadFile("./ClientConfig.json")
	var err = json.Unmarshal(data, &config)
	if err != nil {
		Logs.Loggers().Fatal(err)
	}
	Logs.Loggers().Print("初始化解析客户端配置成功----")
	//为了避免死机重启后有任务还在运行卡流程，加入一个启动服务器检测的功能
	CheckCaseState()
	//启动客户端解析需要请求一次master服务器
	SendStartMess()
	//检查Unity工程是否存在等
	CheckUnityProject()
	address := config.ClientUrl.Ip + ":" + config.ClientUrl.Port
	return address
}

// 检查有没有宕机重启的情况
func CheckCaseState() {
	//检查有没有解析的队列
	//打开文件
	Logs.Loggers().Print("正在检查是否有未完成解析的任务----")
	var newQue []string
	file, err := os.Open("./Analyzing.txt")
	if err != nil {
		Logs.Loggers().Println("文件打开失败 = ", err)
	}
	//及时关闭 file 句柄，否则会有内存泄漏
	defer file.Close()
	//创建一个 *Reader ， 是带缓冲的
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n') //读到一个换行就结束
		if err == io.EOF {                  //io.EOF 表示文件的末尾
			break
		} else if str != "" {
			//还有解析中的任务,直接发送回传消息重新解析\
			newQue = append(newQue, str)
		}
	}
	if len(newQue) != 0 {
		for i := 0; i < len(newQue); i++ {
			var getdata AnalyzeData
			err = json.Unmarshal([]byte(newQue[i]), &getdata)
			if err != nil {
				Logs.Loggers().Print("json转换失败----")
			}
			SendReProfiler(getdata.RawFile, getdata.UUID) //发送失败解析的任务
		}
	}
	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)
	//重新写入空白文件
	write.WriteString("")
	//刷新一下啊
	write.Flush()
}

func CheckUnityProject() {
	Logs.Loggers().Print("正在检查Unity解析模板以及Unity程序是否存在----")
	for i := 0; i < len(config.UnityProjectPath); i++ {
		_, err := os.Stat(config.UnityProjectPath[i].Path)
		if err != nil {
			Logs.Loggers().Fatal("当前解析模板不存在：", config.UnityProjectPath[i].Path)
		}
	}
	for i := 0; i < len(config.UnityPath); i++ {
		_, err := os.Stat(config.UnityPath[i].Path)
		if err != nil {
			Logs.Loggers().Fatal("当前解析程序不存在：", config.UnityProjectPath[i].Path)
		}
	}
	Logs.Loggers().Print("检查完毕，状态完好!!!")
}
