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
	if _, error := fmt.Scanln(&id); error != nil {
		fmt.Print("错误的学员编号输入")
		return
	}
	fmt.Print("请输入学员姓名：")
	if _, error := fmt.Scanln(&name); error != nil {
		fmt.Print("错误的学员姓名输入")
		return
	}
	Class[id] = &Student{id: id, name: name}
}

func Update() {
	var (
		id   int
		name string
	)
	// 获取用户输入
	fmt.Print("请输入学员编号：")
	if _, error := fmt.Scanln(&id); error != nil {
		fmt.Print("错误的学员编号输入")
		return
	}

	// 判断学员编号是否存在
	if _, hasStudent := Class[id]; !hasStudent {
		fmt.Print("该编号学员不存在")
		return
	}

	fmt.Print("请输入学员姓名：")
	if _, error := fmt.Scanln(&name); error != nil {
		fmt.Print("错误的学员编号姓名")
		return
	}

	student := &Student{id: id, name: name}

	Class[id] = student

	Index()
}

// Delete 删除
func Delete() {
	var id int
	// 根据用户输入用的ID删除对应的map数据
	fmt.Print("请输入要删除的学员编号：")
	if _, error := fmt.Scanln(&id); error != nil {
		fmt.Print("错误的学员编号输入")
		return
	}

	if _, isSaved := Class[id]; !isSaved {
		fmt.Print("该编号学员不存在")
		return
	}

	delete(Class, id)

	Index()
}

func main() {
	for {
		fmt.Print(`
操作列表如下：
1. 列表
2. 新增
3. 编辑
4. 删除
5. 退出操作
请输入您要操作的动作：`)
		var input int
		if _, error := fmt.Scanln(&input); error != nil {
			fmt.Print("操作输入错误")
			return
		}
		switch input {
		case 1:
			fmt.Println("列表")
			Index()
		case 2:
			fmt.Println("新增")
			Store()
		case 3:
			fmt.Println("编辑")
			Update()
		case 4:
			fmt.Println("删除")
			Delete()
		case 5:
			os.Exit(1)
		default:
			fmt.Println("输入有误，请重新输入!")
		}

	}
}
