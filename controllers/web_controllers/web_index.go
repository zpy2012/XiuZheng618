package web_controllers

import (
	"github.com/astaxie/beego"
	"XiuZheng618/models"
	"fmt"
	"time"
	"XiuZheng618/controllers/tools"
)

type WebIndexController struct {
	beego.Controller
}

func (this *WebIndexController) Get() {
	this.TplName = "webview/index.html"
}

func (this *WebIndexController) Post() {
	nowTime := tools.TimeNow()
	tm := beego.AppConfig.String("starttime")
	startTime, _ := time.Parse(beego.AppConfig.String("time"), tm)
	t := beego.AppConfig.String("endtime")
	endTime, _ := time.Parse(beego.AppConfig.String("time"), t)
	if nowTime.Before(startTime) {
		fmt.Fprint(this.Ctx.ResponseWriter, "活动还没有开始哦！")
	}else if nowTime.After(endTime) {
		fmt.Fprint(this.Ctx.ResponseWriter, "活动已经结束了哦！")
	}else {
		this.ParseForm(this.Ctx.Input.RequestBody)
		code := this.Ctx.Request.FormValue("code")
		if result, err := models.GetTJackpot618ByCode(code);err != nil {
			fmt.Fprint(this.Ctx.ResponseWriter, "您输入的抽奖码不正确！")
		}else if result.Status == 1 {
			fmt.Fprint(this.Ctx.ResponseWriter, "该抽奖码已经被使用过！")
		}else {
			this.StartSession()
			this.SetSession("code",code)
			fmt.Fprint(this.Ctx.ResponseWriter, "成功")
		}
	}
}