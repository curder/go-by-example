// Go 的通道选择器可以同时等待多个通道操作。Go 协程和通道以及选择器的结合是 Go 的一个强大特性
package main

import (
	"fmt"
	"time"
)

func main() {
	// 将从两个通道中选择
	channelOne := make(chan string)
	channelTwo := make(chan string)

	// 各个通道将在若干时间后接收一个值，这个用来模拟例如并行的 Go 协程中阻塞的 RPC 操作
	go func() {
		time.Sleep(time.Second * 1)
		channelOne <- "one"
	}()
	go func() {
		time.Sleep(time.Second * 2)
		channelTwo <- "two"
	}()

	// 使用 `select` 关键字来同时等待这两个值，并打印各自接收到的值
	for i := 0; i < 2; i++ {
		select {
		case messageOne := <-channelOne:
			fmt.Println("received", messageOne)
		case messageTwo := <-channelTwo:
			fmt.Println("received", messageTwo)
		}
	}
}
