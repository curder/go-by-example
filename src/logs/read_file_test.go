package main

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

const filePath = "./read_file_test.go"

// 通过 io.Read 获取文件内容
func TestGetFileContentByIORead(t *testing.T) {
	fileObject, ok := os.Open(filePath)

	if ok != nil {
		t.Log("打开文件错误")
		return
	}

	defer fileObject.Close() // 延迟关闭

	var tmp = make([]byte, 128) // 构建一个包含128字节的切片
	for {
		n, err := fileObject.Read(tmp)

		if err != nil {
			t.Log("读取错误")
			return
		}

		t.Logf("读取了%d个字节", n)
		t.Log(string(tmp[:n]))

		if n < 128 {
			return
		}
	}
}

// TestGetFileControllerByBufIO 通过 bufio 逐行读取文件内容
func TestGetFileControllerByBufIO(t *testing.T) {
	fileObject, ok := os.Open(filePath)

	if ok != nil {
		t.Log("打开文件错误")
		return
	}

	defer fileObject.Close() // 延迟关闭

	tmp := bufio.NewReader(fileObject)

	for {
		line, err := tmp.ReadString('\n')

		if err == io.EOF {
			t.Log("文件读取完毕～")
			return
		}

		if err != nil {
			t.Log("读取错误")
			return
		}
		t.Log(line)
	}

}

// 通过 ioutil 读取文件内容
func TestGetFileControllerByIOUtil(t *testing.T) {
	if fileObject, ok := ioutil.ReadFile(filePath); ok != nil {
		t.Log("打开或读取文件错误")
		return
	} else {
		t.Log(string(fileObject))
	}
}
