package time

import (
	"testing"
	"time"
)

func TestTimeFormat(t *testing.T) {
	now := time.Now()
	// 时间格式化 .Format()
	// 时间类型有一个自带的方法Format进行格式化，需要注意的是 Go 语言中格式化时间模板不是常见的 `Y-m-d H:M:S`而是使用Go的诞生时间2006年1月2号15点04分（记忆口诀为2006 1 2 3 4）

	// 如果想格式化为`12`小时方式，需指定`PM`
	// 格式化的模板为Go的出生时间2006年1月2号15点04分 Mon Jan

	// 24小时制
	t.Log(now.Format("2006-01-02 15:04:05.000 Mon Jan"))

	// 12小时制
	t.Log(now.Format("2006-01-02 03:04:05.000 PM Mon Jan"))
	t.Log(now.Format("2006/01/02 15:04"))
	t.Log(now.Format("15:04 2006/01/02"))
	t.Log(now.Format("2006/01/02"))
}
