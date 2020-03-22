package main

import (
	"fmt"
	"testing"
)

// Go语言中并没有删除切片元素的专用方法，可以使用切片本身的特性来删除元素。
func TestDeleteSliceItem(t *testing.T) {
	// 从切片中删除元素索引为 2 的元素
	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 = append(s1[:2], s1[3:]...)
	// 从切片a中删除索引为index的元素，操作方法是 a = append(a[:index], a[index+1:]...)
	fmt.Printf("切片中删除索引为 2 的值后，新切片 s1 的值为：%v\n", s1)
}
