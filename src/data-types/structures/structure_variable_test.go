package structures

import (
	"testing"
)

type PersonStruct struct {
	name string
	age  int
}

func TestStructureVariables(t *testing.T) {
	// 1. key-value 形式初始化结构体
	var p1 PersonStruct
	p1.name = "curder"
	p1.age = 29

	t.Log(p1)

	// 2. 使用值列表形式初始化结构体
	var p2 = PersonStruct{
		name: "curder",
		age:  29,
	}
	t.Log(p2)

	// 3. 使用构造函数声明结构体变量
	var p3 = newPersonStruct("curder", 29)
	t.Log(p3)
}

func newPersonStruct(name string, age int) PersonStruct {
	return PersonStruct{
		name: name,
		age:  age,
	}
}
