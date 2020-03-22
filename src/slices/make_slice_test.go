package main

import (
	"fmt"
	"testing"
)

// 使用make创建切片，允许指定切片的长度和容量
func TestMakeSlice(t *testing.T) {
	// 创建一个长度为5，容量是10的 `int` 类型 slice（初始化为零值）。
	s1 := make([]int, 5, 10)
	fmt.Printf("切片s1的值为：%v，长度为：%d，容量为：%d\n", s1, len(s1), cap(s1))

	// 创建一个长度为0，容量是10的 `string` 类型 slice（初始化为零值）。
	s2 := make([]string, 0, 10)
	fmt.Printf("切片s2的值为：%v，长度为：%d，容量为：%d\n", s2, len(s2), cap(s2))

	s3 := make([]string, 1, 10)

	fmt.Printf("切片s3的值为：%v，长度为：%d，容量为：%d\n", s3, len(s3), cap(s3))

}
