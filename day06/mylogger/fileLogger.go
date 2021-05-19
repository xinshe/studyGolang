package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

type FileLogger struct {
	level LogLevel
	basePath string
	filePath string
	fileName string
	errFilePath string
	errFileName string
	maxSize int64

	fileObj *os.File
	errFileObj *os.File
}

func NewFileLogger(logLevel, basePath, fileName string, maxSize int64) *FileLogger {
	level, err := parseLogLevel(logLevel)
	if err != nil {
		fmt.Printf("parse logger level \"%s\" failed, err: %v\n", logLevel, err)
		panic(err)
	}
	fileLogger := &FileLogger{
		level: level,
		basePath: basePath,
		fileName: fileName,
		errFileName: fileName + ".err",
		maxSize: maxSize,
	}
	fileLogger.filePath = path.Join(basePath, fileName)
	fileLogger.errFilePath = path.Join(basePath, fileName + ".err")

	fileObj, err := os.OpenFile(fileLogger.filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("load log file failed! err:%v\n", err)
		panic(err)
	}
	fileLogger.fileObj = fileObj

	errFileObj, err := os.OpenFile(fileLogger.errFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("load err log file failed! err:%v\n", err)
		panic(err)
	}
	fileLogger.errFileObj = errFileObj

	return fileLogger
}

func (f FileLogger) log(level LogLevel, format string, a ...interface{})  {
	if f.level <= level {
		curTime := time.Now().Format("2006-01-02 15:04:05")
		levelString := getLogLevelString(level)
		fileName, funcName, lineNum := getInfo(3)
		msg := fmt.Sprintf(format, a...)
		fmt.Fprintf(f.fileObj, "[%s] [%s] [%s %s:%d] %s\n", curTime, levelString,
			fileName, funcName, lineNum, msg)
		if level >= ERROR {
			fmt.Fprintf(f.errFileObj, "[%s] [%s] [%s %s:%d] %s\n", curTime, levelString,
				fileName, funcName, lineNum, msg)
		}
	}
}

func (f FileLogger) Debug(format string, a ...interface{})  {
	f.log(DEBUG, format, a...)
}

func (f FileLogger) Trace(format string, a ...interface{})  {
	f.log(TRACE, format, a...)
}

func (f FileLogger) Info(format string, a ...interface{})  {
	f.log(INFO, format, a...)
}

func (f FileLogger) Warning(format string, a ...interface{})  {
	f.log(WARNING, format, a...)
}

func (f FileLogger) Error(format string, a ...interface{})  {
	f.log(ERROR, format, a...)
}

func (f FileLogger) Fatal(format string, a ...interface{})  {
	f.log(FATAL, format, a...)
}

