package channels

import (
	"testing"
)

// select多路复用
// 在某些场景下我们需要同时从多个通道接收数据。
// 通道在接收数据时，如果没有数据可以接收将会发生阻塞。

// select的使用类似于switch语句，它有一系列case分支和一个默认的分支。
// 每个case会对应一个通道的通信（接收或发送）过程。select会一直等待，直到某个case的通信操作完成时，就会执行case分支对应的语句。
// 格式如下：
// select{
//    case <-ch1:
//        ...
//    case data := <-ch2:
//        ...
//    case ch3<-data:
//        ...
//    default:
//        默认操作
// }

func TestSimpleChannelSelect(t *testing.T) {
	intChan := make(chan int, 1)

	for i := 0; i < 10; i++ {
		select {
		case val := <-intChan: // 从通道intChan取值并赋给val变量
			t.Log(val)
		case intChan <- i: // 将值存入intChan通道中
		}
	}
}

// 使用select语句能提高代码的可读性。
// 1. 可处理一个或多个channel的发送/接收操作。
// 2. 如果多个case同时满足，select会随机选择一个。
// 3. 对于没有case的select{}会一直等待，可用于阻塞main函数或主进程。
