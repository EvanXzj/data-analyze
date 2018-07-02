package utils

import "flag"

// CmdParams struct
type CmdParams struct {
	LogFilePath,
	RunTimeLogPath string
	RoutineNum int
}

// CmdParamsParse 获取命令行参数
// 调用方法： go run main.go -routineNum 10...
func CmdParamsParse() CmdParams {
	logFilePath := flag.String("logFilePath", "/tmp/logs/dig.log", "file path where to read log data to parse") // this is an address
	routineNum := flag.Int("routineNum", 5, "consumer numble by goroutine")
	runTimeLogPath := flag.String("runTimeLogPath", "/tmp//logs/runtime.log", "this programe runtime log target file path")
	flag.Parse()

	// fmt.Println(flag.Parsed())
	return CmdParams{*logFilePath, *runTimeLogPath, *routineNum}
}
