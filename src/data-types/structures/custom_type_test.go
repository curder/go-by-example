package structures

import (
	"testing"
)

// Go语言中没有“类”的概念，也不支持“类”的继承等面向对象的概念。
// Go语言中通过结构体的内嵌再配合接口比面向对象具有更高的扩展性和灵活性。

// 自定义类型
type myInt int

func TestCustomType(t *testing.T) {
	var n1 myInt = 100
	t.Logf("n1的值为：%v，类型为：%T \n", n1, n1)
}

func TestCustomTypeAddMethod(t *testing.T) {
	var age = myInt(18)
	age.incr()
	t.Log(age)
}

// 给自己定义的类型添加一个方法
func (m *myInt) incr() {
	*m++
}
