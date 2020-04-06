package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// 初始化MySQL连接
func init() {
	var err error
	dns := "root:@tcp(127.0.0.1:3306)/go_by_example"

	db, err = sql.Open("mysql", dns)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
}
