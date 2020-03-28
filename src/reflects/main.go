package main

import (
	"fmt"
	"reflect"
)

type myInt int64

func main() {
	var a *float32 // 指针
	var b myInt    // 自定义类型
	var c rune     // 类型别名
	reflectOf(a)   // type: kind:ptr
	reflectOf(b)   // type:myInt kind:int64
	reflectOf(c)   // type:int32 kind:int32

	type person struct {
		name string
		age  int
	}
	type book struct{ title string }
	var d = person{
		name: "Curder",
		age:  28,
	}
	var e = book{title: "Learning Go Language"}
	reflectOf(d) // type:person kind:struct
	reflectOf(e) // type:book kind:struct

	var f int64 = 100
	reflectSet(&f) // 修改值
	fmt.Println(f) // 200
}

func reflectOf(a interface{}) {
	v := reflect.TypeOf(a)
	fmt.Printf("name: %v，kind: %v\n", v.Name(), v.Kind()) // 变量类型
}

func reflectSet(a interface{}) {
	v := reflect.ValueOf(a)

	// 如果要修改值，需要传递指针，通过 .Elem() 获取指针对应的值
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200)
	}
}
