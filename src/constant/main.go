package main

import "fmt"
import "math"

// 相对于变量，常量是恒定不变的值，用于定义程序运行期间不会改变的那些值。
// 常量的声明和变量声明非常类似，只是把var换成了const，常量在定义的时候必须赋值。

const s string = "constant"

func main() {
	fmt.Println("常量s的值：", s)

	// const可以出现在任何允许var出现的地方
	const n = 500000000

	// 常数表达式可以执行任意精度的运算
	const fl = 3e20 / n
	fmt.Println("常量fl的值：", fl)

	// 数值型常量是没有确定的类型的，直到它们被给定了一个类型，比如说一次显示的类型转化。
	fmt.Println("对常量fl显示的转换成int64：", int64(fl))

	// 当上下文需要时，一个数可以被给定一个类型，比如变量赋值或者函数调用。举个例子，这里的 `math.Sin` 函数需要一个 `float64` 的参数。
	fmt.Println("对常量取Sin值：", math.Sin(n))

	// 多个常量也可以一起声明
	const (
		pi = 3.1415926
		n0 = 1
	)

	fmt.Printf("pi: %f \n", pi)
	fmt.Printf("n0: %d \n", n0)

	// const同时声明多个常量时，如果省略了值则表示和上面一行的值相同。
	const (
		n1 = 100
		n2
		n3
	)
	fmt.Printf("n1: %d \n", n1)
	fmt.Printf("n2: %d \n", n2)
	fmt.Printf("n3: %d \n", n3)

	// iota
	// iota 是go语言的常量计数器，只能在常量的表达式中使用。

	// iota在const关键字出现时将被重置为0。
	// const中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)。
	// 使用iota能简化定义，在定义枚举时很有用。
	const (
		n4 = iota // 0
		n5        // 1
		n6        // 2
	)
	fmt.Printf("n4: %d \n", n4)
	fmt.Printf("n5: %d \n", n5)
	fmt.Printf("n6: %d \n", n6)

	// 使用 _ 忽略某些值
	const (
		n7 = iota // 0
		n8        // 1
		_
		n10 // 3
	)
	fmt.Printf("n7: %d \n", n7)
	fmt.Printf("n8: %d \n", n8)
	fmt.Printf("n10: %d \n", n10)

	// iota 声明时中间插值
	const (
		n11 = iota // 0
		n12 = 100  // 100
		n13 = iota // 2
		n14        // 3
	)
	const n15 = iota // 0
	fmt.Printf("n11: %d \n", n11)
	fmt.Printf("n12: %d \n", n12)
	fmt.Printf("n13: %d \n", n13)
	fmt.Printf("n14: %d \n", n14)
	fmt.Printf("n15: %d \n", n15)

	// 使用 iota 定义字节数
	const (
		_  = iota
		KB = 1 << (10 * iota)
		MB = 1 << (10 * iota)
		GB = 1 << (10 * iota)
		TB = 1 << (10 * iota)
		PB = 1 << (10 * iota)
	)
	fmt.Printf("KB: %d \n", KB)
	fmt.Printf("MB: %d \n", MB)
	fmt.Printf("GB: %d \n", GB)
	fmt.Printf("TB: %d \n", TB)
	fmt.Printf("PB: %d \n", PB)

	// 多个 iota 定义在一行
	const (
		a, b = iota + 1, iota + 2 // 1, 2
		c, d                      // 2, 3
		e, f                      // 3, 4
	)
	fmt.Printf("a: %d \n", a)
	fmt.Printf("b: %d \n", b)
	fmt.Printf("c: %d \n", c)
	fmt.Printf("d: %d \n", d)
	fmt.Printf("e: %d \n", e)
	fmt.Printf("f: %d \n", f)
}
