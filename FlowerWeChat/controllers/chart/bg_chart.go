package chart

import (
	"github.com/astaxie/beego"
	"fmt"
	"FlowerWeChat/models"
	"strconv"
	"FlowerWeChat/controllers/tools"
)

type BgChartController struct {
	beego.Controller
}

var username interface{}
var prepage, _ = beego.AppConfig.Int("prepage")

func (this *BgChartController) Prepare() {
	s := this.StartSession()
	username = s.Get("login")
	fmt.Println(username)
	if username == nil {
		this.Ctx.Redirect(302, "/admin/login")
	}
}

func (c *BgChartController) Get() {
	p := c.Ctx.Request.FormValue("page")
	page, _ := strconv.Atoi(p)
	offset := (page-1)*prepage
	c.Data["UserName"] = username
	var order, sortBy  []string
	order = append(order, "desc")
	sortBy = append(sortBy, "CreatedTime")
	ch, _ := models.GetAllEtiquetteFlowers(nil, nil, sortBy, order, 0, 0)
	if charts, err := models.GetAllEtiquetteFlowers(nil, nil, sortBy, order, int64(offset), int64(prepage)); err == nil {
		res := tools.Paginator(page, prepage, int64(len(ch)))
		c.Data["paginator"] = res
		c.Data["charts"] = charts
	}else {
		res := tools.Paginator(page, prepage, 0)
		c.Data["paginator"] = res
		fmt.Println(err)
	}
	c.TplName = "bgview/chart.html"
}