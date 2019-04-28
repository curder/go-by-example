package main

import "fmt"

func main() {
	// 字符串通过 + 拼接
	fmt.Println("hello" + " world")

	// 整型运算
	fmt.Println("1 + 1 = ", 1+1)

	// 浮点数运算
	fmt.Println("7.0 / 3.0 = ", 7.0/3.0)

	// 布尔型
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

}
