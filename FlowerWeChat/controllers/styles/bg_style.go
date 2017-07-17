package styles

import (
	"github.com/astaxie/beego"
	"fmt"
	"FlowerWeChat/models"
	"strconv"
	"FlowerWeChat/controllers/tools"
)

type BgStyleController struct {
	beego.Controller
}

var username interface{}
var prepage, _ = beego.AppConfig.Int("prepage")

func (this *BgStyleController) Prepare() {
	s := this.StartSession()
	username = s.Get("login")
	fmt.Println(username)
	if username == nil {
		this.Ctx.Redirect(302, "/admin/login")
	}
}

func (c *BgStyleController) Get() {
	p := c.Ctx.Request.FormValue("page")
	page, _ := strconv.Atoi(p)
	offset := (page-1)*prepage
	c.Data["UserName"] = username
	var order, sortBy  []string
	order = append(order, "asc")
	sortBy = append(sortBy, "Code")
	s, _ := models.GetAllStyles(nil, nil, sortBy, order, 0, 0)
	if styles, err := models.GetAllStyles(nil, nil, sortBy, order, int64(offset), int64(prepage)); err == nil {
		res := tools.Paginator(page, prepage, int64(len(s)))
		c.Data["paginator"] = res
		c.Data["styles"] = styles
	}else {
		res := tools.Paginator(page, prepage, 0)
		c.Data["paginator"] = res
		fmt.Println(err)
	}
	c.TplName = "bgview/style.html"
}