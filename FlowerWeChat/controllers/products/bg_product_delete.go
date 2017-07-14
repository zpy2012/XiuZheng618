package products

import (
	"github.com/astaxie/beego"
	"fmt"
	"strconv"
	"FlowerWeChat/models"
	"FlowerWeChat/controllers/tools"
)

type BgProductDeleteController struct {
	beego.Controller
}

func (this *BgProductDeleteController) Prepare() {
	s := this.StartSession()
	username = s.Get("login")
	fmt.Println(username)
	if username == nil {
		this.Ctx.Redirect(302, "/admin/login")
	}
}

func (this *BgProductDeleteController) Post() {
	this.ParseForm(this.Ctx.Input.RequestBody)
	i := this.Ctx.Request.FormValue("id")
	fmt.Println("asdfgh",i)
	id, _ := strconv.Atoi(i)
	if flower, err := models.GetFlowersById(id); err == nil {
		fmt.Println(flower)
		tools.DeleteImageWithPath(flower.ImagePath)
		if err := models.DeleteFlowers(id); err == nil {
			fmt.Fprint(this.Ctx.ResponseWriter, "成功")
		}else {
			fmt.Fprint(this.Ctx.ResponseWriter, "删除失败")
		}
	}else {
		fmt.Fprint(this.Ctx.ResponseWriter, "删除失败")
	}
}