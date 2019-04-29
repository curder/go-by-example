// for 和 range为基本的数据结构提供了迭代的功能。也可以使用这个语法来遍历从通道中取得的值
package main

import "fmt"

func main() {

	// 遍历在 `queue` 通道中的三个值
	queue := make(chan string, 3)
	queue <- "one"
	queue <- "two"
	queue <- "three"
	close(queue)

	// 这个 `range` 迭代从 `queue` 中得到的每个值。
	// 在前面 `close` 了这个通道，这个迭代会在接收完 3 个值之后结束，如果没有 `close` 它，将在这个循环中继续阻塞执行，等待接收第四个值
	for elem := range queue {
		fmt.Println(elem)
	}
}
