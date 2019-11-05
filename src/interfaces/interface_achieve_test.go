package main

import (
	"fmt"
	"testing"
)

// 接口的实现

type Animal interface {
	move() string
	eat(string) string
}

type MiniCat struct {
	name string
	feet int
}

func (m MiniCat) move() string {
	return "猫步~"
}

func (m MiniCat) eat(s string) string {
	return fmt.Sprintf("吃%s~", s)
}

type Chicken struct {
	feet int
}

func (c Chicken) move() string {
	return "鸡动~"
}

func (c Chicken) eat(s string) string { // 实现方法的传参和返回值都需要一致
	return fmt.Sprintf("%s~", s)
}

func TestInterfaceAchieve(t *testing.T) {
	var a1 Animal           // 接口类型变量
	var blueCate = MiniCat{ // 定义一个 MiniCat 类型的变量
		name: "蓝猫",
		feet: 4,
	}

	var kfc = Chicken{
		feet: 2,
	}
	a1 = blueCate
	t.Logf("\n%#v\n猫%s", a1, a1.eat("鱼🐟"))

	a1 = kfc

	t.Logf("\n%#v\nkfc%s", a1, a1.eat("鸡🐔"))

}
