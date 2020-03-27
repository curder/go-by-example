package main

import (
	"fmt"
	tools "github.com/curder/go-by-example/src/packages/practices/calc"
)

func main() {
	fmt.Println(10 / 9)

	fmt.Println("1 + 2 + 3 = ", tools.Sum(1, 2, 3))

	fmt.Println("10 - 9 - 8 = ", tools.Minus(10, 9, 8))

	fmt.Println("1 * 2 * 3 = ", tools.MultiplyBy(1, 2, 3))

	fmt.Println("100 / 5 / 2 = ", tools.Division(100, 5, 2))
}
