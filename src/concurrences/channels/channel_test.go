package channels

import "testing"

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
