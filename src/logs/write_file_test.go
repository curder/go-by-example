package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

// 通过io.Write方式写文件
func TestSetFileContentByIOWrite(t *testing.T) {
	filePath := "./io_write.txt"
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE|os.O_TRUNC, 0644)

	if err != nil {
		fmt.Println("文件发生错误")
		return
	}

	defer f.Close() // 延迟关闭文件

	f.Write([]byte("hello\n"))
	f.WriteString("World\n")
}

// 通过bufio的方式写文件
func TestSetFileContentByBufIO(t *testing.T) {
	filePath := "./buf_io.txt"
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE|os.O_TRUNC, 0644)

	if err != nil {
		fmt.Println("文件发生错误")
		return
	}

	defer f.Close() // 延迟关闭文件

	w := bufio.NewWriter(f)

	w.WriteString("Hello World\n")
	w.Flush() // 将缓存区的内容刷新到磁盘
}

// 通过iO.util的方式写文件
func TestSetFileContentByIOUtil(t *testing.T) {
	filePath := "./io_util.txt"
	str := "hello world"
	err := ioutil.WriteFile(filePath, []byte(str), 0644)

	if err != nil {
		fmt.Printf("文件写入时发生错误")
		return
	}
}
