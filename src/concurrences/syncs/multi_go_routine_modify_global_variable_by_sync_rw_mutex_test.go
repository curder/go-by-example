package syncs

import (
	"sync"
	"testing"
	"time"
)

// 互斥锁是完全互斥的
// 但是有很多实际的场景下是读多写少的，当我们并发的去读取一个资源不涉及资源修改的时候是没有必要加锁的，这种场景下使用读写锁是更好的一种选择。
// 读写锁在Go语言中使用sync包中的RWMutex类型。

// 读写互斥锁
// 读写锁分为两种：读锁和写锁。
// 当一个goroutine获取读锁之后，其他的goroutine如果是获取读锁会继续获得锁，如果是获取写锁就会等待；
// 当一个goroutine获取写锁之后，其他的goroutine无论是获取读锁还是写锁都会等待。

var (
	x           int64
	rwMutexWg   sync.WaitGroup
	lock        sync.Mutex
	rwMutexLock sync.RWMutex
)

// TestGoRoutineByWRMutex
// 需要注意的是读写锁非常适合读多写少的场景，如果读和写的操作差别不大，读写锁的优势就发挥不出来。
func TestGoRoutineByWRMutex(t *testing.T) {
	start := time.Now() // 记录开始时间

	rwMutexWg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			// mutexLock.Lock() // 加互斥锁
			rwMutexLock.Lock() // 加写锁
			x = x + 1
			time.Sleep(10 * time.Millisecond) // 假设读操作耗时10毫秒
			rwMutexLock.Unlock() // 解写锁
			// mutexLock.Unlock() // 解互斥锁
			defer rwMutexWg.Done()
		}()
	}

	rwMutexWg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			// mutexLock.Lock() // 加互斥锁
			rwMutexLock.RLock() // 加读锁
			time.Sleep(time.Millisecond) // 假设读操作耗时1毫秒
			rwMutexLock.RUnlock() // 解读锁
			// mutexLock.Unlock() // 解互斥锁
			defer rwMutexWg.Done()
		}()
	}

	rwMutexWg.Wait()
	end := time.Now()     // 记录结束时间
	t.Log(end.Sub(start)) // 计算开始到结束时间的时间差，使用读写互斥锁运行时间大概在 100+ ms，而使用互斥锁运行时间大概在 1 s
}
