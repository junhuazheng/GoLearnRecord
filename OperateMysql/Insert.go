package main

import (
	"time"
	"fmt"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
)

var (
	dbhost = "52.1.251.11:3306"
	dbusername = "root"
	dbpassword = "gyy981010"
	dbname = "gyy"
)

func main() {
	Insert("zheng", "dev", "1")
	Insert("zheng", "dev", "2")
	Insert("iris", "test", "1")
	Insert("iris", "test", "2")
}

func CheckErr(err error) {
	if err != nil {
		fmt.Println("err: ", err.Error())
		panic(err)
	}
}

//获取sql.DB对象
func GetDB() *sql.DB {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", dbusername, dbpassword, dbhost, dbname))
	CheckErr(err)
	return db
}

//插入数据
func Insert(username, departname, method string) bool {
	db := GetDB()
	defer db.Close()

	if method == "1" {
		_, err := db.Exec("insert into userinfo(username,departname,created) values(?,?,?)",username,departname,time.Now())
		if err != nil {
			fmt.Println("insert err: ", err.Error())
			return false
		}
		fmt.Println("insert success!")
		return true
	}else if method == "2" {
		stmt, err := db.Prepare("INSERT userinfo SET username=?,deparname=?,created=?")
		if err != nil {
			fmt.Println("insert prepare error: ", err.Error())
			return false
		}
		_, err = stmt.Exec(username, departname, time.Now())
		if err != nil {
			fmt.Println("insert exec error: ", err.Error())
			return false
		}
		fmt.Println("insert success!")
		return true
	}
	return false
}