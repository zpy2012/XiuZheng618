package controllers

import (
	"github.com/astaxie/beego"
	"FlowerWeChat/models"
	"FlowerWeChat/controllers/tools"
	"strconv"
	"fmt"
)

type WebDetailController struct {
	beego.Controller
}

var flowers []interface{}

func (c *WebDetailController) Get() {
	c.Data["flowers"] = flowers
	c.TplName = "webview/viewdetails.html"
}

func (this *WebDetailController) Post() {
	fmt.Println("111111111111")
	var operation models.OperationRecord
	this.ParseForm(this.Ctx.Input.RequestBody)
	id := this.Ctx.Request.FormValue("userid")
	uid, _ := strconv.Atoi(id)
	operation.UserId = uid

	e := this.Ctx.Request.FormValue("eflower")
	eFlowerid, _ := strconv.Atoi(e)
	eFlower, _ := models.GetEtiquetteFlowersById(eFlowerid)
	operation.EtiquetteId = eFlower.Title

	sex := this.Ctx.Request.FormValue("sex")
	switch sex {
	case "0":operation.Sex = "男"
	case "1":operation.Sex = "女"
	}

	osex := this.Ctx.Request.FormValue("othersex")
	switch osex {
	case "0":operation.Sex = "男"
	case "1":operation.Sex = "女"
	}

	rel := this.Ctx.Request.FormValue("relations")
	switch rel {
	case "0":operation.Relationship = "男神/女神"
	case "1":operation.Relationship = "朋友"
	case "2":operation.Relationship = "长辈"
	case "3":operation.Relationship = "最佳情侣"
	}
	s := this.Ctx.Request.FormValue("styles")
	styleid, _ := strconv.Atoi(s)
	style, _ := models.GetStylesById(styleid)
	operation.Style = style.Title

	operation.CreatedTime = tools.TimeNow().Format(beego.AppConfig.String("time"))
	fmt.Println(operation)
	models.AddOperationRecord(&operation)
	query := make(map[string]string)
	query["code"] = operation.Relationship+"、"+operation.Style
	if fls, err := models.GetAllFlowers(query, nil, nil, nil, 0, 0); err == nil {
		flowers = fls
		fmt.Fprint(this.Ctx.ResponseWriter,"成功")
	}else {
		fmt.Fprint(this.Ctx.ResponseWriter,"失败")
	}
}