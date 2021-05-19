package mylogger

import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"strings"
)

/*
实现日志库的需求分析
1. 支持往不同的地方输出日志
2. 日志分为几种级别
	Debug
	Trace
	Info
	Warning
	Error
	Fatal
3. 日志要支持开关控制。比如说开发的时候什么级别都能输出，但是上线之后只有Info级别往下的才能输出
4. 完整的日志记录要包含时间、行号、文件名、日志级别、日志信息
5. 日志文件要切割
*/

type LogLevel uint8

const (
	UNKNOWN LogLevel = iota
	DEBUG
	TRACE
	INFO
	WARNING
	ERROR
	FATAL
)

func parseLogLevel(logLevel string) (LogLevel, error) {
	switch logLevel {
	case "debug":
		return DEBUG, nil
	case "trace":
		return TRACE, nil
	case "info":
		return INFO, nil
	case "warning":
		return WARNING, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil
	default:
		err := errors.New("illegal logger level")
		return UNKNOWN, err
	}
}

func getLogLevelString(level LogLevel) string {
	switch level {
	case DEBUG:
		return "DEBUG"
	case TRACE:
		return "TRACE"
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

func getInfo(skip int) (fileName, funcName string, lineNum int) {
	pc, file, lineNum, ok := runtime.Caller(skip)
	if !ok {
		fmt.Printf("runtime caller failed\n")
		return
	}
	fileName = path.Base(file)
	funcName = runtime.FuncForPC(pc).Name()
	funcName = strings.Split(funcName, ".")[1]
	return
}