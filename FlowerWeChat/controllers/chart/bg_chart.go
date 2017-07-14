package chart

import (
	"github.com/astaxie/beego"
	"fmt"
	"FlowerWeChat/models"
)

type BgChartController struct {
	beego.Controller
}

var username interface{}

func (this *BgChartController) Prepare() {
	s := this.StartSession()
	username = s.Get("login")
	fmt.Println(username)
	if username == nil {
		this.Ctx.Redirect(302, "/admin/login")
	}
}

func (c *BgChartController) Get() {
	c.Data["UserName"] = username
	var order, sortBy  []string
	order = append(order, "desc")
	sortBy = append(sortBy, "CreatedTime")
	if charts, err := models.GetAllEtiquetteFlowers(nil, nil, sortBy, order, 0, 0); err == nil {
		c.Data["charts"] = charts
	}else {
		fmt.Println(err)
	}
	c.TplName = "bgview/chart.html"
}