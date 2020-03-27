package main

import (
	"fmt"

	tools "github.com/curder/go-by-example/src/packages/other-package" // 引入自定义包
)

// 包（package）是多个Go源码的集合，是一种高级的代码复用方案
// Go语言为我们提供了很多内置包，如fmt、os、io等。

// 自定义包
// 可以根据自己的需要创建自己的包。一个包可以简单理解为一个存放 .go 文件的文件夹。
// 该文件夹下面的所有go文件都要在代码的第一行添加如下代码，声明该文件归属的包。

func main() {

	fmt.Println(tools.Sum(1, 2, 3, 4, 5, 6, 7, 8, 9))
}
