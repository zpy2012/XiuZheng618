package users

import (
	"github.com/astaxie/beego"
	"fmt"
	"FlowerWeChat/models"
)

type BgUsersController struct {
	beego.Controller
}

var username interface{}

func (this *BgUsersController) Prepare() {
	s := this.StartSession()
	username = s.Get("login")
	fmt.Println(username)
	if username == nil {
		this.Ctx.Redirect(302, "/admin/login")
	}
}

func (c *BgUsersController) Get() {
	c.Data["UserName"] = username
	var order, sortBy  []string
	order = append(order, "asc")
	sortBy = append(sortBy, "CreatedTime")
	if users, err := models.GetAllUsers(nil, nil, sortBy, order, 0, 0); err == nil {
		c.Data["users"] = users
	}else {
		fmt.Println(err)
	}
	c.TplName = "bgview/users.html"
}