package main

import "fmt"

// 递归适合处理那些问题的规模越来越小的场景
// 给定一个数，求一个数的阶乘
// 例如：给定的数是5，那么它的阶乘就是 5 * 4 * 3 * 2 * 1

func f1(n uint64) uint64 {
	if n <= 1 {
		return 1
	}
	return n * f1(n-1)
}

// 可能性的问题
// n个台阶，可以选择一次走一个台阶，也可以走两个台阶，一共有多少种可能
func recursive(n int) int {
	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	return n + recursive(n-1) + recursive(n-2)
}

func main() {
	var n uint64 = 10
	fmt.Printf("%d 的阶乘的值是：%d \n", n, f1(n))

	fmt.Printf("当需要上%d\n", recursive(4))
}
