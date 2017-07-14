package products

import (
	"github.com/astaxie/beego"
	"fmt"
	"FlowerWeChat/models"
)

type BgProductsController struct {
	beego.Controller
}

var username interface{}

func (this *BgProductsController) Prepare() {
	s := this.StartSession()
	username = s.Get("login")
	fmt.Println(username)
	if username == nil {
		this.Ctx.Redirect(302, "/admin/login")
	}
}

func (c *BgProductsController) Get() {
	c.Data["UserName"] = username
	var order, sortBy  []string
	order = append(order, "desc")
	sortBy = append(sortBy, "CreatedTime")
	if flowers, err := models.GetAllFlowers(nil, nil, sortBy, order, 0, 0); err == nil {
		c.Data["flowers"] = flowers
	}else {
		fmt.Println(err)
	}
	c.TplName = "bgview/products.html"
}