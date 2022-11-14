package UnityServer

import (
	"MasterClient/Logs"
	"os"
)

func DeleteFiles(uuid string) {
	filePath := config.FilePath + "/" + uuid
	//删除案例文件夹
	err := os.RemoveAll(filePath)
	if err != nil {
		Logs.Loggers().Print("删除案例文件夹失败----", err)
		return
	}
	Logs.Loggers().Print("清除案例文件夹成功UUID：", uuid)
}
