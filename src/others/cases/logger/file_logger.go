package logger

import (
	"fmt"
	"os"
	"path"
	"time"
)

type FileLogger struct {
	Level           LogLevel // 日志级别
	FilePath        string   // 文件路径
	FileName        string   // 文件名
	FileObject      *os.File // 文件对象
	ErrorFileObject *os.File // 错误文件对象
	MaxFileSize     int64    // 最大文件值
}

// 构造函数
func NewFileLog(l string, filePath string, fileName string, maxFileSize int64) *FileLogger {
	logLevel := parseLogLevel(l)

	f := &FileLogger{
		Level:       logLevel,
		FilePath:    filePath,
		FileName:    fileName,
		MaxFileSize: maxFileSize,
	}

	f.initFile()

	return f
}

// 初始化文件句柄
func (l *FileLogger) initFile() error {
	fullName := path.Join(l.FilePath, l.FileName)
	fileObject, err := os.OpenFile(fullName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("log init fail, error: %v", err)
	}

	errorFileObject, err := os.OpenFile(fullName+".error", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("log init fail, error: %v", err)
	}

	l.FileObject = fileObject
	l.ErrorFileObject = errorFileObject

	return nil
}

// 关闭文件句柄
func (l *FileLogger) closeFile() {
	l.FileObject.Close()
	l.ErrorFileObject.Close()
}

// 判断日志级别是否允许打印日志，如果日志级别大于给定的级别则记录日志，否则不记录日志
func (l *FileLogger) enable(logLevel LogLevel) bool {
	return logLevel >= l.Level
}

// 按文件大小切割，检查文件大小判断是否需要切割
func (l *FileLogger) checkFileSize(fileObject *os.File) bool {
	// 获取文件大小
	fileInfo, _ := fileObject.Stat()

	return fileInfo.Size() >= l.MaxFileSize
}

// 记录日志
func (l *FileLogger) log(logLevel LogLevel, format string, a ...interface{}) {

	if l.enable(logLevel) {
		level := getLogLevel(logLevel)
		now := time.Now()
		fileName, filePath, line := getInfo(3)
		msg := fmt.Sprintf(format, a...)
		if l.checkFileSize(l.FileObject) {
			// 将新打开的文件句柄赋值给 l.fileObject
			l.FileObject = l.splitFile(l.FileObject)
		}

		fmt.Fprintf(l.FileObject, "[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), level, fileName, filePath, line, msg)
		// 如果记录的日志 >= error 级别，需要在err日志中记录
		if logLevel >= ERROR {
			if l.checkFileSize(l.ErrorFileObject) {
				// 将新打开的文件句柄赋值给 l.ErrorFileObject
				l.ErrorFileObject = l.splitFile(l.ErrorFileObject)
			}
			fmt.Fprintf(l.ErrorFileObject, "[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), level, fileName, filePath, line, msg)
		}
	}
}

// 切割文件
func (l *FileLogger) splitFile(file *os.File) (f *os.File) {
	// 获取日志文件名
	originFileInfo, _ := file.Stat()

	// 重命名
	oldFullName := path.Join(l.FilePath, originFileInfo.Name())

	newLogName := originFileInfo.Name() + time.Now().Format(".20060102130405")
	os.Rename(oldFullName, l.FilePath+newLogName)

	// 打开新的文件句柄
	fileObject, _ := os.OpenFile(oldFullName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

	// 关闭当前文件
	file.Close()

	return fileObject
}

// 日志级别 Debug
func (l *FileLogger) Debug(format string, a ...interface{}) {
	l.log(DEBUG, format, a...)
}

// 日志级别 Trace
func (l *FileLogger) Trace(format string, a ...interface{}) {
	l.log(TRACE, format, a...)
}

// 日志级别 Info
func (l *FileLogger) Info(format string, a ...interface{}) {
	l.log(INFO, format, a...)
}

// 日志级别 Warning
func (l *FileLogger) Warning(format string, a ...interface{}) {
	l.log(WARNING, format, a...)
}

// 日志级别 Error
func (l *FileLogger) Error(format string, a ...interface{}) {
	l.log(ERROR, format, a...)
}

// 日志级别 Fatal
func (l *FileLogger) Fatal(format string, a ...interface{}) {
	l.log(FATAL, format, a...)
}
