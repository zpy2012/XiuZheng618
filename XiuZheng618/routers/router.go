package routers

import (
	"XiuZheng618/controllers/bg_controllers"
	"XiuZheng618/controllers/web_controllers"
	"github.com/astaxie/beego"
)

func init() {
    	beego.Router("/admin", &bg_controllers.BgIndexController{})
	beego.Router("/admin/detail", &bg_controllers.BgDetailController{})
	beego.Router("/admin/login", &bg_controllers.BgLoginController{})
	beego.Router("/admin/logout", &bg_controllers.BgLogoutController{})
	beego.Router("/", &web_controllers.WebIndexController{})
	beego.Router("/export", &bg_controllers.ExportController{})
	beego.Router("/start", &web_controllers.WebStartController{})
	beego.Router("/info", &web_controllers.WebInfoController{})
}
