package chart

import (
	"github.com/astaxie/beego"
	"fmt"
	"FlowerWeChat/models"
	"strconv"
	"FlowerWeChat/controllers/tools"
)

type BgChartEditController struct {
	beego.Controller
}

func (this *BgChartEditController) Prepare() {
	s := this.StartSession()
	username = s.Get("login")
	fmt.Println(username)
	if username == nil {
		this.Ctx.Redirect(302, "/admin/login")
	}
}

func (this *BgChartEditController) Get() {
	fmt.Printf("%v", etiquetteFlower)
	this.Data["UserName"] = username
	this.Data["ChartTitle"] = etiquetteFlower.Title
	this.Data["ChartId"] = etiquetteFlower.Id
	this.Data["ChartCode"] = etiquetteFlower.Code
	this.Data["ImagePath"] = etiquetteFlower.ImagePath
	this.TplName = "bgview/chartedit.html"
}

var etiquetteFlower *models.EtiquetteFlowers

func (this *BgChartEditController) Post() {
	this.ParseForm(this.Ctx.Input.RequestBody)
	i := this.Ctx.Request.FormValue("id")
	id, _ := strconv.Atoi(i)
	e := this.Ctx.Request.FormValue("edit")
	edit, _ := strconv.ParseBool(e)
	if edit == true {
		if in, err := models.GetEtiquetteFlowersById(id); err == nil {
			etiquetteFlower = in
			fmt.Fprint(this.Ctx.ResponseWriter,"成功")
		}else {
			fmt.Fprint(this.Ctx.ResponseWriter,"失败")
		}
	}else {
		this.Ctx.Request.ParseMultipartForm(50)
		etiquetteFlower.Title = this.Ctx.Request.FormValue("charttitle")
		file, handler, err := this.Ctx.Request.FormFile("chartimage")
		if err == nil {
			tools.DeleteImageWithPath(etiquetteFlower.ImagePath)
			filePath := beego.AppConfig.String("chartpath")
			if imgpath, err := tools.SaveImage(file,handler,int64(etiquetteFlower.Id),filePath);err == nil {
				etiquetteFlower.ImagePath = imgpath
			}else {
				fmt.Fprint(this.Ctx.ResponseWriter, err)
			}
		}
		if err := models.UpdateEtiquetteFlowersById(etiquetteFlower);err ==nil {
			var result models.Result
			result.Code = "200"
			result.Msg = "上传成功"
			result.JumpUrl = "/admin/chart"
			this.Data["json"] = result
			this.ServeJSON()
			//fmt.Fprint(this.Ctx.ResponseWriter, "/admin")
		}else {
			fmt.Fprint(this.Ctx.ResponseWriter, err)
		}
	}
}