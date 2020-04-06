package mysql

import (
	"fmt"
	"testing"
)

// 创建数据库表
func TestCreateTable(t *testing.T) {
	// 创建表结构
	sqlStr := `
CREATE TABLE IF NOT EXISTS users (
  id bigint(20) NOT NULL AUTO_INCREMENT,
  name varchar(255) NOT NULL,
  age tinyint(4) DEFAULT NULL,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
`

	_, err := db.Exec(sqlStr)
	if err != nil {
		t.Log("create table error", err)
	}
}

// 创建数据行
func TestCreateTableRow(t *testing.T) {
	// 编写查询语句
	sqlStr := `INSERT INTO users(name, age) VALUES (?, ?)`
	// 执行查询语句
	result, err := db.Exec(sqlStr, "curder", 29) // 预处理
	if err != nil {
		panic(err)
		return
	}

	theID, err := result.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get last insert ID failed, err:%v\n", err)
		return
	}

	fmt.Printf("insert success, the id is %d.\n", theID)
}
