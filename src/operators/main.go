package main

import "fmt"

func main() {
	// 运算符

	// 算数运算符
	a := 10
	b := 2

	fmt.Printf("%d + %d = %d\n", a, b, a+b)  // 加法
	fmt.Printf("%d - %d = %d\n", a, b, a-b)  // 减法
	fmt.Printf("%d * %d = %d\n", a, b, a*b)  // 乘法
	fmt.Printf("%d / %d = %d\n", a, b, a/b)  // 除法
	fmt.Printf("%d %% %d = %d\n", a, b, a%b) // 取模

	// ++（自增）和--（自减）在Go语言中是单独的语句，并不是运算符
	a++ // 单独的语句，不能放在 = 的右边进行赋值运算，相当于 a = a + 1
	b-- // 单独的语句，不能放在 = 的右边进行赋值运算，相当于 b = b - 1

	// 关系运算符 返回布尔值
	fmt.Println(a == b)
	fmt.Println(a != b)
	fmt.Println(a >= b)
	fmt.Println(a <= b)
	fmt.Println(a > b)
	fmt.Println(a < b)

	// 逻辑运算符 返回布尔值
	// && 逻辑 AND 运算符。 如果两边的操作数都是 True，则为 true，否则为 false
	if age := 19; age > 18 && age < 50 {
		fmt.Println("上班年龄")
	} else {
		fmt.Println("不用上班")
	}

	// || 逻辑 OR 运算符。 如果两边的操作数有一个 true，则为 true，否则为 false
	if age := 19; age < 18 || age > 51 {
		fmt.Println("不用上班")
	} else {
		fmt.Println("上班年龄")
	}

	// ! 逻辑 NOT 运算符。 如果条件为 true，则为 false，否则为 true。
	isMarried := false
	fmt.Println(!isMarried)

	// 位运算符

	// 位运算符对整数在内存中的二进制位进行操作。
	// 5的二进制值为：101
	// 2的二进制值为： 10

	fmt.Println(5 & 2)  //   0 按位与；  参与运算的两数各对应的二进位相与。（两位均为1才为1）
	fmt.Println(5 | 2)  //   7 按位或；  参与运算的两数各对应的二进位相或。（两位有一个为1就为1）
	fmt.Println(5 ^ 2)  //   7 按位异或；参与运算的两数各对应的二进位相异或，当两对应的二进位相异时，结果为1。（两位不一样则为1）
	fmt.Println(5 << 2) // 20 按位左移；左移n位就是乘以2的n次方。 “a<<b”是把a的各二进位全部左移b位，高位丢弃，低位补0。相当于二进制101左移两位：10100 为 20
	fmt.Println(5 >> 2) //  1 按位右移；右移n位就是乘以2的n次方。相当于二进制101右移2位：1 为 1

	// 赋值运算符
	//  =  简单的赋值运算符，将一个表达式的值赋给一个左值
	// +=  相加后再赋值
	// -=  相减后再赋值
	// *=  相乘后再赋值
	// /=  相除后再赋值
	// %=  求余后再赋值
	// <<= 左移后赋值
	// >>= 右移后赋值
	// &=  按位与后赋值
	// |=  按位或后赋值
	// ^=  按位异或后赋值
}
