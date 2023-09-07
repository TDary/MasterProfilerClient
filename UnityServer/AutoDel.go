package UnityServer

import (
	"MasterClient/Logs"
	"os"
)

//todo:用于定时删除解析后的文件夹
func DeleteFiles() {

	filePath := config.FilePath
	// syscall.Statfs(filePath, stat)
	//删除案例文件夹
	err := os.RemoveAll(filePath)
	if err != nil {
		Logs.Loggers().Print("删除案例文件夹失败----", err)
		return
	}
	Logs.Loggers().Print("清除案例文件夹成功")
}
