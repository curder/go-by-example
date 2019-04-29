// 超时对于一个连接外部资源，或者其它一些需要花费执行时间的操作的程序而言是很重要的
// 得益于通道和 select，在 Go 中实现超时操作是简洁而优雅的
package main

import (
	"fmt"
	"time"
)

func main() {
	// 假如执行一个外部调用，并在 2 秒后通过通道 `channelOne` 返回它的执行结果
	channelOne := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		channelOne <- "result 1"
	}()

	// 这里是使用 `select` 实现一个超时操作
	// `res := <- channelOne` 等待结果，`<-Time.After` 等待超时时间 1 秒后发送的值
	// 由于 `select` 默认处理第一个已准备好的接收操作，如果这个操作超过了允许的 1 秒的话，将会执行超时 case
	select {
	case res := <-channelOne:
		fmt.Println(res)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 1")
	}

	// 如果我允许一个超时时间 3 秒，将会成功的从 `channelTwo` 接收到值，并且打印出结果
	channelTwo := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		channelTwo <- "result 2"
	}()

	select {
	case res := <-channelTwo:
		fmt.Println(res)
	case <-time.After(time.Second * 3):
		fmt.Println("timeout 2")
	}
}
