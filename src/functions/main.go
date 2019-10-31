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
func f6(arg1 string, arg2 ... int) bool {
	fmt.Println(arg1)
	fmt.Println(arg2) // arg2的类型是切片 []int
	return false
}

// 函数中的延迟执行 defer
// Go语言中的defer语句会将其后面跟随的语句进行延迟处理。
// 在defer归属的函数即将返回时，将延迟处理的语句按defer定义的逆序进行执行，也就是说，先被defer的语句最后被执行，最后被defer的语句，最先被执行。
func f7() {
	fmt.Println("执行函数f7开始")
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println("执行函数f7结束")
}

// 函数中return语句在底层并不是原子操作，它分为给返回值赋值和RET指令两步。而defer语句执行的时机就在返回值赋值操作后，RET指令执行前。
// 第一步赋值
// defer
// 第二步真正的return返回
// 函数中存在defer，那么defer执行的时机是在第一步和第二步之间
func f8() int { // 返回值为 5
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f9() (x int) { // 返回值为 6 注意这里的返回值是命名变量 x
	defer func() {
		x++
	}()
	return 5
}

func f10() (y int) { // 返回值为 5
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f11() (x int) { // 返回值是 5
	defer func(x int) {
		x++
	}(x)
	return 5
}

func main() {
	fmt.Printf("1 + 2 = %d\n", Plus(1, 2))
	f7()

	fmt.Println("defer")

	fmt.Println(f8())  // 5
	fmt.Println(f9())  // 6
	fmt.Println(f10()) // 5
	fmt.Println(f11()) // 5
}
