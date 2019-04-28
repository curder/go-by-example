package main

import "fmt"

func main() {
	// 声明字符串
	var a string = "initial"
	fmt.Println(a)

	// 申明整型
	var b, c int = 1, 2
	fmt.Println(b, c)

	// 自定推断类型
	var d = true
	fmt.Println(d)

	// 声明变量且没有给出对应的初始值时，变量将会初始化为 零 值
	var e int
	fmt.Println(e)

	// `:=` 语句是申明并初始化变量的简写，例如下面例子中的 `var f string = "short"`。
	f := "short"
	fmt.Println(f)

}
