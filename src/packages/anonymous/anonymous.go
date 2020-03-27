package anonymous

import "fmt"

// 在Go语言程序执行时导入包语句会自动触发包内部init()函数的调用。
// 需要注意的是： init()函数没有参数也没有返回值。
// init()函数在程序运行时自动被调用执行，不能在代码中主动调用它。

// Go语言包会从main包开始检查其导入的所有包，每个包中又可能导入了其他的包。
// Go编译器由此构建出一个树状的包引用关系，再根据引用顺序决定编译顺序，依次编译这些包的代码。
func init() {
	fmt.Println("the code run in init function.")
}
