package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/evanxzj/data-analyze/utils"
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
	fmt.Println("Server starting...")

	// get params from commman line
	params := utils.CmdParamsParse()

	// 设置runtime log输出文件位置
	logFd, err := os.OpenFile(params.RunTimeLogPath, os.O_CREATE|os.O_WRONLY, 0644)
	if err == nil {
		log.Out = logFd

		defer logFd.Close()
	}
	log.Infof("Params: LogFilePath=%s RoutineNum=%d RunTimeLogPath=%s", params.LogFilePath, params.RoutineNum, params.RunTimeLogPath)

	// 初始化一些channel, 用于数据传递
	var logChannel = make(chan string, 3*params.RoutineNum)

	// 读取日志
	go readFileLineByLine(params, logChannel)

	// // 创建一组日志消费goroutine
	// for i := 0; i < params.RoutineNum; i++ {
	// 	// 消费日志
	// }

	// 防止程序提前退出
	time.Sleep(100 * time.Second)
}

// 读取日志
func readFileLineByLine(parms utils.CmdParams, logChannel chan string) error {
	fd, err := os.Open(parms.LogFilePath)
	if err != nil {
		log.Warningf("readFileLineByLine can't open file:%s", parms.LogFilePath)

		return err
	}
	defer fd.Close()

	count := 0
	bufferRead := bufio.NewReader(fd)

	for {
		line, err := bufferRead.ReadString('\n')

		// 往channel里面写数据
		logChannel <- line

		count++
		if count%(100*parms.RoutineNum) == 0 {
			log.Infof("readFileLineByLine read line counts: %d", count)
		}

		if err != nil {
			if err == io.EOF {
				time.Sleep(3 * time.Second)
				log.Infof("readFileLineByLine wait to read")
			} else {
				log.Warningf("readFileLineByLine read log error")
			}
		}
	}
}
