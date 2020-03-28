package times

import (
	"testing"
	"time"
)

// 定时器
func TestTimer(t *testing.T) {
	// 定时器 Tick
	timer := time.Tick(time.Second) // 每秒钟执行定时器
	for myTimer := range timer {
		t.Log(myTimer)
	}
}
