package main

import (
	"flag"
	"os"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

// 设置log的输出位置和等级
// 初始化函数
func init() {
	log.Out = os.Stdout
	log.SetLevel(logrus.DebugLevel)
}

type cmdParams struct {
	logFilePath,
	runTimeLogPath string
	routineNum int
}

// 获取命令行参数
// 调用方法： go run main.go -routineNum 10...
func cmdParamsParse() cmdParams {
	logFilePath := flag.String("logFilePath", "/Users/zhijian/Public/nginx/logs/dig.log", "file path where to read log data to parse") // this is an address
	routineNum := flag.Int("routineNum", 5, "consumer numble by goroutine")
	runTimeLogPath := flag.String("runTimeLogPath", "/tmp/log", "this programe runtime log target file path")
	flag.Parse()

	return cmdParams{*logFilePath, *runTimeLogPath, *routineNum}
}

func main() {
	log.Println("Server starting...")

	params := cmdParamsParse()

	log.Debugf("Params: logFilePath=%s routineNum=%d runTimeLogPath=%s", params.logFilePath, params.routineNum, params.runTimeLogPath)
}
