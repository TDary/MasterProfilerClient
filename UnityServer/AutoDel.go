package UnityServer

import (
	"MasterClient/Logs"
	"os"
	"path/filepath"
	"time"

	"github.com/shirou/gopsutil/v3/disk"
)

func CheckDiskToFree() {
	for {
		filePath := config.FilePath
		usage, err := disk.Usage(filePath)
		if err != nil {
			Logs.Loggers().Print("无法获取磁盘信息：", err)
			return
		}
		if usage.Free < 52428800 { //小于50MB
			//触发删除源文件功能
			filepath.Walk(filePath, func(path string, info os.FileInfo, err error) error {
				if err != nil {
					Logs.Loggers().Print("当前目录获取失败", err)
					return err
				}
				FileInfo, err := os.Stat(path)
				if err != nil {
					Logs.Loggers().Print("无法获取当前文件状态----", err)
					return err
				}
				modtiTime := FileInfo.ModTime()        //获取当前文件或目录的最后修改时间
				currentTime := time.Now()              //获取当前时间
				duration := currentTime.Sub(modtiTime) //计算时间差
				//判断是否在7天前
				if duration >= 7*24*time.Hour {
					//7天前的 则进行删除
					os.Remove(path)
				}
				return nil
			})
		}
		time.Sleep(1 * time.Hour) //每隔1小时检测一次
	}
}
