package main

// Go变量声明
// 1. 指定变量类型，声明后若不赋值使用默认值，布尔类型默认false，整形和浮点型默认0，字符串默认为空。
// 2. go语言中如果在声明变量时，没有显示的指定变量类型，解析器会根据变量的值自动判断变量类型（类型推导）
// 3. 可以使用 := 的方式声明并赋值，但是变量名不应该是已经声明过的，否则会导致编译出错
// 4. 在 go 语言中申明的变量必须使用，否则编译器无法完成编译

// Go 多变量声明，允许在同一行声明多个变量
import "fmt"

// 声明全局变量
var ga = 1
var gb = "name"
var gc = 3
var (
	gd = 2
	ge = "name"
	gf = 4
)

func main() {
	// 声明变量且没有给出对应的初始值时，变量将会初始化为 零 值
	var a int
	fmt.Println(a)

	// 自定推断类型
	var b = true
	fmt.Println(b)

	// `:=` 语句是申明并初始化变量的简写，例如下例等同于 `var f string = "short"`。注意： := 不能使用在函数外。
	c := "short"
	fmt.Println(c)

	// 声明字符串
	var d string = "initial"
	fmt.Println(d)

	// 多变量声明整型
	var e, f int = 1, 2
	fmt.Println(e, f)

	// 声明多个不同类型的变量
	g, h, i := 10, "name", 11
	fmt.Println(g, h, i)

	// 使用 () 定义多个变量
	var (
		j = 1
		k = 2
		l = 3
	)
	fmt.Println(j, k, l)

	// 打印全局变量
	fmt.Println(ga, gb, gc, gd, ge, gf)

	// 匿名变量，使用 _ 忽略某个值
	var _, message = 404, "Not found!"
	fmt.Println(message)
}
