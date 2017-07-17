package main

import (
	_ "FlowerWeChat/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	beego.SetStaticPath("/admin/static", "static")
	beego.SetStaticPath("/admin/index/static", "static")
	beego.SetStaticPath("/admin/chart/static", "static")
	beego.SetStaticPath("/admin/products/static", "static")
	beego.SetStaticPath("/admin/style/static", "static")
	beego.SetStaticPath("/admin/userlist/static", "static")
	beego.SetStaticPath("/viewdetail/static", "static")
	orm.RegisterDataBase("default", "mysql", "root:12345@tcp(127.0.0.1:3306)/FlowerWeChat")
}

func main() {
	beego.Run()
}