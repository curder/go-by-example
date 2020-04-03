package main

import (
	"fmt"
)

func main() {
	// ➕ 当两边都是字符串时，结果为拼接；当两边都是整型时，结果为加法运算。
	// 字符串通过 + 拼接
	fmt.Println("hello" + " world")

	// 整型 + 运算
	fmt.Println("1 + 1 =", 1+1)

	// 十进制整型
	var i1 = 101
	fmt.Printf("i1十进制值：%d\n", i1)
	fmt.Printf("i1二进制值：%b\n", i1)  // 转换成二进制输出
	fmt.Printf("i1八进制值：%o\n", i1)  // 转换成八进制输出
	fmt.Printf("i1十六进制值：%x\n", i1) // 转换成十六进制输出

	// 八进制整型
	i2 := 077
	fmt.Printf("变量i2的值为: %d\n", i2)

	// 十六进制整型 范围 0 ~ f
	i3 := 0x1234567
	fmt.Printf("变量i3的值为：%d\n", i3)

	// 查看变量类型
	fmt.Printf("变量i3的类型为：%T \n", i3)

	// 声明 int8 类型变量
	i4 := int8(32)
	fmt.Printf("变量i4的类型为：%T\n", i4)

	// 浮点数运算
	fmt.Println("7.0 / 3.0 = ", 7.0/3.0)

	f1 := 1.23
	fmt.Printf("变量f1的类型为: %T\n", f1) // go语言中默认的浮点型为 float64 类型

	f2 := float32(1.23) // 显示的声明float32类型的浮点数
	fmt.Printf("变量f2的类型为：%T\n", f2)

	// 布尔型
	// Go语言中以bool类型进行声明布尔型数据，布尔型数据只有 true（真）和 false（假）两个值。
	// 布尔类型变量的默认值为false。
	// Go 语言中不允许将整型强制转换为布尔型.
	// 布尔型无法参与数值运算，也无法与其他类型进行转换。
	b1 := true
	var b2 bool // 不赋值的话，默认值是 flase
	fmt.Printf("变量b1的类型为：%T\n", b1)
	fmt.Printf("变量b2的类型为：%T，变量b2的值为：%v\n", b2, b2)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
