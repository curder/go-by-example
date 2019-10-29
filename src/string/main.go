package main

import (
	"fmt"
	"strings"
)

func main() {
	// 当需要转义 \ 的时候，可以使用 \\ 转义 \
	s1 := "\\Volumes\\codes\\Go\\src\\github.com\\curder\\go-by-example"
	fmt.Printf("s1: %s \n", s1)

	// 定义多行字符串
	s2 := `
Volumes
	codes
		go
`
	fmt.Printf("s2: %s \n", s2)

	s3 := `\Volumes\codes\Go\src\github.com\curder\go-by-example` // 原样输出字符串
	fmt.Printf("s3: %s \n", s3)

	// 字符串常用操作，计算字符串长度|拼接字符串|分割字符串|判断是否保存某个子字符串|前后缀判断|子字符串位置|join操作

	// 计算字符串长度
	fmt.Printf("字符串s3的长度是：%s \n", len(s3))

	// 拼接字符串 使用 +
	hello := "hello"
	world := "world"
	fmt.Printf(hello + " " + world + "\n")
	// 字符串拼接 使用 Sprintf
	helloWorld := fmt.Sprintf("%s %s \n", hello, world)
	fmt.Printf("helloWorld: %s \n", helloWorld)

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

	// 修改字符串
	s5 := "big"
	byteS5 := []byte(s5)
	byteS5[0] = 'd'
	fmt.Printf("s5: %s\n", byteS5)

	s6 := "白色"
	runeS6 := []rune(s6)
	runeS6[0] = '绿'
	fmt.Printf("s6: %s\n", string(runeS6))

	// 分割字符串
	s7 := "hello world"
	s8 := strings.Split(s7, " ")
	fmt.Println(s8)

	// join操作
	fmt.Println(strings.Join(s8, "|"))

	// 判断是否保存某个子字符串
	fmt.Printf("%v\n", strings.Contains("hello world", "hello"))

	// 前后缀判断
	fmt.Printf("hello world 包含前缀 hello吗？ %v\n", strings.HasPrefix("hello world", "hello"))
	fmt.Printf("hello world 包含后缀 world吗？ %v\n", strings.HasSuffix("hello world", "world"))

	// 查找子字符串位置
	fmt.Printf("hello world 中第一个o出现的索引位置是：%d\n", strings.Index("hello world", "o"))
	fmt.Printf("hello world 中第一个a出现的索引位置是：%d\n", strings.Index("hello world", "a")) // 不存在子字符串默认返回 -1
}
