package Tools

import (
	"MasterClient/Logs"
	"os"
)

func WriteHandFile(data string) {
	//写入到文件中
	fw, err := os.OpenFile("./HandingCase.json", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666) //os.O_TRUNC清空文件重新写入，否则原文件内容可能残留
	if err != nil {
		Logs.Loggers().Print("打开文件失败：", err)
	}
	fw.WriteString(data)
}
