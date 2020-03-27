package other_package

import "fmt"

// 在一个包中引用另外一个包里的标识符（如变量、常量、类型、函数等）时，该标识符必须是对外可见的（public）。
// 在Go语言中只需要将标识符的首字母大写就可以让标识符对外可见。

var var1 = 100 // 首字母小写，外部包不可见，只能在当前包内使用

const Var2 = 1 // 首字母大写外部包可见，可在其他包中使用

type Var3 struct { // 首字母小写，外部包不可见，只能在当前包内使用
	name string
}

// 首字母大写，外部包可见，可在其他包中使用
func Sum(x ...int) int { // Sum 求和
	var sum int

	for _, value := range x {
		sum += value
	}
	return sum
}

func age() { // 首字母小写，外部包不可见，只能在当前包内使用
	var Age = 18 // 函数局部变量，外部包不可见，只能在当前函数内使用
	fmt.Println(Age)
}

// 结构体中的字段名和接口中的方法名如果首字母都是大写，外部包可以访问这些字段和方法。
type Student struct {
	Name  string // 可在包外访问的方法
	class string // 仅限包内访问的字段
}

type Payer interface {
	init() //仅限包内访问的方法
	Pay()  //可在包外访问的方法
}
