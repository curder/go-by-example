package syncs

import (
	"sync"
	"testing"
)

// Go语言中可以使用sync.WaitGroup来实现并发任务的同步。
// sync.WaitGroup有以下几个方法：
//   (wg * WaitGroup) Add(delta int)   计数器+delta
//   (wg *WaitGroup)  Done()           计数器-1
//   (wg *WaitGroup)  Wait()           阻塞直到计数器变为0

// sync.WaitGroup内部维护着一个计数器，计数器的值可以增加和减少。
// 例如当我们启动了N 个并发任务时，就将计数器值增加N。
// 每个任务完成时通过调用Done()方法将计数器减1。
// 通过调用Wait()来等待并发任务执行完，当计数器值为0时，表示所有并发任务已经完成。

// 注意：sync.WaitGroup是一个结构体，传递的时候要传递指针。

func TestGoRoutineWithoutSyncWaitGroup(t *testing.T) {
	go func() { // 开启一个go routine 打印字符串
		t.Log("Hello")
	}()

	t.Log("test is pass")
	// 程序执行结束都不会执行go routine
	// 我们当然可以使用 time.Sleep(d Duration) 来等待go routine异步执行完毕再推出主进程，但是则显然是不符合实际使用场景的。
}

func TestGoRoutineWithSyncWaitGroup(t *testing.T) {
	var wg sync.WaitGroup

	wg.Add(1)

	go func() { // 开启一个go routine 打印字符串
		t.Log("Hello")
		defer wg.Done()
	}()

	wg.Wait()
	t.Log("test is pass")
}
