// 常常需要在某个时刻运行 Go 代码，或者在某段时间间隔内重复运行。Go 的内置 定时器 和 打点器 特性让这写很容易实现
package main

import (
	"fmt"
	"time"
)

func main() {

	// 定时器表示在未来某一时刻的独立事件。告诉定时器需要等待的时间，然后它将提供一个用于通知的通道
	// 这里的定时器将等待 2 秒
	timerOne := time.NewTimer(time.Second * 2)

	// `<-timerOne.C` 直到这个定时器的通道 `C` 明确的发送了定时器失效的值之前，将一直阻塞
	<-timerOne.C
	fmt.Println("Timer 1 expired")

	// 如果需要的仅仅是单纯的等待，需要使用 `time.Sleep`，定时器是有用原因之一就是可以在定时器失效之前，取消这个定时器
	timerTwo := time.NewTimer(time.Second)
	go func() {
		<-timerTwo.C
		fmt.Println("Timer 2 expired")
	}()

	stopTwo := timerTwo.Stop()
	if stopTwo {
		fmt.Println("Timer 2 stopped")
	}
}
