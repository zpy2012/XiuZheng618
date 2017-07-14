package controllers

import "github.com/astaxie/beego"

type BgLogoutController struct {
	beego.Controller
}

func (c *BgLogoutController) Get() {
	s := c.StartSession()
	s.Delete("login")
	c.Ctx.Redirect(302, "/admin/login")
}
