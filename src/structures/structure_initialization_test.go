package structures

import "testing"

// TestStructureInitialization 结构体的初始化
func TestStructureInitialization(t *testing.T) {
	// 声明结构体变量
	var s1 student     // 声明结构体变量
	s1.name = "curder" // 给变量赋值
	s1.age = 29
	t.Log("通过声明结构体变量", s1)
	// 键值对初始化
	var s2 = student{
		name: "curder",
		age:  29,
	}
	t.Log("通过键值对初始化", s2)

	// 值列表初始化
	var s3 = student{
		"curder",
		29,
	}
	t.Log("通过值列表初始化", s3)

	// 通过结构体构造函数
	var s4 = newStudent("curder", 29)
	t.Log("通过结构体构造函数", s4)
}

type student struct {
	name string
	age  int
}

// 自定义结构体构造函数
func newStudent(name string, age int) student {
	return student{name: name, age: age}
}
