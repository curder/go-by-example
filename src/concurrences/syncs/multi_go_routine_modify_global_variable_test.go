package syncs

import (
	"sync"
	"testing"
)

var result int64
var wg sync.WaitGroup

// 开启了两个goroutine去累加变量x的值，这两个goroutine在访问和修改x变量的时候就会存在数据竞争，导致最后的结果与期待的不符
func TestGoRoutine(t *testing.T) {
	wg.Add(1)
	go calculatorSum() // 开启一个go routine
	go calculatorSum() // 再开启一个 go routine
	wg.Wait()
	t.Log(result) // 正确的值应该是 1000000，100 万
}

// 定义一个修改全局变量 result 的函数
func calculatorSum() {
	for i := 0; i < 500000; i++ { // 循环 50 万次
		result = result + 1
	}
	wg.Done()
}
