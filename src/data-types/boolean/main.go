package main

import "fmt"

func main() {
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
