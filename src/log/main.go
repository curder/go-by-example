package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	file, err := os.OpenFile("./log.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open file failed %v \n", err)
		return
	}

	for {
		log.SetOutput(file) // 设置输出到文件句柄，如果不设置则一般输出到终端
		log.Println("test error message")
		time.Sleep(time.Second * 5) // 休眠5秒
	}
}
