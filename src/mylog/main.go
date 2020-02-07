package mylog

import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"strings"
)

//支持不同地方的输出
//日志级别1.Debug 2.Info 3.Warning 4.Error 5.Fatal
// 开关控制
type LogLevel uint8

const (
	//定义日志级别
	UNKNOW LogLevel = iota
	DEBUG
	INFO
	WARNING
	ERROR
	FATAL
)

func parseLevel(levelStr string) (LogLevel, error) {
	switch strings.ToLower(levelStr) {
	case "debug":
		return DEBUG, nil
	case "info":
		return INFO, nil
	case "warning":
		return WARNING, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil
	default:
		err := errors.New("无效的日志级别")
		return UNKNOW, err
	}
}
func getLevelString(lv LogLevel) string {
	switch lv {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	default:
		return "DEBUG"

	}
}
func getInfo(skip int) (file, funcname string, line int) {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		fmt.Println("runtime caller failed")
		return
	}
	funcname = runtime.FuncForPC(pc).Name()
	funcname = strings.Split(funcname, ".")[1]
	file = path.Base(file)
	return
}

func init() {

}
