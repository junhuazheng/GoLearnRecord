package main

import (
	"fmt"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "user:password@test") //对应数据库的用户名和密码
	defer db.Close()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("success")
	}
	rows, err := db.Query("SELECT number FROM squarenum")
	if err != nil {
		panic(err)
		return
	}
	for rows.Next() {
		var name int
		err = rows.Scan(&name)
		if err != nil {
			panic(err)
		}
		fmt.Println(name)
	}
}