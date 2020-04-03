package concurrences

import (
	"runtime"
	"sync"
	"testing"
)

var wg sync.WaitGroup

func TestMaxProcs(t *testing.T) {
	runtime.GOMAXPROCS(2) // 默认CPU的核心数，默认跑满整个CPU

	wg.Add(2)
	go a(t)
	go b(t)

	wg.Wait()
}

func a(t *testing.T) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		t.Logf("A: %d\n", i)
	}
}

func b(t *testing.T) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		t.Logf("B: %d\n", i)
	}
}
