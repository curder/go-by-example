package main

import "fmt"

func main() {
	// ➕ 当两边都是字符串时，结果为拼接；当两边都是整型时，结果为加法运算。
	// 字符串通过 + 拼接
	fmt.Println("hello" + " world")

	// 整型 + 运算
	fmt.Println("1 + 1 =", 1+1)

	// 十进制整形
	var i1 = 101
	fmt.Printf("%d\n", i1)
	fmt.Printf("%b\n", i1) // 转换成二进制输出
	fmt.Printf("%o\n", i1) // 转换成八进制输出
	fmt.Printf("%x\n", i1) // 转换成十六进制输出

	// 八进制整形
	i2 := 077
	fmt.Printf("%d\n", i2)

	// 十六进制整形
	i3 := 0x1234567
	fmt.Printf("%d\n", i3)

	// 查看变量类型
	fmt.Printf("%T \n", i3)

	// 声明 int8 类型变量
	i4 := int8(32)
	fmt.Printf("%T\n", i4)


	// 浮点数运算
	fmt.Println("7.0 / 3.0 = ", 7.0/3.0)

	f1 := 1.23
	fmt.Printf("f1: %T\n", f1) // go语言中默认的浮点型为float64类型

	// 布尔型
	// Go语言中以bool类型进行声明布尔型数据，布尔型数据只有 true（真）和 false（假）两个值。
	// 布尔类型变量的默认值为false。
	// Go 语言中不允许将整型强制转换为布尔型.
	// 布尔型无法参与数值运算，也无法与其他类型进行转换。
	b1 := true
	var b2 bool // 默认值是 flase
	fmt.Printf("%T\n", b1)
	fmt.Printf("%T\n", b2)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
