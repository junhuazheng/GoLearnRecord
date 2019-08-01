package main

import (
	"blog/models"
	_ "blog/routers"
	"github.com/astaxie/beego"
	_"github.com/go-sql-driver/mysql"
)

func init() {
	models.Init()
	beego.BConfig.WebConfig.Session.SessionOn = true
}

func main() {
	beego.Run()
}

