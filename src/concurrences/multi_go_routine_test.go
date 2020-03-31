package concurrences

import (
	"sync"
	"testing"
)

var workGroup sync.WaitGroup

// 使用 sync.WaitGroup 0来实现goroutine的同步
// 执行下面的代码，会发现每次打印的数字的顺序都不一致。
// 这是因为10个goroutine是并发执行的，而goroutine的调度是随机的。
func TestMultiGoRoutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		workGroup.Add(1) // 启动一个goroutine就登记+1
		go multiGoRoutine(t, i)
	}
	workGroup.Wait() // 等待所有登记的goroutine都结束
}

func multiGoRoutine(t *testing.T, i int) {
	t.Log("Hello Multi GoRoutine! ", i, "\n")
	defer workGroup.Done() // goroutine结束就登记-1
}
