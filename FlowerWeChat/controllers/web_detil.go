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
	var operation models.OperationRecord
	id := this.GetString("userid")
	uid, _ := strconv.Atoi(id)
	operation.UserId = uid

	e := this.GetString("eflower")
	eFlowerid, _ := strconv.Atoi(e)
	eFlower, _ := models.GetEtiquetteFlowersById(eFlowerid)
	operation.EtiquetteId = eFlower.Title

	sex := this.GetString("sex")
	switch sex {
	case "0":operation.Sex = "男"
	case "1":operation.Sex = "女"
	}

	osex := this.GetString("othersex")
	switch osex {
	case "0":operation.OtherSex = "男"
	case "1":operation.OtherSex = "女"
	}

	rel := this.GetString("relations")
	switch rel {
	case "0":operation.Relationship = "男神/女神"
	case "1":operation.Relationship = "朋友"
	case "2":operation.Relationship = "长辈"
	case "3":operation.Relationship = "最佳情侣"
	}
	s := this.GetString("styles")
	styleid, _ := strconv.Atoi(s)
	style, _ := models.GetStylesById(styleid)
	operation.Style = style.Title

	operation.CreatedTime = tools.TimeNow().Format(beego.AppConfig.String("time"))
	operation.FlowerCode = operation.Relationship+"、"+operation.Style

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