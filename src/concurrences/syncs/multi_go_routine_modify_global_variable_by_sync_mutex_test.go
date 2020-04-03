package syncs

import (
	"sync"
	"testing"
)

// 并发安全和锁
// 有时候在Go代码中可能会存在多个goroutine同时操作一个资源（临界区），这种情况会发生竞态问题（数据竞态）。
// 类比现实生活中的例子有十字路口被各个方向的的汽车竞争；还有火车上的卫生间被车厢里的人竞争。

var mutexWg sync.WaitGroup
var mutexLock sync.Mutex

// 互斥锁
// 互斥锁是一种常用的控制共享资源访问的方法，它能够保证同时只有一个goroutine可以访问共享资源。
// Go语言中使用sync包的Mutex类型来实现互斥锁。

// 使用互斥锁能够保证同一时间有且只有一个goroutine进入临界区，其他的goroutine则在等待锁；
// 当互斥锁释放后，等待的goroutine才可以获取锁进入临界区，多个goroutine同时等待一个锁时，唤醒的策略是随机的。
func TestGoRoutineByMutex(t *testing.T) {
	mutexWg.Add(2)
	go calculatorSumAndLock() // 开启一个go routine
	go calculatorSumAndLock() // 再开启一个 go routine
	mutexWg.Wait()
	t.Log(result) // 在修改变量时候加锁，修改变量后解锁的操作得到正确的值：1000000，100万
}

// 定义一个修改全局变量 result 的函数，在修改全局的同时上锁
func calculatorSumAndLock() {
	for i := 0; i < 500000; i++ { // 循环 50 万次
		mutexLock.Lock() // 修改全局变量之前上锁
		result = result + 1
		mutexLock.Unlock() // 修改完全局变量解锁
	}
	mutexWg.Done()
}
