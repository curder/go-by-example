package logger

import (
	"testing"
)

// 向文件写日志相关内容
func TestFileLogger(t *testing.T) {
	log := NewFileLog("warning", "./", "go.log", 10*1024*1024) // 文件大小 10M 切割

	for {
		log.Debug("this is a debug log.")
		log.Trace("this is a trace log.")
		log.Info("this is a info log.")
		log.Warning("this is a warning log.")
		log.Error("this is a error log.")
		log.Fatal("this is a fatal log. username: %s, age: %d", "curder", 28)
	}

	defer log.closeFile() // 关闭文件句柄
}
