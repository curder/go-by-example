package mysql

import (
	"fmt"
	"testing"
)

// 查询单行
// 单行查询db.QueryRow()执行一次查询，并期望返回最多一行结果（即Row）。
// QueryRow总是返回非nil的值，直到返回值的Scan方法被调用时，才会返回被延迟的错误。
// func (db *DB) QueryRow(query string, args ...interface{}) *Row
func TestReadSignalRow(t *testing.T) {
	// 声明数据库结构体
	type user struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	var u user

	// 编写查询语句
	sqlStr := `SELECT * FROM users WHERE id = ?;`
	err := db.QueryRow(sqlStr, 1).Scan(&u.Id, &u.Name, &u.Age)
	if err != nil {
		panic(err)
	}

	t.Logf("%#v", u)
}

// 多行查询db.Query()执行一次查询，返回多行结果（即Rows），一般用于执行select命令。
// 参数args表示query中的占位参数。
// func (db *DB) Query(query string, args ...interface{}) (*Rows, error)
func TestMultiRows(t *testing.T) {
	sqlStr := `SELECT * FROM users WHERE age > ?`
	results, err := db.Query(sqlStr, 18)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	defer results.Close() // 非常重要：关闭rows释放持有的数据库链接

	// 声明数据库结构体
	type user struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	// 循环读取结果集中的数据
	for results.Next() {
		var u user
		err := results.Scan(&u.Id, &u.Name, &u.Age)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("id:%d name:%s age:%d\n", u.Id, u.Name, u.Age)
	}
}

// 预处理查询
// 优化MySQL服务器重复执行SQL的方法，可以提升服务器性能，提前让服务器编译，一次编译多次执行，节省后续编译的成本。
// 避免SQL注入问题。
func TestPrepareQueryRows(t *testing.T) {
	sqlStr := "SELECT * FROM users WHERE id > ?"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		t.Logf("prepare failed, err:%v\n", err)
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(0)
	if err != nil {
		t.Logf("query failed, err:%v\n", err)
		return
	}
	defer rows.Close()

	// 声明数据库结构体
	type user struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	// 循环读取结果集中的数据
	for rows.Next() {
		var u user
		err := rows.Scan(&u.Id, &u.Name, &u.Age)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("id:%d name:%s age:%d\n", u.Id, u.Name, u.Age)
	}
}
