package logger

import (
	"testing"
)

// 向终端写日志相关内容
func TestConsoleLogger(t *testing.T) {
	log := NewConsoleLog("error")

	log.Debug("this is a debug log.")
	log.Trace("this is a trace log.")
	log.Info("this is a info log.")
	log.Warning("this is a warning log.")
	log.Error("this is a error log.")
	log.Fatal("this is a fatal log. username: %s, age: %d", "curder", 28)
}
