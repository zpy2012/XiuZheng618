package styles

import (
	"github.com/astaxie/beego"
	"fmt"
	"FlowerWeChat/models"
	"FlowerWeChat/controllers/tools"
	"strings"
)

type BgStyleAddController struct {
	beego.Controller
}

func (this *BgStyleAddController) Prepare() {
	s := this.StartSession()
	username = s.Get("login")
	fmt.Println(username)
	if username == nil {
		this.Ctx.Redirect(302, "/admin/login")
	}
}

func (this *BgStyleAddController) Get() {
	this.Data["UserName"] = username
	this.TplName = "bgview/styleadd.html"
}

func (this *BgStyleAddController) Post()  {
	var style models.Styles
	styles, err := models.GetAllStyles(nil,nil,nil,nil,0,0)
	if err != nil {
		var result models.Result
		result.Code = "201"
		result.Msg = err.Error()
		this.Data["json"] = result
		this.ServeJSON()
		return
	}
	if len(styles) >= 3 {
		var result models.Result
		result.Code = "201"
		result.Msg = "达到数量上限！"
		this.Data["json"] = result
		this.ServeJSON()
		return
	}else {
		this.ParseForm(this.Ctx.Input.RequestBody)
		code := this.Ctx.Request.FormValue("stylecode")
		for _, v := range styles{
			m := v.(models.Styles)
			if strings.EqualFold(m.Code,code) {
				var result models.Result
				result.Code = "201"
				result.Msg = "该编号已存在，请选择其他编号！"
				this.Data["json"] = result
				this.ServeJSON()
				return
			}
		}
		style.Title = this.Ctx.Request.FormValue("styletitle")
		style.IsSelected = 1
		style.Code = code
		style.CreatedTime = tools.TimeNow().Format(beego.AppConfig.String("time"))
		if _, err := models.AddStyles(&style); err == nil {
			var result models.Result
			result.Code = "200"
			result.Msg = "上传成功"
			result.JumpUrl = "/admin/style"
			this.Data["json"] = result
			this.ServeJSON()
		}else {
			var result models.Result
			result.Code = "201"
			result.Msg = err.Error()
			this.Data["json"] = result
			this.ServeJSON()
			return
		}
	}
}