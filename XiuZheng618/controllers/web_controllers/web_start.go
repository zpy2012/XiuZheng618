package web_controllers

import (
	"github.com/astaxie/beego"
	"sync"
	"XiuZheng618/controllers/tools"
	"XiuZheng618/models"
	"fmt"
	"time"
)

type WebStartController struct {
	beego.Controller
}

type Area struct {
	Sum int
	OneStart int
	OneEnd int
	TwoStart int
	TwoEnd int
	ThreeStart int
	ThreeEnd int
	FourStart int
	FourEnd int
	One int
	Two int
	Three int
	Four int
	Lock sync.Mutex
}

var flag int
var s Area

func (this *WebStartController) Get() {
	this.StartSession()
	code := this.GetSession("code")
	if code == nil {
		this.Ctx.Redirect(302, "/")
	}else {
		this.Data["Code"] = code
		this.TplName = "webview/start.html"
	}
}

func (this *WebStartController) Post() {
	this.ParseForm(this.Ctx.Input.RequestBody)
	this.StartSession()
	var r models.AwadResult
	var awd int
	code := this.GetSession("code")
	s.Lock.Lock()
	defer s.Lock.Unlock()
	defer this.DelSession("code")
	if flag == 0 {
		AreaInit(&s)
		flag = 1
	}
	if s.Sum == 0 {
		var errorResult models.ErrorResult
		errorResult.Code = 201
		errorResult.Msg = "奖品已经发放完毕！"
		this.Data["json"] = errorResult
		this.ServeJSON()
		return
	}
	nowTime := tools.TimeNow()
	tm := beego.AppConfig.String("statictime")
	staticTime, _ := time.Parse(beego.AppConfig.String("time"), tm)
	result, _ := models.GetTJackpot618ByCode(code.(string))
	if result.Status == 1 {
		var errorResult models.ErrorResult
		errorResult.Code = 201
		errorResult.Msg = "该抽奖码已经被使用过！"
		this.Data["json"] = errorResult
		this.ServeJSON()
		return
	}
	if nowTime.After(staticTime) && s.One > 0 {
		awd = 1
		result.Award = "一等奖"
		result.Status = 1
		result.UpdateTime = tools.TimeNow().Format(beego.AppConfig.String("time"))
		models.UpdateTJackpot618ById(result)
		ComputeArea(&s, awd)
	}else {
		if s.Sum >= 2 {
			FL: num := tools.GenerateRandomNumber(2, s.Sum)
			fmt.Println(num)
			if num > s.TwoStart && num <= s.TwoEnd {
				if s.Two > 0 {
					awd = 2
					result.Award = "二等奖"
					result.Status = 1
					result.UpdateTime = tools.TimeNow().Format(beego.AppConfig.String("time"))
					models.UpdateTJackpot618ById(result)
					ComputeArea(&s, awd)
				}else {
					goto FL
				}
			}else if num > s.ThreeStart && num <= s.ThreeEnd {
				if s.Three > 0 {
					awd = 3
					result.Award = "三等奖"
					result.Status = 1
					result.UpdateTime = tools.TimeNow().Format(beego.AppConfig.String("time"))
					models.UpdateTJackpot618ById(result)
					ComputeArea(&s, awd)
				}else {
					goto FL
				}
			}else if num > s.FourStart && num <= s.FourEnd {
				if s.Four > 0 {
					awd = 4
					result.Award = "四等奖"
					result.Status = 1
					result.UpdateTime = tools.TimeNow().Format(beego.AppConfig.String("time"))
					models.UpdateTJackpot618ById(result)
					ComputeArea(&s, awd)
				}else {
					goto FL
				}
			}
		}else {
			var errorResult models.ErrorResult
			errorResult.Code = 201
			errorResult.Msg = "奖品已经发放完毕！"
			this.Data["json"] = errorResult
			this.ServeJSON()
			return
		}
	}
	r.Awd = awd
	r.Msg = "成功"
	this.Data["json"] = r
	this.ServeJSON()
}

func ComputeArea(a *Area, awd int)  {
	switch awd {
	case 1:{
		a.Sum -= 1
		a.One -= 1
		a.OneEnd -= 1
		a.TwoStart -= 1
		a.TwoEnd -= 1
		a.ThreeStart -= 1
		a.ThreeEnd -= 1
		a.FourStart -= 1
		a.FourEnd -= 1
	}
	case 2:{
		a.Sum -= 1
		a.Two -= 1
		a.TwoEnd -= 1
		a.ThreeStart -= 1
		a.ThreeEnd -= 1
		a.FourStart -= 1
		a.FourEnd -= 1
	}
	case 3:{
		a.Sum -= 1
		a.Three -= 1
		a.ThreeEnd -= 1
		a.FourStart -= 1
		a.FourEnd -= 1
	}
	case 4:{
		a.Sum -= 1
		a.Four -= 1
		a.FourEnd -= 1
	}
	}
}

//func AreaInit(a *Area)  {
//	a.Sum = 10
//	a.One = 1
//	a.Two = 2
//	a.Three = 3
//	a.Four = 4
//	a.OneStart = 0
//	a.OneEnd = 1
//	a.TwoStart = 1
//	a.TwoEnd = 3
//	a.ThreeStart = 3
//	a.ThreeEnd = 6
//	a.FourStart = 6
//	a.FourEnd = 10
//}

func AreaInit(a *Area)  {
	a.Sum = 1094
	a.One = 1
	a.Two = 100
	a.Three = 400
	a.Four = 593
	a.OneStart = 0
	a.OneEnd = 1
	a.TwoStart = 1
	a.TwoEnd = 101
	a.ThreeStart = 101
	a.ThreeEnd = 501
	a.FourStart = 501
	a.FourEnd = 1094
}