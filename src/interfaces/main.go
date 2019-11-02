// 接口是方法特征的命名集合，它是一种类型，规定了变量有哪些方法
package main

import (
	"fmt"
	"math"
)

// 定义几何体的基本接口
// 只要实现了 area 和 perim 方法的变量都是 geometry 类型
type geometry interface {
	area() float64
	perim() float64
}

// 在我们的例子中，实现 `square` 和 `circle` 方法这个接口
type square struct {
	width, height float64
}
type circle struct {
	radius float64
}

// 要在 Go 中实现一个接口，只需要实现接口中的所有方法。这里让 `square` 实现了 `geometry` 接口
func (s square) area() float64 {
	return s.width * s.height
}
func (s square) perim() float64 {
	return 2*s.width + 2*s.height
}

// `circle` 的实现
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// 如果一个变量是接口类型，可以调用这个被命名的接口中的方法。这里有一个通用的 `measure` 函数，利用这个特性，它可以用在任何 `geometry` 上
func measure(g geometry) {
	fmt.Printf("%#v\n", g)
	fmt.Printf("面积 %f\n", g.area())
	fmt.Printf("周长 %v\n\n", g.perim())
}

func main() {
	s := square{width: 3, height: 4}
	c := circle{radius: 5}

	// 结构体类型 `circle` 和 `square` 都实现了 `geometry` 接口，所以可以使用它们的实例作为 `measure` 的参数
	measure(s)
	measure(c)
}
