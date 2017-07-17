package users

import (
	"github.com/astaxie/beego"
	"fmt"
	"FlowerWeChat/models"
	"strconv"
	"FlowerWeChat/controllers/tools"
)

type BgUserOperationController struct {
	beego.Controller
}

func (this *BgUserOperationController) Prepare() {
	s := this.StartSession()
	username = s.Get("login")
	fmt.Println(username)
	if username == nil {
		this.Ctx.Redirect(302, "/admin/login")
	}
}

func (this *BgUserOperationController) Get() {
	this.ParseForm(this.Ctx.Input.RequestBody)
	p := this.Ctx.Request.FormValue("page")
	page, _ := strconv.Atoi(p)
	offset := (page-1)*prepage
	i := this.Ctx.Request.FormValue("id")
	//id, _ := strconv.Atoi(i)
	quary := make(map[string]string)
	quary["user_id"] = i
	o, _ := models.GetAllOperationRecord(quary, nil, nil, nil, 0, 0)
	operations, err := models.GetAllOperationRecord(quary, nil, nil, nil, int64(offset), int64(prepage))
	if err != nil {
		res := tools.Paginator(page, prepage, 0)
		this.Data["UserName"] = username
		this.Data["operations"] = operations
		this.Data["paginator"] = res
		this.Data["userid"] = i
	}else {
		res := tools.Paginator(page, prepage, int64(len(o)))
		this.Data["UserName"] = username
		this.Data["operations"] = operations
		this.Data["paginator"] = res
		this.Data["userid"] = i
	}
	this.TplName = "bgview/usersoperation.html"
}