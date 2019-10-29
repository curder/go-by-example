package main

import (
	"fmt"
)

func main() {
	for i := 0; i <= 3; {
		fmt.Println(i)
		i = i + 1
	}

	// 经典的初始化/条件/后续形式 `for` 循环
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// 不带条件的 `for` 循环将一直执行，直到在循环体内使用 `break` 或者 `return` 来跳出循环。
	for {
		fmt.Println("loop")
		break
	}

	// for range 循环
	s := "hello中国"
	for i, v := range s {
		fmt.Printf("%d %c\n", i, v)
	}

	println("99乘法表")
	// 99乘法表
	for x := 1; x <= 9 ; x++  {
		for y := 1; y <= x ; y++  {
			fmt.Printf("%d * %d = %d " , x, y, x * y )
		}
		fmt.Println()
	}
}
