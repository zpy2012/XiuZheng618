package index

import (
	"github.com/astaxie/beego"
	"fmt"
	"FlowerWeChat/models"
)

type BgIndexController struct {
	beego.Controller
}

var username interface{}

func (this *BgIndexController) Prepare() {
	s := this.StartSession()
	username = s.Get("login")
	fmt.Println(username)
	if username == nil {
		this.Ctx.Redirect(302, "/admin/login")
	}
}

func (this *BgIndexController) Get() {
	this.Data["UserName"] = username
	var order, sortBy  []string
	order = append(order, "desc")
	sortBy = append(sortBy, "CreatedTime")
	if indexs, err := models.GetAllIndex(nil, nil, sortBy, order, 0, 0); err == nil {
		this.Data["indexs"] = indexs
	}else {
		fmt.Println(err)
	}
	this.TplName = "bgview/index.html"
}
