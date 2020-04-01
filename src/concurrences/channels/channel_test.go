package channels

import (
	"sync"
	"testing"
)

// Go 语言中的通道（channel）是一种特殊的类型。
// 通道像一个传送带或者队列，总是遵循先入先出（First In First Out）的规则，保证收发数据的顺序。
// 每一个通道都是一个具体类型的导管，也就是声明channel的时候需要为其指定元素类型。

// channel是一种类型，一种引用类型。

// 声明通道类型的格式如下
func TestDeclarativeChannel(t *testing.T) {
	// 格式：var 变量 chan 元素类型

	var ch1 chan int      // 声明一个传递整型的通道
	var ch2 chan bool     // 声明一个传递布尔的通道
	var ch3 chan []string // 声明一个传递int切片的通道

	// 通道是引用类型，通道类型的空值是nil
	t.Logf("ch1的类型是：%T，值是：%#v", ch1, ch1)
	t.Logf("ch2的类型是：%T，值是：%#v", ch2, ch2)
	t.Logf("ch3的类型是：%T，值是：%#v", ch3, ch3)
}

// 创建通道
// 声明的通道后需要使用make函数初始化之后才能使用
func TestCreateChannel(t *testing.T) {
	//格式：make(chan 元素类型, [缓冲大小])
	var ch1 = make(chan int, 10)
	t.Logf("ch1的类型是：%T，值是：%#v", ch1, ch1)
}

// 通道的操作
// 发送操作 `ch1<- 1`
// 接受操作 `<- ch1`
// 关闭操作 `close()`
func TestSendAndReceiveChannel(t *testing.T) {
	var wg sync.WaitGroup
	var ch1 = make(chan int, 1) // 声明一个有缓冲大小的通道类型变量
	sendVal := 10               // 定义一个待发送的int值

	ch1 <- sendVal // 给通道中发送一个值
	t.Logf("给通道中发送一个值：%d", sendVal)

	wg.Add(1)
	go func() {
		defer wg.Done()

		receiveVal := <-ch1 // 获取通道里的值
		t.Logf("接收到通道中的值：%d", receiveVal)
	}()

	close(ch1) // 关闭通道
	wg.Wait()
}

// 通道的关闭补充
func TestCloseChannel(t *testing.T) {
	ch1 := make(chan bool, 2)

	ch1 <- true
	t.Logf("第一次往通道中写值%v \n",true)
	ch1 <- true
	t.Logf("第二次往通道中写值%v \n",true)


	close(ch1) // 往通道写完值记得使用close关闭通道，否则会出现deadlock

	x, ok := <- ch1
	t.Logf("第一次从通道中取值%v, %v \n",x,ok)
	x, ok = <- ch1
	t.Logf("第二次从通道中取值%v, %v \n",x,ok)

	x, ok = <- ch1 // 再读就读取不到数据
	t.Logf("第三次从通道中取值%v, %v \n",x,ok)
}

// 练习
// 启动一个goroutine，生成100个数发送到通道ch1
// 启动一个goroutine，从ch1中取值，计算其值的平方放到通道ch2
// 打印通道ch2中的值
func TestSendAndReceiveChannelDemo(t *testing.T) {
	ch1 := make(chan int, 100) // 创建通道ch1
	ch2 := make(chan int, 100) // 创建通道ch2

	go func() { // 开启goroutine将0~100的数发送到ch1中
		for i := 0; i < 100; i++ {
			ch1 <- i
		}
		close(ch1) // 关闭通道
	}()

	go func() { // 开启goroutine从ch1中接收值，并将该值的平方发送到ch2中
		for {   // 循环读取ch1通道数据
			val, ok := <-ch1
			if !ok { // 当ok值状态不为nil的时候取值传入到通道ch2
				break
			}
			ch2 <- val * val // 取值求平方，并放入到通道ch2
		}
		close(ch2) // 关闭通道
	}()

	for ret := range ch2 { // 在主goroutine中从ch2中接收值打印
		t.Log(ret)
	} // 通道关闭后会退出for range循环
}
