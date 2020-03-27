package main

// 包的导入

// 要在代码中引用其他包的内容，需要使用import关键字导入使用的包。
import "testing" // 导入系统包

// 自定义包的导入
// 注意事项：
// import导入语句通常放在文件开头包声明语句的下面。
// 导入的包名需要使用双引号包裹起来。
// 包名是从 `$GOPATH/src/` 后开始计算的，使用 `/` 进行路径分隔。
// Go语言中禁止循环导入包。

// 在导入包名的时候，我们还可以为导入的包设置别名。
// 通常用于导入的包名太长或者导入的包名冲突的情况。
// import 别名 "包的路径"
import (
	tools "github.com/curder/go-by-example/src/packages/other-package"
)

func TestPackageImport(t *testing.T) {
	t.Log(tools.Var2)
}
