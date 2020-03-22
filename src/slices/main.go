package main

import (
	"fmt"
	"sort"
)

// Slice 是 Go 中一个关键的数据类型，是一个比数组更加强大的序列接口
// 不像数组，Slice 的类型仅有它所包含的元素决定（不像数组中还需要元素的个数）。
// 要创建一个长度非零的空Slice，需要使用内建的方法 `make`。
// http://blog.golang.org/2011/01/go-slices-usage-and-internals.html

func main() {
	// 切片的定义
	var s1 []int
	var s2 []string
	fmt.Printf("切片s1是否为空：%v\n", s1 == nil) // true
	fmt.Printf("切片s2是否为空：%v\n", s2 == nil) // true

	// 创建一个长度为3的 `string` 类型 slice（初始化为零值）。
	s := make([]string, 3, 5) // 创建一个3个长度且为字符串类型的切片，容量为5
	fmt.Printf("切片s的值为:%v，长度为%v，容量为%v\n", s, len(s), cap(s))

	// 切片的初始化
	s1 = []int{1, 2, 3}
	s2 = []string{"beijing", "shanghai", "guangzhou", "shenzhen"}
	fmt.Printf("切片s1的值为:%v\n", s1) // false
	fmt.Printf("切片s1是否为空：%v\n", s1 == nil)
	fmt.Printf("切片s2的值为:%v\n", s2) // false
	fmt.Printf("切片s2是否为空：%v\n", s2 == nil)

	// 切片属于引用类型，真正的数据保存在底层数组中。
	ss := []int{1, 2, 3}
	sss := ss
	fmt.Printf("切片ss修改前的值为：%v，切片sss修改前的值为：%v \n", ss, sss)
	ss[0] = 11
	fmt.Printf("切片ss修改后的值为：%v，切片sss修改后的值为：%v \n", ss, sss)

	// 类似设置和获取数组值
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("设置切片s的值:", s)
	fmt.Println("获取切片单元值:", s[2])

	// `len` 返回 slice 的长度
	fmt.Println("长度:", len(s))
	fmt.Printf("判断切片是否为空使用len函数，不能使用切片 == nil判断 \n")

	// `cap` 返回切片的容量，是指底层数组的容量，从第一个元素到最后一个元素的数量
	fmt.Printf("容量：%v\n", cap(s))

	// 由数组得到切片
	a1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s3 := a1[0:5] // 基于一个数组获得切片，左闭右开（左边包含，右边不包含）
	fmt.Printf("切片s3从数组a1获取的值为：%v \n", s3)
	s4 := a1[:4] // => [0:4]
	s5 := a1[2:] // => [2:len(a1)]
	s6 := a1[:]  // => [0:len(a1)]
	fmt.Printf("切片s4从数组a1最左边到第四个索引获取的值为：%v \n", s4)
	fmt.Printf("切片s5从数组a1的第2个索引到最后一个索引获取的值为：%v \n", s5)
	fmt.Printf("切片s6从数组a1的第0个索引到最后一个索引获取的值为：%v \n", s6)
	fmt.Printf("切片s4的长度和容量分别是：%v, %v \n", len(s4), cap(s4))

	// 切片的遍历
	s7 := []int{1, 2, 3}
	for index := 0; index < len(s7); index++ {
		fmt.Printf("for循环：遍历切片第%v个索引的值为：%v \n", index, s7[index])
	}

	for index, value := range []int{1, 2, 3} {
		fmt.Printf("for rang遍历：遍历切片第%v个索引的值为：%v \n", index, value)
	}

	// 作为基本操作的补充，slice 支持比数组更多的操作。

	// Slice 也可以被 `copy`。创建一个空的和 `s` 有相同长度的切片 `c`，并且将 `s` 复制给 `c`。
	s10 := []string{"a", "b", "c"}
	s11 := make([]string, len(s10))
	copy(s11, s10)
	fmt.Println("切片的拷贝，s11的值:", s11)

	// Go语言中并没有删除切片元素的专用方法，可以使用切片本身的特性来删除元素。
	// 例如：从切片中删除元素索引为2的元素
	s12 := []int{1,2,3,4,5,6,7,8,9}
	s12 = append(s12[:2], s12[3:]...)
	fmt.Printf("切片中删除索引为2的值后，新切片s12的值为：%v", s12) // 从切片a中删除索引为index的元素，操作方法是 a = append(a[:index], a[index+1:]...)

	// Slice 支持通过 `slice[low:high]` 语法进行“切片”操作。例如，这里得到一个包含元素 `s[2]`, `s[3]`, `s[4]` 的 slice。
	s13 := s[4:5]
	fmt.Printf("切片s13的值为：%v",s13)

	// slice 从 `s[0]` 到（但是不包含）`s[5]`。
	l := s[:5]
	fmt.Println("sl2:", l)

	// slice 从（包含）`s[2]` 到 slice 的后一个值。
	l = s[2:]
	fmt.Println("sl3:", l)

	// 可以在一行代码中申明并初始化一个 slice 变量。
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// Slice 可以组成多维数据结构。内部的 slice 长度可以不同，这和多位数组不同。
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	// 切片的排序
	s20 := []int{1,9,5,6,7,3,2,8,4}
	sort.Ints(s20)
	fmt.Printf("对切片s20进行排序后的切片s20的值：%v", s20)

}
