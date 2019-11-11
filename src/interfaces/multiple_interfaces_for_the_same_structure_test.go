package main

import "testing"

// 同一个结构体可以实现多个接口

// 结构体嵌套结构体
type MultipleAnimal interface {
	Mover
	Eater
}

type Mover interface {
	move()
}

type Eater interface {
	eat()
}

// 定义结构体
type MultipleCat struct {
	name string
}

// 实现Mover接口的move方法
func (m MultipleCat) move()  {
}

// 实现Eater接口的eat方法
func (m MultipleCat) eat()  {
}

func TestMultipleInterfacesForTheSameStructure(t *testing.T)  {

}
