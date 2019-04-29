// Go 的结构体 是各个字段字段的类型的集合。这在组织数据时非常有用
package main

import "fmt"

// `person` 结构体包含了 `name` 和 `age` 两个字段
type person struct {
	name string
	age  int
}

func main() {
	// 使用语法创建了一个新的结构体元素。
	fmt.Println(person{"personName", 11})

	// 初始化结构体时指定字段名
	fmt.Println(person{age: 11, name: "personName"})

	// 省略的字段如果是整型将被初始化零值，如果是字符串会被初始化为空字符串
	fmt.Println(person{name: "personName"})

	// &符获取结构体的指针
	fmt.Println(&person{name: "personName"})

	// 使用点语法访问结构体字段
	person := person{name: "personName", age: 11}
	fmt.Println(person.name, person.age)

	// 也可以对结构体指针使用`.` - 指针会被自动解引用
	personPoint := &person
	fmt.Println(personPoint.age)

	// 结构体是可变的
	personPoint.age = 12
	fmt.Println(personPoint.age)
}
