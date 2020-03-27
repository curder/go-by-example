package main

import (
	"testing"
)

// 空接口
// 是指没有定义任何方法的接口，因此任何类型都实现了空接口。
// 空接口类型的变量可以存储任意类型的变量。
// interface{} // 空接口
func TestEmptyInterfaceForMap(t *testing.T) {
	//var v1 = map[interface{}]interface{} // error: https://forum.golangbridge.org/t/practical-example-on-empty-interface/5268/3
	v1 := make(map[string]interface{}, 10)
	v1["name"] = "curder"
	v1["age"] = 28
	v1["hobby"] = [...]string{"coding", "music", "study"}
	v1["married"] = false

	t.Logf("\nv1的值为：%#v", v1)
}

func TestEmptyInterfaceForFunction(t *testing.T) {
	// 空接口作为函数参数
	var show = func(a interface{}) {
		t.Logf("类型: %T 值: %v\n", a, a)
	}
	show(false)
	show(nil)

	v2 := make(map[string]interface{}, 10)
	show(v2)
}

// 类型断言
// 空接口可以存储任意类型的值，那我们如何获取其存储的具体数据呢？

func TestEmptyInterfaceTypeAssertion(t *testing.T) {
	var x interface{} // 定义一个空接口类型变量 a
	x = "hello world!"
	_, ok := x.(string)
	if ok {
		t.Log("\n传入的参数是字符串")
	} else {
		t.Log("\n类型断言失败!")
	}

	// 使用switch case判断类型
	switch x.(type) {
	case string:
		t.Log("\n传入的参数是字符串", x)
	case int:
		t.Log("\n传入的参数是整型", x)
	case bool:
		t.Log("\n传入的参数是布尔值", x)
	}
}
