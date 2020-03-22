package main

import (
	"fmt"
	"testing"
)

// 创建一个空的和 `s` 有相同长度的切片 `c`，并且将 `s` 复制给 `c`。
func TestCopySlice(t *testing.T) {
	s1 := []int{1, 2, 3}
	s2 := s1
	s3 := make([]int, 3)
	copy(s3, s1)

	fmt.Println("切片s1的值为：", s1)
	fmt.Println("赋值的切片s2的值为:", s2)
	fmt.Println("切片s3从切片拷贝而来，s3的值:", s3)

	fmt.Println()

	s1[0] = 100
	fmt.Println("修改切片s1索引为0后值为：", s1)
	fmt.Println("切片s2的值也对应的会被修改，结果为:", s2)
	fmt.Println("切片s3从切片拷贝而来，值不会随之修改，s3的值:", s3)
}
