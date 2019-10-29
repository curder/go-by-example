package main

import "fmt"

func main() {
	// 数组是同一种数据类型存放元素的容器，在定义数组的时候必须定义元素类型和存储的长度。
	// 数组 `a` 来存放刚好 5 个 `int`。元素的类型和长度都是数组类型的一部分。
	var a [5]int
	fmt.Println("初始化数组a的值:", a)

	var a1 [2]bool
	var a2 [5]string
	fmt.Printf("a1: %T\na2: %T\n", a1, a2)

	// 数组的初始化
	// 数组默认是零值，对于 `int` 数组来说也就是 `0`；对于`bool`值为false，对于`string`默认值为""。
	a3 := [2]bool{true, true}
	fmt.Printf("a1赋值后: %v \n", a3)

	a4 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9} // 根据初始值自动推算数组长度
	fmt.Printf("a3: %v \n", a4)

	a5 := [...]int{1: 1, 5: 2, 6: 0} // 根据索引初始化数组
	fmt.Printf("a4: %v \n", a5)

	// 多维数组
	// [[1 2] [3 4] [5 6]]
	var a6 [3][2]int
	a6 = [3][2]int{{1, 2}, {3, 4}, {5, 6}}
	fmt.Println("二维数组 a6: ", a6)

	// 数组的存储类型是单一的，但是可以组合这些数据来构造多维的数据结构。
	var a7 [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			a7[i][j] = i + j
		}
	}
	fmt.Println("二维数组 a7: ", a7)

	// 数组的遍历

	// 遍历数组之根据索引遍历
	fmt.Println("*** for 遍历数组 ***")
	cities := [4]string{"beijing", "shanghai", "guangzhou", "shenzhen"}
	for i := 0; i <= len(cities)-1; i++ {
		fmt.Println(cities[i])
	}

	// for range
	fmt.Println("*** for range 遍历数组 ***")
	for _, city := range cities {
		fmt.Println(city)
	}

	// 遍历多维数组
	for _, value := range a6 {
		for key, item := range value {
			fmt.Printf("二维数组的键为：%v，值为: %v \n", key, item)
		}
	}

	// 可以使用 `array[index] = value` 语法来设置数组指定位置的值，或者用 `array[index]` 得到值。
	a[4] = 100
	fmt.Println("a的新值:", a)
	fmt.Println("获取a索引为4的值:", a[4])

	// 使用内置函数 `len` 返回数组的长度
	fmt.Println("数组长度:", len(a))

	// 使用这个语法在一行内初始化一个数组
	b := [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// 数组是值类型

	// 求数组[1, 3, 5, 7, 8]所有元素的和
	a8 := 0
	for _, value := range [...]int{1, 3, 5, 7, 8} {
		a8 += value
	}
	fmt.Printf("数组a8元素的和为：%d \n", a8)

	// 找出数组中和为指定值的两个元素的下标，比如从数组[1, 3, 5, 7, 8]中找出和为8的两个元素的下标分别为(0,3)和(1,2)。
	a9 := [...]int{1, 3, 5, 7, 8}
	for index, value := range a9 {
		for index2, value2 := range a9[index+1:] {
			if (value+value2 == 8) {
				fmt.Printf("%v \n", [...]int{index, index + index2 + 1})
			}
		}
	}

}
