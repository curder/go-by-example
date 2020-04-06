package mysql

import (
	"fmt"
	"testing"
)

// 更新行
func TestUpdate(t *testing.T) {
	sqlStr := "UPDATE users SET age=? WHERE id = ?"
	ret, err := db.Exec(sqlStr, 19, 1)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("update success, affected rows:%d\n", n)
}
