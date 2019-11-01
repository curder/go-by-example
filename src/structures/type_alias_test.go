package structures

import "testing"

// 类型别名
// 类型别名规定：TypeAlias只是Type的别名。
// 本质上TypeAlias与Type是同一个类型。

type youInt = int

func TestTypeAlias(t *testing.T) {
	var n1 youInt
	n1 = 100
	t.Logf("n1的值为：%d，类型为：%T \n", n1, n1)

	// rune 也是一个int32的类型别名
	var n2 rune
	n2 = '中'
	t.Logf("n2的值为：%d，类型为：%T \n", n2, n2)
}
