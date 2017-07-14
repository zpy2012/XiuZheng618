package controllers

import (
	"github.com/astaxie/beego"
	"strconv"
	"FlowerWeChat/models"
	"fmt"
	"html/template"
)

type WebFlowerDetailController struct {
	beego.Controller
}

func (c *WebFlowerDetailController) Get() {
	c.Data["Name"] = flower.Name
	c.Data["ImagePath"] = flower.ImagePath
	c.Data["Description"] = template.HTML(flower.Description)
	c.TplName = "webview/details.html"
}

var flower *models.Flowers

func (this *WebFlowerDetailController) Post() {
	this.ParseForm(this.Ctx.Input.RequestBody)
	id := this.Ctx.Request.FormValue("flowerid")
	flowerid, _ := strconv.Atoi(id)
	if f, err := models.GetFlowersById(flowerid); err == nil {
		flower = f
		fmt.Fprint(this.Ctx.ResponseWriter, "成功")
	}else {
		fmt.Fprint(this.Ctx.ResponseWriter, err)
	}
}