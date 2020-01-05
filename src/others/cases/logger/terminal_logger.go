package logger

import (
	"fmt"
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

// 判断日志级别是否允许打印日志
func (l Logger) enable(logLevel LogLevel) bool {
	return l.Level > logLevel
}

func (l Logger) Debug(msg string) {
	if l.enable(DEBUG) {
		now := time.Now()
		fmt.Printf("[%s] [%s] %s\n", now.Format("2006-01-02 15:04:05"), "DEBUG", msg)
	}
}

func (l Logger) Trace(msg string) {
	if l.enable(TRACE) {
		now := time.Now()
		fmt.Printf("[%s] [%s] %s\n", now.Format("2006-01-02 15:04:05"), "TRACE", msg)
	}
}

func (l Logger) Info(msg string) {
	if l.enable(INFO) {
		now := time.Now()
		fmt.Printf("[%s] [%s] %s\n", now.Format("2006-01-02 15:04:05"), "INFO", msg)
	}
}

func (l Logger) Warning(msg string) {
	if l.enable(WARNING) {
		now := time.Now()
		fmt.Printf("[%s] [%s] %s\n", now.Format("2006-01-02 15:04:05"), "WARNING", msg)
	}
}

func (l Logger) Error(msg string) {
	if l.enable(ERROR) {
		now := time.Now()
		fmt.Printf("[%s] [%s] %s\n", now.Format("2006-01-02 15:04:05"), "ERROR", msg);
	}
}

func (l Logger) Fatal(msg string) {
	if l.enable(FATAL) {
		now := time.Now()
		fmt.Printf("[%s] [%s] %s\n", now.Format("2006-01-02 15:04:05"), "FATAL", msg)
	}
}
