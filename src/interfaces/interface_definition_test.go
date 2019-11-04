package main

import (
	"testing"
)

// 接口是一种类型，定义方式如下：

//type 接口名 interface {
//	方法名1(参数1, 参数2, 参数3...)(返回值1, 返回值2, 返回值3...)
//	方法名2(参数1, 参数2, 参数3...)(返回值1, 返回值2, 返回值3...)
//}

// 定义接口
type Speaker interface {
	speak() string
}

type Dog struct{}
type Cat struct{}
type Person struct{}

// 一个结构体实现了接口中规定的所有的方法，那么这个结构体就实现了这个接口，可以称之为接口类型的结构体。
func (d Dog) speak() string {
	return "wang wang wang~"
}

func (c Cat) speak() string {
	return "miao miao miao~"
}

func (p Person) speak() string {
	return "ying ying ying~"
}

func speak(speaker Speaker) string {
	return speaker.speak()
}

func TestInterfaceDefinition(t *testing.T) {
	var dog Dog
	var cat Cat
	var person Person
	// 使用接口前这样调用
	t.Logf("\n狗声：%s 猫叫：%s 人声：%s", dog.speak(), cat.speak(), person.speak())

	// 使用接口后这样调用
	t.Logf("\n狗声：%s 猫叫：%s 人声：%s", speak(dog), speak(cat), speak(person))
}
