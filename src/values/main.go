package main

import "fmt"

func main() {
	// ➕ 当两边都是字符串时，结果为拼接；当两边都是整型时，结果为加法运算。
	// 字符串通过 + 拼接
	fmt.Println("hello" + " world")

	// 整型 + 运算
	fmt.Println("1 + 1 =", 1+1)

	// 浮点数运算
	fmt.Println("7.0 / 3.0 = ", 7.0/3.0)

	// 布尔型
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

}
