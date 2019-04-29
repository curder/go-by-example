// 有时候想使用和集合的自然排序不同的方法对集合进行排序。例如，想按照字母的长度而不是首字母顺序对字符串排序
package main

import (
	"fmt"
	"sort"
)

// 为了在 Go 中使用自定义函数进行排序，需要一个对应的类型，这里创建一个为内置 `[]string` 类型的别名的 `ByLength` 类型
type ByLength []string

// 在类型中实现了 `sort.Interface` 的 `Len`，`Less` 和 `Swap` 方法，这样可以使用 `sort` 包的通用 `Sort` 方法了，
// `Len` 和 `Swap` 通常在各个类型中都差不多，`Less` 将控制实际的自定义排序逻辑
// 想按字符串长度增加的顺序来排序，所以这里使用了 `len(s[i])` 和 `len(s[j])`
func (s ByLength) Len() int {
	return len(s)
}
func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

// 一切都准备好了，现在可以通过将原始的 `fruits` 切片转型成 `ByLength` 来实现自定义排序。然后对这个转型的切片使用 `sort.Sort` 方法
func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(ByLength(fruits))
	fmt.Println(fruits)
}
