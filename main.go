package main

import (
	"os"

	"github.com/data-analyze/utils"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

// 设置log的输出位置和等级
// 初始化函数
func init() {
	log.Out = os.Stdout
	log.SetLevel(logrus.DebugLevel)
}

func main() {
	log.Println("Server starting...")

	// get params from commman line
	params := utils.CmdParamsParse()
	log.Debugf("Params: LogFilePath=%s RoutineNum=%d RunTimeLogPath=%s", params.LogFilePath, params.RoutineNum, params.RunTimeLogPath)
}
