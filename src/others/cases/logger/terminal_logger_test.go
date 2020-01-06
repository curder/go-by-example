package logger

import (
	"testing"
)

// 向终端写日志相关内容
func TestTerminalLogger(t *testing.T) {
	log := NewLog("error")

	log.Debug("this is a debug log.")
	log.Trace("this is a trace log.")
	log.Info("this is a info log.")
	log.Warning("this is a warning log.")
	log.Error("this is a error log.")
	log.Fatal("this is a fatal log.")
}
