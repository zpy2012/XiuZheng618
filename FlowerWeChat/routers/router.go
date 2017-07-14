package routers

import (
	"FlowerWeChat/controllers/index"
	"FlowerWeChat/controllers/chart"
	"github.com/astaxie/beego"
	"FlowerWeChat/controllers"
	"FlowerWeChat/controllers/products"
	"FlowerWeChat/controllers/styles"
	"FlowerWeChat/controllers/users"
	"FlowerWeChat/controllers/wechat"
)

func init() {
    	beego.Router("/", &controllers.WebIndexController{})
	beego.Router("/viewdetail", &controllers.WebDetailController{})
	beego.Router("/viewdetail/flowerdetail", &controllers.WebFlowerDetailController{})
	beego.Router("/admin", &index.BgIndexController{})
	beego.Router("/admin/login", &controllers.BgLoginController{})
	beego.Router("/admin/logout", &controllers.BgLogoutController{})
	beego.Router("/admin/chart", &chart.BgChartController{})
	beego.Router("/admin/style", &styles.BgStyleController{})
	beego.Router("/admin/products", &products.BgProductsController{})
	beego.Router("/admin/userlist", &users.BgUsersController{})
	beego.Router("/admin/userlist/useroperation", &users.BgUserOperationController{})
	beego.Router("/admin/index/add", &index.BgIndexAddController{})
	beego.Router("/admin/index/edit", &index.BgIndexEditController{})
	beego.Router("/admin/index/delete", &index.BgIndexDeleteController{})
	beego.Router("/admin/chart/add", &chart.BgChartAddController{})
	beego.Router("/admin/chart/edit", &chart.BgChartEditController{})
	beego.Router("/admin/chart/delete", &chart.BgChartDeleteController{})
	beego.Router("/admin/products/add", &products.BgProductAddController{})
	beego.Router("/admin/products/edit", &products.BgProductEditController{})
	beego.Router("/admin/products/delete", &products.BgProductDeleteController{})
	beego.Router("/admin/style/add", &styles.BgStyleAddController{})
	beego.Router("/admin/style/edit", &styles.BgStyleEditController{})
	beego.Router("/admin/style/delete", &styles.BgStyleDeleteController{})
	beego.Router("/wechat/auth", &wechat.WeChatAuthController{}, "*:ProcRequest")
	beego.Router("/wechat/auth/login", &wechat.WeChatAuthLoginController{})
	beego.Router("/wechat/auth/redirect", &wechat.WeChatAuthLoginController{}, "*:Redirect")
	beego.Router("/MP_verify_G5nTSndQ8eFMpvZJ.txt", &wechat.WeChatAuthController{}, "*:File")
}