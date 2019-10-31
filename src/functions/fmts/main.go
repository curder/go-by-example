package main

import "fmt"

// fmt包分为向外输出内容和获取输入内容两大部分。
func main() {
	// Print
	// 获取输出
	// Print系列函数会将内容输出到系统的标准输出
	fmt.Print("hello")

	// 区别在于Print函数直接输出内容，Printf函数支持格式化输出字符串，Println函数会在输出内容的结尾添加一个换行符
	// %T 类型
	// %d 十进制
	// %b 二进制
	// %o 八进制
	// %x 十六进制
	// %c 字符
	// %s 字符片
	// %p 指针
	// %v 值
	// %f 浮点数
	// %t 布尔值
	fmt.Printf("%v", "world")
	fmt.Println()
	fmt.Println("hello world")

	// 获取输入
	// fmt.Scan、fmt.Scanf、fmt.Scanln三个函数，可以在程序运行过程中从标准输入获取用户的输入
	var s string
	fmt.Print("请输入内容：")
	fmt.Scan(&s) // 获取用户的输入
	fmt.Println("输入的内容是：", s)

	var name string
	var age int
	fmt.Print("请输入姓名 年龄：")
	fmt.Scanf("%s %d \n", &name, &age)
	fmt.Printf("name: %s, age: %d\n", name, age)
}
