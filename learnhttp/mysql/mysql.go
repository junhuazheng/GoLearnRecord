package mysql

import (
	"fmt"
	_"github.com/go-sql-driver/mysql"
	"github.com/imoiron/sqlx"
)

type Person struct {
	UserId int `db:"user_id"`
	Username string `db:"username"`
	Sex string `db:"sex"`
	Email string `db:"email"`
}

type Place struct {
	Country string `db:"country"`
	City string `db:"city"`
	TelCode int `db:"telcode"`
}

var Db *sqlx.DB 

func init() {
	database, err := sqlx.Open("mysql", "root:@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println("open mysql failed, ", err)
		return
	}
	Db = database
}

func mysql() {
	r, err := Db.Exec
}