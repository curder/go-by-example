package main

import "fmt"

// Go语言中最常用的流程控制有if 和 for
// 而switch和goto主要是为了简化代码、降低重复代码而生的结构，属于扩展类的流程控制。
func main() {
	if 8%4 == 0 {
		fmt.Println("8能被4整除")
	}

	if 7%2 == 0 {
		fmt.Println("7是偶数")
	} else {
		fmt.Println("7是奇数")
	}

	// 在条件语句之前可以有赋值语句，且变量作用域只在 if 条件判断中；
	// 任何在这里声明的变量都可以在所有的条件分支中使用。
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
	// fmt.Println(num) // 这里无法找到 num 变量
}
