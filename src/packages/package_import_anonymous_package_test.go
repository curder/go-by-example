package main

import "testing"

import _ "github.com/curder/go-by-example/src/packages/anonymous"

// 导入匿名包
// 如果只希望导入包，而不使用包内部的数据时，可以使用匿名导入包。
// 匿名导入的包与其他方式导入的包一样都会被编译到可执行文件中。
func TestImportAnonymousPackage(t *testing.T) {
	t.Log("import anonymous package test")
}
