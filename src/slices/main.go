// Slice 是 Go 中一个关键的数据类型，是一个比数组更加强大的序列接口 http://blog.golang.org/2011/01/go-slices-usage-and-internals.html
package main

import "fmt"

func main() {
	// 不像数组，slice 的类型仅有它所包含的元素决定（不像数组中还需要元素的个数）。要创建一个长度非零的空slice，需要使用内建的方法 `make`。
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
	// 其中一个是内建的 `append`，它返回一个包含了一个或者多个新值的 slice。注意接受返回值由 append 返回的新的 slice 值。
	s8 := []string{"beijing", "shanghai", "guangzhou"}
	s8 = append(s8, "shenzhen")
	s8 = append(s8, "chengdu", "hangzhou")
	fmt.Printf("切片 append 追加操作，其中值为：%v，长度为%v，容量为%v \n", s8, len(s8), cap(s8))

	// 切片追加切片
	s9 := []string{"suzhou", "wuhan", "xian"}
	s8 = append(s8, s9...)
	fmt.Printf("切片 append 追加切片操作，其中值为：%v，长度为%v，容量为%v \n", s8, len(s8), cap(s8))

	// Slice 也可以被 `copy`。这里我们创建一个空的和 `s` 有相同长度的 slice `c`，并且将 `s` 复制给 `c`。
	s10 := make([]string, len(s))
	copy(s10, s)
	fmt.Println("切片的拷贝:", s10)

	// Slice 支持通过 `slice[low:high]` 语法进行“切片”操作。例如，这里得到一个包含元素 `s[2]`, `s[3]`, `s[4]` 的 slice。
	l := s[4:5]
	fmt.Println(l)

	// slice 从 `s[0]` 到（但是不包含）`s[5]`。
	l = s[:5]
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
}