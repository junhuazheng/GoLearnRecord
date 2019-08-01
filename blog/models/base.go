package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
)

/*
initialize the databse connection
register User into orm
*/

func Init() {
	dbhost := beego.AppConfig.String("dbhost")
	dbport := beego.AppConfig.String("dbport")
	dbuser := beego.AppConfig.String("dbuser")
	dbpassword := beego.AppConfig.String("dbpassword")
	dbname := beego.AppConfig.String("dbname")
	if dbport == "" {
		dbport = "3306"
	}
	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + ")/" + dbname + "?charset=utf8&loc=Local"
	orm.RegisterDataBase("default", "mysql", dsn)
	orm.RegisterModel(new(User))
}

//returns the name of the prefixed table
func TableName(str string) string {
	return beego.AppConfig.String("dbprefix") + str
}