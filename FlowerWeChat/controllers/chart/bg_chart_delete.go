package chart

import (
	"github.com/astaxie/beego"
	"fmt"
	"strconv"
	"FlowerWeChat/models"
	"FlowerWeChat/controllers/tools"
)

type BgChartDeleteController struct {
	beego.Controller
}

func (this *BgChartDeleteController) Prepare() {
	s := this.StartSession()
	username = s.Get("login")
	fmt.Println(username)
	if username == nil {
		this.Ctx.Redirect(302, "/admin/login")
	}
}

func (this *BgChartDeleteController) Post() {
	this.ParseForm(this.Ctx.Input.RequestBody)
	i := this.Ctx.Request.FormValue("id")
	id, _ := strconv.Atoi(i)
	if index, err := models.GetEtiquetteFlowersById(id); err == nil {
		tools.DeleteImageWithPath(index.ImagePath)
		if err := models.DeleteEtiquetteFlowers(id); err == nil {
			fmt.Fprint(this.Ctx.ResponseWriter, "成功")
		}else {
			fmt.Fprint(this.Ctx.ResponseWriter, "删除失败")
		}
	}else {
		fmt.Fprint(this.Ctx.ResponseWriter, "删除失败")
	}
}