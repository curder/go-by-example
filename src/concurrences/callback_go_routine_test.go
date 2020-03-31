package concurrences

import (
	"testing"
	"time"
)

// goroutine使用匿名函数
func TestSignalGoRoutineCallback(t *testing.T) {
	t.Log("hello")
	go func() {
		t.Log("Hello GoRoutine! \n")
	}()
	time.Sleep(time.Second) // 进程休眠 1S

	t.Log("end")
}

// goroutine使用匿名函数
func TestMultiGoRoutineCallback(t *testing.T) {
	t.Log("hello")

	for i := 0; i < 10; i++ {
		go func(i int) {
			t.Log("Hello Multi GoRoutine! \n", i)
		}(i) // 这里需要使用函数传递参数，而不是for循环中国呢的i参数
	}
	time.Sleep(time.Second) // 进程休眠 1S

	t.Log("end")
}
