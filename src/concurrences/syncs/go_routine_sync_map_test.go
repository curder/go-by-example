package syncs

import (
	"strconv"
	"sync"
	"testing"
)

// Go语言中内置的map不是并发安全的。
// 如下代码，当多次并发读写map达到一定数量时，会出现：fatal error: concurrent map writes 错误。
// 修复这种问题可以通过写加锁的方式，比如 sync.RWMutex。
// 但是go针对这种经常使用的操作map的方式提供了开箱即用的 sync.Map
func TestGoRoutineSetAndGetMapValue(t *testing.T) {
	var wg sync.WaitGroup
	m := make(map[string]interface{})

	setValue := func(key string, value interface{}) {
		m[key] = value
	}

	getValue := func(key string) interface{} {
		return m[key]
	}

	wg.Add(2)
	for i := 0; i < 2; i++ { // 当增加并发数量则会出现报错
		go func(n int) {
			defer wg.Done()

			key := strconv.Itoa(n)

			setValue(key, n)

			t.Logf("key = %s, value = %v", key, getValue(key))
		}(i)
	}

	wg.Wait()
}

// 通过 sync.Map 的方式对读写map操作并发安全
func TestGoRoutineSetMapValueBySync(t *testing.T) {
	var wg sync.WaitGroup
	m := sync.Map{} // 开箱即用表示不用像内置的map一样使用make函数初始化就能直接使用

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(n int) {
			defer wg.Done()

			key := strconv.Itoa(n)
			m.Store(key, n) // 存储map

			value, _ := m.Load(key) // 通过key获取map的值
			t.Logf("key = %s, value = %v", key, value)
		}(i)
	}

	wg.Wait()
}
