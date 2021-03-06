package main

import "fmt"

// 变量、指针地址、指针变量、取地址、取值的相互关系和特性如下：
// 	1. 对变量进行取地址（&）操作，可以获得这个变量的指针变量。
// 	2. 指针变量的值是指针地址。
// 	3. 对指针变量进行取值（*）操作，可以获得指针变量指向的原变量的值。

func main() {
	// Go 支持指针，允许在程序中通过引用传递值或者数据结构。
	// & 取变量所在内存地址
	n := 1
	t := &n

	fmt.Printf("获取变量n的内存地址：%v，类型：%T \n", t, t)

	// * 根据地址取值
	v := *t
	fmt.Printf("获取指针t的值：%v，类型：%T \n", v, v)

	// 取地址操作符 & 和取值操作符 * 是一对互补操作符，& 取出地址，* 根据地址取出地址指向的值。

	var p1 *int    // nil 指针
	p2 := new(int) // new 函数申请一个内存地址
	fmt.Printf("声明一个空指针p1为：%v，声明一个零值的指针p2为：%v\n", p1, p2)
	fmt.Printf("获取修改p2前的值：%v\n", *p2)
	*p2 = 10
	fmt.Printf("获取修改p2后的值：%v\n", *p2)

	// new与make的区别
	//   二者都是用来做内存分配的。
	//   make 只用于 slice 、 map 以及 channel 的初始化，返回的还是这三个引用类型本身；
	//   而 new 用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针。
}
