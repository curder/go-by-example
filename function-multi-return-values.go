// Go 内建_多返回值_ 支持。这个特性在 Go 语言中经常被用到，例如用来同时返回一个函数的结果和错误信息。

package main

import "fmt"

// `(int, string)` 在这个函数中标志着这个函数返回 1 个 `int` 和 1 个 `string`
func vals() (int, string) {
	return 404, "Not Found!"
}

func main() {
	// 这里通过_多赋值_ 操作来使用这两个不同的返回值。
	code, message := vals()

	fmt.Println(code)
	fmt.Println(message)

	// 如果仅仅需要返回值的一部分的话，你可以使用空白定义符 `_`。
	_, messages := vals()
	fmt.Println(messages)
}
