package models

import (
	"regexp"
	"strings"

	"github.com/astaxie/beego/orm"
	_"github.com/go-sql-driver/mysql"
)

var db orm.Ormer

type MovieInfo struct {
	Id int64
	Movie_id int64
	Movie_name string 
	Movie_pic string
	Movie_director string
	
}