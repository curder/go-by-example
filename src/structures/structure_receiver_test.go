package structures

import (
	"testing"
)

// TestStructureReceiver 结构体接收者
func TestStructureReceiver(t *testing.T) {
	var person = person{
		name: "curder",
		age:  29,
	}

	t.Log(person)

	person.pointIncr() // 传递的是指针，可以修改结构体变量
	t.Log(person)

	person.valueIncr() // 传递的是值拷贝，无法修改结构体变量的值
	t.Log(person)
}

// person 结构体
type person struct {
	name string
	age  int
}

// 指针接收者
func (p *person) valueIncr() {
	p.age++
}

// 值接收者
func (p person) pointIncr() {
	p.age++
}
