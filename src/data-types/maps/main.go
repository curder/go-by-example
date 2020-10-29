package main

import "fmt"

// Go 语言中提供的映射关系容器为 map，其内部使用散列表（hash）实现。
// map 是一种无序的基于 key-value 的数据结构。
// Go 语言中的 map 是引用类型，必须初始化才能使用。

func main() {
	// 创建一个空 map，需要使用内建的 `make`: `make(map[key-type]val-type)`
	m1 := make(map[string]int)
	fmt.Printf("声明一个空map值为：%v, 类型：%T\n", m1, m1)

	// 设置值
	m1["key1"] = 7
	m1["key2"] = 13
	m1["key3"] = 22

	fmt.Printf("赋值后的新map的值为:%v\n", m1)

	// 获取值
	fmt.Printf("map中键key2对应的值: %v\n", m1["key2"])

	// 获取 map 的长度
	fmt.Printf("map的长度为: %d\n", len(m1))

	// 删除map中对应的key->value
	delete(m1, "key1")
	fmt.Printf("删除后的map的值为:%v\n", m1)
	delete(m1, "keyNotExists") // 删除一个不存在的key

	// 访问一个不存在key的map值，返回的将会是一个零值
	fmt.Printf("访问一个map 中不存在的value3键，得到:%v\n", m1["k3"])

	// 判断map中key是否存在
	_, keyExists := m1["keyExists"]
	fmt.Printf("获取map中一个不存在的key的值为:%v\n", keyExists)

	// map的遍历
	for key, value := range m1 {
		fmt.Printf("遍历获取map的键值对，其中key: %v, value: %v \n", key, value)
	}

	// 定义并初始化一个非空的map
	m2 := map[string]int{"foo": 1, "bar": 2}
	fmt.Printf("\n获取m2的值为:%v\n\n", m2)

	// 值为 map 类型的切片
	var mapSlice = make([]map[string]string, 2)
	for index, value := range mapSlice {
		fmt.Printf("mapSlice的索引: %d 值: %v\n", index, value)
	}
	fmt.Println("初始化后的mapSlice")
	// 对切片中的 map 元素进行初始化
	mapSlice[0] = make(map[string]string, 10)
	mapSlice[0]["name"] = "curder"
	mapSlice[0]["password"] = "111111"
	mapSlice[0]["address"] = "beijing"
	for index, value := range mapSlice {
		fmt.Printf("mapSlice的索引:%d 值: %v\n", index, value)
	}

	// 值为切片类型的 map
	var sliceMap = make(map[string][]string, 2)
	fmt.Println(sliceMap)
	fmt.Println("初始化之后的sliceMap")
	key := "China"
	value, ok := sliceMap[key]
	if !ok {
		value = make([]string, 0, 2)
	}
	value = append(value, "beijing", "shanghai")
	sliceMap[key] = value
	fmt.Println(sliceMap)
}
