package structures

import (
	"testing"
)

// 结构体指针

type NewPerson struct {
	name string
	age  int
}

func f1(p NewPerson) { // 传参数是值拷贝
	p.name = "luo" // 修改的是函数内部拷贝的副本 NewPerson
}

func f2(p *NewPerson) { // 结构体指针参数
	// (*p).name = "curder" // 根据内存地址找到原变量，赋值给原变量
	p.name = "curder" // p 根据指针找对应的值，而无需通过对应的指针找到对应的变量，语法糖
}

func TestStructurePointer(t *testing.T) {
	var v1 NewPerson
	v1.name = "guest"
	f1(v1) // 调用 f1 函数不会修改 v1 的 name 值
	t.Logf("\n调用f1函数后v1的值为：%#v \n", v1)

	f2(&v1)
	t.Logf("\n调用f2函数后v1的值为：%#v \n", v1)

	// 结构体指针
	var v2 = new(NewPerson)
	(*v2).name = "curder"
	v2.age = 28
	t.Logf("\nv2结构体指针%#v\n", v2) // v2 为结构体指针

	// 使用 key -> value 类型初始化结构体指针
	var v3 = &NewPerson{
		name: "luo",
		age:  28,
	} // 声明结构体变量且初始化
	t.Logf("\nv3结构体指针%#v\n", v3)
}
