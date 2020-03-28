package times

import (
	"testing"
	"time"
)

func TestTimeDuration(t *testing.T) {
	now := time.Now()
	// 时间间隔 Before() After() Equal()
	t.Log(time.Second) // 获取一秒钟

	// 可以使用 `Add` 将时间后移一个时间间隔，或者使用一个 `-` 来将时间前移一个时间间隔
	t.Log(now.Add(24 * time.Hour)) // 当前之间添加24小时，即一天后
	// 这些方法来比较两个时间，分别测试一下是否是之前，之后或者是同一时刻，精确到秒
	t.Log(now.Before(now))
	t.Log(now.After(now))
	t.Log(now.Equal(now))
}
