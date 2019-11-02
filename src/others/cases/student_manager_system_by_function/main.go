package main

import (
	"fmt"
	"os"
)

var Class = make(map[int]*Student)

type Student struct {
	id   int
	name string
}

// Index 列表
func Index() {
	for id, student := range Class {
		fmt.Printf("学员编号：%d 学员姓名：%s\n", id, student.name)
	}
}

// Store 新增
func Store() {
	var (
		id   int
		name string
	)
	// 获取用户输入
	fmt.Print("请输入学员编号：")
	fmt.Scanln(&id)
	fmt.Print("请输入学员姓名：")
	fmt.Scanln(&name)
	Class[id] = &Student{id: id, name: name}
}

// Delete 删除
func Delete() {
	var id int
	// 根据用户输入用的ID删除对应的map数据
	fmt.Print("请输入要删除的学员编号：")
	fmt.Scanln(&id)

	delete(Class, id)

	Index()
}

func main() {
	for {
		fmt.Print(`
操作列表如下：
1. 列表
2. 新增
3. 删除
4. 退出操作
请输入您要操作的动作：`)
		var input int
		fmt.Scanln(&input)
		switch input {
		case 1:
			fmt.Println("列表")
			Index()
		case 2:
			fmt.Println("新增")
			Store()
		case 3:
			fmt.Println("删除")
			Delete()
		case 4:
			os.Exit(1)
		default:
			fmt.Println("输入有误，请重新输入!")
		}

	}
}
