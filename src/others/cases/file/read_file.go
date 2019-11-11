package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 打开文件

func main() {
	filePath := "./read_file.go"

	// 通过io.Read获取文件内容
	//getFileContentByIORead(filePath)

	// 通过bufio获取文件内容
	//getFileControllerByBufIO(filePath)

	// 通过ioutil读取文件内容
	getFileControllerByIOUtil(filePath)
}

func getFileContentByIORead(filePath string) {
	fileObject, ok := os.Open(filePath)

	if ok != nil {
		fmt.Println("打开文件错误")
		return
	}

	defer fileObject.Close() // 延迟关闭

	var tmp = make([]byte, 128) // 构建一个包含128字节的切片
	for {
		n, err := fileObject.Read(tmp)

		if err != nil {
			fmt.Println("读取错误")
		}

		fmt.Printf("读取了%d个字节", n)
		fmt.Println(string(tmp[:n]))

		if n < 128 {
			return
		}
	}
}

func getFileControllerByBufIO(filePath string) {
	fileObject, ok := os.Open(filePath)

	if ok != nil {
		fmt.Println("打开文件错误")
		return
	}

	defer fileObject.Close() // 延迟关闭

	tmp := bufio.NewReader(fileObject)

	for {
		line, err := tmp.ReadString('\n')

		if err == io.EOF {
			fmt.Println("文件读取完毕～")
			return
		}

		if err != nil {
			fmt.Println("读取错误")
			return
		}
		fmt.Print(line)
	}
}

func getFileControllerByIOUtil(filePath string) {
	fileObject, ok := ioutil.ReadFile(filePath)

	if ok != nil {
		fmt.Println("打开或读取文件错误")
		return
	}

	fmt.Println(string(fileObject))

}
