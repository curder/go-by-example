package main

import "fmt"

// for 是 go 语言唯一的循环结构，
func main()  {
	i := 0
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// 经典的初始化/条件/后续形式 `for` 循环
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// 不带条件的 `for` 循环将一直执行，直到在循环体内使用了 `break` 或者 `return` 来跳出循环。
	for {
		fmt.Println("loop")
		break
	}
}
