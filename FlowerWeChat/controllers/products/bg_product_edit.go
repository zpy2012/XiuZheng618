package products

import (
	"github.com/astaxie/beego"
	"fmt"
	"strconv"
	"FlowerWeChat/controllers/tools"
	"FlowerWeChat/models"
)

type BgProductEditController struct {
	beego.Controller
}

func (this *BgProductEditController) Prepare() {
	s := this.StartSession()
	username = s.Get("login")
	fmt.Println(username)
	if username == nil {
		this.Ctx.Redirect(302, "/admin/login")
	}
}

func (this *BgProductEditController) Get() {
	fmt.Printf("%v", flower)
	this.Data["UserName"] = username
	this.Data["FlowerName"] = flower.Name
	this.Data["FlowerId"] = flower.Id
	this.Data["flowerdescription"] = flower.Description
	this.Data["ImagePath"] = flower.ImagePath
	this.TplName = "bgview/productedit.html"
}

var flower *models.Flowers

func (this *BgProductEditController) Post() {
	this.ParseForm(this.Ctx.Input.RequestBody)
	i := this.Ctx.Request.FormValue("id")
	id, _ := strconv.Atoi(i)
	e := this.Ctx.Request.FormValue("edit")
	edit, _ := strconv.ParseBool(e)
	if edit == true {
		if in, err := models.GetFlowersById(id); err == nil {
			flower = in
			fmt.Fprint(this.Ctx.ResponseWriter,"成功")
		}else {
			fmt.Fprint(this.Ctx.ResponseWriter,"失败")
		}
	}else {
		this.Ctx.Request.ParseMultipartForm(50)
		flower.Name = this.Ctx.Request.FormValue("flowername")
		flower.Description = this.Ctx.Request.FormValue("flowerdescription")
		file, handler, err := this.Ctx.Request.FormFile("flowerimage")
		if err == nil {
			tools.DeleteImageWithPath(flower.ImagePath)
			filePath := beego.AppConfig.String("flowerpath")
			if imgpath, err := tools.SaveImage(file,handler,int64(flower.Id),filePath);err == nil {
				flower.ImagePath = imgpath
			}else {
				fmt.Fprint(this.Ctx.ResponseWriter, err)
			}
		}
		if err := models.UpdateFlowersById(flower);err ==nil {
			var result models.Result
			result.Code = "200"
			result.Msg = "上传成功"
			result.JumpUrl = "/admin/products"
			this.Data["json"] = result
			this.ServeJSON()
		}else {
			fmt.Fprint(this.Ctx.ResponseWriter, err)
		}
	}
}