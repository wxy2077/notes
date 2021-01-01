/*
* @Time    : 2021-01-01 17:39
* @Author  : CoderCharm
* @File    : curd_test.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :
**/
package demo_mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
	"time"
)

type Employee struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	City string `json:"city"`
}

var emp Employee
var empList []Employee

func dbConn() *sql.DB {
	db, err := sql.Open("mysql", "root:Admin12345-@tcp(172.16.137.129:3306)/temp_db_date?timeout=90s&collation=utf8mb4_unicode_ci")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section. https://github.com/go-sql-driver/mysql#important-settings
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db
}

func SelectDB(db *sql.DB) []Employee {
	rows, err := db.Query("SELECT * FROM employee ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {
		err = rows.Scan(&emp.Id, &emp.Name, &emp.City)
		if err != nil {
			panic(err.Error())
		}
		empList = append(empList, emp)
	}
	return empList
}

func InsertDB(name, city string, db *sql.DB) (lastId int64, affectId int64) {
	rows, err := db.Exec("INSERT INTO employee(`name`, city) VALUE (?, ?)", name, city)
	if err != nil {
		panic("数据插入错误")
	}

	lastId, err = rows.LastInsertId()
	affectId, err = rows.RowsAffected()

	return lastId, affectId

}

func UpdateDB(id int, name string, city string, db *sql.DB) (lastId, affectId int64) {

	rows, err := db.Exec("UPDATE employee SET name=?, city=? WHERE id=?", name, city, id)
	if err != nil {
		panic(fmt.Sprintf("更新错误 %s", err))
	}

	lastId, err = rows.LastInsertId()
	affectId, err = rows.RowsAffected()

	return lastId, affectId
}

func DeleteDB(id int, db *sql.DB) (lastId, affectId int64) {
	//rows, err := db.Exec("UPDATE employee SET name=?, city=? WHERE id=?", name, city, id)
	rows, err := db.Exec("DELETE FROM employee WHERE id=?", id)
	if err != nil {
		panic(fmt.Sprintf("更新错误 %s", err))
	}
	//rows.LastInsertId()
	//lastId := rows.LastInsertId()
	//rows.RowsAffected()

	lastId, err = rows.LastInsertId()
	affectId, err = rows.RowsAffected()

	return lastId, affectId
}

func TestCURD(t *testing.T) {
	db := dbConn()

	defer db.Close()

	res := SelectDB(db)

	fmt.Println(fmt.Sprintf("查询结果: %+v \n", res))

	//lastId, affectId := InsertDB("小明", "wuhan", db)
	//fmt.Println(fmt.Sprintf("插入结果rowsId: %d", lastId))
	//fmt.Println(fmt.Sprintf("影响函数: %d", affectId))

	//lastId, affectId := UpdateDB(3,"小明", "湖北", db)
	//fmt.Println(fmt.Sprintf("插入结果rowsId: %d", lastId))
	//fmt.Println(fmt.Sprintf("影响函数: %d", affectId))

	lastId, affectId := DeleteDB(5, db)
	fmt.Println(fmt.Sprintf("插入结果rowsId: %d", lastId))
	fmt.Println(fmt.Sprintf("影响函数: %d", affectId))

}
