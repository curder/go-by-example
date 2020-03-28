package times

import (
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	// 获取当前时间
	now := time.Now()

	t.Log(now) // 当前时间对象
	t.Log(now.Year()) // 获取年
	t.Log(now.Month()) // 获取月
	t.Log(now.Day()) // 获取日
	t.Log(now.Hour()) // 获取时
	t.Log(now.Minute()) // 获取分
	t.Log(now.Second()) // 获取秒

	// 时间戳
	t.Log(now.Unix()) // 时间戳
	t.Log(now.UnixNano()) // 时间戳纳秒数

	// 时间格式转换 time.Unix()
	ret := time.Unix(now.Unix(), 0) // 将时间戳转为时间格式
	t.Logf("%d-%02d-%02d %02d:%02d:%02d\n", ret.Year(), ret.Month(), ret.Day(), ret.Hour(), ret.Minute(), ret.Second())

}
