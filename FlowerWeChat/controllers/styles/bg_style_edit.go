package styles

import (
	"github.com/astaxie/beego"
	"fmt"
	"FlowerWeChat/models"
	"strconv"
)

type BgStyleEditController struct {
	beego.Controller
}

func (this *BgStyleEditController) Prepare() {
	s := this.StartSession()
	username = s.Get("login")
	fmt.Println(username)
	if username == nil {
		this.Ctx.Redirect(302, "/admin/login")
	}
}

func (this *BgStyleEditController) Get() {
	this.Data["UserName"] = username
	this.Data["StyleTitle"] = style.Title
	this.Data["StyleId"] = style.Id
	this.Data["StyleCode"] = style.Code
	this.TplName = "bgview/styleedit.html"
}

var style *models.Styles

func (this *BgStyleEditController) Post() {
	this.ParseForm(this.Ctx.Input.RequestBody)
	i := this.Ctx.Request.FormValue("id")
	id, _ := strconv.Atoi(i)
	e := this.Ctx.Request.FormValue("edit")
	edit, _ := strconv.ParseBool(e)
	if edit == true {
		if in, err := models.GetStylesById(id); err == nil {
			style = in
			fmt.Fprint(this.Ctx.ResponseWriter,"成功")
		}else {
			fmt.Fprint(this.Ctx.ResponseWriter,"失败")
		}
	}else {
		this.ParseForm(this.Ctx.Input.RequestBody)
		style.Title = this.Ctx.Request.FormValue("styletitle")
		if err := models.UpdateStylesById(style);err ==nil {
			var result models.Result
			result.Code = "200"
			result.Msg = "上传成功"
			result.JumpUrl = "/admin/style"
			this.Data["json"] = result
			this.ServeJSON()
		}else {
			fmt.Fprint(this.Ctx.ResponseWriter, err)
		}
	}
}