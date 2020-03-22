package main

import (
	"fmt"
)

func main() {

	// 经典的初始化/条件/后续形式 `for` 循环
	for j := 7; j <= 9; j++ {
		fmt.Printf("循环结构体中j的值为：%d\n", j)
	}

	// for变种，将退出条件放在 {} 中
	for i := 0; i <= 3; {
		fmt.Println(i)
		i = i + 1
	}

	// 不带条件的 `for` 循环将一直执行，直到在循环体内使用 `break` 来跳出循环 或者 `continue` 跳过此处循环，继续下一次循环。
	for {
		fmt.Println("loop")
		break
	}

	// for range 循环
	s := "hello中国"
	for i, v := range s {
		fmt.Printf("字符串变量s %s，索引为：%d 值为：%c\n", s, i, v)
	}

	println("使用for循环打印 99 乘法表")
	// 99乘法表
	for x := 1; x <= 9; x++ {
		for y := 1; y <= x; y++ {
			fmt.Printf("%d * %d = %d ", x, y, x*y)
		}
		fmt.Println()
	}

	// 跳出多层循环的两种写法
	for i := 0; i < 10; i++ {
		var flag bool
		for j := 'A'; j <= 'Z'; j++ {
			if j == 'C' {
				flag = true
				break
			}
			fmt.Printf("%d - %c \n", i, j)
		}
		if flag {
			fmt.Println()
			break
		}
	}

	// 使用 goto - label 方式达到同样的效果
	for i := 0; i < 10; i++ {
		for j := 'A'; j <= 'Z'; j++ {
			if j == 'C' {
				goto LABEL
			}
			fmt.Printf("%d - %c \n", i, j)
		}
	}
LABEL:
	fmt.Printf("for over")
}
