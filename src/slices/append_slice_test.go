package main

import (
	"fmt"
	"testing"
)

// `append`是内建的slice操作，它返回一个包含了一个或者多个新值的 slice。
// 注意：必须使用原切片变量接收返回值由 append 操作返回的新的切片值。
func TestAppendSlice(t *testing.T) {
	s1 := []string{"beijing", "shanghai", "guangzhou"}
	fmt.Printf("切片s1的值为:%v，长度为%v，容量为%v\n", s1, len(s1), cap(s1))

	s1 = append(s1, "shenzhen")
	fmt.Printf("s1切片 append 追加操作后的值为:%v，长度为%v，容量为%v\n", s1, len(s1), cap(s1))

	// 切片追加切片
	newSlice := []string{"wuhan", "xian", "suzhou"}
	s1 = append(s1, newSlice...)
	fmt.Printf("s1切片 append 追加操作后的值为:%v，长度为%v，容量为%v\n", s1, len(s1), cap(s1))
}
