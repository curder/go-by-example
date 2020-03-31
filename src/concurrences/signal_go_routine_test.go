package concurrences

import (
	"testing"
	"time"
)

// Go语言中使用goroutine非常简单，只需要在调用函数的时候在前面加上go关键字，就可以为一个函数创建一个goroutine。
func TestSignalGoRoutine(t *testing.T) {
	t.Log("hello")

	go signalGoRoutine(t) // 启动另外一个goroutine去执行hello函数

	time.Sleep(time.Second) // 进程休眠 1S

	t.Log("end")
}

func signalGoRoutine(t *testing.T) {
	t.Log("Hello Signal GoRoutine!")
}
