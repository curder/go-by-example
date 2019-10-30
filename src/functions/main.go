package main

import "fmt"

// 求和函数
func Plus(a int, b int) int {
	return a + b // Go 需要明确的返回值，它不会自动返回最后一个表达式的值
}

// 无返回值
func f1(arg1 int, arg2 int) {
	fmt.Println(arg1 + arg2)
}

// 多个返回值函数
func f2() (code int, message string) { // 多个返回值必须使用 () 包裹起来
	return 404, "Not Found!"
}

// 没有参数没有返回值
func f3() {
	fmt.Println("No Arguments & No Return.")
}

// 返回值参数命名函数，命名的参数相当于在函数中声明了一个变量
func f4(arg1 int, arg2 int) (sum int) {
	sum = arg1 + arg2
	return
}

// 参数类型简写，当参数中有连续多个参数的类型一致时可以将非最后一个参数的类型省略
func f5(arg1, arg2 int, arg3, arg4, arg5 string, arg6, arg7, arg8 bool) bool {
	return false
}

// 可变长参数
func f6(arg1 string, arg2... int) bool {
	fmt.Println(arg1)
	fmt.Println(arg2) // arg2的类型是切片 []int
	return false
}

func main() {
	fmt.Printf("1 + 2 = %d\n", Plus(1, 2))
}
