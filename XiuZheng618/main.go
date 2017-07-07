package main

import (
	_ "XiuZheng618/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	beego.SetStaticPath("/admin/static", "static")
	beego.SetStaticPath("/admin/detail/static", "static")
	orm.RegisterDataBase("default", "mysql", "zpy:12345@tcp(127.0.0.1:3306)/XiuZheng")
}

func main() {
	beego.Run()
}
