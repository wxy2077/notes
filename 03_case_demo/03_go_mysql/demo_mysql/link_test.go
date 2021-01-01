/*
* @Time    : 2021-01-01 15:25
* @Author  : CoderCharm
* @File    : mysql01_test.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :

https://www.golangprograms.com/example-of-golang-crud-using-mysql-from-scratch.html

**/
package demo_mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
	"time"
)

func TestMySQL(t *testing.T) {
	db, err := sql.Open("mysql", "root:Admin12345-@tcp(172.16.137.129:3306)/temp_db_date?timeout=90s&collation=utf8mb4_unicode_ci")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section. https://github.com/go-sql-driver/mysql#important-settings
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	selDB, err := db.Query("SELECT * FROM employee ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(selDB.Next())

	defer db.Close()

}
