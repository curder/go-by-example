package others

import (
	"fmt"
	"strconv"
	"testing"
)

// TestTransformStringToInt 字符串转整型
func TestTransformStringToIntByParseInt(t *testing.T) {
	i := "54321"
	result, _ := strconv.ParseInt(i, 10, 64) // 从字符串解析成整型
	t.Logf("\n转换前的类型是：%T，值是：%#v；\n转换后的类型是：%T，值是：%#v\n", i, i, result, result)
}

// TestTransformStringToIntByAtoi 字符串转整型
func TestTransformStringToIntByAtoi(t *testing.T) {
	i := "54321"
	result, _ := strconv.Atoi(i)
	t.Logf("\n转换前的类型是：%T，值是：%#v；\n转换后的类型是：%T，值是：%#v\n", i, i, result, result)
}

// TestTransformIntToString 整型转字符串
func TestTransformIntToStringBySprintf(t *testing.T) {
	var i int32 = 54321
	result := fmt.Sprintf("%d", i)
	t.Logf("\n转换前的类型是：%T，值是：%#v；\n转换后的类型是：%T，值是：%#v\n", i, i, result, result)
}

// TestTransformStringToIntByItoa 整型转字符串
func TestTransformStringToIntByItoa(t *testing.T) {
	i := 54321
	result := strconv.Itoa(i)
	t.Logf("\n转换前的类型是：%T，值是：%#v；\n转换后的类型是：%T，值是：%#v\n", i, i, result, result)
}

// TestTransformStringToBool 字符串解析成布尔值
func TestTransformStringToBool(t *testing.T) {
	i := "true"
	result, _ := strconv.ParseBool(i)
	t.Logf("\n转换前的类型是：%T，值是：%#v；\n转换后的类型是：%T，值是：%#v\n", i, i, result, result)
}

// TestTransformStringToFloat 字符串解析成浮点数
func TestTransformStringToFloat(t *testing.T) {
	i := "5.4321"
	result, _ := strconv.ParseFloat(i, 64)
	t.Logf("\n转换前的类型是：%T，值是：%#v；\n转换后的类型是：%T，值是：%#v\n", i, i, result, result)
}
