package UnityServer

import (
	"MasterClient/Logs"
	"MasterClient/Tools"
	"bufio"
	"io"
	"os"
)

func CheckUnityVersion(data AnalyzeData) string {
	for i := 0; i < len(config.UnityPath); i++ {
		if data.UnityVersion == config.UnityPath[i].Version {
			return config.UnityPath[i].Path
		}
	}
	return ""
}

// 成功启动的话将进行写入进行中队列，独特的队列
func SuccessBegin(data AnalyzeData, num int) {
	//todo:将其保存为二进制序列化文件
	data.AnalyzeNum = num
	transinfo := Tools.StructToMap(data)
	infodata := Tools.MapToString(transinfo)
	lck.Lock() //加互斥锁
	defer lck.Unlock()
	//文件为空
	file, err := os.OpenFile("./Analyzing.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		Logs.Loggers().Println("文件打开失败", err)
	}
	//及时关闭file句柄
	defer file.Close()
	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)
	write.WriteString(infodata)
	//Flush将缓存的文件真正写入到文件中
	write.Flush()
}

// 解析完毕的话去除文件中进行的
func SuccessAnalyze(data AnalyzeData) {
	transinfo := Tools.StructToMap(data)
	infodata := Tools.MapToString(transinfo)
	var newQue []string
	lck.Lock() //加互斥锁
	defer lck.Unlock()
	//打开文件
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
		} else if str == infodata {
			Logs.Loggers().Print("正在删除该解析记录(已完成该任务解析)----")
			continue
		}
		newQue = append(newQue, str)
	}
	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)
	for i := 0; i < len(newQue); i++ {
		write.WriteString(newQue[i])
	}
	//Flush将缓存的文件真正写入到文件中
	write.Flush()
	RleaseUnityProject(data.AnalyzeNum) //解析完毕将工程释放出来
}
