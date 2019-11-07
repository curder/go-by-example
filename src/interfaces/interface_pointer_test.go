package main

import "testing"

// 使用指针接受者实现接口方法
// 使用值接收者实现的接口方法，结构体变量类型和指针类型变量都能存储
// 使用指针接收者只能存储指针类型变量

type PointerAnimal interface {
	speak() string
}

type PointerCat struct {
	name string
}
type PointerDog struct {
	name string
}

func (p PointerCat) speak() string {
	return "miao miao miao~"
}

func (p *PointerDog) speak() string {
	return "wang wang wang~"
}

func TestInterfacePointerTest(t *testing.T) {
	var animal PointerAnimal

	cat := PointerCat{
		name: "Kitty",
	}
	dog := PointerDog{
		name: "gig",
	}

	animal = &cat // 值接收者即能赋值指针也能将值赋值给接口
	t.Logf("\ncat结构体能够赋值给animal接口, %#v", animal)

	animal = &dog
	t.Logf("\ndog结构体能够赋值给animal接口, %#v", animal)
}
