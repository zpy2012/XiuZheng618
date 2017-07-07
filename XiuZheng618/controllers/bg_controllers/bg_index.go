package bg_controllers

import (
	"github.com/astaxie/beego"
	"XiuZheng618/models"
)

type BgIndexController struct {
	beego.Controller
}

func (this *BgIndexController) Prepare() {
	s := this.StartSession()
	username = s.Get("login")
	if username == nil {
		this.Ctx.Redirect(302, "/admin/login")
	}
}

func (c *BgIndexController) Get() {
	c.Data["UserName"] = username
	var order, sortBy  []string
	order = append(order, "desc")
	sortBy = append(sortBy, "UpdateTime")

	query := make(map[string]string)
	query["Status"] = "1"

	record, _ := models.GetAllTJackpot618(query, nil, sortBy, order, 0, 10)
	c.Data["recods"] = record
	c.Data["sum"] = len(record)

	query["Award"] = "一等奖"
	r1, _ := models.GetAllTJackpot618(query, nil, sortBy, order, 0, 0)
	c.Data["one"] = len(r1)

	query["Award"] = "二等奖"
	r2, _ := models.GetAllTJackpot618(query, nil, sortBy, order, 0, 0)
	c.Data["two"] = len(r2)

	query["Award"] = "三等奖"
	r3, _ := models.GetAllTJackpot618(query, nil, sortBy, order, 0, 0)
	c.Data["three"] = len(r3)

	query["Award"] = "四等奖"
	r4, _ := models.GetAllTJackpot618(query, nil, sortBy, order, 0, 0)
	c.Data["four"] = len(r4)

	c.TplName = "bgview/index.html"
}