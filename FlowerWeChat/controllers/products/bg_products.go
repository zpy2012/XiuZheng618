package products

import (
	"github.com/astaxie/beego"
	"fmt"
	"FlowerWeChat/models"
	"strconv"
	"FlowerWeChat/controllers/tools"
)

type BgProductsController struct {
	beego.Controller
}

var username interface{}
var prepage, _ = beego.AppConfig.Int("prepage")

func (this *BgProductsController) Prepare() {
	s := this.StartSession()
	username = s.Get("login")
	fmt.Println(username)
	if username == nil {
		this.Ctx.Redirect(302, "/admin/login")
	}
}

func (c *BgProductsController) Get() {
	p := c.Ctx.Request.FormValue("page")
	page, _ := strconv.Atoi(p)
	offset := (page-1)*prepage
	c.Data["UserName"] = username
	var order, sortBy  []string
	order = append(order, "desc")
	sortBy = append(sortBy, "CreatedTime")
	f, _ := models.GetAllFlowers(nil, nil, sortBy, order, 0, 0)
	if flowers, err := models.GetAllFlowers(nil, nil, sortBy, order, int64(offset), int64(prepage)); err == nil {
		res := tools.Paginator(page, prepage, int64(len(f)))
		c.Data["paginator"] = res
		c.Data["flowers"] = flowers
	}else {
		res := tools.Paginator(page, prepage, 0)
		c.Data["paginator"] = res
		fmt.Println(err)
	}
	c.TplName = "bgview/products.html"
}