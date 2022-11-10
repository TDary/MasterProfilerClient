package UnityServer

import (
	"MasterClient/Logs"
	"os"
	"os/exec"
)

//启动解析进程
func StartAnalyze(data AnalyzeData) {
	//判断unity版本然后进行选取
	if data.UnityVersion == "" {

	}
	Unity_Name := "F:/Unity2021.3.2f1/Unity.exe " //需要启动的程序名,Unity.exe的具体目录
	argu := "-quit -batchmode -nographics "
	argu = argu + "-projectPath " + config.UnityProjectPath + " "
	argu = argu + "-executeMethod Entrance.EntranceParseBegin "
	argu = argu + "-logFile E:/TestFiles/1668009600.log "
	argu = argu + "-rawPath E:/TestFiles/1668009600.raw "
	argu = argu + "-csvPath E:/TestFiles/1668009600.csv "
	argu = argu + "-funjsonPath E:/TestFiles/1668009600raw_funjson.json "
	argu = argu + "-funrowjsonPath E:/TestFiles/1668009600raw_funrow.json "
	argu = argu + "-funrenderrowjsonPath E:/TestFiles/1668009600raw_renderrow.json "
	argu = argu + "-funhashPath E:/TestFiles/1668009600raw_funhash.json "
	argu = argu + "-shieldSwitch false "
	argu = Unity_Name + argu
	cmd := exec.Command("cmd.exe", "/c", "start "+argu)
	er := cmd.Run()
	if er != nil { // 运行命令
		Logs.Loggers().Print(er)
	}
}

func Test2() {
	Unity_Name := "F:/Unity2021.3.2f1/Unity.exe " //需要启动的程序名,Unity.exe的具体目录
	argu := "-quit -batchmode -nographics "
	argu = argu + "-projectPath " + config.UnityProjectPath + " "
	argu = argu + "-executeMethod Entrance.EntranceParseBegin "
	argu = argu + "-logFile E:/TestFiles/1668009600.log "
	argu = argu + "-rawPath E:/TestFiles/1668009600.raw "
	argu = argu + "-csvPath E:/TestFiles/1668009600.csv "
	argu = argu + "-funjsonPath E:/TestFiles/1668009600raw_funjson.json "
	argu = argu + "-funrowjsonPath E:/TestFiles/1668009600raw_funrow.json "
	argu = argu + "-funrenderrowjsonPath E:/TestFiles/1668009600raw_renderrow.json "
	argu = argu + "-funhashPath E:/TestFiles/1668009600raw_funhash.json "
	argu = argu + "-shieldSwitch false "
	argu = Unity_Name + argu
	TestRun(argu)
}

func TestRun(argu string) {
	for {
		cmd := exec.Command("cmd.exe", "/c", "start "+argu)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		er := cmd.Start()
		if er != nil { // 运行命令
			Logs.Loggers().Println("start oasis erir [-] Error: ", os.Stderr)
			Logs.Loggers().Print(er)
		}
		Logs.Loggers().Println("启动成功，当前进程ID:", cmd.Process.Pid)
		cmd.Wait()
	}
}
