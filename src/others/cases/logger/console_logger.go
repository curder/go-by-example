package logger

import (
	"fmt"
	"path"
	"runtime"
	"strings"
	"time"
)

type ConsoleLogger struct {
	Level LogLevel
}

func NewConsoleLog(l string) ConsoleLogger {
	level := parseLogLevel(l)
	return ConsoleLogger{Level: level}
}

// 获取执行的文件名，函数名和行号
func getInfo(skip int) (fileName string, funcName string, line int) {
	// file 文件名；line 行；ok 状态
	pc, file, line, ok := runtime.Caller(skip)

	if !ok {
		fmt.Printf("runtime.Caller() faild, error: %v \n", ok)
	}
	fileName = path.Base(file)

	funcName = strings.Split(path.Base(runtime.FuncForPC(pc).Name()), ".")[1]

	return
}

// 判断日志级别是否允许打印日志
func (l ConsoleLogger) enable(logLevel LogLevel) bool {
	return logLevel >= l.Level
}

// 记录日志
func (l ConsoleLogger) log(logLevel LogLevel, format string, a ...interface{}) {
	if l.enable(logLevel) {
		level := getLogLevel(logLevel)
		now := time.Now()
		fileName, filePath, line := getInfo(3)
		msg := fmt.Sprintf(format, a...)
		fmt.Printf("[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), level, fileName, filePath, line, msg)
	}
}

func (l ConsoleLogger) Debug(format string, a ...interface{}) {
	l.log(DEBUG, format, a...)
}

func (l ConsoleLogger) Trace(format string, a ...interface{}) {
	l.log(TRACE, format, a...)
}

func (l ConsoleLogger) Info(format string, a ...interface{}) {
	l.log(INFO, format, a...)
}

func (l ConsoleLogger) Warning(format string, a ...interface{}) {
	l.log(WARNING, format, a...)
}

func (l ConsoleLogger) Error(format string, a ...interface{}) {
	l.log(ERROR, format, a...)
}

func (l ConsoleLogger) Fatal(format string, a ...interface{}) {
	l.log(FATAL, format, a...)
}
