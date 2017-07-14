package styles

import (
	"github.com/astaxie/beego"
	"fmt"
	"FlowerWeChat/models"
)

type BgStyleController struct {
	beego.Controller
}

var username interface{}

func (this *BgStyleController) Prepare() {
	s := this.StartSession()
	username = s.Get("login")
	fmt.Println(username)
	if username == nil {
		this.Ctx.Redirect(302, "/admin/login")
	}
}

func (c *BgStyleController) Get() {
	c.Data["UserName"] = username
	var order, sortBy  []string
	order = append(order, "asc")
	sortBy = append(sortBy, "Code")
	if styles, err := models.GetAllStyles(nil, nil, sortBy, order, 0, 0); err == nil {
		c.Data["styles"] = styles
	}else {
		fmt.Println(err)
	}
	c.TplName = "bgview/style.html"
}