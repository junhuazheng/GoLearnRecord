mysql操作

go-sql-driver/mysql
go操作mysql的驱动包

使用
package
import (
    "database/sql"
    _"github.com/go-sql-driver/mysql"
)

数据库
在mysql中创建一张测试表，sql如下:
CREATE TABLE `userinfo` (
    `uid` INT(10) NOT NULL AUTO_INCREMENT,
    `username` VARCHAR(64) NULL DEFAULT NULL,
    `departname` VARCHAR(64) NULL DEFAULT NULL,
    `created` DATE NULL DEFAULT NULL,
    PRIMARY KEY(`uid`)
)

连接
db, err := sql.Open("mysql", "用户名:密码@tcp(IP:端口)/数据库?charset=utf8")

insert
有两种方法
1、直接使用Exec函数添加
result, er := db.Exec("INSERT INTO userinfo(username, departname,created) VALUES (?, ?, ?)", "zjh", "dev", "2019-06-05")

2、收线使用Prepare获得stmt，然后调用Exec添加
stmt, err := db.Prepare("INSERT userinfo SET username=?,departname=?,created=?")
res, err := stmt.Exec("iris", "test", "2019-06-05")

另外一个经常用到的功能，获得刚刚添加数据的自增ID
id, err := res,LastInsertId()

