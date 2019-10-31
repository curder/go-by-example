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

// 函数的作用域
// 1. 如果在函数中找到了变量的定义，则使用函数内的变量；否则使用全局的变量；如果全局变量都没有，则报错。
// 2. 函数体内定义的变量，无法在函数体外使用。
// 3. if条件判断、for循环、switch语句定义的变量，只能在语句块中使用。
var v1 = 1

func f12() {
	// v1 := 111
	fmt.Println(v1)
}

// 函数类型
func f13() {
	fmt.Println("Hello World")
}
func f14() int {
	return 1
}

// 函数作为参数类型
func f15(x func() int) {
	var ret = x()
	fmt.Println(ret)
}

// 函数作为返回值
func f16(a, b int) int {
	return a + b
}
func f17(x func() int) func(int, int) int {
	return f16
}

func main() {
	fmt.Printf("1 + 2 = %d\n", Plus(1, 2))
	f7()

	fmt.Println("defer")

	fmt.Println(f8())  // 5
	fmt.Println(f9())  // 6
	fmt.Println(f10()) // 5
	fmt.Println(f11()) // 5
	f12()              // 1

	v13 := f13 // 将函数赋值给变量，函数变量
	fmt.Printf("f13的类型是：%T, f14的类型是：%T \n", v13, f14)

	f15(f14) // 传递函数类型参数

	fmt.Printf("函数f17返回的是函数类型的值：%T\n", f17(f14))
}
