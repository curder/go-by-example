// 命令行参数是指定程序运行参数的一个常见方式。例如，go run hello.go，程序 go 使用了 run 和 hello.go 两个参数
package main

import (
	"fmt"
	"os"
)

func main() {

	// `os.Args` 提供原始命令行参数访问功能。注意，切片中的第一个参数是该程序的路径，并且 `os.Args[1:]`保存所有程序的的参数
	argsWithProgram := os.Args
	argsWithoutProgram := os.Args[1:]

	// 可以使用标准的索引位置方式取得单个参数的值
	arg := os.Args[3]

	fmt.Println(argsWithProgram)
	fmt.Println(argsWithoutProgram)
	fmt.Println(arg)
}

/*
go build command-line-arguments.go

./command-line-arguments a b c d
*/
