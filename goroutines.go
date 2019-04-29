// Go 语言中使用go关键字来创建并启动一个协程，协程是一种轻量级的线程，占用系统资源更少
package main

import "fmt"

// _Go 协程_ 在执行上来说是轻量级的线程

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	// 假设有一个函数叫做 `f(s)`。我们使用一般的方式调用并同时运行
	f("direct")

	// 使用 `go f(s)` 在一个 Go 协程中调用这个函数。这个新的 Go 协程将会并行的执行这个函数调用
	go f("goroutine")

	// 也可以为匿名函数启动一个 Go 协程
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// 现在这两个 Go 协程在独立的 Go 协程中异步的运行，所以我们需要等它们执行结束
	// 这里的 `Scanln` 代码需要我们在程序退出前按下任意键结束
	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}
