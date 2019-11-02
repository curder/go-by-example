package structures

import (
	"testing"
)

// 结构体
// Go语言中的基础数据类型可以表示一些事物的基本属性，当需要表达一个事物的全部或部分属性时，此时再用单一的基本数据类型明显就无法满足需求。
// Go语言提供了一种自定义数据类型，可以封装多个基本数据类型，这种数据类型叫结构体，英文名称struct。
// 也就是我们可以通过struct来定义自己的类型了。
// Go语言中通过struct来实现面向对象。

//   type 类型名 struct {
//      字段名 字段类型
//      字段名 字段类型
//      …
//   }

type Person struct {
	name   string
	gender string
	age    int
	hobby  []string
}

// 类型名：标识自定义结构体的名称，在同一个包内不能重复。
// 字段名：表示结构体字段名。结构体中的字段名必须唯一。
// 字段类型：表示结构体字段的具体类型。

func TestStructureDefinition(t *testing.T) {
	var v1 Person      // 声明 NewPerson 类型的变量 v1
	v1.name = "curder" // 赋值
	v1.age = 28
	v1.gender = "男"
	v1.hobby = []string{"coding", "music", "movies"}

	t.Logf("结构体v1的类型是：%T\n名字：%s \n年龄：%d\n性别：%s\n爱好：%#v\n", v1, v1.name, v1.age, v1.gender, v1.hobby) // 获取结构体值
}

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

// 匿名字段
// 适用于字段比较少的场景
func TestAnonymousFieldStructure(t *testing.T) {
	type Person struct {
		int
		string
	}

	p1 := Person{1, "curder"}

	t.Logf("\n匿名结构体p1的值为：%#v \n通过类型为key获取值，int的值为：%#v string的值为：%#v\n", p1, p1.int, p1.string)
}

// 嵌套结构体

func TestNestedStructure(t *testing.T) {

	type Address struct {
		province string
		city     string
	}

	type Person struct {
		name    string
		age     int
		Address // 匿名嵌套结构体 Address
	}

	type Company struct {
		name    string
		address Address
	}

	p1 := Person{
		name: "curder",
		age:  28,
		Address: Address{
			province: "jx",
			city:     "zs",
		},
	}

	t.Logf("\n结构体详细值：%#v\n", p1)
	t.Logf("\n省份：%v\n", p1.province) // 通过p1.province访问结构体内容是一个go语言的语法糖，需要保证是匿名嵌套结构体
}
