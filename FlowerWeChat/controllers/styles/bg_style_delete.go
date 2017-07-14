package styles

import (
	"github.com/astaxie/beego"
	"fmt"
	"strconv"
	"FlowerWeChat/models"
)

type BgStyleDeleteController struct {
	beego.Controller
}

func (this *BgStyleDeleteController) Prepare() {
	s := this.StartSession()
	username = s.Get("login")
	fmt.Println(username)
	if username == nil {
		this.Ctx.Redirect(302, "/admin/login")
	}
}

func (this *BgStyleDeleteController) Post() {
	this.ParseForm(this.Ctx.Input.RequestBody)
	i := this.Ctx.Request.FormValue("id")
	id, _ := strconv.Atoi(i)
	if err := models.DeleteStyles(id); err == nil {
		fmt.Fprint(this.Ctx.ResponseWriter, "成功")
	}else {
		fmt.Fprint(this.Ctx.ResponseWriter, "删除失败")
	}
}