// Go 的 sort 包实现了内置和用户自定义数据类型的排序功能
package main

import (
	"fmt"
	"sort"
)

func main() {

	// 排序方法是正对内置数据类型的；
	// 注意排序是原地更新的，所以他会改变给定的序列并且不返回一个新值
	strings := []string{"c", "a", "b"}
	sort.Strings(strings)
	fmt.Println("Strings:", strings)

	// 一个 `int` 排序的例子
	integers := []int{7, 2, 4}
	sort.Ints(integers)
	fmt.Println("Integers:   ", integers)

	// 我们也可以使用 `sort` 来检查一个序列是不是已经是排好序的
	s := sort.IntsAreSorted(integers)
	fmt.Println("Sorted: ", s)
}
