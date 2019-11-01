package structures

import "testing"

// 匿名结构体

func TestAnonymousStructure(t *testing.T) {
	var s1 struct {
		name string
		age  int
	}

	s1.name = "curder"
	s1.age = 28

	t.Logf("\n匿名结构体的类型：%T\n姓名：%s\n年龄：%d\n", s1, s1.name, s1.age)
}
