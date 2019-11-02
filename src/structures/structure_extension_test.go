package structures

import (
	"fmt"
	"testing"
)

type Animal struct {
	name string
}

// Move 移动方法
func (a Animal) Move() string {
	return fmt.Sprintf("%v 在动", a.name)
}

type Dog struct {
	feet uint8
	Animal // 匿名嵌套结构体，在结构体 Animal 中定义的方法 Dog 结构体提也拥有了
}

// Speak
func (d Dog) Speak() string {
	return fmt.Sprintf("%v is wang wang wang~", d.name)
}

// 结构体模拟"类的继承"
func TestMockClassExtensionByStructure(t *testing.T) {
	a1 := Dog{
		feet:   4,
		Animal: Animal{name: "kitt"},
	}

	t.Logf("%#v", a1)
	t.Log(a1.Move())
	t.Log(a1.Speak())
}
