package main

import "fmt"
import "time"

func main() {
	i := 2
	// 基本的 switch 结构
	switch i {
	case 1:
		fmt.Println("一")
	case 2:
		fmt.Println("二")
	case 3:
		fmt.Println("三")
	}

	// 或者将变量的定义也一并写在switch结构体，这样结构体外就不会存在多余的变量
	switch i:= 3; i {
	case 1:
		fmt.Println("一")
	case 2:
		fmt.Println("二")
	case 3:
		fmt.Println("三")
	}

	// 在一个 `case` 语句中，你可以使用逗号来分隔多个表达式。在这个例子中，我们很好的使用了可选的`default` 分支。
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("今天是周末")
	default:
		fmt.Println("今天是工作日")
	}

	// 不带表达式的 `switch` 是实现 if/else 逻辑的另一种方式。这里展示了 `case` 表达式是如何使用非常量的。
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("上午")
	default:
		fmt.Println("下午")
	}
}
