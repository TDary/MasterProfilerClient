package UnityServer

import (
	"MasterClient/RabbitMqServer"
	"archive/zip"
	"encoding/gob"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

//检查Unity版本
func CheckUnityVersion(data AnalyzeData) string {
	for i := 0; i < len(config.UnityPath); i++ {
		if data.UnityVersion == config.UnityPath[i].Version {
			return config.UnityPath[i].Path
		}
	}
	return ""
}

//序列化写入文件
func writeGob(filePath string, object interface{}) error {
	file, err := os.Create(filePath)
	if err == nil {
		encoder := gob.NewEncoder(file)
		encoder.Encode(object)
	}
	file.Close()
	return err
}

//反序列化
func readGob(filePath string, object interface{}) error {
	file, err := os.Open(filePath)
	if err == nil {
		decoder := gob.NewDecoder(file)
		err = decoder.Decode(object)
	}
	file.Close()
	return err
}

//获取当前解析状态
func GetAnalyzeProjState() MachineState {
	for i := 0; i < len(config.UnityProjectPath); i++ {
		if config.UnityProjectPath[i].State {
			m_State.State = "idle"
			return m_State
		}
	}
	m_State.State = "busy"
	return m_State
}

//获取空闲状态的解析工程进程
func GetIdleAnalyzer() int {
	num := 0
	for i := 0; i < len(config.UnityProjectPath); i++ {
		if config.UnityProjectPath[i].State {
			num++
		}
	}
	return num
}

//任务轮转通信 检测
func TaskTransfer() {
	m_State.Ip = config.ClientUrl.Ip
	m_State.State = "idle" //刚开始赋值为空闲状态
	for {
		if GetAnalyzeProjState().State == "busy" {
			//将任务进行轮转至其他解析器上
			idleTable := SendToGetAnalyzer()
			if len(idleTable) > 0 {
				isNoque := false
				for i := 0; i < len(idleTable); i++ {
					if idleTable[i].State == "idle" && idleTable[i].Num > 0 {
						for j := 0; j < idleTable[i].Num; j++ {
							taskPath := "./AnalyTask/ParseQue"
							data := RabbitMqServer.GetData(taskPath)
							if data == "" {
								//队列已经空了，没有任务了。现在要进行打断退出
								isNoque = true
								break
							}
							SendAnalyzeToOther(data, idleTable[i].Ip, taskPath) //轮转任务
						}
						if isNoque {
							//队列已空，退出轮转状态
							break
						}
					}
				}
			}
		}
		time.Sleep(10 * time.Minute) //每隔十分钟检测一次
	}
}

//解压zip文件
func ExtractZip(zipFile string, targetFolder string) error {
	reader, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer reader.Close()
	for _, file := range reader.File {
		// 获取相对路径
		relPath := strings.TrimPrefix(file.Name, filepath.Dir(file.Name))

		// 拼接目标文件路径
		targetPath := filepath.Join(targetFolder, relPath)

		if file.FileInfo().IsDir() {
			err := os.MkdirAll(targetPath, os.ModePerm)
			if err != nil {
				return err
			}
			continue
		}

		srcFile, err := file.Open()
		if err != nil {
			return err
		}
		defer srcFile.Close()

		destFile, err := os.OpenFile(targetPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}
		defer destFile.Close()

		_, err = io.Copy(destFile, srcFile)
		if err != nil {
			return err
		}
	}

	return nil
}
