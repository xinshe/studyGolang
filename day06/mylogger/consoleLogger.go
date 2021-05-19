package mylogger

import (
	"fmt"
	"time"
)

type ConsoleLogger struct {
	level LogLevel
}

func NewConsoleLogger(logLevel string) *ConsoleLogger {
	level, err := parseLogLevel(logLevel)
	if err != nil {
		fmt.Printf("parse logger level \"%s\" failed, err: %v\n", logLevel, err)
		panic(err)
	}
	return &ConsoleLogger{
		level: level,
	}
}

func (c ConsoleLogger) log(level LogLevel, format string, a ...interface{})  {
	if c.level <= level {
		curTime := time.Now().Format("2006-01-02 15:04:05")
		fileName, funcName, lineNum := getInfo(3)
		msg := fmt.Sprintf(format, a...)
		fmt.Printf("[%s] [%s] [%s %s:%d] %s\n", curTime, getLogLevelString(level),
			fileName, funcName, lineNum, msg)
	}
}

func (c ConsoleLogger) Debug(format string, a ...interface{})  {
	c.log(DEBUG, format, a...)
}

func (c ConsoleLogger) Trace(format string, a ...interface{})  {
	c.log(TRACE, format, a...)
}

func (c ConsoleLogger) Info(format string, a ...interface{})  {
	c.log(INFO, format, a...)
}

func (c ConsoleLogger) Warning(format string, a ...interface{})  {
	c.log(WARNING, format, a...)
}

func (c ConsoleLogger) Error(format string, a ...interface{})  {
	c.log(ERROR, format, a...)
}

func (c ConsoleLogger) Fatal(format string, a ...interface{})  {
	c.log(FATAL, format, a...)
}
