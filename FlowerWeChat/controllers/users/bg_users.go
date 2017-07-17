package users

import (
	"github.com/astaxie/beego"
	"fmt"
	"FlowerWeChat/models"
	"strconv"
	"FlowerWeChat/controllers/tools"
)

type BgUsersController struct {
	beego.Controller
}

var username interface{}
var prepage, _ = beego.AppConfig.Int("prepage")

func (this *BgUsersController) Prepare() {
	s := this.StartSession()
	username = s.Get("login")
	fmt.Println(username)
	if username == nil {
		this.Ctx.Redirect(302, "/admin/login")
	}
}

func (c *BgUsersController) Get() {
	p := c.Ctx.Request.FormValue("page")
	page, _ := strconv.Atoi(p)
	offset := (page-1)*prepage
	c.Data["UserName"] = username
	var order, sortBy  []string
	order = append(order, "asc")
	sortBy = append(sortBy, "CreatedTime")
	u, _ := models.GetAllUsers(nil, nil, sortBy, order, 0, 0)
	if users, err := models.GetAllUsers(nil, nil, sortBy, order, int64(offset), int64(prepage)); err == nil {
		res := tools.Paginator(page, prepage, int64(len(u)))
		c.Data["paginator"] = res
		c.Data["users"] = users
	}else {
		res := tools.Paginator(page, prepage, 0)
		c.Data["paginator"] = res
		fmt.Println(err)
	}
	c.TplName = "bgview/users.html"
}