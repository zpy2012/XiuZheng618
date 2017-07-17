package index

import (
	"github.com/astaxie/beego"
	"fmt"
	"FlowerWeChat/models"
	"strconv"
	"FlowerWeChat/controllers/tools"
)

type BgIndexController struct {
	beego.Controller
}

var username interface{}
var prepage, _ = beego.AppConfig.Int("prepage")

func (this *BgIndexController) Prepare() {
	s := this.StartSession()
	username = s.Get("login")
	fmt.Println(username)
	if username == nil {
		this.Ctx.Redirect(302, "/admin/login")
	}
}

func (this *BgIndexController) Get() {
	p := this.Ctx.Request.FormValue("page")
	page, _ := strconv.Atoi(p)
	offset := (page-1)*prepage
	this.Data["UserName"] = username
	var order, sortBy  []string
	order = append(order, "desc")
	sortBy = append(sortBy, "CreatedTime")
	i, _ := models.GetAllIndex(nil, nil, sortBy, order, 0, 0)
	if indexs, err := models.GetAllIndex(nil, nil, sortBy, order, int64(offset), int64(prepage)); err == nil {
		res := tools.Paginator(page, prepage, int64(len(i)))
		this.Data["paginator"] = res
		this.Data["indexs"] = indexs
	}else {
		res := tools.Paginator(page, prepage, 0)
		this.Data["paginator"] = res
		fmt.Println(err)
	}
	this.TplName = "bgview/index.html"
}
