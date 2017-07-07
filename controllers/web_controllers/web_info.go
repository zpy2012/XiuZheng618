package web_controllers

import (
	"github.com/astaxie/beego"
	"XiuZheng618/models"
	"time"
	"fmt"
	"XiuZheng618/controllers/tools"
)

type WebInfoController struct {
	beego.Controller
}

func (this *WebInfoController) Get() {
	this.StartSession()
	code := this.GetSession("code")
	if code == nil {
		this.Ctx.Redirect(302, "/")
	}else {
		this.TplName = "webview/info.html"
	}
}

func (this *WebInfoController) Post() {
	nowTime := tools.TimeNow()
	t := beego.AppConfig.String("endtime")
	endTime, _ := time.Parse(beego.AppConfig.String("time"), t)
	if nowTime.After(endTime) {
		fmt.Fprint(this.Ctx.ResponseWriter, "活动已经结束了哦！")
	}else {
		this.StartSession()
		this.ParseForm(this.Ctx.Input.RequestBody)
		code := this.GetSession("code").(string)
		result, _ := models.GetTJackpot618ByCode(code)
		result.Province = this.Ctx.Request.FormValue("province")
		result.Area = this.Ctx.Request.FormValue("city")
		result.UserName = this.Ctx.Request.FormValue("name")
		result.Telephone = this.Ctx.Request.FormValue("phone")
		result.TrugstoreAddress = this.Ctx.Request.FormValue("address")
		result.ZipCode = this.Ctx.Request.FormValue("zipcode")
		models.UpdateTJackpot618ById(result)
	}
	this.Ctx.Redirect(302, "/start")
}