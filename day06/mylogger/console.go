package mylogger

import (
	"errors"
	"fmt"
	"time"
)

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

type MyLogger struct {
	level LogLevel
}

func NewLogger(logLevel string) *MyLogger {
	level, err := parseLogLevel(logLevel)
	if err != nil {
		fmt.Printf("parse logger level \"%s\" failed, err: %v\n", logLevel, err)
		panic("exit...")
	}
	return &MyLogger{
		level: level,
	}
}

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

func (l MyLogger) Debug(msg string)  {
	if l.level <= DEBUG {
		curTime := time.Now().Format("2006-01-02 15:04:05")
		fmt.Printf("[%s] [Debug] %s\n", curTime, msg)
	}
}

func (l MyLogger) Trace(msg string)  {
	if l.level <= TRACE {
		curTime := time.Now().Format("2006-01-02 15:04:05")
		fmt.Printf("[%s] [Trace] %s\n", curTime, msg)
	}
}

func (l MyLogger) Info(msg string)  {
	if l.level <= INFO {
		curTime := time.Now().Format("2006-01-02 15:04:05")
		fmt.Printf("[%s] [Info] %s\n", curTime, msg)
	}
}

func (l MyLogger) Warning(msg string)  {
	if l.level <= WARNING {
		curTime := time.Now().Format("2006-01-02 15:04:05")
		fmt.Printf("[%s] [Warning] %s\n", curTime, msg)
	}
}

func (l MyLogger) Error(msg string)  {
	if l.level <= ERROR {
		curTime := time.Now().Format("2006-01-02 15:04:05")
		fmt.Printf("[%s] [Error] %s\n", curTime, msg)
	}
}

func (l MyLogger) Fatal(msg string)  {
	if l.level <= FATAL {
		curTime := time.Now().Format("2006-01-02 15:04:05")
		fmt.Printf("[%s] [Fatal] %s\n", curTime, msg)
	}
}
