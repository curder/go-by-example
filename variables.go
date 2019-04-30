package main

// Go变量声明的三种方式
// 1. 指定变量类型，声明后若不赋值使用默认值
// 2. 根据变量值自动判断变量类型（类型推导）
// 3. 使用 := 的方式声明并赋值，但是变量名不应该是已经声明过的，否则会导致编译出错

// Go 多变量声明，在同一行声明多个变量
import "fmt"

func main() {
	// 声明变量且没有给出对应的初始值时，变量将会初始化为 零 值
	var a int
	fmt.Println(a)

	// 自定推断类型
	var b = true
	fmt.Println(b)

	// `:=` 语句是申明并初始化变量的简写，例如下例等同于 `var f string = "short"`。
	c := "short"
	fmt.Println(c)

	// 声明字符串
	var d string = "initial"
	fmt.Println(d)

	// 多变量声明整型
	var e, f int = 1, 2
	fmt.Println(e, f)

	// 声明多个不同类型的变量
	g, h, i := 10, "name", 11
	fmt.Println(g, h, i)
}
