package syncs

// 在编程的很多场景下我们需要确保某些操作在高并发的场景下只执行一次，例如只加载一次配置文件、只关闭一次通道等。
// Go语言中的sync包中提供了一个针对只执行一次场景的解决方案–sync.Once。
// sync.Once只有一个Do方法，其签名如下：
// func (o *Once) Do(f func()) {}

import (
	"sync"
	"testing"
)

// sync.Once 是在代码运行中需要的时候执行，且只执行一次
func TestSyncOnce(t *testing.T) {
	var once sync.Once

	onceBody := func() { // 定义一个匿名函数供 sync.Once 使用
		t.Logf("sync.Once")
	}

	done := make(chan struct{})

	for i := 0; i < 10; i++ {
		go func() {
			once.Do(onceBody)
			done <- struct{}{}
			t.Log("write channel.")
		}()
	}
	for i := 0; i < 10; i++ {
		<-done
		t.Log("read channel.")
	}
}
