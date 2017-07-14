package index

import (
	"github.com/astaxie/beego"
	"fmt"
	"strconv"
	"FlowerWeChat/models"
	"FlowerWeChat/controllers/tools"
)

type BgIndexDeleteController struct {
	beego.Controller
}

func (this *BgIndexDeleteController) Prepare() {
	s := this.StartSession()
	username = s.Get("login")
	fmt.Println(username)
	if username == nil {
		this.Ctx.Redirect(302, "/admin/login")
	}
}

func (this *BgIndexDeleteController) Post() {
	this.ParseForm(this.Ctx.Input.RequestBody)
	i := this.Ctx.Request.FormValue("id")
	id, _ := strconv.Atoi(i)
	if index, err := models.GetIndexById(id); err == nil {
		tools.DeleteImageWithPath(index.ImagePath)
		if err := models.DeleteIndex(id); err == nil {
			fmt.Fprint(this.Ctx.ResponseWriter, "成功")
		}else {
			fmt.Fprint(this.Ctx.ResponseWriter, "删除失败")
		}
	}else {
		fmt.Fprint(this.Ctx.ResponseWriter, "删除失败")
	}
}