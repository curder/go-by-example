package calc

// Sum 求和
func Sum(a ...int) int {
	var result int

	for _, val := range a {
		result += val
	}

	return result
}

// Minus 求差
func Minus(a ...int) int {
	var result int

	for key, val := range a {
		if key == 0 {
			result = val
		} else {
			result -= val
		}
	}

	return result
}

// MultiplyBy 乘法
func MultiplyBy(a ...int) int {
	var result = 1

	for _, val := range a {
		result *= val
	}

	return result
}

// Division 除法
// 注意在go语言中 10 / 3 得到的结果是：3
func Division(a ...int) interface{} {
	var result int

	for key, val := range a {
		if val == 0 {
			panic("error params")
		}

		if key == 0 {
			result = val
		} else {
			result /= val
		}
	}

	return result
}
