package users

import (
	"github.com/astaxie/beego"
	"fmt"
	"FlowerWeChat/models"
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

func (c *BgUserOperationController) Get() {
	c.Data["UserName"] = username
	c.Data["operations"] = UserOperations
	c.TplName = "bgview/usersoperation.html"
}

var UserOperations []interface{}

func (this *BgUserOperationController) Post() {
	this.ParseForm(this.Ctx.Input.RequestBody)
	i := this.Ctx.Request.FormValue("id")
	//id, _ := strconv.Atoi(i)
	quary := make(map[string]string)
	quary["user_id"] = i
	operations, err := models.GetAllOperationRecord(quary, nil, nil, nil, 0, 0)
	if err != nil {
		fmt.Fprint(this.Ctx.ResponseWriter, err)
	}else {
		UserOperations = operations
		fmt.Fprint(this.Ctx.ResponseWriter, "成功")
	}
}