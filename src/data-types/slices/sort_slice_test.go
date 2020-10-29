package main

import (
	"fmt"
	"sort"
	"testing"
)

// 切片的排序
func TestSortSlice(t *testing.T) {
	s := []int{100, 2, 9, 5, 6, 7, 3, 2, 8, 4, 111}
	sort.Ints(s) // 对整形排序
	fmt.Printf("对切片s进行排序后的值：%v\n", s)

	s2 := []string{"A", "D", "C", "B", "G", "F", "E"}
	sort.Strings(s2) // 对字符串排序
	fmt.Printf("对切片s2进行排序后的值：%v\n", s2)
}
