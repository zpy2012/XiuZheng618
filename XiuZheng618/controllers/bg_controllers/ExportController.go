package bg_controllers

import (
	"github.com/astaxie/beego"
	"XiuZheng618/controllers/tools"
	"XiuZheng618/models"
	"fmt"
)

type ExportController struct {
	beego.Controller
}

func (this *ExportController) Get() {
	quary := make(map[string]string)
	quary["Status"] = "1"
	data, _ := models.GetAllTJackpot618(quary,nil,nil,nil,0,0)
	if success, err := tools.ExportExcelWithData(data); success == false {
		fmt.Fprint(this.Ctx.ResponseWriter, err)
	}else {
		fmt.Fprint(this.Ctx.ResponseWriter, success)
	}
}