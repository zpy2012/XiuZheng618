package bg_controllers

import (
	"github.com/astaxie/beego"
	"XiuZheng618/models"
	"fmt"
	"strings"
	"XiuZheng618/controllers/tools"
)

type BgLoginController struct {
	beego.Controller
}

var username interface{}

func (c *BgLoginController) Get() {
	Prepare()
	c.Data["UserName"] = username
	c.TplName = "bgview/login.html"
}

func (c *BgLoginController) Post ()  {
	c.ParseForm(c.Ctx.Input.RequestBody)
	username := c.Ctx.Request.FormValue("username")
	password := c.Ctx.Request.FormValue("password")
	password = tools.EncryptionPassWord(password, "zpy")
	manager, err := models.GetManagerByName(username)
	if err != nil {
		fmt.Fprint(c.Ctx.ResponseWriter, err)
	}
	if manager == nil {
		fmt.Fprint(c.Ctx.ResponseWriter, "用户不存在")
	}else if !strings.EqualFold(manager.Password, password) {
		fmt.Fprint(c.Ctx.ResponseWriter, "密码错误")
	}else {
		s := c.StartSession()
		s.Set("login", manager.Username)
		//c.Ctx.Redirect(302, "/admin")
		fmt.Fprint(c.Ctx.ResponseWriter, "登陆成功")
	}
}