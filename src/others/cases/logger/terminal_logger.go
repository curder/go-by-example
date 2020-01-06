package logger

import (
	"fmt"
	"path"
	"runtime"
	"strings"
	"time"
)

// 错误等级
type LogLevel uint16

const (
	DEBUG LogLevel = iota
	TRACE
	INFO
	WARNING
	ERROR
	FATAL
)

type Logger struct {
	Level LogLevel
}

func NewLog(level string) Logger {
	logLevel := parseLogLevel(level)
	return Logger{Level: logLevel}
}

// 分析日志级别，返回LogLevel
func parseLogLevel(s string) LogLevel {
	s = strings.ToLower(s)
	switch s {
	case "debug":
		return DEBUG
	case "trace":
		return TRACE
	case "info":
		return INFO
	case "warning":
		return WARNING
	case "error":
		return ERROR
	case "fatal":
		return FATAL
	default:
		return DEBUG
	}
}

func getLogLevel(logLevel LogLevel) string  {
	switch logLevel {
	case DEBUG:
		return "DEBUG"
	case TRACE:
		return "TRADE"
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

// 判断日志级别是否允许打印日志
func (l Logger) enable(logLevel LogLevel) bool {
	return logLevel >= l.Level
}

// 获取执行的文件名，函数名和行号
func getInfo(skip int) (fileName string, funcName string, line int){
	// file 文件名；line 行；ok 状态
	pc, file, line, ok := runtime.Caller(skip)

	if !ok {
		fmt.Printf("runtime.Caller() faild, error: %v \n", ok)
	}
	fileName = path.Base(file)

	funcName = strings.Split(path.Base(runtime.FuncForPC(pc).Name()), ".")[1]

	return
}

// 记录日志
func (l Logger) log(logLevel LogLevel, msg string) {
	level := getLogLevel(logLevel)
	now := time.Now()
	fileName, filePath, line := getInfo(3)
	fmt.Printf("[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), level, fileName, filePath, line, msg)
}

func (l Logger) Debug(msg string) {
	if l.enable(DEBUG) {
		l.log(DEBUG, msg)
	}
}

func (l Logger) Trace(msg string) {
	if l.enable(TRACE) {
		l.log(TRACE, msg)
	}
}

func (l Logger) Info(msg string) {
	if l.enable(INFO) {
		l.log(INFO, msg)
	}
}

func (l Logger) Warning(msg string) {
	if l.enable(WARNING) {
		l.log(WARNING, msg)
	}
}

func (l Logger) Error(msg string) {
	if l.enable(ERROR) {
		l.log(ERROR, msg)
	}
}

func (l Logger) Fatal(msg string) {
	if l.enable(FATAL) {
		l.log(FATAL, msg)
	}
}
