// range 迭代各种各样的数据结构。让我们来看看如何在已经学过的数据结构上使用 `range`
package main

import (
	"fmt"
)

func main() {
	// 这里我们使用 `range` 来统计一个 slice 的各元素的值之和。数组也可以采用这种方法。
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// `range` 在数组和 slice 中都同样提供每个项的索引和值。上面我们不需要索引，所以我们使用 `_` 来忽略它。有时候需要这个索引。
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// `range` 在 map 中迭代键值对。
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// `range` 在字符串中迭代 unicode 编码。第一个返回值是 `go` 的起始字节位置，然后第二个是 `go` 自己。
	for i, c := range "中国" {
		fmt.Println(i, c)
	}
}
