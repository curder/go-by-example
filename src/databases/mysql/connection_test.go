package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

// 使用MySQL之前需要下载驱动，使用命令下载：go get -u github.com/go-sql-driver/mysql

// 使用MySQL驱动
// func Open(driverName, dataSourceName string) (*DB, error)
func TestConnection(t *testing.T) {
	dataSourceName := "root:@tcp(127.0.0.1:3306)/go_by_example" // "user:password@tcp(127.0.0.1:3306)/dbname"

	db, err := sql.Open("mysql", dataSourceName) // 仅校验dataSourceName格式是否正确
	if err != nil {
		panic(err.Error())
	}

	err = db.Ping() // 发送一个校验的包
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
}
