// 如何使用 Go 协程和通道实现一个 工作池
package main

import (
	"fmt"
	"time"
)

// 将要在多个并发实例中支持的任务。这些执行者将从 `jobs` 通道接收任务，并且通过 `results` 发送对应的结果。
// 让每个任务间隔 1s 来模仿一个耗时的任务
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}

func main() {

	// 为了使用 worker 工作池并且收集他们的结果，我们需要 2 个通道
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// 这里启动了 3 个 worker，初始是阻塞的，因为还没有传递任务
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// 这里发送 9 个 `jobs`，然后 `close` 这些通道来表示这就是所有的任务

	for j := 1; j <= 9; j++ {
		jobs <- j
	}
	close(jobs)

	// 最后，收集所有这些任务的返回值
	for a := 1; a <= 9; a++ {
		<-results
	}
}
