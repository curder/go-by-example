package main

import (
	"fmt"
	"strings"
)

// Go语言中的字符串以原生数据类型出现。
// 使用字符串就像使用其他原生数据类型（int、bool、float32、float64等）一样。


func main() {
	// 当需要转义 \ 的时候，可以使用 \\ 转义 \
	s1 := "\\Volumes\\codes\\Go\\src\\github.com\\curder\\go-by-example"
	fmt.Printf("字符串s1的值为: %s \n", s1)

	// 定义多行字符串
	s2 := `
Volumes
	codes
		go
	`
	fmt.Printf("字符串s2的值为: %s \n", s2)

	s3 := `\Volumes\codes\Go\src\github.com\curder\go-by-example` // 原样输出字符串
	fmt.Printf("字符串s3的值为: %s \n", s3)

	// 字符串常用操作，计算字符串长度|拼接字符串|分割字符串|判断是否保存某个子字符串|前后缀判断|子字符串位置|join操作

	// 计算字符串长度
	fmt.Printf("字符串s3的长度是：%d \n", len(s3))

	// 拼接字符串 使用 +
	hello := "hello"
	world := "world"
	fmt.Printf(hello + " " + world + "\n")
	// 字符串拼接 使用 Sprintf 返回字符串变量
	helloWorld := fmt.Sprintf("%s %s \n", hello, world)
	fmt.Printf("变量helloWorld的值为: %s \n", helloWorld)

	// 遍历字符串
	s4 := "hello你好"
	for i := 0; i < len(s4); i++ { // byte
		fmt.Printf("%v(%c) ", s4[i], s4[i])
	}
	fmt.Println()
	for _, r := range s4 { // rune
		fmt.Printf("%v(%c) ", r, r)
	}
	fmt.Println()

	// 修改字符串，字符串不能直接修改
	s5 := "big"
	byteS5 := []byte(s5) // 将字符串s5强制转换成byte类型的切片
	byteS5[0] = 'd'
	fmt.Printf("字符串变量s5修改后的值为: %s\n", byteS5)

	s6 := "白色"
	runeS6 := []rune(s6) // 将字符串s6强制转换成rune类型的切片
	runeS6[0] = '绿'
	fmt.Printf("字符串变量s6修改后的值为: %s\n", string(runeS6))

	// 分割字符串
	s7 := "hello world"
	s8 := strings.Split(s7, " ")
	fmt.Printf("使用空格分割s7字符串的值为：%v，类型为：%T\n" , s8, s8)

	// join操作
	fmt.Println("使用｜合并字符串：", strings.Join(s8, "｜"))

	// 判断是否保存某个子字符串
	fmt.Printf("变量s7是否包含'hello'子字符串：%v\n", strings.Contains(s7, "hello"))

	// 前后缀判断
	fmt.Printf("变量s7前缀是 hello 吗？ %v\n", strings.HasPrefix(s7, "hello"))
	fmt.Printf("变量s7后缀是 world 吗？ %v\n", strings.HasSuffix(s7, "world"))

	// 查找子字符串位置
	fmt.Printf("hello world 中第一个o出现的索引位置是：%d\n", strings.Index(s7,"o"))
	fmt.Printf("hello world 中第一个a出现的索引位置是：%d\n", strings.Index(s7, "a")) // 不存在子字符串默认返回 -1
}
